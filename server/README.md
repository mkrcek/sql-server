# ToolBag database


**Remember:** 
You need to capitalize to have a variable, type, or func exported from a package


# create database
  
```
CREATE DATABASE devices;
```

### list databases
```
\l
```

### connect to a database
```
\c devices
```

### drop (remove, delete) database
```
DROP DATABASE <database name>;
```

# create table - auto increment ID
```
CREATE TABLE mydevices (
   deviceID SERIAL PRIMARY KEY     NOT NULL,
   deviceName           TEXT    NOT NULL,
   deviceLocation       TEXT    NOT NULL,
   deviceIP             TEXT    NOT NULL,
   type                 TEXT    NOT NULL,
   deviceBoard          TEXT    NOT NULL,
   deviceSwVersion      TEXT    NOT NULL,
   targetServer         TEXT    NOT NULL,
   httpPort             TEXT    NOT NULL,
   note                 TEXT    NOT NULL
);
```


### create table - unique ID
```
CREATE TABLE devices (
   deviceID INT PRIMARY KEY     NOT NULL,
   deviceName           TEXT    NOT NULL,
   deviceLocation       TEXT    NOT NULL,
   deviceIP             TEXT    NOT NULL,
   type                 TEXT    NOT NULL,
   deviceBoard          TEXT    NOT NULL,
   deviceSwVersion      TEXT    NOT NULL,
   targetServer         TEXT    NOT NULL,
   httpPort             TEXT    NOT NULL,
   note                 TEXT    NOT NULL
);
```

### show tables in a database (list down)
```
\d
```

### show details of a table
```
\d <table name>
```

### drop (remove, delete) a table 
delete
```
DROP TABLE <table name>;
```
### rename a column in a database table 
```
ALTER TABLE mydevices RENAME COLUMN type TO deviceType;
```
           



# insert a record - auto increment
```
INSERT INTO mydevices (deviceName, deviceLocation, deviceIP, deviceType, deviceBoard, deviceSwVersion, targetServer, httpPort, note) VALUES ('Garage Controller','Garage','192.168.0.44','Arduino','','2017-11-28','192.168.0.18','9090','super device');
```
and - unique manual ID
```
INSERT INTO devices (deviceID, deviceName, deviceLocation, deviceIP, deviceType, deviceBoard, deviceSwVersion, targetServer, httpPort, note) VALUES (1, 'Garage Controller','Garage','192.168.0.44','Arduino','','2017-11-28','192.168.0.18','9090','super device');
```

### list records in a table
```
SELECT * FROM devices;
```
a) autoincrement
```
INSERT INTO mydevices (deviceName, deviceLocation, deviceIP, deviceType, deviceBoard, deviceSwVersion, targetServer, httpPort, note) 
VALUES 
('Room Lights','House','192.168.0.45','Arduino','','2017-11-28','192.168.0.18','9091','Ovlada svetla v mistnosti'),
('House ','House','192.168.0.46','Arduino','','2017-11-29','192.168.0.18','9091','A blika celym domem');

```
b) manual
```
INSERT INTO devices (deviceID, deviceName, deviceLocation, deviceIP, deviceType, deviceBoard, deviceSwVersion, targetServer, httpPort, note) 
VALUES 
(2, 'Room Lights','House','192.168.0.45','Arduino','','2017-11-28','192.168.0.18','9091','Ovlada svetla v mistnosti'),
(3, 'House ','House','192.168.0.46','Arduino','','2017-11-29','192.168.0.18','9091','A blika celym domem');

```

# update

syntax
```
UPDATE table
SET col1 = val1, col2 = val2, ..., colN = valN
WHERE <condition>;
```

```
SELECT * FROM devices;
```

####Examples:
```
UPDATE mydevices SET deviceName = 'REMOTE Garage Controller' WHERE deviceID = 3;
```
add to all rows at the tables
```
UPDATE mydevices SET deviceBoard = 'RobotDyn Wifi D1R2';
```


## order by
```
SELECT * FROM mydevices ORDER BY deviceId;
```
```
SELECT * FROM mydevices ORDER BY -deviceId;
```

# delete

syntax
```
DELETE FROM table
WHERE <condition>;
```

```
SELECT * FROM mydevices;
```

```
DELETE FROM mydevices WHERE deviceId = 4;
```

**WARNING: this deletes all records:**
```
DELETE FROM mydevices;
```

# users & privileges

## see current user
```
SELECT current_user;
```

## details of users
```
\du
```

## create user
```
CREATE USER bond WITH PASSWORD 'password';
```

## grant privileges
```
GRANT ALL PRIVILEGES ON DATABASE devices to bond;
```
to DATABESE name