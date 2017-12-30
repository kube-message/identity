#!/usr/bin/env bash
set -e

if [[ $# -lt 1 ]]; then
    echo "Usage: build_and_push_image.sh -t IMAGE_VERSION [-p]" >&2
    exit 1
fi

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/user-svc main.go

while getopts "v:p" OPTION
do
    case $OPTION in
        v | --version)
            new_tag=app-${OPTARG}
            echo "Using tag: app-$new_tag"
            docker build -t user:"$new_tag" .
            docker tag user:"$new_tag" imauld/user:"$new_tag"
            echo "Built image: user:$new_tag"
            ;;
        p | --push)
            echo "Pushing image to DockerHub: user:$new_tag"
            docker push imauld/user:$new_tag
            ;;
    esac
done
