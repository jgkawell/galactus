# registry

This service functions as a registry for other services. Providing a way for other services to register themselves and to get notified when other services are registered or inactive.

## Idea

Process:

- On bootup, the regsitry service will connect to its DB (Postgres) and create a table for storing the service values
- It will expose a gRPC API that will have the following methods: `Register()` (anything else?)
- The other services at bootup will connect to the registry service and register themselves.
  - They will send all their information (listed below) to the registry service
  - The registry service will return the ports that the service should listen on (e.g. gRPC=8090 when remote, generated value when local)
  - The registry service will send a notification to any subscribed services that the service is now active (e.g. eventstore, commandhandler, etc.)
- When a service wishes to make a connection to another service, it will first send a request to the registry service asking for connection information. If the service is registered and active, the registry will return the connection information.
- The registry is responsible for keeping track of the services that are registered and active. It will do this by polling a health endpoint OR checking the service status from k8s (only when remote).
  - Question: If a service is inactive for a while, should the registry still keep track of it? Should it unregister it so we don't keep polling the health endpoint?
- The registry is responsible for ALL RabbitMQ configuration. All the services do is connection to precreated Exchange/Queues/Topics.

Data model:

- Each entry will be all the information needed to understand the interface and status of a service. Basically, if you query the registry for a service, you will get all the information to be able to make sync (gRPC/HTTP) or async (broker) requests to it.
- Each entry will have the following fields:
  - `name`: The name of the service.
  - `version`: The version of the service.
  - `description`: A description of the service.
  - `status`: The status of the service.
  - `address`: The address of the service.
  - `apis`: This is a repeated list of structs that have the following fields:
    - `name`: The name of the API (gRPC or HTTP)
    - `port`: The port of the API.
  - `brokers`: This is a repeated list of structs that have the following fields:
    - `exchange`: The exchange the service is listening on.
    - `routing_key`: The routing key the service is listening for on the above exchange.
    - `type`: The type of listener (queue or topic).
