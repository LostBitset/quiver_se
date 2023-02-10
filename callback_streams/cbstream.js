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
	conlog(`(Output will be saved as "./${new_filename}")`);
	readFile(filename, (err, contents_buf) => {
		if (err) throw err;
		let contents = contents_buf.toString();
		conlog(`Read file into memory (${contents.length} bytes).`)
		let estree = parseScript(contents);
		conlog(`Parsed via seafox (estree.type == "${estree.type}").`);
		conlog(`Starting instrumentation...`);
		writeInstrumented(estree, new_filename);
		conlog(`All done. Instrumented version saved as "./${new_filename}".`);
	});
}

function writeInstrumented(estree, outfile) {
	conlog("[TODO everything]");
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

