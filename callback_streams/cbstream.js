// A program that instruments Node.js programs to emit a sequence of which callbacks run
// This isn't guarenteed to be perfectly accurate, as it doesn't have access to the runtime itself
// Instead, it exploits the (admittedly very strange) "nextTickQueue" to check whether a function
// call is a new callback being invoked or not

import seafox from "seafox";
const { parseModule } = seafox;

import { strict as assert } from "node:assert";

import { readFile, writeFile } from "node:fs";

var debug;

function conlog(...args) {
	if (debug) console.log("INFO[cbstream@node] ", ...args);
	return undefined;
}

function main(filename) {
	let new_filename = filename.replace(/\.js$/, ".INSTRUMENTED-cbstream.js");
	conlog(`Starting (cbstream) instrumentation of file "./${filename}"...`);
	readFile(filename, (err, contents_buf) => {
		if (err) throw err;
		let contents = contents_buf.toString();
		conlog(`Read file into memory (${contents.length} bytes).`)
		let estree = parseModule(contents, {loc: true});
		conlog(`Parsed via seafox (estree.type == "${estree.type}").`);
		conlog(`Starting instrumentation...`);
		let instrumented = instrument(contents, estree, new_filename);
		writeFile(new_filename, instrumented, (err) => {
			if (err) throw err;
			conlog(`All done. Instrumented version saved as "./${new_filename}".`);
		});
	});
}

function replaceIndexRange(str, start, end, repl) {
	return str.substring(0, start) + repl + str.substring(end);
}

const INSTRUMENTATION_OUTER_TEMPLATE = `
// This code has been instrumented by cbstream.js
// to emit the stream of callbacks that are run

// bgn decl-prefix (static)

let _Q$hCb = false;
let _Q$sen = 0;

function _Q$cbH(i) {
	_Q$sen = (_Q$sen + i) % 8;
}

function _Q$cCb(i) {
	let h = _Q$hCb;
	_Q$hCb = true;
	if (!h) _Q$cbH(i);
	process.nextTick(() => {
		_Q$hCb = false;
	});
}

function _Q$end() {
	if (_Q$sen == 0) eval("");
}

// end decl-prefix (static)
// bgn entry-point (wraps-instrumented)

function _Q$ent() {
<%=SCRIPT%>
}

// end entry-point (wraps-instrumented)
// bgn main-rescue (static)

try {
	_Q$cCb(-1); // virtual callback "top"
	_Q$ent();
	_Q$end();
} catch (e) {
	_Q$cbH(-2); // virtual callback "fail"
	_Q$end();
	throw e;
}

// end main-rescue (static)

`;

function instrument(contents, estree) {
	console.log(Array.from(estreeSubObjects(estree.body)));
	let code = contents;
	let ims = "// bgn imports-raw (raw)\n\n";
	for (const [start, end] of estreeImports(estree)) {
		let im = code.substring(start, end);
		code = replaceIndexRange(code, start, end, "");
		ims += `${im}\n`;
	}
	ims += "\n// end imports-raw (raw)\n\n";
	let cb_id = 0;
	for (const [inject] of estreeBlockFunctions(estree)) {
		code = replaceIndexRange(
			code,
			inject,
			inject,
			injectionForBlockFunction(cb_id),
		);
		cb_id++;
	}
	for (const [wrap_start, wrap_end] of estreeBlockFunctions(estree)) {
		code = replaceIndexRange(
			code,
			wrap_start,
			wrap_end,
			wrapForValueFunction(
				code.substring(wrap_start, wrap_end),
				cb_id,
			),
		);
		cb_id++;
	}
	return ims + INSTRUMENTATION_OUTER_TEMPLATE.replace("<%=SCRIPT%>", code);
}

function* estreeImports(estree) {
	for (const i in estree.body) {
		let tl = estree.body[i];
		if (tl.type == "ImportDeclaration") {
			let { start, end } = tl;
			yield [start, end];
		}
	}
}

function isEstreeSubObject(estree_value) {
	if (typeof estree_value !== "object") return false;
	if (!estree_value) return false;
	return estree_value.hasOwnProperty("start") && estree_value.hasOwnProperty("end");
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
		if (!sub.hasOwnProperty("type")) continue;
		if (sub.type == "ArrowFunctionExpression") {
			if (sub.body.type == "BlockStatement") {
				let inject = sub.body.start;
				yield [inject];
			}
		} else if (sub.type == "FunctionDeclaration") {
			if (sub.hasOwnProperty("generator") && !sub.generator) {
				let inject = sub.body.start;
				yield [inject];
			}
		}
	}
}

function* estreeValueFunctions(estree) {
	for (const sub of estreeSubObjects(estree.body)) {
		if (!sub.hasOwnProperty("type")) continue;
		if (sub.type == "ArrowFunctionExpression") {
			if (sub.body.type !== "BlockStatement") {
				let wrap_start = sub.body.start;
				let wrap_end = sub.body.end;
				yield [wrap_start, wrap_end];
			}
		}
	}
}

function injectionForBlockFunction(id) {
	return `_Q$cCb(${id});`;
}

function wrapForValueFunction(orig, id) {
	return `{ ${injectionForBlockFunction(id)} return (${orig}); }`;
}

// @UnitTest
assert.equal(parseModule("console.log(42);").body[0].expression.type, "CallExpression");

// @PEBCAK
assert(process.argv.length == 3 || process.argv.length == 4);

let argument = process.argv[2];
debug = (process.argv[3] || "") == "debug";

// @PEBCAK
assert.notEqual(argument.match(/\.js$/), null)

main(argument)

