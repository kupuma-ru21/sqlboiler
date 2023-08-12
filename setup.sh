#!/usr/local/bin/bash

set -eu

readonly DBFILE_NAME="mygraphql.db"

# Create DB file
if [ ! -e ${DBFILE_NAME} ];then
  echo ".open ${DBFILE_NAME}" | sqlite3
fi

# Create DB Tables
echo "creating tables..."
sqlite3 ${DBFILE_NAME} "
PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS Users(\
    ID TEXT NOT NULL UNIQUE,\
    Username VARCHAR (127) NOT NULL UNIQUE,\
    Password VARCHAR (127) NOT NULL,\
    PRIMARY KEY (ID)\
);

CREATE TABLE IF NOT EXISTS Links(\
    ID TEXT NOT NULL UNIQUE,\
    Title VARCHAR (255) ,\
    Address VARCHAR (255) ,\
    UserID INT ,\
    FOREIGN KEY (UserID) REFERENCES Users(ID) ,\
    PRIMARY KEY (ID)\
)
"