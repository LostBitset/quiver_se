const EventEmitter = require("node:events");

var ev = new EventEmitter();

var sym__x = "X:Real";
var sym__y = "Y:Real";

var z = sym__x;
var a = false;

ev.on("first", function() {
    if (z == sym_y + 1 && a) {
        throw 'todo: something clever';
    }
    if (z < sym__y) {
        z = z + 2;
        ev.emit("second");
    }
});

ev.on("second", function() {
    if (z === sym__y && !a) {
        ev.emit("third");
    } else {
        ev.emit("first");
    }
});

ev.on("third", function() {
    z = z - 1;
    if (z != 2) {
        a = true;
    }
    ev.emit("first");
});

if (sym__x < sym__y) {
    ev.emit("first");
}
