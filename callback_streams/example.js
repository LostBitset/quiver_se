import fs from 'fs';

fs.readFile('something.txt', 'utf8', (err, contents) => {
	if (err !== null) {
		throw new Error(err.toString());
	} else {
		fs.readFile('something2.txt', 'utf8', (err, contents2) => {
			if (err !== null) {
				throw new Error(err.toString());
			} else {
				let eq = (x, y) => (x === y);
				console.log(eq(contents, contents2));
			}
		});
	}
});

