// Send (PathCondition) messages over the EIDIN protocol

// do not remove the following comment
// JALANGI DO NOT INSTRUMENT

const eidin = require("./EIDIN/proto_js/eidin_pbjs");

const { CallbackStreamSeperator } = require("./callback_stream_sep");

function sendEIDINPathCondition(cgiid_map, free_funs, pc) {
    let spc = [];
    let curr_cb_id = null;
    let curr_segment = [];
    for (const elem of pc) {
        if (elem instanceof CallbackStreamSeperator) {
            if (curr_cb_id !== null) {
                let new_cb_id = makeCallbackId(cgiid_map[elem.cgiid]);
                spc.push(eidin.PathConditionSegment.fromObject({
                    this_callback_id: curr_cb_id,
                    next_callback_id: new_cb_id,
                    partial_pc: curr_segment,
                }));
                curr_cb_id = new_cb_id;
                curr_segment = [];
            }
        } else if (elem instanceof Array) {
            curr_segment.push(eidin.SMTConstraint.fromObject({
                constraint: elem[0],
                assertion_value: elem[1],
            }));
        } else if (typeof elem === "string") {
            curr_segment.push(eidin.SMTConstraint.fromObject({
                constraint: `@__RAW__${elem}`,
                assertion_value: true,
            }));
        }
    }
    spc.push(eidin.PathConditionSegment.fromObject({
        this_callback_id: curr_cb_id,
        partial_pc: curr_segment,
    }));
    let msg = eidin.PathCondition.fromObject({
        freeFuns: free_funs.map(([fun_name, sort]) => {
            return eidin.SMTFreeFun.fromObject({
                name: fun_name,
                arg_sorts: [],
                ret_sort: sort,
            });
        }),
        pc: spc,
    });
}

module.exports = {
    sendEIDINPathCondition,
};
