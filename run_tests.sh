#!/bin/bash

# Subdirectory: lib
echo "[run_tests.sh] Enter subdirectory lib..."
cd lib
go test -v || exit 1
cd ..
echo "[run_tests.sh] Leave subdirectory lib. (pass)"

# Subdirectory: z3_test
echo "[run_tests.sh] Enter subdirectory z3_test..."
cd z3_test
./z3_test.sh || exit 1
cd ..
echo "[run_tests.sh] Leave subdirectory z3_test. (pass)"

# Subdirectory: callback_streams
echo "[run_tests.sh] Enter subdirectory callback_streams..."
cd callback_streams
./test.sh || exit 1
cd ..
echo "[run_tests.sh] Leave subdirectory callback_streams. (pass)"

echo "[run_tests.sh] All tests in repo passed. #[PASSED_ALL]"

