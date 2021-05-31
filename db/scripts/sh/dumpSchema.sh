#!/bin/bash

sqlite3 ../traveldistances.db <<EOF
.schema
EOF
