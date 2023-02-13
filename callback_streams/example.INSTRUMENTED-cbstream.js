// bgn imports-raw (raw)

import fs from 'fs';

// end imports-raw (raw)


// This code has been instrumented by cbstream.js
// to emit the stream of callbacks that are run

// bgn decl-prefix (static)

let _Q$hCb = false; // shorthand: has callback

// shorthand: callback handler
function _Q$cbH(i) {
	console.log("[CBSTREAM] Got: " + i);
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
	console.log("[CBSTREAM] Done.");
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


fs.readFile('something.txt', 'utf8', (err, contents) => {":::MAGIC@js_concolic/arg-names|||err, contents";_Q$cCb(0);
	if (err !== null) {
		throw new Error(err.toString());
	} else {
		fs.readFile('something2.txt', 'utf8', (err, contents2) => {":::MAGIC@js_concolic/arg-names|||err, contents2";_Q$cCb(1);
			if (err !== null) {
				throw new Error(err.toString());
			} else {
				let eq = (x, y) => { ":::MAGIC@js_concolic/arg-names|||x, y";_Q$cCb(2); return ((x === y)); };
				console.log(eq(contents, contents2));
			}
		});
	}
});


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

