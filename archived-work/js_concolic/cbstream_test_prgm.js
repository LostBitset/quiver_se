// A simple test program to make sure concolic execution is emitting callback streams properly

// bgn SHOULD BE ADDED
function _Q$xnH(e) { throw e; }
function _Q$ent() {
// end SHOULD BE ADDED

var fs = require("node:fs");

fs.readFile("something.txt", "utf-8", function (err, _data) {
    "!!MAGIC@js_concolic/src-range=1:2";
    "!!MAGIC@js_concolic/idents=err:fs:readFile";
    if (err !== null) {
        fs.readFile("something2.txt", "utf-8", function (err, _data) {
            "!!MAGIC@js_concolic/src-range=34:35";
            "!!MAGIC@js_concolic/idents=err:_data:console:log";
            if (err !== null) {
                console.log("Both files do not exist.");
            }
        });
    }
});

// bgn SHOULD BE ADDED
}
try {
    _Q$ent();
} catch (e) {
    _Q$xnH(e);
}
// end SHOULD BE ADDED