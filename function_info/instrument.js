// A program that instruments functions inside Node.js programs with standalone
// string literals containing the following information:
// 1. Where the function occurs in the source code (start and end bytes)
// 2. All of the identifiers that the function accesses
// These are encoded in the following form:
// 1. "!!MAGIC@function_info/src-range=<start>:<end>";
// 2. "!!MAGIC@function_info/idents=<ident0>:<ident1>:<ident2>:<...>";
// It also wraps the code in a function, and passes exceptions to a handler

import seafox from "seafox";
const { parseModule } = seafox;

import { strict as assert } from "node:assert";

import { readFile, writeFile } from "node:fs";

var debug;
var stats_idents = 0;
var stats_funs = 0;
var stats_subobjects = 0;

function conlog(...args) {
	if (debug) console.log("INFO[function_info] ", ...args);
	return undefined;
}

function main(filename) {
	let new_filename = filename.replace(/\.js$/, "._fninf.js");
	conlog(`Starting (function_info) instrumentation of file "./${filename}"...`);
	readFile(filename, (err, contents_buf) => {
		if (err) throw err;
		let contents = contents_buf.toString();
		conlog(`Read file into memory (${contents.length} bytes).`)
		let estree = parseModule(contents, {loc: true});
		conlog(`Parsed via seafox (estree.type == "${estree.type}").`);
		conlog(`Starting instrumentation...`);
		let instrumented = instrument(contents, estree, new_filename);
		conlog(`Added information to ${stats_funs} functions, and saw a total of ${stats_idents} identifiers.`);
		conlog(`A total of ${stats_subobjects} ESTree subobjects were walked through.`);
		conlog(`Instrumented code generated (${Buffer.byteLength(instrumented, "utf8")} bytes).`);
		writeFile(new_filename, instrumented, (err) => {
			if (err) throw err;
			conlog(`All done. Instrumented version saved as "./${new_filename}".`);
		});
	});
}

function replaceIndexRange(str, start, end, repl) {
	return str.substring(0, start) + repl + str.substring(end);
}

function instrument(contents, estree) {
	let code = contents;
	let offset = 0;
	for (const [rinject, f_estree] of estreeBlockFunctions(estree)) {
		let inject = rinject - offset;
		let lbefore = code.length;
		code = replaceIndexRange(
			code,
			inject + 1,
			inject + 1,
			injectionForBlockFunction(f_estree),
		);
		offset += lbefore - code.length;
		stats_funs++;
	}
	return INSTRUMENTATION_OUTER_TEMPLATE.replace("<%=SCRIPT%>", code);
}

function isEstreeSubObject(estree_value) {
	if (typeof estree_value != "object") return false;
	if (!estree_value) return false;
	if (!estree_value.hasOwnProperty("type")) return false;
	if (estree_value.type == "Identifier") return true;
	return estree_value.hasOwnProperty("start") && estree_value.hasOwnProperty("end");
}

function* estreeSubObjectsOfType(estree_obj, t) {
	if (estree_obj.hasOwnProperty("type") && estree_obj.type === t) {
		yield estree_obj;
	}
	for (const [_, v] of Object.entries(estree_obj)) {
		if (isEstreeSubObject(v)) {
			yield* estreeSubObjectsOfType(v, t);
		} else if (v instanceof Array) {
			for (const v_sub of v) {
				if (isEstreeSubObject(v_sub)) {
					yield* estreeSubObjectsOfType(v_sub, t);
				}
			}
		}
	}
}

function* estreeSubObjects(estree_obj) {
	yield estree_obj;
	for (const [_, v] of Object.entries(estree_obj)) {
		if (isEstreeSubObject(v)) {
			yield* estreeSubObjects(v);
		} else if (v instanceof Array) {
			for (const v_sub of v) {
				if (isEstreeSubObject(v_sub)) {
					yield* estreeSubObjects(v_sub);
				}
			}
		}
	}
}

function* estreeBlockFunctions(estree) {
	for (const sub of estreeSubObjects(estree.body)) {
		stats_subobjects++;
		if (!sub.hasOwnProperty("type")) continue;
		if (sub.type == "FunctionExpression") {
			if (sub.body.type == "BlockStatement") {
				let inject = sub.body.start;
				yield [inject, sub];
			}
		} else if (sub.type == "FunctionDeclaration") {
			if (sub.hasOwnProperty("generator") && !sub.generator) {
				let inject = sub.body.start;
				yield [inject, sub];
			}
		}
	}
}

function injectionForBlockFunction(estree) {
	return `
	"!!MAGIC@js_concolic/src-range=${estree.start}:${estree.end}";
	`;
}

// @UnitTest
assert.equal(parseModule("console.log(42);").body[0].expression.type, "CallExpression");

// @PEBCAK
assert(process.argv.length == 3 || process.argv.length == 4);

let argument = process.argv[2];
debug = (process.argv[3] || "") == "debug";

// @PEBCAK
assert.notEqual(argument.match(/\.js$/), null)

const INSTRUMENTATION_OUTER_TEMPLATE = `
// This code has been instrumented by function_info/instrument.js

// bgn decl-prefix (static)
function _Q$xnH(e) { throw e; }
// end decl-prefix (static)

// bgn entry-point (has-script)
function _Q$ent() {

<%=SCRIPT%>

}
// end entry-point (has-script)

// bgn main-rescue (actual-entry-point)
try {
	_Q$ent();
} catch (e) {
	_Q$xnH(e);
}
// end main-rescue (actual-entry-point)

`;

main(argument)

