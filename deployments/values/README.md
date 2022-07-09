# values

Place environment specific deployment configuration values in this directory.

Some notes to keep in mind:

- Values in `all.yaml` are applied to all namespaces.
- Values matching the pattern {{env}}.yaml are applied per deployment stage and can override values in `all.yaml`.
- The `developer-namespace.yaml` values are used by the top-level Makefile during deployments to dev namespaces from a development machine only.
