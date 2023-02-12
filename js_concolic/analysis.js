// A Jalangi2 Analysis
// @!REQUIRES_CONTEXT jalangi2

// do not remove the following comment
// JALANGI DO NOT INSTRUMENT

const { ConcolicValue, apConcolic } = require("./concolic_entities");

const { cToBool, ctUnary, ctBinary } = require("./concolic_functions");

function conlog(...args) {
	console.log("[jalangi2:analysis]", ...args);
}

// lkk = lo kelpanka ku = playground = sandbox

// @extern(jalangi2).analysis_iife
(function (lkk) {

    // Free funs as an array of [name, sort] pairs
    var free_funs = [];
    
    // Path condition as an array of [expr, bool] pairs
    var pc = [];

    // Logs
    var logs = [];

	// @extern(jalangi2).analysis_object
	lkk.analysis = {

        /*conditional: function (_iid, result_uncoerced) {
            logs.push("conditional");
            let result = apConcolic(cToBool, result_uncoerced);
            let actual = result;
            if (result instanceof ConcolicValue) {
                actual = result.ccr;
                if (result.sym !== null) {
                    pc.push([this.sym, actual]);
                }
            }
            return {
                result: actual,
            };
        },*/

        /*literal: function (_iid, val) {
            logs.push("literal");
            if (typeof val === "function") {
                return {
                    result: val,
                };
            }
            return {
                result: ConcolicValue.fromConcrete(val),
            };
        },

        binaryPre: function (_iid, op, left, right) {
            return { op, left, right, skip: true };
        },

        binary: function (_iid, op, left, right) {
            return {
                result: apConcolic(ctBinary[op], left, right),
            };
        },

        unaryPre: function (_iid, op, left) {
            return { op, left, skip: true };
        },

        unary: function (_iid, op, left) {
            return {
                result: apConcolic(ctUnary[op], left),
            };
        },

        endExecution: function () {
            for (const log of logs) {
                conlog(log);
            }
            conlog("Ended. ");
            console.log(JSON.stringify({
                free_funs,
                pc,
            }));
        },*/

    };

}(J$));
