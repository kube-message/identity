set -e
set -x

location="$( dirname $0 )"
# pushd $location
# ./scripts/build_application.sh
docker build -t imauld/user-server:0.0.1 .
# popd