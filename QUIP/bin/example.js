
// This code has been instrumented by function_info/instrument.js

// bgn decl-prefix (static)
function _Q$xnH(e) { throw e; }
// end decl-prefix (static)

// bgn entry-point (has-script)
function _Q$ent() {

var sym__x = "X:Real";
var sym__y = "Y:Real";

var z = sym__x;

function onFirst() {
	"!!MAGIC@js_concolic/src-range=64:153";
	"!!MAGIC@js_concolic/idents=z:sym__y:onSecond"
	
    if (z < sym__y) {
        z = z + 1;
        onSecond();
    }
}

function onSecond() {
	"!!MAGIC@js_concolic/src-range=155:238";
	"!!MAGIC@js_concolic/idents=z:onFirst"
	
    if (z == 3) {
        throw 'oof';
    }
    onFirst();
}

onFirst();


}
// end entry-point (has-script)

// bgn main-rescue (actual-entry-point)
try {
	_Q$ent();
} catch (e) {
	_Q$xnH(e);
}
// end main-rescue (actual-entry-point)

