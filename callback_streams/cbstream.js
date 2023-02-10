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
		let estree = parseModule(contents);
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
	let code = contents;
	let cb_id = 0;
	for (const [start, end, ...rest] of estreeFunctions(estree)) {
		code = replaceIndexRange(
			code,
			start,
			end,
			instrumentFunction(...rest, cb_id),
		);
		cb_id++;
	}
	return INSTRUMENTATION_OUTER_TEMPLATE.replace("<%=SCRIPT%>", code);
}

function* estreeImports(estree) {
	conlog('TODO!!!');
}

function* estreeFunctions(estree) {
	conlog('TODO!!!');
}

function instrumentFunction(orig, estree, id) {
	conlog('TODO!!!');
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

