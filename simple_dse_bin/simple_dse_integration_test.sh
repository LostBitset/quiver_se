#!/bin/bash

TMP_ANALYZER_OUTPUT="$(mktemp)"
ANALYZER_PID_FILE="$(mktemp)"

echo "[simple_dse_integration_test] Starting Analyzer Process..."
(
    while read line; do
        echo $line >>$TMP_ANALYZER_OUTPUT
    done < <(
        cd ../js_concolic/analyze_bin
        go run . dse_test_prgm_st.js m_EC551E7efjBuHARc4MsPgg==_ 2>&1
    )
    echo "$!" >$ANALYZER_PID_FILE
) &

echo "[simple_dse_integration_test] Starting Orchestration Process..."
(
    go run . m_EC551E7efjBuHARc4MsPgg==_
) &
ORCHESTRATION_PID="$!"

sleep 0.5

echo "[simple_dse_integration_test] Starting the cycle..."
cp ../js_concolic/analyze_bin/empty_Analyze.eidin.bin ../js_concolic/.eidin-run/Analyze/m_EC551E7efjBuHARc4MsPgg==_emptyAnalyze.eidin.bin
echo "[simple_dse_integration_test] Cycle started, waiting 2s."

sleep 2

echo "[simple_dse_integration_test] Cleaning up..."
kill -9 $(cat $ANALYZER_PID_FILE)
kill -9 $ORCHESTRATION_PID
sleep 0.5
echo "[simple_dse_integration_test] Done cleaning up."

if cat $TMP_ANALYZER_OUTPUT | grep "Crash? ... Yeah, burn? ... Make a wish." >/dev/null; then
    echo "[simple_dse_integration_test] Test passed."
else
    echo "[simple_dse_integration_test] TEST FAILED!"
    SHOULD_FAIL=1
fi

rm $TMP_ANALYZER_OUTPUT
