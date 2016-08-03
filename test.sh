#!/bin/bash
curl localhost:8000
curl 0.0.0.0:8000
curl 127.0.0.1:8000
status=$(curl --write-out "%{http_code}\n" --silent --output /dev/null localhost:8000)

if [ $status -ne 200 ]; then
  exit 1
fi

exit 0
