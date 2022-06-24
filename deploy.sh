#!/bin/bash

sh create-cert.sh
cat ./validate-delete-namespace.yaml | sh ./patch-webhook-ca.sh > ./webhook.yaml

echo "Creating k8s admission deployment"
kubectl apply -f deployment.yaml
kubectl apply -f webhook.yaml