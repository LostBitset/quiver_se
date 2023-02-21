#!/bin/bash

cd ../../simple_dse_bin
echo "[QUIP:run_simple_dse.sh] Simple DSE about to run in $(pwd)."
echo "[QUIP:run_simple_dse.sh] Arguments: " $@
go run . $@
echo "[QUIP:run_simple_dse.sh] SIMPLE DSE ORCHESTRATION PROCESS TERMINATED"
rm ../.eidin-run/PathCondition/$1*EIDIN-SIGNAL-STOP
