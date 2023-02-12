// A Jalangi2 Analysis
// @!REQUIRES_CONTEXT jalangi2

// do not remove the following comment
// JALANGI DO NOT INSTRUMENT

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

	// @extern(jalangi2).analysis_object
	lkk.analysis = {

        conditional: function (_iid, result) {
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

        literal: function (_iid, val) {
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

        invokeFunPre: function (_iid, f, base, args) {
            return { f, base, args, skip: true };
        },

        invokeFun: function (_iid, f, _base, args, result) {
            if (f.name === "C$symbol") {
                let [name, sort] = args;
                let free_fun = [name, sort];
                free_funs.push(free_fun);
                return ConcolicValue.fromFreeFun(free_fun);
            } else {
                return {
                    result: result,
                };
            }
        },

        endExecution: function () {
            console.log(JSON.stringify({
                free_funs,
                pc,
            }));
        },

    };

}(J$));
