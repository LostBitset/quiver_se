// A program that instruments functions inside Node.js programs with standalone
// string literals containing the following information:
// 1. Where the function occurs in the source code (start and end bytes)
// 2. All of the identifiers that the function accesses
// These are encoded in the following form:
// 1. "!!MAGIC@function_info/src-range=<start>:<end>";
// 2. "!!MAGIC@function_info/idents=<ident0>:<ident1>:<ident2>:<...>";

