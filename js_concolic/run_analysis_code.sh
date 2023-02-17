#!/bin/bash

cd "$(git rev-parse --show-toplevel)/js_concolic"

./jalangi2_analyse.sh --analysis analysis.js $@

