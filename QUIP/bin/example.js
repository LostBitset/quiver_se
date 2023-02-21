
// This code has been instrumented by function_info/instrument.js

// bgn decl-prefix (static)
function _Q$xnH(e) { throw e; }
// end decl-prefix (static)

// bgn entry-point (has-script)
function _Q$ent() {


// This code has been instrumented by function_info/instrument.js

// bgn decl-prefix (static)
function _Q$xnH(e) {
	"!!MAGIC@js_concolic/src-range=96:127";
	"!!MAGIC@js_concolic/idents=e"
	 throw e; }
// end decl-prefix (static)

// bgn entry-point (has-script)
function _Q$ent() {
	"!!MAGIC@js_concolic/src-range=189:566";
	"!!MAGIC@js_concolic/idents=EventEmitter:require:sym__x:sym__y:z:demo:on:emit"
	

const EventEmitter = require("node:events");

var sym__x = "X:Real";
var sym__y = "Y:Real";

var z;

const demo = new EventEmitter();

demo.on("first", () => {
    if (z < sym__y) {
        z = z + 1;
        demo.emit("second");
    }
});

demo.on("second", () => {
    if (z == 3) {
        throw 'oof';
    }
    demo.emit("first");
});

z = sym__x;


}
// end entry-point (has-script)

// bgn main-rescue (actual-entry-point)
try {
	_Q$ent();
} catch (e) {
	_Q$xnH(e);
}
// end main-rescue (actual-entry-point)



}
// end entry-point (has-script)

// bgn main-rescue (actual-entry-point)
try {
	_Q$ent();
} catch (e) {
	_Q$xnH(e);
}
// end main-rescue (actual-entry-point)

