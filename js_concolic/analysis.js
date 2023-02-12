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

    var pc = [];

	// @extern(jalangi2).analysis_object
	lkk.analysis = {

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

        invokeFunPre: function (iid, f, base, args) {},

    };

}(J$));
