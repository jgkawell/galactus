# How to install shawarma

Links:

- [shawarma](https://github.com/CenterEdge/shawarma)
- [cert-manager](https://github.com/jetstack/cert-manager)

## Overview

This folder contains the necessary files to install `shawarma` to a k8s cluster:

- `cert-manager.yaml`: this manages TLS between the API server and the webhook.
- `k8s-sidecar-injector.yaml`: this sets up the webhook service that will be used to inject the sidecar into into pods automatically
- `rbac.yaml`: this configures RBAC rules for `shawarma`

## Steps

**NOTE**: Steps 1 and 2 only need to be run once per cluster. Steps 3 and 4 must be run once for every namespace where you want shawarma to be injected as a sidecar to each pod.

1. Install `cert-manager`: `kubectl apply -f cert-manager.yaml`
2. Install the webhook service: `kubectl apply -f k8s-sidecar-injector.yaml`
3. Apply RBAC rules: `kubectl apply -f rbac.yaml`
    - **NOTE**: These RBAC rules need to be applied to all namespaces where you want the `shawarma` sidecar to be injected. In other words, you need to edit the `namespace` values in the `rbac.yaml` file to match the namespace you want to apply the RBAC rules to.
4. Label the namespace where you want the sidecar injected. This should match the namespace you applied the RBAC rules to in step 3. Run: `kubectl label namespace <namespace> shawarma-injection=enabled`
