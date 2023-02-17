// Send (PathCondition) messages over the EIDIN protocol

// do not remove the following comment
// JALANGI DO NOT INSTRUMENT

const { eidin } = require("./EIDIN/proto_js/eidin_pbjs");

const { v4: uuidv4} = require("uuid");

const { CallbackStreamSeperator } = require("./callback_stream_sep");

const { writeFileSync } = require("node:fs");

function sendEIDINPathCondition(cgiid_map, free_funs, pc) {
    let spc = [];
    let curr_cb_id = null;
    let curr_segment = [];
    for (const elem of pc) {
        if (elem instanceof CallbackStreamSeperator) {
            let new_cb_id = makeCallbackId(elem.cgiid, cgiid_map);
            if (curr_cb_id !== null) {
                spc.push(eidin.PathConditionSegment.fromObject({
                    thisCallbackId: curr_cb_id,
                    nextCallbackId: new_cb_id,
                    partialPc: curr_segment,
                }));
            }
            curr_cb_id = new_cb_id;
            curr_segment = [];
        } else if (elem instanceof Array) {
            curr_segment.push(eidin.SMTConstraint.fromObject({
                constraint: elem[0],
                assertionValue: elem[1],
            }));
        } else if (typeof elem === "string") {
            curr_segment.push(eidin.SMTConstraint.fromObject({
                constraint: `@__RAW__${elem}`,
                assertionValue: true,
            }));
        }
    }
    spc.push(eidin.PathConditionSegment.fromObject({
        thisCallbackId: curr_cb_id,
        partialPc: curr_segment,
    }));
    let msg = eidin.PathCondition.fromObject({
        freeFuns: free_funs.map(([fun_name, sort]) => {
            return eidin.SMTFreeFun.fromObject({
                name: fun_name,
                argSorts: [],
                retSort: sort,
            });
        }),
        segmentedPc: spc,
    });
    let msg_buffer = eidin.PathCondition.encode(msg).finish();
    sendEIDINMessage(msg_buffer);
}

function makeCallbackId(cgiid, cgiid_map) {
    if (cgiid === "__top__") {
        return eidin.CallbackId.fromObject({
            bytesStart: 0,
            bytesEnd: 0,
        });
    }
    let src_range = cgiid_map[cgiid];
    let [str_start, str_end] = src_range.split(":");
    let start = Number.parseInt(str_start);
    let end = Number.parseInt(str_end);
    return eidin.CallbackId.fromObject({
        bytesStart: start,
        bytesEnd: end,
    });
}

function sendEIDINMessage(msg) {
    let filename = `m_${uuidv4()}.eidin.bin`;
    let filepath = `.eidin-run/PathCondition/${filename}`;
    console.log(msg);
    writeFileSync(filepath, msg, { encoding: 'binary', flags: 'wb' }, err => {
        if (err) throw err;
        console.log("[js_concolic@node] [jalangi2:analysis:callback] Results written. ");
    });
}

module.exports = {
    sendEIDINPathCondition,
};
