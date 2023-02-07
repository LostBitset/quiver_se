// A program that instruments Node.js programs to emit a sequence of which callbacks run
// This isn't guarenteed to be perfectly accurate, as it doesn't have access to the runtime itself
// Instead, it exploits the (admittedly very strange) "nextTickQueue" to check whether a function
// call is a new callback being invoked or not

import seafox from "seafox";
const { parseScript } = seafox;

import { strict as assert } from "node:assert";

import { readFile } from "node:fs";

function main(filename) {
	let new_filename = filename.replace(/\.js$/, ".INSTRUMENTED-cbstream.js");
	console.log(`Starting (cbstream) instrumentation of file "./${filename}"...`);
	console.log(`(This is not intended to be a destructive operation, output will be saved as "./${new_filename}")`);
	readFile(filename, (err, contents_buf) => {
		if (err) throw err;
		let contents = contents_buf.toString();
		console.log(`Read file into memory (${contents.length} bytes).`)
		let estree = parseScript(contents);
		console.log(`Parsed via seafox (estree.type == "${estree.type}").`);
		console.log(`Starting instrumentation...`);
		writeInstrumented(estree, new_filename);
		console.log(`All done. Instrumented version saved as "./${new_filename}".`);
	});
}

function writeInstrumented(estree, outfile) {
	console.log("[TODO everything]");
}

// @UnitTest
assert.equal(parseScript("console.log(42);").body[0].expression.type, "CallExpression");

// @PEBCAK
assert.equal(process.argv.length, 3)

let argument = process.argv[2];

// @PEBCAK
assert.notEqual(argument.match(/\.js$/), null)

main(argument)

