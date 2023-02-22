
// This code has been instrumented by function_info/instrument.js

// bgn decl-prefix (static)
function _Q$xnH(e) { throw e; }
// end decl-prefix (static)

// bgn entry-point (has-script)
function _Q$ent() {

var EventEmitter = require("node:events");

var ev = new EventEmitter();

var sym__x = "X:Real";
var sym__y = "Y:Real";

var z = sym__x;
var a = false;



function onFirst() {
	"!!MAGIC@js_concolic/src-range=155:333";
	"!!MAGIC@js_concolic/idents=z:sym__y:a:setImmediate:onSecond";
	
    if (z === sym__y + 1 && a) {
        throw 'Stickerbrush? Really?';
    }
    if (z < sym__y) {
        z = z + 2;
        setImmediate(onSecond)
    }
}

function onSecond() {
	"!!MAGIC@js_concolic/src-range=335:469";
	"!!MAGIC@js_concolic/idents=z:sym__y:a:setImmediate:onThird:onFirst";
	
    if (z === sym__y && !a) {
        setImmediate(onThird);
    } else {
        setImmediate(onFirst);
    }
}

function onThird() {
	"!!MAGIC@js_concolic/src-range=471:577";
	"!!MAGIC@js_concolic/idents=z:a:setImmediate:onFirst";
	
    z = z - 1;
    if (z != 2) {
        a = true;
    }
    setImmediate(onFirst);
}

if (sym__x < sym__y) {
    setImmediate(onFirst);
}


}
// end entry-point (has-script)

// bgn main-rescue (actual-entry-point)
try {
	_Q$ent();
} catch (e) {
	_Q$xnH(e);
}
// end main-rescue (actual-entry-point)

