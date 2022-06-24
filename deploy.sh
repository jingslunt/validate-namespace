#!/bin/bash

echo "Creating certificates"
mkdir certs
openssl req -nodes -new -x509 -keyout certs/ca.key -out certs/ca.crt -subj "/CN=Admission Controller Demo"
openssl genrsa -out certs/validate-delete-namespace-tls.key 2048
openssl req -new -key certs/validate-delete-namespace-tls.key -subj "/CN=validate-delete-namespace.default.svc.cluster.local" | openssl x509 -req -CA certs/ca.crt -CAkey certs/ca.key -CAcreateserial -out certs/validate-delete-namespace-tls.crt

echo "Creating k8s Secret"
kubectl create secret tls validate-delete-namespace-tls \
    --cert "certs/validate-delete-namespace-tls.crt" \
    --key "certs/validate-delete-namespace-tls.key"

echo "Creating k8s admission deployment"
kubectl create -f deployment.yaml

echo "Creating k8s webhooks for demo"
CA_BUNDLE=$(cat certs/ca.crt | base64 | tr -d '\n')
sed -e 's@${CA_BUNDLE}@'"$CA_BUNDLE"'@g' <"validate-delete-namespace.yaml" | kubectl create -f -