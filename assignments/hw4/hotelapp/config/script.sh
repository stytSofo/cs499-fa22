#!/bin/bash
for node in 0 1 2 3 4 5 6 7 8 9 10
	do
		scp nginx.conf node"$node":/tmp/
	done