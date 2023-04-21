#!/bin/bash

echo "[SMTLib_2VA/run_tests.sh] Enter subdirectory lib..."
cd lib
go test -v || exit 1
cd ..
echo "[SMTLib_2VA/run_tests.sh] Leave subdirectory lib. (pass)"

echo "[SMTLib_2VA/run_tests.sh] Enter subdirectory bin..."
cd bin
./test.sh || exit 1
cd ..
echo "[SMTLib_2VA/run_tests.sh] Leave subdirectory bin. (pass)"

echo "[SMTLib_2VA/run_tests.sh] All tests (in subproject SMTLib_2VA) passed. "

