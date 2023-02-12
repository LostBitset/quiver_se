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
	lkk.analysis = {};

}(J$));

