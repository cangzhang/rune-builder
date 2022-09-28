#!/usr/bin/env bash

args="$@"

repo="cangzhang/data-maker"
url=$(curl --silent "https://api.github.com/repos/$repo/releases/latest" | jq -r '.assets | map(select(.name | contains("linux")))[0].browser_download_url')
wget -q "$url" -O data-maker
chmod +x data-maker
./data-maker -v
./data-maker "$args"

npm install -g @jsdevtools/npm-publish

dir=$(pwd)
echo "$dir"

cd output/ || echo "cd <output/> failed"
for d in */; do
  echo "$d"
  cd "$d" || echo "cd <$d> failed"
  count=$(find . -type f -maxdepth 1 |  wc -l | xargs)
  if [ "$count" -gt 10 ]; then
    npm-publish --access public --token "$token"
  else
    echo "skipped publish $d"
  fi
  cd ..
done

cd "$dir"

curl ip.sb
