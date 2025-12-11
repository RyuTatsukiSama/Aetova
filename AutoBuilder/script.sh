#!/bin/bash

LOGFILE="./logs/build_server.log"

mkdir -p "$(dirname "$LOGFILE")"

git fetch

difference=$(git diff origin/main --name-only Go/server)
if [ -n "$difference" ]; then
	echo "//--------- $(date '+%Y-%m-%d %H:%M:%S') AUTO BUILD LAUNCH---------//" >> "$LOGFILE"
	echo "$difference" >> "$LOGFILE"
	git pull
	echo "$(../Go/server/build_server.sh)" >> "$LOGFILE"
fi