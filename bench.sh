#!/bin/sh
iterations=$1
url=$2
echo "Running $iterations iterations for curl $url"
totaltime=0.0
for run in $(seq 1 $iterations)
do
 time=$(curl -X POST $url \
    -s -o /dev/null -w "%{time_total}")
 totaltime=$(echo "$totaltime" + "$time" | bc)
done
avgtimeMs=$(echo "scale=4; 1000*$totaltime/$iterations" | bc)
echo "Averaged $avgtimeMs ms in $iterations iterations"
echo "Total Time $totaltime"

