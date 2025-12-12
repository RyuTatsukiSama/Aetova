#!/bin/bash

LOGFILE="./logs/build_server.log"

mkdir -p "$(dirname "$LOGFILE")"

echo "//--------- $(date '+%Y-%m-%d %H:%M:%S') AUTO BUILD LAUNCH---------//" >> "$LOGFILE"

git fetch

difference=$(git diff origin/main --name-only Go/server)
if [ -n "$difference" ]; then
	echo "$difference" >> "$LOGFILE"
	git pull
	echo "$(../Go/server/build_server.sh)" >> "$LOGFILE"
else
	echo "There is no difference" >> "$LOGFILE"
fi