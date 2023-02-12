#!/bin/bash

JALANGI2_PATH="$(pwd)/node_modules/jalangi2/src/js/commands/jalangi.js"
JALANGI2_DEFAULT_ARGS="--inlineIID --inlineSource"

node $JALANGI2_PATH $JALANGI2_DEFAULT_ARGS $@

