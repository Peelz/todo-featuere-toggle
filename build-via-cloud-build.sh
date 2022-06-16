#!/usr/bin/env bash

artifactRegion="asia-southeast1-docker.pkg.dev"
gcpProjectName="tdg-dh-truehealth-core-nonprod"
repository="cossack-docker"
artifactName="poc-go-todo-service"

tag="latest"
if [[ ! -z "$GO_PIPELINE_LABEL" ]]; then
    tag="$GO_PIPELINE_LABEL"
fi

imageName="${artifactRegion}/${gcpProjectName}/${repository}/${artifactName}"

cat deployment/cloudbuild.yaml |
    sed s/{{.tag}}/${tag}/g |
    sed s/{{.image}}/${imageName}/g >cloudbuild.yaml

gcloud builds submit
