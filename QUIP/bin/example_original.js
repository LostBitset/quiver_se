const EventEmitter = require("node:events");

var sym__x = "X:Real";
var sym__y = "Y:Real";

var z;

const demo = new EventEmitter();

demo.on("first", function() {
    if (z < sym__y) {
        z = z + 1;
        demo.emit("second");
    }
});

demo.on("second", function() {
    if (z == 3) {
        throw 'oof';
    }
    demo.emit("first");
});

z = sym__x;
