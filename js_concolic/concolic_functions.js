// Concolic functions, that can be represented as both
// concrete and symbolic transformations

// do not remove the following comment
// JALANGI DO NOT INSTRUMENT

const cToBool = new ConcolicFunction(
    x => x ? true : false,
    x => {
        return {
            "Bool": [x[0], "Bool"],
            "Real": [`(not (= ${x[0]} 0.0))`, "Bool"],
            "~undefined": ["false", "Bool"],
            "~null": ["false", "Bool"],
            "~float.nan": ["false", "Bool"],
            "~float.posinf": ["true", "Bool"],
            "~float.neginf": ["true", "Bool"],
        }[x[1]];
    }
)

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
            if (
                true
                && (x[1] === "Real" || x[1] === "Bool")
                && (y[1] === "Real" || y[1] === "Bool")
            ) {
                let lhs = (x[1] === "Bool")
                    ? `(ite ${x[0]} 1.0 0.0)`
                    : x[0];
                let rhs = (y[1] === "Bool")
                    ? `(ite ${y[0]} 1.0 0.0)`
                    : y[0];
                return [`(${opSMT} ${lhs} ${rhs})`, "Bool"]
            }
            return undefined;
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
            return [`(ite ${lhs} ${y} ${x})`, "Bool"];
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
            return [`(ite ${lhs} ${x} ${y})`, "Bool"];
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
    "%": new ConcolicFunction((x, y) => x % y, (_, _) => undefined),
    ">>": new ConcolicFunction((x, y) => x >> y, (_, _) => undefined),
    "<<": new ConcolicFunction((x, y) => x << y, (_, _) => undefined),
    ">>>": new ConcolicFunction((x, y) => x >>> y, (_, _) => undefined),
    "&": new ConcolicFunction((x, y) => x & y, (_, _) => undefined),
    "|": new ConcolicFunction((x, y) => x | y, (_, _) => undefined),
    "^": new ConcolicFunction((x, y) => x ^ y, (_, _) => undefined),
    "instanceof": new ConcolicFunction((x, y) => x instanceof y, (_, _) => undefined),
    "in": new ConcolicFunction((x, y) => x in y, (_, _) => undefined),
};
