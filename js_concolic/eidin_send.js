// Send (PathCondition) messages over the EIDIN protocol

// do not remove the following comment
// JALANGI DO NOT INSTRUMENT

const { eidin } = require("./EIDIN/proto_js/eidin_pbjs");

const { v4: uuidv4} = require("uuid");

const { CallbackStreamSeperator } = require("./callback_stream_sep");

const { writeFileSync } = require("node:fs");

const crypto = require("node:crypto");

function sendEIDINPathCondition(cgiid_map, cgiid_map_idents, free_funs, pc) {
    let spc = [];
    let curr_cb_id = null;
    let curr_segment = [];
    for (const elem of pc) {
        if (elem instanceof CallbackStreamSeperator) {
            let new_cb_id = makeCallbackId(elem.cgiid, cgiid_map, cgiid_map_idents);
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

function makeCallbackId(cgiid, cgiid_map, cgiid_map_idents) {
    if (cgiid === "__top__") {
        return eidin.CallbackId.fromObject({
            bytesStart: 0,
            bytesEnd: 0,
            usedFreeFuns: usedFreeFunsFromObject(
                cgiid_map_idents[cgiid]
            ),
        });
    }
    let src_range = cgiid_map[cgiid];
    let [str_start, str_end] = src_range.split(":");
    let start = Number.parseInt(str_start);
    let end = Number.parseInt(str_end);
    return eidin.CallbackId.fromObject({
        bytesStart: start,
        bytesEnd: end,
        usedFreeFuns: usedFreeFunsFromObject(
            cgiid_map_idents[cgiid]
        ),
    });
}

function usedFreeFunsFromObject(obj) {
    return Object.entries(obj)
        .map(([fun_name, sort]) => eidin.SMTFreeFun.fromObject({
            name: fun_name,
            argSorts: [],
            retSort: sort,
        }));
}

function md5(s) {
    return crypto.createHash("md5")
        .update(s, "binary")
        .digest("base64");
}

function sendEIDINMessage(msg) {
    let infile;
    if (process.argv.length > 1 && process.argv[process.argv.length - 2] === "json-model") {
        infile = process.argv[process.argv.length - 3];
    } else {
        infile = process.argv[process.argv.length - 1];
    }
    let infile_hash = md5(infile);
    console.log(infile_hash);
    let filename = `m_${infile_hash}_${uuidv4()}.eidin.bin`;
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
