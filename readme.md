# Advent of Code, in Go

I've attempted [Advent of Code](https://adventofcode.com) every year since it's creation in a variety of languages and with varying degrees of success, but at the time of writing, I'm all about [Go](https://go.dev).

This repo is my attempt to split out the Go attempts (my current language of choice) from all the other languages I've used (which are [still available elsewhere](https://github.com/vaughany/advent_of_code) and create one application per year which handles all the days of the challenges.

One of the things I did not do early on was to always parse the input files - sometimes, with the more complex files, it was easier to dump it all into a string or two within the code itself and work with that.  But now, I have the input files per-year in `./cmd/[year]/inputs/` and will endeavour to parse the file no matter how complex it is. However, this is not a 'parse the input file' challenge but a problem solving challenge, so I will focus on that primarily and parsing overly-complex input files as a secondary consideration.
