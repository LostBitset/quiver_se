#!/bin/bash

# MUST BE RUN THROUGH BASH FOR SOME STUPID REASON I DONT UNDERSTAND

TMP_ANALYZER_OUTPUT="$(mktemp)"

echo "[simple_dse_integration_test] Starting Analyzer Process..."
(
    cd ../js_concolic/analyze_bin
    go run . dse_test_prgm_st.js m_EC551E7efjBuHARc4MsPgg==_ | tee $TMP_ANALYZER_OUTPUT
    echo "[simple_dse_integration_test] ANALYZER PROCESS TERMINATED"
    rm ../js_concolic/.eidin-run/Analyze/m_EC551E7efjBuHARc4MsPgg==__EIDIN-SIGNAL-STOP
) &

echo "[simple_dse_integration_test] Starting Orchestration Process..."
(
    go run . m_EC551E7efjBuHARc4MsPgg==_
    echo "[simple_dse_integration_test] ORCHESTRATION PROCESS TERMINATED"
    rm ../js_concolic/.eidin-run/PathCondition/m_EC551E7efjBuHARc4MsPgg==__EIDIN-SIGNAL-STOP
) &

sleep 0.5

echo "[simple_dse_integration_test] Starting the cycle..."
cp ../js_concolic/analyze_bin/empty_Analyze.eidin.bin ../js_concolic/.eidin-run/Analyze/m_EC551E7efjBuHARc4MsPgg==_emptyAnalyze.eidin.bin
echo "[simple_dse_integration_test] Cycle started, waiting 2s."

sleep 2

echo "[simple_dse_integration_test] Terminating EIDIN processes..."
touch ../js_concolic/.eidin-run/Analyze/m_EC551E7efjBuHARc4MsPgg==__EIDIN-SIGNAL-STOP
touch ../js_concolic/.eidin-run/PathCondition/m_EC551E7efjBuHARc4MsPgg==__EIDIN-SIGNAL-STOP

if cat $TMP_ANALYZER_OUTPUT | grep "Crash? ... Yeah, burn? ... Make a wish." >/dev/null; then
    echo "[simple_dse_integration_test] Test passed."
else
    echo "[simple_dse_integration_test] TEST FAILED!"
fi

rm $TMP_ANALYZER_OUTPUT

sleep 0.5

