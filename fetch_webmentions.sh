#!/usr/bin/env bash

webmentionskey="crV19oTq-qnH4pHtaxnWJA"

# Fetch mentions from the API
response=$(curl -s "https://webmention.io/api/mentions.jf2?token=${webmentionskey}&per-page=1000")

# Parse the JSON response to get the children
newMentions=$(echo "$response" | jq '.children')

# Output the new mentions
echo "$newMentions" >> static/webmentions.json

