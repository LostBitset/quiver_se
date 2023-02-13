// A simple test program to make sure concolic execution is working

function main(magic_number) {
    ":::MAGIC@js_concolic/arg-names|||magic_number";

    console.log("start");

    var sym__X = "X:Real";

    var yo = sym__X < magic_number;

    if (!yo) {
        throw 'Crash? ... Yeah, burn? ... Make a wish.';
    }

}

main(42);
main(44);
