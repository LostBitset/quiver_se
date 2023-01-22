#!/bin/bash

echo "[z3_test] Attempting to use z3 binary to generate MUS..."

example_out=`z3 -smt2 example.smt2 | paste -s -d','`
desired_out_1='unsat,(ga_1 ga_0)'
desired_out_2='unsat,(ga_0 ga_1)'
if [[ "$example_out" == "$desired_out_1" || "$example_out" == "$desired_out_2" ]]; then
	echo "PASS"
else
	echo "/!\\ FAIL /!\\"
	echo "Expected (1):  $desired_out_1"
	echo "Expected (2):  $desired_out_2"
	echo "Actual:        $example_out"
	echo "(Both expected values would have been acceptable)"
	exit 1
fi

