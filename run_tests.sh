#!/bin/bash

# Subdirectory: z3_test
echo "[run_tests.sh] Enter subdirectory z3_test..."
cd z3_test
./z3_test.sh || exit 1
cd ..
echo "[run_tests.sh] Leave subdirectory z3_test. (pass)"

# Subdirectory: js-ir
echo "[run_tests.sh] Enter subdirectory js-ir..."
cd js-ir
./test.sh || exit 1
cd ..
echo "[run_tests.sh] Leave subdirectory js-ir. (pass)"

# Subdirectory: SMTLib_2VA
echo "[run_tests.sh] Enter subdirectory SMTLib_2VA..."
cd SMTLib_2VA
./run_tests.sh || exit 1
cd ..
echo "[run_tests.sh] Leave subdirectory SMTLib_2VA. (pass)"

# Subdirectory: lib_synthetic
echo "[run_tests.sh] Enter subdirectory lib_synthetic..."
cd go_project/lib_synthetic
go test -v || exit 1
cd ../..
echo "[run_tests.sh] Leave subdirectory lib_synthetic. (pass)"

# Subdirectory: lib
echo "[run_tests.sh] Enter subdirectory lib..."
cd go_project/lib
go build
for RUNINDEX in {1..5}
do
	echo "subdirectory lib / TEST RUN #$RUNINDEX"
	go test -v || exit 1
done
cd ../..
echo "[run_tests.sh] Leave subdirectory lib. (pass)"

echo "[run_tests.sh] All tests in repo passed. #[PASSED_ALL]"

