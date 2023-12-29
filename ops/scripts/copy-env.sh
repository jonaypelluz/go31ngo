#!/bin/sh
if [[ -f ".env" ]]; then
    rm .env
fi

cp -n ops/docker/go/.env.development .env

