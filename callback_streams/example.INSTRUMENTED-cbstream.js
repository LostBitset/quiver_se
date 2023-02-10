
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
console.log(1 + 1);


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

