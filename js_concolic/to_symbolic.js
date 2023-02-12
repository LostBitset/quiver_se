// Conversions from JS objects to symbolic representations

// do not remove the following comment
// JALANGI DO NOT INSTRUMENT

// Z3 can't model floating-point arithmetic with
// Infinity, -Infinity, or NaN. 
const MODEL_POS_INFTY = "~*posinf*~";
const MODEL_NEG_INFTY = "~*neginf*~"
const MODEL_NAN = "~*nan*~";

function symbolicConstant(x) {
    if (typeof x === "number") {
        if (Number.isFinite(x)) {
            return x + (
                Number.isInteger(x)
                    ? ".0"
                    : ""
            );
        } else {
            if (Number.isNaN(x)) {
                return MODEL_NAN;
            } else if (x === Infinity) {
                return MODEL_POS_INFTY;
            } else if (x === -Infinity) {
                return MODEL_NEG_INFTY;
            } else {
                return null;
            }
        }
    } else if (typeof x === "boolean") {
        return x + "";
    }
}
