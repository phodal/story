#!/usr/bin/env bash

cd languages

antlr -Dlanguage=Go -visitor -listener Feature.g4 -o feature
