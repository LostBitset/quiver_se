// A simple test program to make sure concolic execution is working

// bgn SHOULD BE ADDED
function _Q$rXn(e) { throw e; }
function _Q$ent() {
// end SHOULD BE ADDED

var f;

function pre() {
    f = false;
}

function main(magic_number) {
    console.log("start");

    var sym__X = "X:Real";

    var yo = sym__X < magic_number;

    if (yo === f) {
        throw 'Crash? ... Yeah, burn? ... Make a wish.';
    }

}

pre();
main(42);

// bgn SHOULD BE ADDED
}
try {
    _Q$ent();
} catch (e) {
    if (e instanceof ReferenceError) {
        _Q$rXn(e);
    }
}
// end SHOULD BE ADDED
