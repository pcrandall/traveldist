CREATE TABLE IF NOT EXISTS "CHANGE" (
	"id"	INTEGER,
	"shuttle"	TEXT NOT NULL,
	"distance"	INTEGER NOT NULL,
	"timestamp"	TEXT,
	"notes"	TEXT, uuid TEXT,
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE sqlite_sequence(name,seq);
CREATE TABLE IF NOT EXISTS "DISTANCES" (
	"id"	INTEGER,
	"shuttle"	TEXT NOT NULL,
	"distance"	INTEGER NOT NULL,
	"timestamp"	TEXT,
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE VIEW "shoe_travel" AS SELECT 
  DISTANCES.shuttle as t1_shuttle, 
  DISTANCES.distance as t1_distance, 
  max(CHANGE.distance) as t2_distance,
  (DISTANCES.distance-CHANGE.distance)/1000 as shoe_travel,
  max(DISTANCES.timestamp) as t1_timestamp,
  CHANGE.timestamp as t2_timestamp,
  (SELECT
  printf("%03d", julianday(DISTANCES.timestamp)-julianday(CHANGE.timestamp))) as days_installed,
  CHANGE.notes as notes,
  CHANGE.uuid as uuid
FROM
  DISTANCES
INNER JOIN CHANGE on CHANGE.shuttle = DISTANCES.shuttle
GROUP BY t1_shuttle
ORDER BY shoe_travel DESC
/* shoe_travel(t1_shuttle,t1_distance,t2_distance,shoe_travel,t1_timestamp,t2_timestamp,days_installed,notes,uuid) */;
CREATE VIEW "clean_shoe_travel" AS select 
t1_shuttle as Shuttle,
t1_timestamp as Last_Updated,
shoe_travel as Shoe_Travel,
ltrim(days_installed,'0') as Days_Installed,
t1_distance as Shoes_Last_Distance,
t2_distance as Shoes_Change_Distance,
t2_timestamp as Shoes_Last_Changed,
notes as Notes,
uuid as UUID
from shoe_travel
/* clean_shoe_travel(Shuttle,Last_Updated,Shoe_Travel,Days_Installed,Shoes_Last_Distance,Shoes_Change_Distance,Shoes_Last_Changed,Notes,UUID) */;
