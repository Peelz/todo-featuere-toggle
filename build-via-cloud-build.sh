#!/usr/bin/env bash

echo "building $GO_PIPELINE_LABEL"

artifactRegion="asia-southeast1-docker.pkg.dev"
gcpProjectName="tdg-dh-truehealth-core-nonprod"
repository="cossack-docker"
artifactName="poc-go-todo-service"
tag="latest"
if [[ ! -z "$GO_PIPELINE_LABEL" ]]; then
    tag="$GO_PIPELINE_LABEL"
fi
prefixArtifactName="${artifactRegion}/${gcpProjectName}/${repository}/${artifactName}"

cat deployment/cloudbuild.yaml

gcloud builds submit \
    --config=build
    --suppress-logs 