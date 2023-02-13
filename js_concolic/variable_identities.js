// A class to represent unique identifiers for variable-associated
// values

// do not remove the following comment
// JALANGI DO NOT INSTRUMENT

class VarIdent {
    static next_id = 0;

    constructor() {
        this.id = VarIdent.next_id;
        VarIdent.next_id++;
    }

    toString() {
        return `IDENT${this.id}`;
    }
}

module.exports = {
    VarIdent,
};
