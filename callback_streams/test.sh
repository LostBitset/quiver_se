#!/bin/bash

node cbstream.js example.js debug

INSTRUMENTED_OUTPUT=$(mktemp)

node example.INSTRUMENTED-cbstream.js 2>/dev/null | grep "^\[CBSTREAM\]" >$INSTRUMENTED_OUTPUT

if diff test_output_failure.txt $INSTRUMENTED_OUTPUT; then
	echo "[cbstream-test] [ok] test intent:output_failure"
else
	echo "[cbstream-test] [!!] test intent:output_failure FAILED!"
	exit 1
fi

touch something.txt something2.txt

node example.INSTRUMENTED-cbstream.js 2>/dev/null | grep "^\[CBSTREAM\]" >$INSTRUMENTED_OUTPUT

if diff test_output_normal.txt $INSTRUMENTED_OUTPUT; then
	echo "[cbstream-test] [ok] test intent:output_normal"
else
	echo "[cbstream-test] [!!] test intent:output_normal FAILED!"
	exit 1
fi

rm something.txt something2.txt

rm $INSTRUMENTED_OUTPUT

