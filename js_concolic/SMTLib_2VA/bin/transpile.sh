#!/bin/bash

IN_FILENAME="$1"
OUT_FILENAME="${IN_FILENAME::-7}.TRANSPILED-orig_smt2va.smt2"

echo "[SMTLib_2VA -> SMTLib_v2] Transpiling..."
echo "[SMTLib_2VA -> SMTLib_v2] Input file is at $IN_FILENAME."

go run . $@

echo "[SMTLib_2VA -> SMTLib_v2] Actual transpilation complete, cleaning up..."

awk -i inplace NF $OUT_FILENAME

echo "[SMTLib_2VA -> SMTLib_v2] Transpilation complete."
echo "[SMTLib_2VA -> SMTLib_v2] Output file is at $OUT_FILENAME"
