# A Format for SMT Queries involving Lexically Scoped Internal Bindings

## 0. Abstract

This document defines the `SMTLib_2VA` format, which is a superset of `SMTLibv2` that can be transpiled to the latter. The `SMTLib_2VA` format adds support for variable-like definitions, that can be defined, redefined, and are lexically scoped. 

## 1. New S-expressions

The following new s-expressions are added to the `SMTLib_2VA` language:

| Syntax                                | Type       | Description                                                                 |
|---------------------------------------|------------|-----------------------------------------------------------------------------|
| `(*/enter-scope/*)`                   | Statement  | Enter a new lexical scope.                                                  |
| `(*/leave-scope/*)`                   | Statement  | Leave a lexical scope.                                                      |
| `(*/decl-var/* name)`                 | Statement  | Declares a variable to be defined in the current scope.                     |
| `(*/write-var/* name *{{value}}*)`    | Statement  | Write to an already-declared variable.                                      |
| `(*/read-var/* name)`                 | Expression | Read from a variable.                                                       |
| `(*/is-defined?/* name)`              | Expression | Evaluates to a `Bool` stating whether or not the variable has been defined. |

The somewhat strange syntax `*{{value}}*` is referred to as a "capture", and used to enable fast transpilation. Specifically, once strings are accounted for, the grammar for each of these new constructs becomes regular. Regex engines, based on Thompson NFAs, can then search in linear time. 

## 2. New Symbol Types

The following new symbol types are added to the `SMTLib_2VA` language:

| Syntax             | Type                | Shorthand Name | Description   |
|--------------------|---------------------|----------------|---------------|
| `**variable`       | Variable Name (new) | `name`         | A variable.   |

