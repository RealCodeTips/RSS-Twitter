#!/bin/sh

cd $(dirname $BASH_SOURCE)

GIT_COMMIT=$(git rev-parse HEAD)
if [ -f build/commit.dat ]; then
  LAST_COMMIT=$(cat build/commit.dat)
fi

if [ "$GIT_COMMIT" == "$LAST_COMMIT" ]; then
  echo "Nothing new"
  exit 0
fi

rm -rf build

mkdir -p build/darwin
mkdir -p build/linux

OLD_VER=$(awk '/^const version =/ {print $0}' version.go)
NEW_VER=$(echo $OLD_VER | awk -F'[ .]' '/^const version =/ {print $1,$2,$3,$4"."$5"."$6+1"\""}')

sed -i .bak "s/$OLD_VER/$NEW_VER/g" version.go

VER=$(awk -F'[ ."]' '/^const version =/ {print $5"."$6"."$7}' version.go)

git add version.go

git commit -m "New version - $VER"

GIT_COMMIT=$(git rev-parse HEAD)

go build -o build/darwin/rss_feeder -ldflags="-X main.commit=${GIT_COMMIT}"
env GOOS=linux GOARCH=amd64 go build -o build/linux/rss_feeder -ldflags="-X main.commit=${GIT_COMMIT}"

echo $GIT_COMMIT > build/commit.dat