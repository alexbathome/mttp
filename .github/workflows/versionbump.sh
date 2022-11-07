#!/usr/bin/bash

# Read the version
version=$(cat version)
versionString="mttp/v$version"

# Push up the new version to GH
git tag $versionString
git push origin $versionString

# Verify with go list
GOPROXY=proxy.golang.org go list -m github.com/alexbathome/mttp/mttp@v$version

# Update the go.pkg
curl https://sum.golang.org/lookup/github.com/alexbathome/mttp/mttp@v$version
