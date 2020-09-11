#!/bin/bash

# Prints the changelog for just the latest tag
git-chglog -o LATESTRELEASE.md $(git describe --tags $(git rev-list --tags --max-count=1)) 

# Prints the changelog for all unreleased changes. "UPCOMING" needs to be an unused tag
git-chglog --next-tag UPCOMING -o UPCOMINGRELEASE.md UPCOMING 