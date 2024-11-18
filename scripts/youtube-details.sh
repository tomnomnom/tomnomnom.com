#!/usr/bin/env bash

if [ -z "$YOUTUBE_KEY" ]; then
    echo "YOUTUBE_KEY not set"
    exit 1
fi

while read VIDEO_ID; do
    curl -s "https://www.googleapis.com/youtube/v3/videos?part=snippet&id=$VIDEO_ID&key=$YOUTUBE_KEY" | jq '{
        id: .items[0].id,
        published: .items[0].snippet.publishedAt,
        title: .items[0].snippet.localized.title,
        channel: .items[0].snippet.channelTitle,
        channelId: .items[0].snippet.channelId,
    }'
done | jq --slurp
