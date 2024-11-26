#!/usr/bin/env bash

ROOT_DIR=$(realpath "$(dirname $0)/..")

mkdir -p $ROOT_DIR/static/channel-icons


cat $ROOT_DIR/static/videos.json | jq -r '.[] | .channelId' |

while read CHANNEL_ID; do
	ICON_FILE="$ROOT_DIR/static/channel-icons/$CHANNEL_ID.jpg"

	if [ -f "$ICON_FILE" ]; then
		echo "Skipping $CHANNEL_ID (file exists)"
		continue
	fi

    ICON_URL=$(curl "https://www.googleapis.com/youtube/v3/channels?part=snippet&id=$CHANNEL_ID&fields=items%2Fsnippet%2Fthumbnails&key=$YOUTUBE_KEY" | jq -r '.items[0].snippet.thumbnails.default.url')
	echo "Fetching icon for $CHANNEL_ID"
	curl -s "$ICON_URL" -o "$ICON_FILE"
done
