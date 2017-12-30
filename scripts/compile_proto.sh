#!/usr/bin/env bash
set -e

protoc -I ~/repos/kube/proto ~/repos/kube/proto/identity.proto --go_out=plugins=grpc:proto
