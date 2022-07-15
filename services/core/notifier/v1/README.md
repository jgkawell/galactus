# Notifier
A service that is used to publish notifications to connected user clients. The high level objective is to
persist a long running connection with a users client and publish notifications to the user.

[Example Implementation](https://lucid.app/lucidchart/invitations/accept/inv_b2c1e17f-bf8a-4533-93c3-05518e61b47a?viewport_loc=844%2C-757%2C2832%2C2033%2C0_0)

# Model
If a consumer, or service would like to send a notification to a user. The below `Event` should be published
to the event store
```proto
// Notification Event
message NotificationDeliveryRequested {
  string user_id = 1;
  atlas.notifier.v1.Notification notification = 3;
}
```

The notification body is simple. First, we have a notification type that currently is a string that the web client
can use to process the delivered notification. Second, the data value is the message/notification that will be sent
to the client.
```proto
// A notification sent to the web client
message Notification {
  // notification type is the name of the type that is sent in the data value. This is recommended
  // given the `web-client` may want to check the message type and perform specific actions. Generally
  // this would be an enum, but given that fact we don't want redeploy this service very often it's best
  // to keep the type as a string.
  string notification_type = 1;

  // Data for each notification will be a message found in this package that
  // is json encoded as a struct type. This main advantage to this is we will not need to
  // redeploy the server, and client when a new notification is added, the `struct` type
  // in javascript is an object
  // google.protobuf.Struct data = 2;
  string data = 2;
}
```

# Interface
The focus on this service is to manage users client connections, and deliver notifications to the connected
client. It's not to create chat rooms, or manage streams of data around what notifications a user should receive.
It's designed to have a simple interface, and send messages to the client.

```proto
// Notifier is a service that can push a notification to a connected `web-client`.
// A map of `session_id`'s and `user_id`'s is maintained to know how a notification
// should be sent. Using this interface requires that the web-client has issued
// a `StartUserSession` command against the `CommandHandler`
//
// For a connection to be established with the `notification` service, the users `session_id` and `user_id` are required 
// 
// Given these two values the service is required to maintain a mapping between events, channels, and
// users. This will allow the service to push messages to the correct clients.
service Notifier {
  rpc Connect(ConnectionRequest) returns (stream Notification) {}
}
```
