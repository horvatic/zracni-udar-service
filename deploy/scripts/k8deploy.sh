#!/bin/bash


if [ -z "$1" ]
then
    echo "Must input a correct short git sha"
    exit 1
fi


if [ -z "$2" ]
then
    echo "Must input a correct deployment id"
    exit 1
fi

if [ -z "$3" ]
then
    echo "Must input a correct environment"
    exit 1
fi

SHORT_GIT_SHA="$1"
FOLDER_ID="$2"
ENV="$3"

cd /deployments/work/$FOLDER_ID

export TAG="$SHORT_GIT_SHA"

envsubst < deploy/base/deployment.yaml > tempbase
cat tempbase > deploy/base/deployment.yaml
rm tempbase

export TAG=""
export MONGO_CONNECTION_STRING="$(cat /deployments/secrets/zracni-udar-service/$ENV/MONGO_CONNECTION_STRING)"
export MONGO_DATABASE="$(cat /deployments/secrets/zracni-udar-service/$ENV/MONGO_DATABASE)"
export MONGO_COLLECTION="$(cat /deployments/secrets/zracni-udar-service/$ENV/MONGO_COLLECTION)"
export FRONT_END_HOST="$(cat /deployments/secrets/zracni-udar-service/$ENV/FRONT_END_HOST)"
export GITHUB_PAT="$(cat /deployments/secrets/zracni-udar-service/$ENV/GITHUB_PAT)"

envsubst '${MONGO_CONNECTION_STRING} ${MONGO_DATABASE} ${MONGO_COLLECTION} ${FRONT_END_HOST} ${GITHUB_PAT}' < "deploy/$ENV/secret.yaml" > tempsecret
cat tempsecret > "deploy/$ENV/secret.yaml"
rm tempsecret

microk8s kubectl apply -k "deploy/$ENV" -n "$ENV"

sleep 10s

microk8s kubectl wait --for=condition=ready pod -l app=zus -n "$ENV" --timeout=10m

cd /deployments

rm -rf /deployments/work/$FOLDER_ID
