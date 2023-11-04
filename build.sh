#!/bin/bash

THISYEAR=$(date +%Y)
STARTYEAR=2015
# STARTYEAR="${THISYEAR}"

for (( YEAR="${STARTYEAR}"; YEAR<="${THISYEAR}"; YEAR++ )); do
    echo -e "\e[1mBuilding Linux ${YEAR}...\e[0m"
    env GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-s -w" -a -o "bin/aoc-${YEAR}" "./cmd/${YEAR}"

    echo -e "\e[1mBuilding Windows ${YEAR}...\e[0m"
    env GOOS=windows GOARCH=amd64 go build -trimpath -ldflags "-s -w" -a -o "bin/aoc-${YEAR}.exe" "./cmd/${YEAR}"

    echo -e "\e[1mDone ${YEAR}.\e[0m\n"
done

echo -e "\e[1mAll done.\e[0m\n"
sleep 1
ls -ghl bin/

echo
file bin/*
