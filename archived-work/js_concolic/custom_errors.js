// A file for custom error types

// do not remove the following comment
// JALANGI DO NOT INSTRUMENT

class TerminatingAfterSingleCallback extends Error {
    constructor() {
        super("NOT A FAILURE - Terminating after single callback.");
    }
}

module.exports = {
    TerminatingAfterSingleCallback,
}
