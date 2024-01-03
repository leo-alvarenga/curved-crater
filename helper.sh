#!/bin/bash

PRODUCT=$1
DB="analytics.db"

if [[ -z "$PRODUCT" ]];
then
    sqlite3 "$DB" "SELECT * FROM Event"
    exit 0
fi

sqlite3 "$DB" "SELECT * FROM Event WHERE product = '$PRODUCT'"

printf "\nEvent count for $PRODUCT: "
sqlite3 "$DB" "SELECT Count(*) FROM Event WHERE product = '$PRODUCT'"
