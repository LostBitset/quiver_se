#!/bin/bash

FSTDOUT="$(mktemp)"

time go test -v -run [Rr]eal >$FSTDOUT

if [[ `cat $FSTDOUT` =~ "unexpected call to" ]]
then
    echo "@SIGNAL(OKAY)"
fi

cat $FSTDOUT

rm $FSTDOUT
