#!/bin/bash
for conn in 100 200 300 400 500 600 700 800 900 1000
	do
		./wrk -t10 -c"$conn" -d30s -R2000 -L -s ./scripts/hotel-reservation/mixed-workload_type_1.lua http://node1:80
	done