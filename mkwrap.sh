#!/bin/sh

if [ $GOARCH = amd64 ]; then
    GOCHAR=6
elif [ $GOARCH = 386 ]; then
    GOCHAR=8
else 
    GOCHAR=5
fi

export GOCHAR

make $*
