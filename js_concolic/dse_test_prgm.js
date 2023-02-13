// A simple test program to make sure concolic execution is working

// bgn SHOULD BE ADDED
function __assert_not_defined(e) { throw e; }
try {
// end SHOULD BE ADDED

var f;

function pre() {
    f = false;
}

function main(magic_number) {
    console.log("start");

    var sym__X = "X:Real";

    var yo = sym__X < magic_number;

    if (yo === fbaka) {
        throw 'Crash? ... Yeah, burn? ... Make a wish.';
    }

}

pre();
main(42);

// bgn SHOULD BE ADDED
} catch (e) {
    if (e instanceof ReferenceError) {
        __assert_not_defined(e);
    }
}
// end SHOULD BE ADDED
