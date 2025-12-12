#!/bin/bash

LOGFILE="./logs/build_server.log"

mkdir -p "$(dirname "$LOGFILE")"

echo "//--------- $(date '+%Y-%m-%d %H:%M:%S') AUTO BUILD LAUNCH---------//" >> "$LOGFILE"

git fetch

difference=$(git diff origin/main --name-only ../Go/server)
if [ -n "$difference" ]; then
	echo "$difference" >> "$LOGFILE"
	git pull
	sh ../Go/server/build_server.sh >> "$LOGFILE" 2>&1
else
	echo "There is no difference" >> "$LOGFILE"
fi

echo "//--------- $(date '+%Y-%m-%d %H:%M:%S') AUTO BUILD END---------//" >> "$LOGFILE"