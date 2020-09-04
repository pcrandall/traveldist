#! /bin/bash

URL=$1
NAV=$2

echo ${NAV}

curl -s http://${URL}/srm1TravelDistanceList.html | \
     sort -r | grep -m 1 -i total | \
     awk -F : '{print $5}' | \
     awk '{print "  X-Axis: " $1 "    Y-Axis: " $2 }'; \
     echo ""
