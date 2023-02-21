#!/bin/bash

# Arguments:
# 1 - The name of the program
# 2 - The output name that will be changed, usually this is just .js
#     replaced with ._fninf.js

cd ../../function_info

node instrument.js $1 debug

mv $2 $1
