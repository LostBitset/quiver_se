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
    // or just strings for special stuff
    var pc = [];

    // Logs
    var logs = [];

    // Whether or not we have identified the current callback
    let cbstreamHasCallback = false;

    // Claim a particular, currently executing function as the current callback
    function cbstreamClaimCallback(id) {
        let had = cbstreamHasCallback;
        cbstreamHasCallback = true;
        if (!had) cbstreamOnCallback(id);
        process.nextTick(() => {
            cbstreamHasCallback = false;
        });
    }

    // Handle the discovery of a new callback from the cbstream process
    function cbstreamOnCallback(id) {
        logs.push(`[cbstream::CALLBACK_TRANSITION] Transitioned to ${id}.`);
        let obj = {
            cgiid: id, // Creation Global Instruction IDentifier
        };
        pc.push(`;; @@CALLBACK-TRANSITION[${JSON.stringify(obj)}]`);
    }

    // Claim the entry point / top callback
    cbstreamClaimCallback("__top__");

	// @extern(jalangi2).analysis_object
	lkk.analysis = {

        conditional: function (_iid, result_uncoerced) {
            let result = apConcolic(cToBool, result_uncoerced);
            let actual = result;
            if (result instanceof ConcolicValue) {
                actual = result.ccr;
                if (result.sym[0] !== actual.toString()) {
                    // Don't bother adding blatantly obvious tautologies
                    // to the path condition
                    if (result.sym !== null) {
                        pc.push([result.sym[0], actual]);
                    }
                }
            }
            return {
                result: actual,
            };
        },

        read: function (_iid, name, val) {
            let result = val;
            if (val instanceof ConcolicValue && !name.startsWith("sym__")) {
                result = new ConcolicValue(
                    val.ccr,
                    [`(*/read-var/* **jsvar_${name})`, val.sym[1]],
                );
            }
            return {
                result,
            };
        },

        write: function (_iid, name, val, lhs) {
            let result = val;
            if (name.startsWith("_Q$") || name.startsWith("dseunwrapped__")) {
                if (result instanceof ConcolicValue) {
                    result = result.ccr;
                }
            }
            if (name.startsWith("sym__") && lhs === undefined) {
                let [fun, sort] = val.ccr.split(":");
                free_funs.push([fun, sort]);
                result = ConcolicValue.fromFreeFun([fun, sort]);
            } else {
                if (val instanceof ConcolicValue) {
                    pc.push(`(*/write-var/* **jsvar_${name} ${val.sym[0]})`);
                }
            }
            return {
                result,
            };
        },

        functionEnter: function (_iid, _f, _dis, _args) {
            pc.push("(*/enter-scope/*)")
        },

        functionExit: function (_iid) {
            pc.push("(*/leave-scope/*)");
        },

        declare: function (_iid, name, val, isArgument) {
            if (
                (val === undefined || val instanceof ConcolicValue)
                && !name.startsWith("sym__")
            ) {
                if (val !== undefined && val.ccr === undefined) {
                    return {
                        result: val,
                    };
                }
                pc.push(`(*/decl-var/* **jsvar_${name})`);
                if (isArgument) {
                    if (val !== undefined) {
                        pc.push(`(*/write-var/* **jsvar_${name} ${val.sym[0]})`);
                    }
                }
            }
            return {
                result: val,
            };
        },

        literal: function (iid, val) {
            // Special cases
            if (typeof val === "function") {
                logs.push(`Marked fn ${val.name?`"${val.name}"`:"<anon>"} as instrumented.`);
                val["C$_INSTRUMENTED"] = true;
                return {
                    result: function (...args) {
                        cbstreamClaimCallback(lkk.getGlobalIID(iid));
                        return (val)(...args);
                    },
                };
            }
            // Make (conc|symb)olic otherwise
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
            if (f.hasOwnProperty("name") && f.name === "_Q$xnH") {
                let exn = args[0];
                if (exn instanceof ReferenceError) {
                    let varName = exn.message.split(" ")[0];
                    pc.push([`(*/is-defined?/* ${varName})`, false]);
                    return { f, base, args, skip: false };
                }
            }
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
                logs.push("Concretized target of field write.");
                baseCcr = baseCcr.ccr;
            }
            if (offsetCcr instanceof ConcolicValue) {
                logs.push("Concretized property of field write.");
                offsetCcr = offsetCcr.ccr;
            }
            if (valCcr instanceof ConcolicValue) {
                logs.push("Concretized value of field write.");
            }
            return {
                base: baseCcr,
                offset: offsetCcr,
                val: valCcr,
                skip: false,
            };
        },

        _throw: function (_iid, val) {
            if (val instanceof ConcolicValue) {
                logs.push("Concretized value thrown.");
                return {
                    result: val.ccr,
                };
            } else {
                return {
                    result: val,
                };
            }
        },

        _with: function (_iid, val) {
            if (val instanceof ConcolicValue) {
                logs.push("Concretized target of with statement.");
                return {
                    result: val.ccr,
                };
            } else {
                return {
                    result: val,
                };
            }
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
