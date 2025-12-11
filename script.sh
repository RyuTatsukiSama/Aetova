#!/bin/bash

LOGFILE="Logs/log.log"

git fetch

difference=$(git diff origin/main --name-only Go/server)
if [ -n "$difference" ]; then
	git pull
	echo "$difference" >> $"LOGFILE"
	echo "$(../Go/server/build_server.sh)" >> "$LOGFILE"
else
	echo "There is no diff" >> "$LOGFILE"
fi

read