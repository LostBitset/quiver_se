#!/bin/bash

echo "[SMTLib_2VA-transpiler-test] Running test..."

./transpile.sh example.smt2va log-info

if diff target.smt2 example.TRANSPILED-orig_smt2va.smt2; then
    echo "[SMTLib_2VA-transpiler-test] Test passed."
else
    echo "[SMTLib_2VA-transpiler-test] TEST FAILED!!!"
    exit 1
fi
