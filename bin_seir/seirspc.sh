#!/bin/bash

# Yes, this is incredibly slow because we restart SBT every time
# Yes, there is a way to avoid doing this
# No, I am not going to deal with that until I actually need to

ORIG_PWD="$(pwd)"
if [[ $1 ==  /* ]]; then
    TARGET_PATH="$1"
else
    TARGET_PATH="$ORIG_PWD/$1"
fi

cd ../js-ir
sbt "run $TARGET_PATH" | tail -n2 | head -n1
