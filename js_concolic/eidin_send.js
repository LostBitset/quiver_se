// Send (PathCondition) messages over the EIDIN protocol

// do not remove the following comment
// JALANGI DO NOT INSTRUMENT

const eidin = require("./EIDIN/proto_js/eidin_pbjs")

function sendEIDINPathCondition(cgiid_map, free_funs, pc) {
    let msg = eidin.PathCondition.fromObject({
        freeFuns: free_funs.map(([fun_name, sort]) => {
            return eidin.SMTFreeFun.fromObject({
                name: fun_name,
                arg_sorts: [],
                ret_sort: sort,
            });
        }),
        pc: "TODO",
    });
}

module.exports = {
    sendEIDINPathCondition,
};
