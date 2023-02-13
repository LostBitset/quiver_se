// A program that instruments Node.js programs to emit a sequence of which callbacks run
// This isn't guarenteed to be perfectly accurate, as it doesn't have access to the runtime itself
// Instead, it exploits the (admittedly very strange) "nextTickQueue" to check whether a function
// call is a new callback being invoked or not

import seafox from "seafox";
const { parseModule } = seafox;

import { strict as assert } from "node:assert";

import { readFile, writeFile } from "node:fs";

var debug;
var s_imports_moved = 0;
var s_injected = 0;
var s_wrapped = 0;
var s_subobjects = 0;

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
		conlog(`Walked through ESTree repr and found ${s_subobjects} subobjects.`);
		conlog(`Moved ${s_imports_moved} import(s), injected into ${s_injected} fn(s), and wrapped ${s_wrapped} value a-fn(s). `);
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
	let ims = "// bgn imports-raw (raw)\n\n";
	for (const [rstart, rend] of estreeImports(estree)) {
		let [start, end] = [rstart - offset, rend - offset];
		let im = code.substring(start, end);
		let lbefore = code.length;
		code = replaceIndexRange(code, start, end, "");
		offset += lbefore - code.length;
		ims += `${im}\n`;
		s_imports_moved++;
	}
	ims += "\n// end imports-raw (raw)\n\n";
	let cb_id = 0;
	for (const [rinject, fn_estree] of estreeBlockFunctions(estree)) {
		let inject = rinject - offset;
		let lbefore = code.length;
		code = replaceIndexRange(
			code,
			inject + 1,
			inject + 1,
			injectionForBlockFunction(
				cb_id,
				argNamesMagic(
					fn_estree,
					(start, end) => code.substring(start - offset, end - offset),
				),
			),
		);
		offset += lbefore - code.length;
		cb_id++;
		s_injected++;
	}
	for (const [rwrap_start, rwrap_end, fn_estree] of estreeValueFunctions(estree)) {
		let [wrap_start, wrap_end] = [rwrap_start - offset, rwrap_end - offset];
		while (code[wrap_start - 1] == "(") {
			wrap_start--;
			wrap_end++;
		}
		code = replaceIndexRange(
			code,
			wrap_start,
			wrap_end,
			wrapForValueFunction(
				code.substring(wrap_start, wrap_end),
				cb_id,
				argNamesMagic(
					fn_estree,
					(start, end) => code.substring(start - offset, end - offset),
				),
			),
		);
		cb_id++;
		s_wrapped++;
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
	if (typeof estree_value != "object") return false;
	if (!estree_value) return false;
	if (!estree_value.hasOwnProperty("type")) return false;
	if (estree_value.type == "Identifier") return false;
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
		s_subobjects++;
		if (!sub.hasOwnProperty("type")) continue;
		if (sub.type == "ArrowFunctionExpression") {
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

function* estreeValueFunctions(estree) {
	for (const sub of estreeSubObjects(estree.body)) {
		if (!sub.hasOwnProperty("type")) continue;
		if (sub.type == "ArrowFunctionExpression") {
			if (sub.body.type != "BlockStatement") {
				let wrap_start = sub.body.start;
				let wrap_end = sub.body.end;
				yield [wrap_start, wrap_end, sub];
			}
		}
	}
}

function injectionForBlockFunction(id, magic) {
	return `${magic || ""}${magic ? ";" : ""}_Q$cCb(${id});`;
}

function wrapForValueFunction(orig, id, magic) {
	return `{ ${magic || ""}${magic ? ";" : ""}${injectionForBlockFunction(id)} return (${orig}); }`;
}

function argNamesMagic(estree, substringFn) {
	return `":::MAGIC@js_concolic/arg-names|||${argNamesMagicInner(estree, substringFn)}"`;
}

function argNamesMagicInner(estree, substringFn) {
	switch (estree.type) {
		case "ArrowFunctionExpression":
			if (estree.params.length <= 0) {
				return "";
			} else {
				let params_start = estree.params[0].start;
				let params_end = estree.params[estree.params.length-1].end;
				return (substringFn)(params_start, params_end);
			}
		case "FunctionDeclaration":
			//...
			break;
		default:
			return null;
	}
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
// This code has been instrumented by cbstream.js
// to emit the stream of callbacks that are run

// bgn decl-prefix (static)

let _Q$hCb = false; // shorthand: has callback

// shorthand: callback handler
function _Q$cbH(i) {
	${debug?"console.log(\"[CBSTREAM] Got: \" + i);":""}
}

// shorthand: claim callback
function _Q$cCb(i) {
	let h = _Q$hCb;
	_Q$hCb = true;
	if (!h) _Q$cbH(i);
	process.nextTick(() => {
		_Q$hCb = false;
	});
}

function _Q$end() {
	${debug?"console.log(\"[CBSTREAM] Done.\");":""}
}

// shorthand: exception handler
function _Q$xnH(e, d) {
	_Q$cbH(-2); // virtual callback "fail"
	_Q$end();
	if (!d) process.removeListener("uncaughtException", _Q$xnH);
	throw e;
}

// end decl-prefix (static)
// bgn entry-point (wraps-instrumented)

// shorthand: entry point
function _Q$ent() {
<%=SCRIPT%>
}

// end entry-point (wraps-instrumented)
// bgn main-rescue (static)

// note: transient, this binding will remove itself
// to observe errors without redirecting them
process.on("uncaughtException", _Q$xnH);

process.on("beforeExit", _Q$end);

// note: actual entry point for instrumented code
try {
	_Q$cCb(-1); // virtual callback "top"
	_Q$ent();
} catch (e) {
	_Q$xnH(e, true);
}

// end main-rescue (static)

`;

main(argument)

