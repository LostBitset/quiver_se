// A simple program that just prints out a JSON representation of
// the output of the seafox ESTree parser

import seafox from 'seafox';
const { parseModule } = seafox;

import { strict as assert } from 'node:assert';

import { readFile } from 'node:fs';

// @PEBCAK
assert(process.argv.length == 3 || process.argv.length == 4);

readFile(process.argv[2], (err, contents_buf) => {
	if (err) throw err;
	let contents = contents_buf.toString();
	let estree;
	if (process.argv.length > 3 && process.argv[3] == "loc") {
		estree = parseModule(contents, {loc: true});
	} else {
		estree = parseModule(contents);
	}
	console.log(JSON.stringify(estree));
});

