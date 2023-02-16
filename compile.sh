#!/bin/bash

function compile-tracer-proto() {
  make compile-tracer
  go get -u all
}

if (( $# > 0 )) && (( $1 == "tracer"));
then
 compile-tracer-proto
fi