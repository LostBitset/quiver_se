// A program that instruments Node.js programs to emit a sequence of which callbacks run
// This isn't guarenteed to be perfectly accurate, as it doesn't have access to the runtime itself
// Instead, it exploits the (admittedly very strange) "nextTickQueue" to check whether a function
// call is a new callback being invoked or not

import seafox from "seafox";
const { parseScript } = seafox;

import { strict as assert } from "node:assert";

import { readFile } from "node:fs";

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
		let estree = parseScript(contents);
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

const INSTRUMENTATION_PREFIX = `
// begin INSTRUMENTATION_PREFIX

// [TODO everything]

// end INSTRUMENTATION_PREFIX
`;

function instrument(contents, estree) {
	let code = `${INSTRUMENTATION_PREFIX}\n${contents}`;
	let cb_id = 0;
	estreeFunctions(estree).forEach(([start, end, orig]) => {
		code = replaceIndexRange(
			code,
			start,
			end,
			instrumentFunction(orig, cb_id),
		);
		cb_id++;
	});
	return code;
}

// @UnitTest
assert.equal(parseScript("console.log(42);").body[0].expression.type, "CallExpression");

// @PEBCAK
assert(process.argv.length == 3 || process.argv.length == 4);

let argument = process.argv[2];
debug = (process.argv[3] || "") == "debug";

// @PEBCAK
assert.notEqual(argument.match(/\.js$/), null)

main(argument)

