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

# Subdirectory: js_concolic
echo "[run_tests.sh] Enter subdirectory js_concolic..."
cd js_concolic
./test.sh || exit 1
cd ..
echo "[run_tests.sh] Leave subdirectory js_concolic. (pass)"

# Subdirectory: js_concolic/SMTLib_2VA/lib
echo "[run_tests.sh] Enter subdirectory js_concolic/SMTLib_2VA/lib..."
cd js_concolic/SMTLib_2VA/lib
go test -v || exit 1
cd ../../..
echo "[run_tests.sh] Leave subdirectory js_concolic/SMTLib_2VA/lib. (pass)"

echo "[run_tests.sh] All tests in repo passed. #[PASSED_ALL]"

