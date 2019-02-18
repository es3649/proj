#!/bin/bash

if [ -e runme ]; then
    echo "rm runme"
	rm runme
fi

if ls *.o 1> /dev/null 2>&1; then
    echo "rm *.o"
	rm *.o
fi

if [ -e *.zip ]; then
    echo "mv *.zip tmp/"
	mv *.zip tmp/
fi