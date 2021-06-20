#!/bin/bash
godog -f cucumber | json-to-messages > messages.json

TARGET=$(http "https://messages.cucumber.io/api/reports" "Authorization: Bearer $CUCUMBER_PUBLISH_TOKEN" --headers | grep Location | sed "s/Location: //" | tr -d '\r')

http -v PUT $TARGET @./messages.json
