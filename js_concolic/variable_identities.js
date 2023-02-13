// A class to represent unique identifiers for variable-associated
// values

// do not remove the following comment
// JALANGI DO NOT INSTRUMENT

class VarIdent {
    constructor(name, scopeType) {
        this.name = name;
        this.scopeType = scopeType;
    }

    toString() {
        return `**${this.scopeType}/jsvar_${this.name}`;
    }
}

module.exports = {
    VarIdent,
};
