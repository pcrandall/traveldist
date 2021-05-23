#!/bin/bash
input=$1
table=$2

for f in $input
do
    sqlite3 -separator '|' traveldistances.db ".import $f $table"
done
