// Conversions from JS objects to symbolic representations

// do not remove the following comment
// JALANGI DO NOT INSTRUMENT

function symbolicConstant(x) {
    if (x === undefined) {
        return ["~", "~undefined"];
    } else if (x === null) {
        return ["~", "~null"];
    } else if (typeof x === "number") {
        if (Number.isFinite(x)) {
            return [
                x + (
                    Number.isInteger(x)
                        ? ".0"
                        : ""
                ),
                "Real"
            ];
        } else {
            if (Number.isNaN(x)) {
                return ["~", "~float.nan"];
            }
            if (x === Infinity) {
                return ["~", "~float.posinf"];
            }
            if (x === -Infinity) {
                return ["~", "~float.neginf"]
            }
        }
    } else if (typeof x === "boolean") {
        return [x + "", "Bool"];
    }
    return null;
}

module.exports = {
    symbolicConstant,
};
