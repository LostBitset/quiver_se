#!/bin/bash

echo "[js_concolic-test] Testing Jalangi2 (not the actual DSE engine)..."

EXAMPLE_JALANGI2_OUTPUT=$(mktemp)

./jalangi2_analyse.sh --analysis test_analysis.js test_analysis_prgm.js 2>&1 >$EXAMPLE_JALANGI2_OUTPUT

if diff test_analysis_output.txt $EXAMPLE_JALANGI2_OUTPUT; then
	echo "[js_concolic-test] [ok] Jalangi2 appears to work properly (example analysis passed)."
else
	echo "[js_concolic-test] [!!] Jalangi2 example analysis FAILED!"
	exit 1
fi

echo "[js_concolic-test] Testing (of) Jalangi2 completed."

echo "[js_concolic-test] WARNING!!! ACTUAL DSE ENGINE HAS NOT YET BEEN IMPLEMENTED!!!"

