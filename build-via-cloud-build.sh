#!/usr/bin/env bash

artifactRegion="asia-southeast1-docker.pkg.dev"
gcpProjectName="tdg-dh-truehealth-core-nonprod"
repository="cossack-docker"
artifactName="poc-go-todo-service"
tag="$1"
prefixArtifactName="${asia-southeast1-docker.pkg.dev}/${gcpProjectName}/${repository}/${artifactName}"

gcloud builds submit \
    --tag "${prefixArtifactName}:${tag}"
