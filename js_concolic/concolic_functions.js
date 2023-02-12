// Concolic functions, that can be represented as both
// concrete and symbolic transformations

// do not remove the following comment
// JALANGI DO NOT INSTRUMENT

const { ConcolicFunction } = require("./concolic_entities");

const cToBool = new ConcolicFunction(
    x => x ? true : false,
    x => {
        return [
            {
                "Bool": x[0],
                "Real": `(not (= ${x[0]} 0.0))`,
                "~undefined": "false",
                "~null": "false",
                "~float.nan": "false",
                "~float.posinf": "true",
                "~float.neginf": "true",
            }[x[1]],
            "Bool",
        ];
    }
);

const cToReal = new ConcolicFunction(
    x => x ? true : false,
    x => {
        return {
            "Bool": [`(ite ${x[0]} 1.0 0.0)`, "Real"],
            "Real": x,
            "~undefined": ["~", "~float.nan"],
            "~null": ["0.0", "Real"],
            "~float.nan": x,
            "~float.posinf": x,
            "~float.neginf": x,
        }[x[1]];
    }
);

const ctUnary = {
    "!": new ConcolicFunction(
        x => !x,
        x => {
            return {
                "Bool": [`(not ${x[0]})`, "Bool"],
                "Real": [`(= ${x[0]} 0.0)`, "Bool"],
            }[x[1]];
        },
    ),
    "+": new ConcolicFunction(
        x => +x,
        x => {
            return {
                "Bool": [`(ite ${x[0]} 1.0 0.0)`, "Real"],
                "Real": x,
            }[x[1]];
        }
    ),
    "-": new ConcolicFunction(
        x => -x,
        x => {
            return {
                "Bool": [`(ite ${x[0]} -1.0 0.0)`, "Real"],
                "Real": [`(* -1.0 ${x[0]})`, "Real"],
            }[x[1]];
        }
    ),
    "~": new ConcolicFunction(x => ~x, _ => undefined),
    "typeof": new ConcolicFunction(x => typeof x, _ => undefined),
};

function makeArithmeticBinary(opSMT, ccrOp) {
    return new ConcolicFunction(
        ccrOp,
        (x, y) => {
            let lhs = cToReal.symOp(x);
            let rhs = cToReal.symOp(y);
            if (lhs === undefined || rhs === undefined) return undefined;
            if (lhs === null || rhs === null) return null;
            if (lhs[1] === "~float.nan" || rhs[1] === "~float.nan") {
                return [(opSMT[0] === "!") + "", "Bool"];
            }
            if (ccrOp === "<" || ccrOp === "<=") {
                if (lhs[1] === "~float.posinf") {
                    // Posinf is only ever less than posinf
                    if (lhs[2] === "~float.posinf") {
                        return [(ccrOp === "<=") + "", "Bool"];
                    }
                    return ["false", "Bool"];
                }
                if (lhs[2] === "~float.posinf") {
                    // All numbers are less than posinf except posinf
                    if (lhs[1] === "~float.posinf") {
                        return [(ccrOp === "<=") + "", "Bool"];
                    }
                    return ["true", "Bool"];
                }
            }
            return [`(${opSMT} ${lhs[0]} ${rhs[0]})`, "Bool"];
        },
    );
}

const ctBinary = {
    "==": new ConcolicFunction(
        (x, y) => x == y,
        (x, y) => {
            if (x[1] === y[1] && x[1][0] !== "~" && y[1][0] !== "~") {
                return [`(= ${x[0]} ${y[0]})`, "Bool"];
            }
            return undefined;
        },
    ),
    "===": new ConcolicFunction(
        (x, y) => x === y,
        (x, y) => {
            if (x[1] === y[1] && x[1][0] !== "~" && y[1][0] !== "~") {
                return [`(= ${x[0]} ${y[0]})`, "Bool"];
            }
            return undefined;
        },
    ),
    "!=": new ConcolicFunction(
        (x, y) => x != y,
        (x, y) => {
            if (x[1] === y[1] && x[1][0] !== "~" && y[1][0] !== "~") {
                return [`(not (= ${x[0]} ${y[0]}))`, "Bool"];
            }
            return undefined;
        },
    ),
    "!==": new ConcolicFunction(
        (x, y) => x !== y,
        (x, y) => {
            if (x[1] === y[1] && x[1][0] !== "~" && y[1][0] !== "~") {
                return [`(not (= ${x[0]} ${y[0]}))`, "Bool"];
            }
            return undefined;
        },
    ),
    "&&": new ConcolicFunction(
        (x, y) => x && y,
        (x, y) => {
            if (x[1] === "Bool" && y[1] === "Bool") {
                return [`(and ${x[0]} ${y[0]})`, "Bool"];
            }
            let lhs = cToBool.symOp(x);
            if (lhs === undefined) return undefined;
            if (lhs === null) return null;
            return [`(ite ${lhs[0]} ${y} ${x})`, "Bool"];
        },
    ),
    "||": new ConcolicFunction(
        (x, y) => x || y,
        (x, y) => {
            if (x[1] === "Bool" && y[1] === "Bool") {
                return [`(or ${x[0]} ${y[0]})`, "Bool"];
            }
            let lhs = cToBool.symOp(x);
            if (lhs === undefined) return undefined;
            if (lhs === null) return null;
            return [`(ite ${lhs[0]} ${x} ${y})`, "Bool"];
        },
    ),
    "<": makeArithmeticBinary("<", (x, y) => x < y),
    ">": makeArithmeticBinary(">", (x, y) => x > y),
    "<=": makeArithmeticBinary("<=", (x, y) => x <= y),
    ">=": makeArithmeticBinary(">=", (x, y) => x >= y),
    "+": makeArithmeticBinary("+", (x, y) => x + y),
    "-": makeArithmeticBinary("-", (x, y) => x - y),
    "*": makeArithmeticBinary("*", (x, y) => x * y),
    "/": makeArithmeticBinary("/", (x, y) => x / y),
    "%": new ConcolicFunction((x, y) => x % y, (_x, _y) => undefined),
    ">>": new ConcolicFunction((x, y) => x >> y, (_x, _y) => undefined),
    "<<": new ConcolicFunction((x, y) => x << y, (_x, _y) => undefined),
    ">>>": new ConcolicFunction((x, y) => x >>> y, (_x, _y) => undefined),
    "&": new ConcolicFunction((x, y) => x & y, (_x, _y) => undefined),
    "|": new ConcolicFunction((x, y) => x | y, (_x, _y) => undefined),
    "^": new ConcolicFunction((x, y) => x ^ y, (_x, _y) => undefined),
    "instanceof": new ConcolicFunction((x, y) => x instanceof y, (_x, _y) => undefined),
    "in": new ConcolicFunction((x, y) => x in y, (_x, _y) => undefined),
};

module.exports = {
    cToBool, ctUnary, ctUnary,
};
