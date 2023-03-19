var sym__x = 2;
var sym__y = 4;

var z = sym__x;
var a = false;



function onFirst() {
    if (z === sym__y + 1 && a) {
        throw 'Stickerbrush? Really?';
    }
    if (z < sym__y) {
        z = z + 2;
        console.log("e2")
        setImmediate(onSecond)
    }
}

function onSecond() {
    if (z === sym__y && !a) {
        console.log("e3");
        setImmediate(onThird);
    } else {
        console.log("e4");
        setImmediate(onFirst);
    }
}

function onThird() {
    z = z - 1;
    if (z != 2) {
        a = true;
        console.log("e6");
    } else {
        console.log("e5");
    }
    setImmediate(onFirst);
}

if (sym__x < sym__y) {
    console.log("e1");
    setImmediate(onFirst);
}
