
// This code has been instrumented by function_info/instrument.js

// bgn decl-prefix (static)
function _Q$xnH(e) { throw e; }
// end decl-prefix (static)

// bgn entry-point (has-script)
function _Q$ent() {

var fs = require('node:fs');

function eq(x, y) {
	"!!MAGIC@js_concolic/src-range=30:68";
	
	return x === y;
}

fs.readFile('something.txt', 'utf8', function (err, contents) {
	"!!MAGIC@js_concolic/src-range=107:397";
	
	if (err !== null) {
		throw new Error(err.toString());
	} else {
		fs.readFile('something2.txt', 'utf8', function (err, contents2) {
	"!!MAGIC@js_concolic/src-range=240:390";
	
			if (err !== null) {
				throw new Error(err.toString());
			} else {
				console.log(eq(contents, contents2));
			}
		});
	}
});


}
// end entry-point (has-script)

// bgn main-rescue (actual-entry-point)
try {
	_Q$ent();
} catch (e) {
	_Q$xnH(e);
}
// end main-rescue (actual-entry-point)

