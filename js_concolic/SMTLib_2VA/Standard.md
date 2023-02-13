# A Format for SMT Queries involving Scoped Internal Bindings

## 0. Abstract

This document defines the `SMTLib_2VA` format, which is a superset of `SMTLibv2` that can be transpiled to the latter. The `SMTLib_2VA` format adds support for variable-like definitions, that can be defined, redefined, and are either lexically or globally scoped. 

## 1. New S-expressions

The following new s-expressions are added to the `SMTLib_2VA` language:

| Syntax                          | Type       | Description                    |
|---------------------------------|------------|--------------------------------|
| `(*/write-var-gl/* name value)` | Statement  | Write to a variable globally.  |
| `(*/write-var-le/* name value)` | Statement  | Write to a variable lexically. |
| `(*/read-var/*)`                | Expression | Read from a variable.          |
| `(*/enter-scope/*)`             | Statement  | Enter a new lexical scope.     |
| `(*/leave-scope/*)`             | Statement  | Leave a lexical scope.         |

