#!/bin/bash

# Yes, this is incredibly slow because we restart SBT every time
# Yes, there is a way to avoid doing this
# No, I am not going to deal with that until I actually need to

ORIG_PWD="$(pwd)"
cd ../js-ir
sbt "run $ORIG_PWD/$1" | tail -n2 | head -n1
