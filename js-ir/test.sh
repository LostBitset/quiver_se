#!/bin/bash

echo "Running scala tests..."
sbt test || exit 1
echo "Running scala tests...OK"

echo "Running main test..."
sbt "run test_input.json" | grep -Ff test_output_expected.json || exit 1
echo "Running main test...OK"

echo "[js-ir/test.sh] ALL TESTS PASSED"
