CREATE TABLE IF NOT EXISTS "DISTANCES" (
	"id"	INTEGER,
	"shuttle"	TEXT NOT NULL,
	"distance"	INTEGER NOT NULL,
	"timestamp"	TEXT,
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE sqlite_sequence(name,seq);
CREATE TABLE IF NOT EXISTS "CHANGE" (
	"id"	INTEGER,
	"shuttle"	TEXT NOT NULL,
	"distance"	INTEGER NOT NULL,
	"timestamp"	TEXT,
	"notes"	TEXT,
	"uuid"	TEXT,
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "SHOE_CHECK" (
	"id"	INTEGER,
	"shuttle"	TEXT NOT NULL,
	"distance"	INTEGER NOT NULL,
	"timestamp"	TEXT,
	"notes"	TEXT,
	"uuid"	TEXT,
	"wear"	REAL,
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE VIEW "view_travel" AS SELECT 
  DISTANCES.shuttle as t1_shuttle, 
  max(DISTANCES.distance) as t1_distance, 
  CHANGE.distance as t2_distance,
  (max(DISTANCES.distance)-max(CHANGE.distance))/1000 as shoe_travel,
  max(DISTANCES.timestamp) as t1_timestamp,
  max(CHANGE.timestamp) as t2_timestamp,
  (SELECT
  printf("%03d", julianday(max(DISTANCES.timestamp))-julianday(max(CHANGE.timestamp)))) as days_installed,
  CHANGE.notes as notes,
  CHANGE.uuid as uuid
FROM
  DISTANCES
INNER JOIN CHANGE on CHANGE.shuttle = DISTANCES.shuttle
GROUP BY t1_shuttle
ORDER BY shoe_travel DESC
/* view_travel(t1_shuttle,t1_distance,t2_distance,shoe_travel,t1_timestamp,t2_timestamp,days_installed,notes,uuid) */;
CREATE VIEW "view_change" AS
SELECT t1_shuttle AS Shuttle,
    t1_timestamp AS Last_Updated,
    shoe_travel AS Shoe_Travel,
    ltrim(days_installed, '0') AS Days_Installed,
    t1_distance AS Shoes_Last_Distance,
    t2_distance AS Shoes_Change_Distance,
    t2_timestamp AS Shoes_Last_Changed,
    notes AS Notes,
    uuid AS UUID
FROM view_travel
/* view_change(Shuttle,Last_Updated,Shoe_Travel,Days_Installed,Shoes_Last_Distance,Shoes_Change_Distance,Shoes_Last_Changed,Notes,UUID) */;
CREATE VIEW "view_check" AS
SELECT view_change.Shuttle AS Shuttle,
    view_change.Shoes_Change_Distance AS Zero_Distance,
    SHOE_CHECK.distance AS Last_Check_Distance,
    CASE
        WHEN SHOE_CHECK.distance > view_change.Shoes_Last_Changed + 1500000 THEN SHOE_CHECK.distance + 500000
        WHEN view_change.Shoes_Last_Distance > view_change.Shoes_Change_Distance + 1500000 THEN view_change.Shoes_Change_Distance + 1500000
        ELSE 0
    END Check_Trigger,
    view_change.Shoes_Last_Distance AS Current_Distance,
    CASE
        WHEN (
            view_change.Shoes_Last_Distance > (
                CASE
                    WHEN SHOE_CHECK.distance > view_change.Shoes_Last_Changed + 1500000 THEN SHOE_CHECK.distance + 500000
                    WHEN view_change.Shoes_Last_Distance > view_change.Shoes_Change_Distance + 1500000 THEN view_change.Shoes_Change_Distance + 1500000
                    ELSE 0
                END
            )
        ) THEN 'TRUE'
        ELSE 'FALSE'
    END Check_Shoes,
    SHOE_CHECK.timestamp AS Last_Check_Timestamp,
    SHOE_CHECK.notes AS Last_Check_Notes,
    SHOE_CHECK.uuid AS Last_Check_UUID,
    SHOE_CHECK.wear AS Last_Check_Wear
FROM SHOE_CHECK
    INNER JOIN view_change ON view_change.shuttle = SHOE_CHECK.shuttle
GROUP BY SHOE_CHECK.shuttle
ORDER BY Check_Shoes DESC,
    -- Ends with 11, 12, 13, 14; make level 4 top to level 1 bottom 
    CASE
        WHEN SHOE_CHECK.shuttle LIKE '%11' THEN 3
        WHEN SHOE_CHECK.shuttle LIKE '%12' THEN 2
        WHEN SHOE_CHECK.shuttle LIKE '%13' THEN 1
        WHEN SHOE_CHECK.shuttle LIKE '%14' THEN 0
    END,
    SHOE_CHECK.shuttle DESC
/* view_check(Shuttle,Zero_Distance,Last_Check_Distance,Check_Trigger,Current_Distance,Check_Shoes,Last_Check_Timestamp,Last_Check_Notes,Last_Check_UUID,Last_Check_Wear) */;
