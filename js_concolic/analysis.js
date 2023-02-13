// A Jalangi2 Analysis
// @!REQUIRES_CONTEXT jalangi2

// do not remove the following comment
// JALANGI DO NOT INSTRUMENT

const { ConcolicValue, apConcolic } = require("./concolic_entities");

const { cToBool, ctUnary, ctBinary } = require("./concolic_functions");

function conlog(...args) {
	console.log("[js_concolic@node] [jalangi2:analysis]", ...args);
}

// lkk = lo kelpanka ku = playground = sandbox

// @extern(jalangi2).analysis_iife
(function (lkk) {

    // Free funs as an array of [name, sort] pairs
    var free_funs = [];
    
    // Path condition as an array of [expr, bool] pairs
    var pc = [];

    // Functions that have been instrumented as {[fn]: true}
    var instrumented_fns = {};

    // Logs
    var logs = [];

	// @extern(jalangi2).analysis_object
	lkk.analysis = {

        conditional: function (_iid, result_uncoerced) {
            let result = apConcolic(cToBool, result_uncoerced);
            let actual = result;
            if (result instanceof ConcolicValue) {
                actual = result.ccr;
                if (result.sym !== null) {
                    pc.push([result.sym[0], actual]);
                }
            }
            return {
                result: actual,
            };
        },

        write: function (_iid, name, val, lhs) {
            let result = val;
            if (name.startsWith("sym__") && lhs === undefined) {
                let [fun, sort] = val.ccr.split(":");
                free_funs.push([fun, sort]);
                result = ConcolicValue.fromFreeFun([fun, sort]);
            }
            return {
                result,
            };
        },

        literal: function (_iid, val) {
            if (typeof val === "function") {
                logs.push(`Marked fn "${val.name}" as instrumented.`);
                val["C$_INSTRUMENTED"] = true;
                return {
                    result: val,
                };
            }
            let result = ConcolicValue.fromConcrete(val);
            return {
                result,
            };
        },

        binaryPre: function (_iid, op, left, right) {
            return { op, left, right, skip: true };
        },

        binary: function (_iid, op, left, right) {
            let result = apConcolic(ctBinary[op], left, right);
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

        // begin CONCRETIZED

        invokeFunPre: function (_iid, f, base, args) {
            if (!f.hasOwnProperty("C$_INSTRUMENTED")) {
                // Concretize calls that have not been instrumented
                logs.push(`Concretized call to external fn "${base}.${f.name}".`);
                let baseCcr = (base instanceof ConcolicValue) ? base.ccr : base;
                for (let i = 0; i < args.length; i++) {
                    if (args[i] instanceof ConcolicValue) {
                        args[i] = args[i].ccr;
                    }
                }
                return {
                    f,
                    base: baseCcr,
                    args,
                    skip: false,
                };
            }
            return { f, base, args, skip: false };
        },

        forInObject: function (_iid, val) {
            if (val instanceof ConcolicValue) {
                logs.push("Concretized target of for-in loop.")
                return {
                    result: val.ccr,
                };
            } else {
                return {
                    result: val,
                };
            }
        },

        getFieldPre: function (_iid, base, offset) {
            let baseCcr = base;
            let offsetCcr = offset;
            if (baseCcr instanceof ConcolicValue) {
                logs.push("Concretized target of field access.")
                baseCcr = baseCcr.ccr;
            }
            if (offsetCcr instanceof ConcolicValue) {
                logs.push("Concretized property of field access.")
                offsetCcr = offsetCcr.ccr;
            }
            return {
                base: baseCcr,
                offset: offsetCcr,
                skip: false,
            };
        },

        putFieldPre: function(_iid, base, offset, val) {
            let baseCcr = base;
            let offsetCcr = offset;
            let valCcr = val;
            if (baseCcr instanceof ConcolicValue) {
                logs.push("Concretized target of field write.")
                baseCcr = baseCcr.ccr;
            }
            if (offsetCcr instanceof ConcolicValue) {
                logs.push("Concretized property of field write.")
                offsetCcr = offsetCcr.ccr;
            }
            if (valCcr instanceof ConcolicValue) {
                logs.push("Concretized value of field write.")
            }
            return {
                base: baseCcr,
                offset: offsetCcr,
                val: valCcr,
                skip: false,
            };
        },

        // end CONCRETIZED

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
