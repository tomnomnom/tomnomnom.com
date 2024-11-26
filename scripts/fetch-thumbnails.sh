#!/usr/bin/env bash

ROOT_DIR=$(realpath "$(dirname $0)/..")

mkdir -p $ROOT_DIR/static/thumbnails


cat $ROOT_DIR/static/videos.json | jq -r '.[] | .id' |

while read VIDEO_ID; do
	THUMB_FILE="$ROOT_DIR/static/thumbnails/$VIDEO_ID.jpg"

	if [ -f "$THUMB_FILE" ]; then
		echo "Skipping $VIDEO_ID (file exists)"
		continue
	fi

	echo "Fetching thumbnail for $VIDEO_ID"
	curl -s "https://i.ytimg.com/vi/$VIDEO_ID/mqdefault.jpg" -o "$THUMB_FILE"
done
