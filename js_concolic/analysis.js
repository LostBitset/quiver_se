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

        conditional: function (_iid, result_uncoerced) {
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
        },

        write: function (_iid, name, val, lhs) {
            let result = val;
            logs.push(name);
            logs.push("write --^ --v");
            logs.push(val);
            if (name.startsWith("sym__") && lhs === undefined) {
                let [fun, sort] = val.split(":");
                free_funs.push([fun, sort]);
                result = ConcolicValue.fromFreeFun([fun, sort]);
            }
            return {
                result,
            };
        },

        literal: function (_iid, val) {
            if (typeof val === "function") {
                return {
                    result: val,
                };
            }
            let result = ConcolicValue.fromConcrete(val);
            logs.push("literal --v");
            logs.push(result);
            return {
                result,
            };
        },

        binaryPre: function (_iid, op, left, right) {
            return { op, left, right, skip: true };
        },

        binary: function (_iid, op, left, right) {
            let result = apConcolic(ctBinary[op], left, right);
            logs.push(left);
            logs.push(right);
            logs.push("binary --^ --^ --v");
            logs.push(result);
            return {
                result,
            };
        },

        unaryPre: function (_iid, op, left) {
            return { op, left, skip: true };
        },

        unary: function (_iid, op, left) {
            let result = apConcolic(ctUnary[op], left);
            return {
                result,
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
        },

    };

}(J$));
