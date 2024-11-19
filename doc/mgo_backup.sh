#!/bin/bash
source /etc/profile

echo ""
echo ""
echo "start backup weedone db"

DIR=`date +%d-%m-%y`
DEST=/home/backup/db/$DIR
mkdir $DEST

echo $DEST
mongodump -h localhost:27017 -d weedone -o $DEST
mongodump -h localhost:27017 -d insou -o $DEST

echo "backup finish :)"
title_temp="Auto:_Mongo_Backup_Finish_[DataBase:weedone]" 
title="${title_temp}_${DIR}" 

curl -X POST -H 'Content-type: application/json' --data '{"text":"'${title}'"}' https://hooks.slack.com/services/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