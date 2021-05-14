#!/bin/sh
sqlite3 ~/.dnote/dnote.db <<EOF
select * from books;
select * from notes;
EOF
