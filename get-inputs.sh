#!/bin/bash

# Load cookie from .cookie file
if [[ -f .cookie ]]; then
    COOKIE=$(cat .cookie | tr -d '\n\r ')
    if [[ -z "${COOKIE}" ]]; then
        echo "Error: .cookie file is empty" >&2
        exit 1
    fi
else
    echo "Error: .cookie file not found" >&2
    echo "Please create a .cookie file with your AoC session cookie" >&2
    exit 1
fi

WHOAMI=$(whoami)
UNAME=$(uname -n)
THISYEAR=$(( $(date +%Y) - 1 ))

for (( YEAR=2015; YEAR<="${THISYEAR}"; YEAR++ )); do
    for DAY in {1..25}; do
        URL="https://adventofcode.com/${YEAR}/day/${DAY}/input"
        OUTPUT="cmd/${YEAR}/inputs/$(printf "%02d" "${DAY}").txt"

        echo "Downloading ${YEAR} day ${DAY} from ${URL}"
        curl -A "${WHOAMI}-${UNAME} via curl" "${URL}" --progress-bar --cookie "session=${COOKIE}" -o "${OUTPUT}"

        sleep 5
    done
done
