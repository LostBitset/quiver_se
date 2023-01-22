#!/bin/bash

# Subdirectory: lib
cd lib
go test -v
cd ..

# Subdirectory: z3_test
cd z3_test
./z3_test.sh
cd ..

