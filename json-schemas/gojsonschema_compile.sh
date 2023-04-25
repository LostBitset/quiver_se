#!/bin/bash

MODULENAME="genseirschema"
MODULEDIR="gen_seir_schema"

gojsonschema -p $MODULENAME qse_dse_inp.schema.json > $MODULEDIR/generated_schema_inp.go
gojsonschema -p $MODULENAME qse_dse_out.schema.json > $MODULEDIR/generated_schema_out.go

