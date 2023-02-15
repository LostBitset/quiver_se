// Send (PathCondition) messages over the EIDIN protocol

// do not remove the following comment
// JALANGI DO NOT INSTRUMENT

const protobuf = require("protobufjs");

var eidin_pc_possibly_uninit;
var eidin_pc_onload_callbacks = [];

protobuf.load("EIDIN/proto/eidin.proto", (err, root) => {
    if (err) throw err;
    eidin_pc_possibly_uninit = root.lookupType("eidin.PathCondition");
    for (const cb of eidin_pc_onload_callbacks) {
        cb(eidin_pc_possibly_uninit);
    }
    eidin_pc_onload_callbacks = [];
});

function onEIDINLoadPC(cb) {
    if (eidin_pc_possibly_uninit === undefined) {
        eidin_pc_onload_callbacks.push(cb);
    } else {
        cb(eidin_pc_possibly_uninit);
    }
}

function sendEIDINPathCondition(free_funs, pc) {
    // TODO
}
