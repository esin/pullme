#!/bin/bash
dt=$(date +%y%m%d%H%M)
docker build -t pullme:latest -t pullme:$dt -f Dockerfile .
if [ $? -ne 0 ]; then
  exit 1;
fi

docker tag pullme:latest es1n/pullme:latest
docker tag pullme:$dt es1n/pullme:$dt
docker push es1n/pullme:$dt
docker push es1n/pullme:latest