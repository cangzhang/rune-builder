#!/usr/bin/env bash

for d in output/*/ ; do
    npx --yes tsx scripts/change-pkg-name.js "${d}package.json"
#    npm-publish --access public --token "$token" --dry-run
done
