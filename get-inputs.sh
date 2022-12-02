#!/bin/bash

COOKIE=your-aoc-session-cookie-here

for YEAR in {2015..2022}; do
    for DAY in {1..25}; do
        echo "Downloading $YEAR day $DAY:"
        curl -A "$(whoami)-$(uname -n) via curl" https://adventofcode.com/$YEAR/day/$DAY/input --progress-bar --cookie "session=$COOKIE" -o inputs/$YEAR/$(printf "%02d" $DAY).txt
        sleep 5
        echo
    done
done