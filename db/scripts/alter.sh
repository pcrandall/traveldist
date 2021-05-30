#!/bin/bash

sqlite3 ../traveldistances.db <<EOF
ALTER TABLE CHANGE
RENAME COLUMN uuid TO "uuid";
EOF
