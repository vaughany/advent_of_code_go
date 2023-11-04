#!/bin/bash

COOKIE=your-aoc-session-cookie-here
WHOAMI=$(whoami)
UNAME=$(uname -n)

for YEAR in {2015..2023}; do
    for DAY in {1..25}; do
        echo "Downloading ${YEAR} day ${DAY}:"
        OUTPUT="cmd/${YEAR}/inputs/$(printf "%02d" "${DAY}").txt"
        curl -A "${WHOAMI}-${UNAME} via curl" "https://adventofcode.com/${YEAR}/day/${DAY}/input" --progress-bar --cookie "session=${COOKIE}" -o "${OUTPUT}"
        sleep 5
        echo
    done
done