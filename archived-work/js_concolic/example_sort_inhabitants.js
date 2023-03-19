// Example inhabitants of SMT sorts used for concretization of symbols

// do not remove the following comment
// JALANGI DO NOT INSTRUMENT

const exampleSortInhabitants = {
    "Real": 0.0,
    "Bool": false,
    "~undefined": undefined,
    "~null": null,
    "~float.nan": NaN,
    "~float.posinf": Infinity,
    "~float.neginf": -Infinity,
};

module.exports = {
    exampleSortInhabitants,
};
