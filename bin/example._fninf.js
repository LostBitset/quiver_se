
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

ev.on("first", function() {
	"!!MAGIC@js_concolic/src-range=168:332";
	"!!MAGIC@js_concolic/idents=z:sym_y:a:sym__y:ev:emit"
	
    if (z == sym_y + 1 && a) {
        throw 'Stickerbrush? Really?';
    }
    if (z < sym__y) {
        z = z + 2;
        ev.emit("second");
    }
});

ev.on("second", function() {
	"!!MAGIC@js_concolic/src-range=352:467";
	"!!MAGIC@js_concolic/idents=z:sym__y:a:ev:emit"
	
    if (z === sym__y && !a) {
        ev.emit("third");
    } else {
        ev.emit("first");
    }
});

ev.on("third", function() {
	"!!MAGIC@js_concolic/src-range=486:579";
	"!!MAGIC@js_concolic/idents=z:a:ev:emit"
	
    z = z - 1;
    if (z != 2) {
        a = true;
    }
    ev.emit("first");
});

if (sym__x < sym__y) {
    ev.emit("first");
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

