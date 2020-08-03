#!/bin/bash

function getToken() {
    echo `date` "Generating token"
    export TOKEN="$(curl --silent --header 'GET' "https://auth.docker.io/token?service=registry.docker.io&scope=repository:es1n/pullme:pull" | jq -r '.token' )"
}

while true; do
    exitCode=$(curl --silent --header "Authorization: Bearer ${TOKEN}" -o /dev/null -s -w "%{http_code}\n" https://registry-1.docker.io/v2/es1n/pullme/manifests/latest)
    if [[ "$exitCode" != "200" ]]; then
        getToken
    fi

    sleep 0.1
done