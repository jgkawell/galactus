# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [options/gorm.proto](#options_gorm-proto)
    - [AutoServerOptions](#gorm-AutoServerOptions)
    - [BelongsToOptions](#gorm-BelongsToOptions)
    - [ExtraField](#gorm-ExtraField)
    - [GormFieldOptions](#gorm-GormFieldOptions)
    - [GormFileOptions](#gorm-GormFileOptions)
    - [GormMessageOptions](#gorm-GormMessageOptions)
    - [GormTag](#gorm-GormTag)
    - [HasManyOptions](#gorm-HasManyOptions)
    - [HasOneOptions](#gorm-HasOneOptions)
    - [ManyToManyOptions](#gorm-ManyToManyOptions)
    - [MethodOptions](#gorm-MethodOptions)
  
    - [File-level Extensions](#options_gorm-proto-extensions)
    - [File-level Extensions](#options_gorm-proto-extensions)
    - [File-level Extensions](#options_gorm-proto-extensions)
    - [File-level Extensions](#options_gorm-proto-extensions)
    - [File-level Extensions](#options_gorm-proto-extensions)
  
- [core/aggregates/v1/events.proto](#core_aggregates_v1_events-proto)
    - [Event](#core-aggregates-v1-Event)
  
- [core/aggregates/v1/registry.proto](#core_aggregates_v1_registry-proto)
    - [Consumer](#core-aggregates-v1-Consumer)
    - [Protocol](#core-aggregates-v1-Protocol)
    - [Registration](#core-aggregates-v1-Registration)
  
    - [ConsumerKind](#core-aggregates-v1-ConsumerKind)
    - [ProtocolKind](#core-aggregates-v1-ProtocolKind)
    - [ServiceStatus](#core-aggregates-v1-ServiceStatus)
  
- [core/commandhandler/v1/commandhandler.proto](#core_commandhandler_v1_commandhandler-proto)
    - [ApplyCommandRequest](#core-commandhandler-v1-ApplyCommandRequest)
    - [ApplyCommandResponse](#core-commandhandler-v1-ApplyCommandResponse)
  
    - [CommandHandler](#core-commandhandler-v1-CommandHandler)
  
- [core/eventstore/v1/eventstore.proto](#core_eventstore_v1_eventstore-proto)
    - [CreateRequest](#core-eventstore-v1-CreateRequest)
    - [CreateResponse](#core-eventstore-v1-CreateResponse)
  
    - [EventStore](#core-eventstore-v1-EventStore)
  
- [core/notifier/v1/notifier.proto](#core_notifier_v1_notifier-proto)
    - [ConnectionRequest](#core-notifier-v1-ConnectionRequest)
    - [Heartbeat](#core-notifier-v1-Heartbeat)
    - [Notification](#core-notifier-v1-Notification)
  
    - [NotificationType](#core-notifier-v1-NotificationType)
  
    - [Notifier](#core-notifier-v1-Notifier)
  
- [core/registry/v1/registry.proto](#core_registry_v1_registry-proto)
    - [ConnectionRequest](#core-registry-v1-ConnectionRequest)
    - [ConnectionResponse](#core-registry-v1-ConnectionResponse)
    - [ConsumerRequest](#core-registry-v1-ConsumerRequest)
    - [ConsumerResponse](#core-registry-v1-ConsumerResponse)
    - [ProtocolRequest](#core-registry-v1-ProtocolRequest)
    - [ProtocolResponse](#core-registry-v1-ProtocolResponse)
    - [RegisterRequest](#core-registry-v1-RegisterRequest)
    - [RegisterResponse](#core-registry-v1-RegisterResponse)
  
    - [ServiceStatus](#core-registry-v1-ServiceStatus)
  
    - [Registry](#core-registry-v1-Registry)
  
- [generic/events/v1/commands.proto](#generic_events_v1_commands-proto)
- [generic/events/v1/notifications.proto](#generic_events_v1_notifications-proto)
    - [NotificationDelivered](#generic-events-v1-NotificationDelivered)
    - [NotificationDeliveryRequested](#generic-events-v1-NotificationDeliveryRequested)
  
    - [NotificationEventCode](#generic-events-v1-NotificationEventCode)
  
- [generic/events/v1/system.proto](#generic_events_v1_system-proto)
    - [SystemError](#generic-events-v1-SystemError)
  
    - [SystemErrorCode](#generic-events-v1-SystemErrorCode)
    - [SystemEventCode](#generic-events-v1-SystemEventCode)
  
- [todo/aggregates/v1/todo.proto](#todo_aggregates_v1_todo-proto)
    - [Todo](#todo-aggregates-v1-Todo)
  
    - [TodoStatus](#todo-aggregates-v1-TodoStatus)
  
- [generic/events/v1/todo.proto](#generic_events_v1_todo-proto)
    - [TodoCreated](#generic-events-v1-TodoCreated)
    - [TodoCreationFailed](#generic-events-v1-TodoCreationFailed)
  
    - [TodoEventCode](#generic-events-v1-TodoEventCode)
  
- [generic/events/v1/events.proto](#generic_events_v1_events-proto)
    - [EventType](#generic-events-v1-EventType)
  
    - [AggregateType](#generic-events-v1-AggregateType)
  
- [gogo.proto](#gogo-proto)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
    - [File-level Extensions](#gogo-proto-extensions)
  
- [todo/todo/v1/todo.proto](#todo_todo_v1_todo-proto)
    - [CreateTodoRequest](#todo-todo-v1-CreateTodoRequest)
    - [CreateTodoResponse](#todo-todo-v1-CreateTodoResponse)
  
    - [Todo](#todo-todo-v1-Todo)
  
- [types/type.proto](#types_type-proto)
    - [InetValue](#gorm-types-InetValue)
    - [JSONValue](#gorm-types-JSONValue)
    - [TimeOnly](#gorm-types-TimeOnly)
    - [UUID](#gorm-types-UUID)
    - [UUIDValue](#gorm-types-UUIDValue)
  
- [Scalar Value Types](#scalar-value-types)



<a name="options_gorm-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## options/gorm.proto



<a name="gorm-AutoServerOptions"></a>

### AutoServerOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| autogen | [bool](#bool) |  |  |
| txn_middleware | [bool](#bool) |  |  |
| with_tracing | [bool](#bool) |  |  |






<a name="gorm-BelongsToOptions"></a>

### BelongsToOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| foreignkey | [string](#string) |  |  |
| foreignkey_tag | [GormTag](#gorm-GormTag) |  |  |
| association_foreignkey | [string](#string) |  |  |
| association_autoupdate | [bool](#bool) |  |  |
| association_autocreate | [bool](#bool) |  |  |
| association_save_reference | [bool](#bool) |  |  |
| preload | [bool](#bool) |  |  |






<a name="gorm-ExtraField"></a>

### ExtraField



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| name | [string](#string) |  |  |
| tag | [GormTag](#gorm-GormTag) |  |  |
| package | [string](#string) |  |  |






<a name="gorm-GormFieldOptions"></a>

### GormFieldOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag | [GormTag](#gorm-GormTag) |  |  |
| drop | [bool](#bool) |  |  |
| has_one | [HasOneOptions](#gorm-HasOneOptions) |  |  |
| belongs_to | [BelongsToOptions](#gorm-BelongsToOptions) |  |  |
| has_many | [HasManyOptions](#gorm-HasManyOptions) |  |  |
| many_to_many | [ManyToManyOptions](#gorm-ManyToManyOptions) |  |  |
| reference_of | [string](#string) |  |  |






<a name="gorm-GormFileOptions"></a>

### GormFileOptions







<a name="gorm-GormMessageOptions"></a>

### GormMessageOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ormable | [bool](#bool) |  |  |
| include | [ExtraField](#gorm-ExtraField) | repeated |  |
| table | [string](#string) |  |  |
| multi_account | [bool](#bool) |  |  |






<a name="gorm-GormTag"></a>

### GormTag



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| column | [string](#string) |  |  |
| type | [string](#string) |  |  |
| size | [int32](#int32) |  |  |
| precision | [int32](#int32) |  |  |
| primary_key | [bool](#bool) |  |  |
| unique | [bool](#bool) |  |  |
| default | [string](#string) |  |  |
| not_null | [bool](#bool) |  |  |
| auto_increment | [bool](#bool) |  |  |
| index | [string](#string) |  |  |
| unique_index | [string](#string) |  |  |
| embedded | [bool](#bool) |  |  |
| embedded_prefix | [string](#string) |  |  |
| ignore | [bool](#bool) |  |  |
| foreignkey | [string](#string) |  |  |
| association_foreignkey | [string](#string) |  |  |
| many_to_many | [string](#string) |  |  |
| jointable_foreignkey | [string](#string) |  |  |
| association_jointable_foreignkey | [string](#string) |  |  |
| association_autoupdate | [bool](#bool) |  |  |
| association_autocreate | [bool](#bool) |  |  |
| association_save_reference | [bool](#bool) |  |  |
| preload | [bool](#bool) |  |  |






<a name="gorm-HasManyOptions"></a>

### HasManyOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| foreignkey | [string](#string) |  |  |
| foreignkey_tag | [GormTag](#gorm-GormTag) |  |  |
| association_foreignkey | [string](#string) |  |  |
| position_field | [string](#string) |  |  |
| position_field_tag | [GormTag](#gorm-GormTag) |  |  |
| association_autoupdate | [bool](#bool) |  |  |
| association_autocreate | [bool](#bool) |  |  |
| association_save_reference | [bool](#bool) |  |  |
| preload | [bool](#bool) |  |  |
| replace | [bool](#bool) |  |  |
| append | [bool](#bool) |  |  |
| clear | [bool](#bool) |  |  |






<a name="gorm-HasOneOptions"></a>

### HasOneOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| foreignkey | [string](#string) |  |  |
| foreignkey_tag | [GormTag](#gorm-GormTag) |  |  |
| association_foreignkey | [string](#string) |  |  |
| association_autoupdate | [bool](#bool) |  |  |
| association_autocreate | [bool](#bool) |  |  |
| association_save_reference | [bool](#bool) |  |  |
| preload | [bool](#bool) |  |  |
| replace | [bool](#bool) |  |  |
| append | [bool](#bool) |  |  |
| clear | [bool](#bool) |  |  |






<a name="gorm-ManyToManyOptions"></a>

### ManyToManyOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| jointable | [string](#string) |  |  |
| foreignkey | [string](#string) |  |  |
| jointable_foreignkey | [string](#string) |  |  |
| association_foreignkey | [string](#string) |  |  |
| association_jointable_foreignkey | [string](#string) |  |  |
| association_autoupdate | [bool](#bool) |  |  |
| association_autocreate | [bool](#bool) |  |  |
| association_save_reference | [bool](#bool) |  |  |
| preload | [bool](#bool) |  |  |
| replace | [bool](#bool) |  |  |
| append | [bool](#bool) |  |  |
| clear | [bool](#bool) |  |  |






<a name="gorm-MethodOptions"></a>

### MethodOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_type | [string](#string) |  |  |





 

 


<a name="options_gorm-proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| field | GormFieldOptions | .google.protobuf.FieldOptions | 52119 |  |
| file_opts | GormFileOptions | .google.protobuf.FileOptions | 52119 |  |
| opts | GormMessageOptions | .google.protobuf.MessageOptions | 52119 | ormable will cause orm code to be generated for this message/object |
| method | MethodOptions | .google.protobuf.MethodOptions | 52119 |  |
| server | AutoServerOptions | .google.protobuf.ServiceOptions | 52119 |  |

 

 



<a name="core_aggregates_v1_events-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## core/aggregates/v1/events.proto



<a name="core-aggregates-v1-Event"></a>

### Event



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the table primary key |
| received_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | the time the event create request was received |
| published_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | the time the event was been published to it&#39;s respective topic/queue |
| transaction_id | [string](#string) |  | uuid generated by the caller (e.g. commandhandler) and received by the client (e.g. react app) |
| aggregate_type | [string](#string) |  | the string identifier (enum value) of the aggregate type this event belongs to NOTE: this is simply a string to keep the eventstore from depending on changing types |
| event_type | [string](#string) |  | the string identifier of the event type NOTE: this is simply a string to keep the eventstore from depending on changing types |
| event_code | [string](#string) |  |  |
| aggregate_id | [string](#string) |  | the id of the aggregate that this event belongs to |
| event_data | [string](#string) |  | data representing the state of the system that this event encapsulates it is saved as a json string so that the eventstore can be completey agnostic to the the data structure of the system (and thus be a static service) this data MUST be able to be unmarshalled into a Proto message type |





 

 

 

 



<a name="core_aggregates_v1_registry-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## core/aggregates/v1/registry.proto



<a name="core-aggregates-v1-Consumer"></a>

### Consumer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the table primary key |
| routing_key | [string](#string) |  | routing key (this should match with an event type on the above aggregate) |
| kind | [ConsumerKind](#core-aggregates-v1-ConsumerKind) |  | consumer kind is whether the consumer is a queue (unicast - 1:N queue to consumer) or topic (multicast - 1:1 queue to consumer) |






<a name="core-aggregates-v1-Protocol"></a>

### Protocol



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the table primary key |
| kind | [ProtocolKind](#core-aggregates-v1-ProtocolKind) |  | the api kind |
| version | [string](#string) |  | the api version (e.g. &#34;v1&#34; or &#34;v2&#34;) |
| port | [int32](#int32) |  | the api port (e.g. 8080 or 8090) |






<a name="core-aggregates-v1-Registration"></a>

### Registration



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the table primary key |
| name | [string](#string) |  | NOTE: name &#43; version must be unique the service name |
| version | [string](#string) |  | the service version |
| description | [string](#string) |  | plain text description of the service |
| address | [string](#string) |  | the service address in the environment. e.g. &#34;http://localhost:8080&#34; when local or &#34;eventstore:8090&#34; when remote |
| status | [ServiceStatus](#core-aggregates-v1-ServiceStatus) |  | the service current status |
| protocols | [Protocol](#core-aggregates-v1-Protocol) | repeated | the protocols this service exposes |
| consumers | [Consumer](#core-aggregates-v1-Consumer) | repeated | the consumer configuration of the service |





 


<a name="core-aggregates-v1-ConsumerKind"></a>

### ConsumerKind


| Name | Number | Description |
| ---- | ------ | ----------- |
| CONSUMER_KIND_INVALID | 0 |  |
| CONSUMER_KIND_QUEUE | 1 |  |
| CONSUMER_KIND_TOPIC | 2 |  |



<a name="core-aggregates-v1-ProtocolKind"></a>

### ProtocolKind


| Name | Number | Description |
| ---- | ------ | ----------- |
| PROTOCOL_KIND_INVALID | 0 |  |
| PROTOCOL_KIND_GRPC | 1 |  |
| PROTOCOL_KIND_HTTP | 2 |  |



<a name="core-aggregates-v1-ServiceStatus"></a>

### ServiceStatus
deregister vs unregister reference: https://grammarhow.com/unregister-vs-deregister/

| Name | Number | Description |
| ---- | ------ | ----------- |
| SERVICE_STATUS_INVALID | 0 |  |
| SERVICE_STATUS_REGISTERED | 1 |  |
| SERVICE_STATUS_DEREGISTERED | 2 |  |
| SERVICE_STATUS_HEALTHY | 3 |  |
| SERVICE_STATUS_UNHEALTHY | 4 |  |


 

 

 



<a name="core_commandhandler_v1_commandhandler-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## core/commandhandler/v1/commandhandler.proto



<a name="core-commandhandler-v1-ApplyCommandRequest"></a>

### ApplyCommandRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| aggregate_type | [string](#string) |  | the string identifier (enum value) of the aggregate type this event belongs to NOTE: this is simply a string to keep the eventstore from depending on changing types |
| event_type | [string](#string) |  | map of all event types NOTE: this is simply a string to keep the eventstore from depending on changing types |
| event_code | [string](#string) |  | TODO: is this the way we want to route things? |
| aggregate_id | [string](#string) |  | the id of the aggregate that this command belongs to |
| command_data | [string](#string) |  | data representing the change to the state of the system that this command encapsulates it is saved as a json string so that the core services can be completey agnostic to the the data structure of the system (and thus be static services) this data MUST be able to be unmarshalled into a Proto message type |






<a name="core-commandhandler-v1-ApplyCommandResponse"></a>

### ApplyCommandResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| transaction_id | [string](#string) |  |  |





 

 

 


<a name="core-commandhandler-v1-CommandHandler"></a>

### CommandHandler


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Apply | [ApplyCommandRequest](#core-commandhandler-v1-ApplyCommandRequest) | [ApplyCommandResponse](#core-commandhandler-v1-ApplyCommandResponse) | this is ASYNCHRONOUS and will only return a transaction ID. the client should listen for a completed event on the notifier service |

 



<a name="core_eventstore_v1_eventstore-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## core/eventstore/v1/eventstore.proto



<a name="core-eventstore-v1-CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| aggregate_type | [string](#string) |  | the string identifier (enum value) of the aggregate type this event belongs to NOTE: this is simply a string to keep the eventstore from depending on changing types |
| event_type | [string](#string) |  | map of all event types NOTE: this is simply a string to keep the eventstore from depending on changing types |
| event_code | [string](#string) |  | TODO: is this the way we want to route things? |
| aggregate_id | [string](#string) |  | the id of the aggregate that this event belongs to |
| event_data | [string](#string) |  | data representing the state of the system that this event encapsulates it is saved as a json string so that the eventstore can be completey agnostic to the the data structure of the system (and thus be a static service) this data MUST be able to be unmarshalled into a Proto message type |






<a name="core-eventstore-v1-CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | event id |





 

 

 


<a name="core-eventstore-v1-EventStore"></a>

### EventStore


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#core-eventstore-v1-CreateRequest) | [CreateResponse](#core-eventstore-v1-CreateResponse) | Create a new event in the event store |

 



<a name="core_notifier_v1_notifier-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## core/notifier/v1/notifier.proto



<a name="core-notifier-v1-ConnectionRequest"></a>

### ConnectionRequest
ConnectionRequest - is used for a client, to connect and receive `Notifications` from processed events in the system.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actor_id | [string](#string) |  |  |
| client_id | [string](#string) |  | `client_id` is generated by the client and if not present the request will be denied. |






<a name="core-notifier-v1-Heartbeat"></a>

### Heartbeat
Heartbeat - A message sent on a consistent time interval maintaining the users session, and expiration deadline.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session_id | [string](#string) |  | users current session id |
| expiration_deadline | [int64](#int64) |  | Time when the `UsersSession` will expire if another `Heartbeat` is not sent before. |
| client_id | [string](#string) |  | client id |






<a name="core-notifier-v1-Notification"></a>

### Notification
A notification send to the web client


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| notification_type | [NotificationType](#core-notifier-v1-NotificationType) |  | notification type is the name of the type that is sent in the data value. This is recommended given the `web-client` may want to check the message type and perform specific actions. |
| data | [string](#string) |  | Data for each notification will be a message found in this package that is json encoded as a struct type. This main advantage to this is we will not need to redeploy the server, and client when a new notification is added, the `struct` type in javascript is an object |
| transaction_id | [string](#string) |  | UUID providing traceability all the way through the system to client |





 


<a name="core-notifier-v1-NotificationType"></a>

### NotificationType
NotificationType - A code that communicates to a integrated client the message type that is being sent.
TODO: check to see if you can add to this enum without having to redeploy the notifier service

| Name | Number | Description |
| ---- | ------ | ----------- |
| INVALID | 0 |  |
| HEARTBEAT | 1 | NOTE: add enum values here as needed |


 

 


<a name="core-notifier-v1-Notifier"></a>

### Notifier
`Connect` - A rpc method that can push a notification to a connected `web-client`.
A map of `session_id`&#39;s and `user_id`&#39;s is maintained to know how a notification
should be sent. Using this interface requires that the web-client has issued
a `StartUserSession` command against the `CommandHandler`

For a connection to be established with the `notification` service, the users `session_id` and `user_id` are required 

Given these two values the service is required to maintain a mapping between events, channels, and
users. This will allow the service to push messages to the correct clients.

`ConnectInternal` - Maintains a map of all `MinionD` connections so it can receive a notification, and push
`MinionD` specific messages

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Connect | [ConnectionRequest](#core-notifier-v1-ConnectionRequest) | [Notification](#core-notifier-v1-Notification) stream | Connect a web-client |

 



<a name="core_registry_v1_registry-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## core/registry/v1/registry.proto



<a name="core-registry-v1-ConnectionRequest"></a>

### ConnectionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| version | [string](#string) |  | this is the major version of the service to connect to (v1, v2, etc.) |
| type | [core.aggregates.v1.ProtocolKind](#core-aggregates-v1-ProtocolKind) |  |  |






<a name="core-registry-v1-ConnectionResponse"></a>

### ConnectionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| port | [int32](#int32) |  |  |
| status | [ServiceStatus](#core-registry-v1-ServiceStatus) |  |  |






<a name="core-registry-v1-ConsumerRequest"></a>

### ConsumerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order | [int32](#int32) |  |  |
| kind | [core.aggregates.v1.ConsumerKind](#core-aggregates-v1-ConsumerKind) |  |  |
| aggregate_type | [string](#string) |  |  |
| event_type | [string](#string) |  |  |
| event_code | [string](#string) |  |  |






<a name="core-registry-v1-ConsumerResponse"></a>

### ConsumerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order | [int32](#int32) |  |  |
| kind | [core.aggregates.v1.ConsumerKind](#core-aggregates-v1-ConsumerKind) |  |  |
| routing_key | [string](#string) |  |  |
| exchange | [string](#string) |  |  |
| queue_name | [string](#string) |  |  |






<a name="core-registry-v1-ProtocolRequest"></a>

### ProtocolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order | [int32](#int32) |  |  |
| kind | [core.aggregates.v1.ProtocolKind](#core-aggregates-v1-ProtocolKind) |  |  |






<a name="core-registry-v1-ProtocolResponse"></a>

### ProtocolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order | [int32](#int32) |  |  |
| kind | [core.aggregates.v1.ProtocolKind](#core-aggregates-v1-ProtocolKind) |  |  |
| port | [int32](#int32) |  |  |






<a name="core-registry-v1-RegisterRequest"></a>

### RegisterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| domain | [string](#string) |  | this needs to match the k8s namespace for remote deployments |
| version | [string](#string) |  | address = domain.name:port

should be semver |
| description | [string](#string) |  |  |
| protocols | [ProtocolRequest](#core-registry-v1-ProtocolRequest) | repeated |  |
| consumers | [ConsumerRequest](#core-registry-v1-ConsumerRequest) | repeated |  |






<a name="core-registry-v1-RegisterResponse"></a>

### RegisterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| protocols | [ProtocolResponse](#core-registry-v1-ProtocolResponse) | repeated |  |
| consumers | [ConsumerResponse](#core-registry-v1-ConsumerResponse) | repeated |  |





 


<a name="core-registry-v1-ServiceStatus"></a>

### ServiceStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| SERVICE_STATUS_INVALID | 0 |  |
| SERVICE_STATUS_HEALTHY | 1 |  |
| SERVICE_STATUS_UNHEALTHY | 2 |  |


 

 


<a name="core-registry-v1-Registry"></a>

### Registry


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Register | [RegisterRequest](#core-registry-v1-RegisterRequest) | [RegisterResponse](#core-registry-v1-RegisterResponse) | Register registers a new microservice with the registry. |
| Connection | [ConnectionRequest](#core-registry-v1-ConnectionRequest) | [ConnectionResponse](#core-registry-v1-ConnectionResponse) | Connection returns the connection info for a microservice. Example: Service A wishes to call Service B, A calls registry.Connection(B) which returns the connection info for B. If B is not available or not registered, an error will be returned. |

 



<a name="generic_events_v1_commands-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## generic/events/v1/commands.proto


 

 

 

 



<a name="generic_events_v1_notifications-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## generic/events/v1/notifications.proto



<a name="generic-events-v1-NotificationDelivered"></a>

### NotificationDelivered







<a name="generic-events-v1-NotificationDeliveryRequested"></a>

### NotificationDeliveryRequested
NotificationDeliveryRequested is an event used to send a message to an actor connected the `notifier` service.
`Multicast`, is the default delivery type and a `actor_id` is required. If `Unicast` is desired (i.e Sending a notification
to only one client) Then a `client_id` should also be provided.
   {
     &#34;actor_id&#34;: &#34;cffbbfa8-1a7e-4b64-af2e-345654b37aa7&#34;,
     &#34;client_id&#34;: &#34;07925e22-3eee-4931-aea9-19fc621fd825&#34;,
     &#34;notification&#34;: &#34;&lt;NOTIFICATION_MESSAGE&gt;&#34;
   }


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actor_id | [string](#string) |  | `actor_id` is the identifier of a actor, a message should be sent to (think phone number). For example, if it&#39;s a user connected to the notifier service with the web-client then `actor_id` is equal to the `user_id` of the user. Using the `actor_id` instead of a specific `user_id` field allows for many differnt types of client connections to the notifier and gives the system a common way to send data to those connected clients whitout having to change the underlying datastructure when adding new clients. |
| client_id | [string](#string) |  | optional, specify only if `unicast` to one client is desired. if empty, `multicast` to all clients associated with the `actor_id`` will be used. |
| notification | [core.notifier.v1.Notification](#core-notifier-v1-Notification) |  | notification is the data payload that will be sent the client. |





 


<a name="generic-events-v1-NotificationEventCode"></a>

### NotificationEventCode


| Name | Number | Description |
| ---- | ------ | ----------- |
| NOTIFICATION_EVENT_CODE_INVALID | 0 |  |
| NOTIFICATION_EVENT_CODE_DELIVERY_REQUESTED | 1 |  |
| NOTIFICATION_EVENT_CODE_DELIVERED | 2 |  |


 

 

 



<a name="generic_events_v1_system-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## generic/events/v1/system.proto



<a name="generic-events-v1-SystemError"></a>

### SystemError



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [SystemErrorCode](#generic-events-v1-SystemErrorCode) |  |  |





 


<a name="generic-events-v1-SystemErrorCode"></a>

### SystemErrorCode


| Name | Number | Description |
| ---- | ------ | ----------- |
| SYSTEM_ERROR_CODE_INVALID | 0 |  |
| SYSTEM_ERROR_CODE_FAILED_EVENT_PUBLISH | 1 |  |
| SYSTEM_ERROR_CODE_FAILED_EVENT_SAVED | 2 |  |
| SYSTEM_ERROR_CODE_FAILED_EVENT_FORWARD | 3 |  |
| SYSTEM_ERROR_CODE_MALFORMED_EVENT_DATA | 4 |  |



<a name="generic-events-v1-SystemEventCode"></a>

### SystemEventCode


| Name | Number | Description |
| ---- | ------ | ----------- |
| SYSTEM_EVENT_CODE_INVALID | 0 |  |
| SYSTEM_EVENT_CODE_ERROR | 1 |  |


 

 

 



<a name="todo_aggregates_v1_todo-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## todo/aggregates/v1/todo.proto



<a name="todo-aggregates-v1-Todo"></a>

### Todo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | values |
| title | [string](#string) |  |  |
| description | [string](#string) |  |  |
| status | [TodoStatus](#todo-aggregates-v1-TodoStatus) |  | example timestamp attribute google.protobuf.Timestamp scheduled_time = 5 [(gorm.field).tag = {type: &#34;timestamp&#34;}]; |





 


<a name="todo-aggregates-v1-TodoStatus"></a>

### TodoStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| TODO_STATUS_INVALID | 0 |  |
| COMPLETE | 1 |  |
| INCOMPLETE | 2 |  |


 

 

 



<a name="generic_events_v1_todo-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## generic/events/v1/todo.proto



<a name="generic-events-v1-TodoCreated"></a>

### TodoCreated



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| todo | [todo.aggregates.v1.Todo](#todo-aggregates-v1-Todo) |  |  |






<a name="generic-events-v1-TodoCreationFailed"></a>

### TodoCreationFailed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| todo | [todo.aggregates.v1.Todo](#todo-aggregates-v1-Todo) |  |  |
| error | [string](#string) |  |  |





 


<a name="generic-events-v1-TodoEventCode"></a>

### TodoEventCode


| Name | Number | Description |
| ---- | ------ | ----------- |
| TODO_EVENT_CODE_INVALID | 0 |  |
| TODO_EVENT_CODE_CREATED | 1 |  |
| TODO_EVENT_CODE_DELETED | 2 |  |


 

 

 



<a name="generic_events_v1_events-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## generic/events/v1/events.proto



<a name="generic-events-v1-EventType"></a>

### EventType
map of all event types, add to it as more event types are added to the application


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| system_code | [SystemEventCode](#generic-events-v1-SystemEventCode) |  |  |
| notification_code | [NotificationEventCode](#generic-events-v1-NotificationEventCode) |  |  |
| todo_event_code | [TodoEventCode](#generic-events-v1-TodoEventCode) |  |  |





 


<a name="generic-events-v1-AggregateType"></a>

### AggregateType
map of all aggregate types, add to it as more aggregates are added to the application

| Name | Number | Description |
| ---- | ------ | ----------- |
| AGGREGATE_TYPE_INVALID | 0 |  |
| AGGREGATE_TYPE_SYSTEM | 1 |  |
| AGGREGATE_TYPE_NOTIFICATION | 2 |  |
| AGGREGATE_TYPE_TODO | 3 |  |


 

 

 



<a name="gogo-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gogo.proto


 

 


<a name="gogo-proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| enum_customname | string | .google.protobuf.EnumOptions | 62023 |  |
| enum_stringer | bool | .google.protobuf.EnumOptions | 62022 |  |
| enumdecl | bool | .google.protobuf.EnumOptions | 62024 |  |
| goproto_enum_prefix | bool | .google.protobuf.EnumOptions | 62001 |  |
| goproto_enum_stringer | bool | .google.protobuf.EnumOptions | 62021 |  |
| enumvalue_customname | string | .google.protobuf.EnumValueOptions | 66001 |  |
| castkey | string | .google.protobuf.FieldOptions | 65008 |  |
| casttype | string | .google.protobuf.FieldOptions | 65007 |  |
| castvalue | string | .google.protobuf.FieldOptions | 65009 |  |
| customname | string | .google.protobuf.FieldOptions | 65004 |  |
| customtype | string | .google.protobuf.FieldOptions | 65003 |  |
| embed | bool | .google.protobuf.FieldOptions | 65002 |  |
| jsontag | string | .google.protobuf.FieldOptions | 65005 |  |
| moretags | string | .google.protobuf.FieldOptions | 65006 |  |
| nullable | bool | .google.protobuf.FieldOptions | 65001 |  |
| stdduration | bool | .google.protobuf.FieldOptions | 65011 |  |
| stdtime | bool | .google.protobuf.FieldOptions | 65010 |  |
| wktpointer | bool | .google.protobuf.FieldOptions | 65012 |  |
| benchgen_all | bool | .google.protobuf.FileOptions | 63016 |  |
| compare_all | bool | .google.protobuf.FileOptions | 63029 |  |
| description_all | bool | .google.protobuf.FileOptions | 63014 |  |
| enum_stringer_all | bool | .google.protobuf.FileOptions | 63022 |  |
| enumdecl_all | bool | .google.protobuf.FileOptions | 63031 |  |
| equal_all | bool | .google.protobuf.FileOptions | 63013 |  |
| face_all | bool | .google.protobuf.FileOptions | 63005 |  |
| goproto_enum_prefix_all | bool | .google.protobuf.FileOptions | 63002 |  |
| goproto_enum_stringer_all | bool | .google.protobuf.FileOptions | 63021 |  |
| goproto_extensions_map_all | bool | .google.protobuf.FileOptions | 63025 |  |
| goproto_getters_all | bool | .google.protobuf.FileOptions | 63001 |  |
| goproto_registration | bool | .google.protobuf.FileOptions | 63032 |  |
| goproto_sizecache_all | bool | .google.protobuf.FileOptions | 63034 |  |
| goproto_stringer_all | bool | .google.protobuf.FileOptions | 63003 |  |
| goproto_unkeyed_all | bool | .google.protobuf.FileOptions | 63035 |  |
| goproto_unrecognized_all | bool | .google.protobuf.FileOptions | 63026 |  |
| gostring_all | bool | .google.protobuf.FileOptions | 63006 |  |
| marshaler_all | bool | .google.protobuf.FileOptions | 63017 |  |
| messagename_all | bool | .google.protobuf.FileOptions | 63033 |  |
| onlyone_all | bool | .google.protobuf.FileOptions | 63009 |  |
| populate_all | bool | .google.protobuf.FileOptions | 63007 |  |
| protosizer_all | bool | .google.protobuf.FileOptions | 63028 |  |
| sizer_all | bool | .google.protobuf.FileOptions | 63020 |  |
| stable_marshaler_all | bool | .google.protobuf.FileOptions | 63019 |  |
| stringer_all | bool | .google.protobuf.FileOptions | 63008 |  |
| testgen_all | bool | .google.protobuf.FileOptions | 63015 |  |
| typedecl_all | bool | .google.protobuf.FileOptions | 63030 |  |
| unmarshaler_all | bool | .google.protobuf.FileOptions | 63018 |  |
| unsafe_marshaler_all | bool | .google.protobuf.FileOptions | 63023 |  |
| unsafe_unmarshaler_all | bool | .google.protobuf.FileOptions | 63024 |  |
| verbose_equal_all | bool | .google.protobuf.FileOptions | 63004 |  |
| benchgen | bool | .google.protobuf.MessageOptions | 64016 |  |
| compare | bool | .google.protobuf.MessageOptions | 64029 |  |
| description | bool | .google.protobuf.MessageOptions | 64014 |  |
| equal | bool | .google.protobuf.MessageOptions | 64013 |  |
| face | bool | .google.protobuf.MessageOptions | 64005 |  |
| goproto_extensions_map | bool | .google.protobuf.MessageOptions | 64025 |  |
| goproto_getters | bool | .google.protobuf.MessageOptions | 64001 |  |
| goproto_sizecache | bool | .google.protobuf.MessageOptions | 64034 |  |
| goproto_stringer | bool | .google.protobuf.MessageOptions | 64003 |  |
| goproto_unkeyed | bool | .google.protobuf.MessageOptions | 64035 |  |
| goproto_unrecognized | bool | .google.protobuf.MessageOptions | 64026 |  |
| gostring | bool | .google.protobuf.MessageOptions | 64006 |  |
| marshaler | bool | .google.protobuf.MessageOptions | 64017 |  |
| messagename | bool | .google.protobuf.MessageOptions | 64033 |  |
| onlyone | bool | .google.protobuf.MessageOptions | 64009 |  |
| populate | bool | .google.protobuf.MessageOptions | 64007 |  |
| protosizer | bool | .google.protobuf.MessageOptions | 64028 |  |
| sizer | bool | .google.protobuf.MessageOptions | 64020 |  |
| stable_marshaler | bool | .google.protobuf.MessageOptions | 64019 |  |
| stringer | bool | .google.protobuf.MessageOptions | 67008 |  |
| testgen | bool | .google.protobuf.MessageOptions | 64015 |  |
| typedecl | bool | .google.protobuf.MessageOptions | 64030 |  |
| unmarshaler | bool | .google.protobuf.MessageOptions | 64018 |  |
| unsafe_marshaler | bool | .google.protobuf.MessageOptions | 64023 |  |
| unsafe_unmarshaler | bool | .google.protobuf.MessageOptions | 64024 |  |
| verbose_equal | bool | .google.protobuf.MessageOptions | 64004 |  |
| gogoproto_import | bool | .google.protobuf.FileOptions | 63027 |  |

 

 



<a name="todo_todo_v1_todo-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## todo/todo/v1/todo.proto



<a name="todo-todo-v1-CreateTodoRequest"></a>

### CreateTodoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payload | [todo.aggregates.v1.Todo](#todo-aggregates-v1-Todo) |  | WHY DO WE NEED A LEADING . HERE? |






<a name="todo-todo-v1-CreateTodoResponse"></a>

### CreateTodoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [todo.aggregates.v1.Todo](#todo-aggregates-v1-Todo) |  | WHY DO WE NEED A LEADING . HERE? |





 

 

 


<a name="todo-todo-v1-Todo"></a>

### Todo


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateTodoRequest](#todo-todo-v1-CreateTodoRequest) | [CreateTodoResponse](#todo-todo-v1-CreateTodoResponse) |  |

 



<a name="types_type-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## types/type.proto



<a name="gorm-types-InetValue"></a>

### InetValue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  |  |






<a name="gorm-types-JSONValue"></a>

### JSONValue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  |  |






<a name="gorm-types-TimeOnly"></a>

### TimeOnly



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [uint32](#uint32) |  |  |






<a name="gorm-types-UUID"></a>

### UUID



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  |  |






<a name="gorm-types-UUIDValue"></a>

### UUIDValue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  |  |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

