#!/bin/bash
godog -f cucumber | json-to-messages > messages.json

TARGET=$(curl -s -D - -o /dev/null -H "Authorization: Bearer $CUCUMBER_PUBLISH_TOKEN" "https://messages.cucumber.io/api/reports" | grep -iF location | sed "s/location: //I" | tr -d '\r')

curl -v --upload-file messages.json "$TARGET"
