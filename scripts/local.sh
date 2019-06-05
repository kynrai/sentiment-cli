#!/usr/bin/env bash

if [ ! -r .env ]; then
  echo 'You need a .env file first. Example:'
  echo 'ENV=VALUE'
  exit 1
fi


export $(cat .env | xargs)
sentiment-cli
