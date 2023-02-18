import fs from 'fs';

function eq(x, y) {
	return x === y;
}

fs.readFile('something.txt', 'utf8', function (err, contents) {
	if (err !== null) {
		throw new Error(err.toString());
	} else {
		fs.readFile('something2.txt', 'utf8', function (err, contents2) => {
			if (err !== null) {
				throw new Error(err.toString());
			} else {
				console.log(eq(contents, contents2));
			}
		});
	}
});
