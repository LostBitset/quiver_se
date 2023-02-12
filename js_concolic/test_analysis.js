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

	var branches = {};

	// @extern(jalangi2).analysis_object
	lkk.analysis = {
		
		conditional: function (iid, result) {
			let id = lkk.getGlobalIID(iid);
			if (!branches.hasOwnProperty(id)) {
				branches[id] = { t: 0, f: 0 };
			}
			if (result) {
				branches[id].t++;
			} else {
				branches[id].f++;
			}
		},

		endExecution: function () {
			for (const [id, info] of Object.entries(branches)) {
				let loc = lkk.iidToLocation(id);
				conlog(`At ${loc}, the condition was true ${info.t} time(s) and false ${info.f} time(s).`);
			}
		},

	};

}(J$));

