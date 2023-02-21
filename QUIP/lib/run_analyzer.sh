#!/bin/bash

cd ../../js_concolic/analyze_bin
go run . $1 $2
echo "[QUIP:run_analyzer.sh] ANALYZER PROCESS TERMINATED"
rm ../.eidin-run/Analyze/$2*EIDIN-SIGNAL-STOP
