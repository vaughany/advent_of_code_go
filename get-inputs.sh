#!/bin/bash

COOKIE=your-aoc-session-cookie-here
WHOAMI=$(whoami)
UNAME=$(uname -n)
THISYEAR=$(date +%Y)

for (( YEAR=2015; YEAR<="${THISYEAR}"; YEAR++ )); do
    for DAY in {1..25}; do
        URL="https://adventofcode.com/${YEAR}/day/${DAY}/input"
        OUTPUT="cmd/${YEAR}/inputs/$(printf "%02d" "${DAY}").txt"

        echo "Downloading ${YEAR} day ${DAY} from ${URL}"
        curl -A "${WHOAMI}-${UNAME} via curl" "${URL}" --progress-bar --cookie "session=${COOKIE}" -o "${OUTPUT}"

        sleep 5
    done
done
