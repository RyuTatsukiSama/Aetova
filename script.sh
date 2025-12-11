#!/bin/bash

git fetch

difference=$(git diff origin/main --name-only Go/server)
if [ -n "$difference" ]; then
	git pull
	./Go/server/build_server.sh
else
	echo "There is no diff"
fi

read