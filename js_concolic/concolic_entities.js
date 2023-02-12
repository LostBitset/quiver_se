// A representation of values coupled with a symbolic representation

// do not remove the following comment
// JALANGI DO NOT INSTRUMENT

const { symbolicConstant } = require("./to_symbolic");

const { exampleSortInhabitants } = require("./example_sort_inhabitants");

// Symbolic values should be stored as
// [expr_string, sort_string] pairs. 
class ConcolicValue {
    // ccr = concrete, sym = symbolic
    // Symbolic can be null, indicating that only
    // the concrete value is available. 
    // This is *not* concretization, that's when
    // the concrete value used as the symbolic value. 
    constructor(ccr, sym) {
        this.ccr = ccr;
        this.sym = sym;
    }

    static fromConcrete(ccr) {
        return new ConcolicValue(
            ccr,
            symbolicConstant(ccr),
        );
    }

    static fromFreeFun([name, sort]) {
        return new ConcolicValue(
            exampleSortInhabitants[sort],
            name,
        );
    }
}

class ConcolicFunction {
    // ccrOp = concrete operation
    // symOp = symbolic operation
    constructor(ccrOp, symOp) {
        this.ccrOp = ccrOp;
        this.symOp = symOp;
    }
}

// Takes either a normal function or a
// ConcolicFunction, concretizing in the
// first case and applying to both components
// in the second. 
function apConcolic(fn, ...args) {
    if (fn instanceof ConcolicFunction) {
        let args_concolic = args.map(x => {
            if (x instanceof ConcolicValue) {
                return x;
            } else {
                return ConcolicValue.fromConcrete(x);
            }
        });
        if (args.some(x => x.sym === null)) {
            let args_concrete = args.map(x => {
                if (x instanceof ConcolicValue) {
                    return x.ccr;
                } else {
                    return x;
                }
            });
            return new ConcolicValue.fromConcrete(
                (fn)(...args_concrete)
            );
        } else {
            let ccr_ret = fn.ccrOp(...(
                args_concolic.map(x => x.ccr)
            ));
            let sym_ret = fn.symOp(...(
                args_concolic.map(x => x.sym)
            ));
            if (sym_ret === undefined) {
                return ConcolicValue.fromConcrete(ccr_ret);
            } else {
                return new ConcolicValue(ccr_ret, sym_ret);
            }
        }
    } else {
        let args_concrete = args.map(x => {
            if (x instanceof ConcolicValue) {
                return x.ccr;
            } else {
                return x;
            }
        });
        return new ConcolicValue.fromConcrete(
            (fn)(...args_concrete)
        );
    }
}

module.exports = {
    ConcolicValue, ConcolicFunction, apConcolic,
};
