#!/bin/bash

for threads in 1 2 4 8 16 32 64 128; 
do
	echo "./client node1 8080 /local/websearch/ISPASS_PAPER_QUERIES_100K 1000 "$threads" onlyHits.jsp 1 1 /tmp/out 1 2> /dev/null"
done
