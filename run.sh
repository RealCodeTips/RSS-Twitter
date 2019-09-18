#!/bin/sh

cd $(dirname "$0")

go run -race main.go rss.go twitter.go