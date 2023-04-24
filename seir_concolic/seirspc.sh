#!/bin/bash

ORIG_PWD="$(pwd)"
cd ../js-ir
sbt "run $ORIG_PWD/$1" | tail -n2 | head -n1
