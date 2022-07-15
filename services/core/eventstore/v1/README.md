# EventStore
A grpc service integrated with mongo, and RabbitMQ for storing, and publishing system events.

## Problem Statement
Currently all the events the system should know about are not logged, in a consolidated place
other than the logging system which is much more verbose. 

Domain Event: Is a record of some business-significant occurrence in a bounded context.

## Service Interface
```proto
service EventStore {
  // Query events for using the provided key/val for lookup
  rpc QueryEvents(ListEventRequest) returns (ListEventResponse) {}

  // Create a new event in the event store
  rpc CreateEvent(CreateEventRequest) returns (CreateEventResponse) {}
}
```

## Model
```proto
// NOTE - Check the protorepo for the exact model. This is only a reference.

// Event - Is the model of the event store
message Event {
  // uuid generated in the grpc handler 
  string event_id = 1;

  // UnixMilli -> Number of milliseconds elapsted since 00:00:00 UTC on 1 Jan 1970
  // this timestamp is set in the grpc handler
  int64 timestamp = 2;

  // event_template is currently a string b/c all events types are unkown. Long term
  // we wold like to make this an enum.
  string event_type = 3;
  string aggregate_id = 4;
  string aggregate_type = 5; 

  // data used to augment the state of the system
  // NOTE: Currently we will use a json string, that will need to serialize to to a specific type
  string event_data = 6;

  // UnixMilli -> Number of milliseconds elapsted since 00:00:00 UTC on 1 Jan 1970
  // this timestamp is set when the event has been published to it's respective topic/queue
  int64 published_date = 7;
}
```

## Run
Their are two ways of running this service. First, you can deploy it to your namespace. Or you can run it locally as long as mongo is running on localhost:27017.
```shell
# from the root of galactus, deploy to remote namespace
$ make eventstore

# from the root of the eventstore service
$ go run main.go
```

