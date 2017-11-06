set -x
set -e

docker build -t imauld/user-server:base-0.0.1 -f Dockerfile.base .
