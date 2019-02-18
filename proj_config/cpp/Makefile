SHELL := /bin/bash

all: build

build: objects link
	# evaluate changes
	# rebuild graph if needed
	# compile changed objects

build-raw:
	# redo this with xargs
	g++ --std=c++17 -Wall -Werror *.cpp -o runme

objects:
	# evaluate changes, 
	# rebuild dependency graph if needed
	# remove .o files for changed files and their dependents
	# compile those files
	# link
	./tools/objects.sh

link:
	g++ -std=c++17 *.o -g -o runme

clean:
	./tools/cleanup.sh

test: build
	./tools/test.sh

debug:
	# open the debugger??