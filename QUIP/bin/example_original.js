var sym__x = "X:Real";
var sym__y = "Y:Real";

var z = sym__x;

function onFirst() {
    if (z < sym__y) {
        z = z + 1;
        onSecond();
    }
}

function onSecond() {
    if (z == 3) {
        throw 'oof';
    }
    onFirst();
}

onFirst();
