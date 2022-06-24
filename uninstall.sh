#!/bin/bash

echo "卸载"
kubectl delete -f deployment.yaml
kubectl delete -f webhook.yaml
kubectl delete secret validate-delete-namespace-tls