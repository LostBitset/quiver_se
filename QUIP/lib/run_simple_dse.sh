#!/bin/bash

cd ../../simple_dse_bin
go run . $@
echo "[QUIP:run_simple_dse.sh] SIMPLE DSE ORCHESTRATION PROCESS TERMINATED"
rm ../.eidin-run/PathCondition/$1*EIDIN-SIGNAL-STOP
