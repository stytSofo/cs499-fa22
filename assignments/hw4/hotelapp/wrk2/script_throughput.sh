#!/bin/bash
for throughput in 1000 2000 3000 4000 5000 6000 7000 8000 9000 10000
	do
		./wrk -t10 -c200 -d30s -R"$throughput" -L -s ./scripts/hotel-reservation/mixed-workload_type_1.lua http://node1:80
	done