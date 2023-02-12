// A simple test program to make sure concolic execution is working

function main() {

    console.log("start");

    var sym__X = "X:Real";

    if (sym__X > 42) {
        throw 'Crash? ... Yeah, burn? ... Make a wish.';
    }

}

main();
