#!/bin/bash

for f in $(ls *.cpp | cut -d'.' -f1); do
    if [ ! -e $f.o ]; then
        echo "Recompiling $f.cpp"
        g++ -std=c++17 -Wall -Werror $f.cpp -g -c -o $f.o || exit 1
    fi
done