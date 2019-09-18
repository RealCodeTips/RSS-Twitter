# RSS-Twitter

## Introduction
Quick repository to hook one or more RSS feeds into a Twitter bot

## Set-up
- Create a `.env` file with the following Twitter credentials:
  - CONSUMER_KEY
  - CONSUMER_SECRET
  - ACCESS_TOKEN
  - ACCESS_TOKEN_SECRET
- Add the RSS feeds you want to use in `config.json`

## Running
A `run.sh` script has been provided to run this script locally.

## Building
A `build.sh` script has been provided which will build Linux and Darwin executables under a `/build` directory. 