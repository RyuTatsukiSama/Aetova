#!/bin/bash

LOGFILE="$HOME/Aetova/Aetova/AutoBuilder/logs/build_server.log"

mkdir -p "$(dirname "$LOGFILE")"

echo "//--------- $(date '+%Y-%m-%d %H:%M:%S') AUTO BUILD LAUNCH---------//" >> "$LOGFILE"
echo >> "$LOGFILE"

git fetch

difference=$(git diff origin/main --name-only ../Go/server)
if [[ -n "$difference" ]]; then
	echo "$difference" >> "$LOGFILE"
	git pull
	(cd ../Go/server && sh build_server.sh) >> "$LOGFILE" 2>&1
else
	echo "There is no difference" >> "$LOGFILE"
fi

echo >> "$LOGFILE"
echo "//--------- $(date '+%Y-%m-%d %H:%M:%S') AUTO BUILD END---------//" >> "$LOGFILE"
echo >> "$LOGFILE"
echo >> "$LOGFILE"