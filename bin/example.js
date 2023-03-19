var sym__x = "X:Real";
var sym__y = "Y:Real";

var z = sym__x;
var a = false;

function onFirst() {
    if (z === sym__y + 1 && a) {
        throw 'Stickerbrush? Really?';
    }
    if (z < sym__y) {
        z = z + 2;
        setImmediate(onSecond)
    }
}

function onSecond() {
    if (z === sym__y && !a) {
        setImmediate(onThird);
    } else {
        setImmediate(onFirst);
    }
}

function onThird() {
    z = z - 1;
    if (z != 2) {
        a = true;
    }
    setImmediate(onFirst);
}

if (sym__x < sym__y) {
    setImmediate(onFirst);
}
