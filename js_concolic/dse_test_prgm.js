// A simple test program to make sure concolic execution is working

var f = true;

function pre() {
    f = false;
}

function main(magic_number) {
    ":::MAGIC@js_concolic/arg-names|||magic_number";

    console.log("start");

    var sym__X = "X:Real";

    var yo = sym__X < magic_number;

    if (yo === f) {
        throw 'Crash? ... Yeah, burn? ... Make a wish.';
    }

}

pre();
main(42);
