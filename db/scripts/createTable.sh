#!/bin/bash

sqlite3 ./traveldistances.db <<EOF
CREATE TABLE "CHANGE" (
	"id"	INTEGER,
	"shuttle"	TEXT NOT NULL,
	"distance"	INTEGER NOT NULL,
	"timestamp"	TEXT,
	"notes"	TEXT,
	PRIMARY KEY("id" AUTOINCREMENT)
);
EOF


sqlite3 ./traveldistances.db <<EOF
CREATE TABLE "DISTANCES" (
	"id"	INTEGER NOT NULL,
	"shuttle"	TEXT NOT NULL,
	"distance"	INTEGER NOT NULL,
	"timestamp"	TEXT NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT)
);
EOF
