# `/functions`

This directory is for all functions that are a part of the application. Although `galactus` is mainly designed around Kubernetes microservices, functions like AWS Lambda and Azure Functions are a common tool and can be written here alongside the microservices so they can share common tooling and modules. For example, a function here can use the same gRPC interface defined under `/api` and can use common modules defined under `/pkg`.
