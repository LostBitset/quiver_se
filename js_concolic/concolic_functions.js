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
};
