#!/bin/bash

rm reprdigraph_pngout.png || true
fdp reprdigraph.dot -Tpng >reprdigraph_pngout.png
