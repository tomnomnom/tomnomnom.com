#!/usr/bin/env bash
mkdir -p static/thumbnails
while read VIDEO_ID; do
	curl -s "https://i.ytimg.com/vi/$VIDEO_ID/sddefault.jpg" -o "static/thumbnails/$VIDEO_ID.jpg"
done
