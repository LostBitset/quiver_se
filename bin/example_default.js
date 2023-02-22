var EventEmitter = require("node:events");

var ev = new EventEmitter();

var sym__x = 0;
var sym__y = 0;

var z = sym__x;
var a = false;



function onFirst() {
    if (z === sym__y + 1 && a) {
        throw 'Stickerbrush? Really?';
    }
    if (z < sym__y) {
        z = z + 2;
        ev.emit("second");
    }
}

function onSecond() {
    if (z === sym__y && !a) {
        ev.emit("third");
    } else {
        ev.emit("first");
    }
}

function onThird() {
    z = z - 1;
    if (z != 2) {
        a = true;
    }
    ev.emit("first");
}

ev.on("first", onFirst);
ev.on("second", onSecond);
ev.on("third", onThird);

if (sym__x < sym__y) {
    ev.emit("first");
}
