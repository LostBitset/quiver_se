#!/bin/bash

# Subdirectory: lib
echo "[run_tests.sh] Enter subdirectory lib..."
cd lib
go build
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

# Subdirectory: js_concolic/EIDIN
echo "[run_tests.sh] Enter subdirectory js_concolic/EIDIN..."
echo "[run_tests.sh] Recompiling probuf schema..."
cd js_concolic/EIDIN
./compile_proto.sh || exit 1
cd ../..
echo "[run_tests.sh] Leave subdirectory js_concolic/EIDIN. (pass)"

# Subdirectory: js_concolic/analyze_bin
echo "[run_tests.sh] Enter subdirectory js_concolic/analyze_bin..."
echo "[run_tests.sh] Recompiling go source..."
cd js_concolic/analyze_bin
go build
cd ../..
echo "[run_tests.sh] Leave subdirectory js_concolic/analyze_bin. (pass)"

# Subdirectory: js_concolic/SMTLib_2VA/lib
echo "[run_tests.sh] Enter subdirectory js_concolic/SMTLib_2VA/lib..."
cd js_concolic/SMTLib_2VA/lib
go build
go test -v || exit 1
cd ../../..
echo "[run_tests.sh] Leave subdirectory js_concolic/SMTLib_2VA/lib. (pass)"

# Subdirectory: js_concolic/SMTLib_2VA/bin
echo "[run_tests.sh] Enter subdirectory js_concolic/SMTLib_2VA/bin..."
cd js_concolic/SMTLib_2VA/bin
go build
./test.sh || exit 1
cd ../../..
echo "[run_tests.sh] Leave subdirectory js_concolic/SMTLib_2VA/bin. (pass)"

# Subdirectory: simple_dse_bin
echo "[run_tests.sh] Enter subdirectory simple_dse_bin..."
cd simple_dse_bin
go build
bash ./simple_dse_integration_test.sh || exit 1
cd ..
echo "[run_tests.sh] Leave subdirectory simple_dse_bin. (pass)"

echo "[run_tests.sh] All tests in repo passed. #[PASSED_ALL]"