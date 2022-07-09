// source: core/eventstore/v1/eventstore.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
goog.exportSymbol('proto.core.eventstore.v1.CreateEventRequest', null, global);
goog.exportSymbol('proto.core.eventstore.v1.CreateEventResponse', null, global);
goog.exportSymbol('proto.core.eventstore.v1.Event', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.core.eventstore.v1.CreateEventRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.core.eventstore.v1.CreateEventRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.core.eventstore.v1.CreateEventRequest.displayName = 'proto.core.eventstore.v1.CreateEventRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.core.eventstore.v1.CreateEventResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.core.eventstore.v1.CreateEventResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.core.eventstore.v1.CreateEventResponse.displayName = 'proto.core.eventstore.v1.CreateEventResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.core.eventstore.v1.Event = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.core.eventstore.v1.Event, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.core.eventstore.v1.Event.displayName = 'proto.core.eventstore.v1.Event';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.core.eventstore.v1.CreateEventRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.core.eventstore.v1.CreateEventRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.core.eventstore.v1.CreateEventRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.eventstore.v1.CreateEventRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    event: (f = msg.getEvent()) && proto.core.eventstore.v1.Event.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.core.eventstore.v1.CreateEventRequest}
 */
proto.core.eventstore.v1.CreateEventRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.core.eventstore.v1.CreateEventRequest;
  return proto.core.eventstore.v1.CreateEventRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.core.eventstore.v1.CreateEventRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.core.eventstore.v1.CreateEventRequest}
 */
proto.core.eventstore.v1.CreateEventRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.core.eventstore.v1.Event;
      reader.readMessage(value,proto.core.eventstore.v1.Event.deserializeBinaryFromReader);
      msg.setEvent(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.core.eventstore.v1.CreateEventRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.core.eventstore.v1.CreateEventRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.core.eventstore.v1.CreateEventRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.eventstore.v1.CreateEventRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getEvent();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.core.eventstore.v1.Event.serializeBinaryToWriter
    );
  }
};


/**
 * optional Event event = 1;
 * @return {?proto.core.eventstore.v1.Event}
 */
proto.core.eventstore.v1.CreateEventRequest.prototype.getEvent = function() {
  return /** @type{?proto.core.eventstore.v1.Event} */ (
    jspb.Message.getWrapperField(this, proto.core.eventstore.v1.Event, 1));
};


/**
 * @param {?proto.core.eventstore.v1.Event|undefined} value
 * @return {!proto.core.eventstore.v1.CreateEventRequest} returns this
*/
proto.core.eventstore.v1.CreateEventRequest.prototype.setEvent = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.core.eventstore.v1.CreateEventRequest} returns this
 */
proto.core.eventstore.v1.CreateEventRequest.prototype.clearEvent = function() {
  return this.setEvent(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.core.eventstore.v1.CreateEventRequest.prototype.hasEvent = function() {
  return jspb.Message.getField(this, 1) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.core.eventstore.v1.CreateEventResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.core.eventstore.v1.CreateEventResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.core.eventstore.v1.CreateEventResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.eventstore.v1.CreateEventResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    isPublished: jspb.Message.getBooleanFieldWithDefault(msg, 2, false)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.core.eventstore.v1.CreateEventResponse}
 */
proto.core.eventstore.v1.CreateEventResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.core.eventstore.v1.CreateEventResponse;
  return proto.core.eventstore.v1.CreateEventResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.core.eventstore.v1.CreateEventResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.core.eventstore.v1.CreateEventResponse}
 */
proto.core.eventstore.v1.CreateEventResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIsPublished(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.core.eventstore.v1.CreateEventResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.core.eventstore.v1.CreateEventResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.core.eventstore.v1.CreateEventResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.eventstore.v1.CreateEventResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getIsPublished();
  if (f) {
    writer.writeBool(
      2,
      f
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.core.eventstore.v1.CreateEventResponse.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.eventstore.v1.CreateEventResponse} returns this
 */
proto.core.eventstore.v1.CreateEventResponse.prototype.setId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional bool is_published = 2;
 * @return {boolean}
 */
proto.core.eventstore.v1.CreateEventResponse.prototype.getIsPublished = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 2, false));
};


/**
 * @param {boolean} value
 * @return {!proto.core.eventstore.v1.CreateEventResponse} returns this
 */
proto.core.eventstore.v1.CreateEventResponse.prototype.setIsPublished = function(value) {
  return jspb.Message.setProto3BooleanField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.core.eventstore.v1.Event.prototype.toObject = function(opt_includeInstance) {
  return proto.core.eventstore.v1.Event.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.core.eventstore.v1.Event} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.eventstore.v1.Event.toObject = function(includeInstance, msg) {
  var f, obj = {
    eventId: jspb.Message.getFieldWithDefault(msg, 1, ""),
    receivedDate: (f = msg.getReceivedDate()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    publishedDate: (f = msg.getPublishedDate()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    transactionId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    publish: jspb.Message.getBooleanFieldWithDefault(msg, 5, false),
    eventType: jspb.Message.getFieldWithDefault(msg, 17, 0),
    aggregateType: jspb.Message.getFieldWithDefault(msg, 18, 0),
    aggregateId: jspb.Message.getFieldWithDefault(msg, 19, ""),
    eventData: jspb.Message.getFieldWithDefault(msg, 20, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.core.eventstore.v1.Event}
 */
proto.core.eventstore.v1.Event.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.core.eventstore.v1.Event;
  return proto.core.eventstore.v1.Event.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.core.eventstore.v1.Event} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.core.eventstore.v1.Event}
 */
proto.core.eventstore.v1.Event.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setEventId(value);
      break;
    case 2:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setReceivedDate(value);
      break;
    case 3:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setPublishedDate(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setTransactionId(value);
      break;
    case 5:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setPublish(value);
      break;
    case 17:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setEventType(value);
      break;
    case 18:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setAggregateType(value);
      break;
    case 19:
      var value = /** @type {string} */ (reader.readString());
      msg.setAggregateId(value);
      break;
    case 20:
      var value = /** @type {string} */ (reader.readString());
      msg.setEventData(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.core.eventstore.v1.Event.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.core.eventstore.v1.Event.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.core.eventstore.v1.Event} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.eventstore.v1.Event.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getEventId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getReceivedDate();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getPublishedDate();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getTransactionId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getPublish();
  if (f) {
    writer.writeBool(
      5,
      f
    );
  }
  f = message.getEventType();
  if (f !== 0) {
    writer.writeInt64(
      17,
      f
    );
  }
  f = message.getAggregateType();
  if (f !== 0) {
    writer.writeInt64(
      18,
      f
    );
  }
  f = message.getAggregateId();
  if (f.length > 0) {
    writer.writeString(
      19,
      f
    );
  }
  f = message.getEventData();
  if (f.length > 0) {
    writer.writeString(
      20,
      f
    );
  }
};


/**
 * optional string event_id = 1;
 * @return {string}
 */
proto.core.eventstore.v1.Event.prototype.getEventId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.eventstore.v1.Event} returns this
 */
proto.core.eventstore.v1.Event.prototype.setEventId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional google.protobuf.Timestamp received_date = 2;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.core.eventstore.v1.Event.prototype.getReceivedDate = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 2));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.core.eventstore.v1.Event} returns this
*/
proto.core.eventstore.v1.Event.prototype.setReceivedDate = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.core.eventstore.v1.Event} returns this
 */
proto.core.eventstore.v1.Event.prototype.clearReceivedDate = function() {
  return this.setReceivedDate(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.core.eventstore.v1.Event.prototype.hasReceivedDate = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional google.protobuf.Timestamp published_date = 3;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.core.eventstore.v1.Event.prototype.getPublishedDate = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 3));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.core.eventstore.v1.Event} returns this
*/
proto.core.eventstore.v1.Event.prototype.setPublishedDate = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.core.eventstore.v1.Event} returns this
 */
proto.core.eventstore.v1.Event.prototype.clearPublishedDate = function() {
  return this.setPublishedDate(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.core.eventstore.v1.Event.prototype.hasPublishedDate = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional string transaction_id = 4;
 * @return {string}
 */
proto.core.eventstore.v1.Event.prototype.getTransactionId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.eventstore.v1.Event} returns this
 */
proto.core.eventstore.v1.Event.prototype.setTransactionId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional bool publish = 5;
 * @return {boolean}
 */
proto.core.eventstore.v1.Event.prototype.getPublish = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 5, false));
};


/**
 * @param {boolean} value
 * @return {!proto.core.eventstore.v1.Event} returns this
 */
proto.core.eventstore.v1.Event.prototype.setPublish = function(value) {
  return jspb.Message.setProto3BooleanField(this, 5, value);
};


/**
 * optional int64 event_type = 17;
 * @return {number}
 */
proto.core.eventstore.v1.Event.prototype.getEventType = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 17, 0));
};


/**
 * @param {number} value
 * @return {!proto.core.eventstore.v1.Event} returns this
 */
proto.core.eventstore.v1.Event.prototype.setEventType = function(value) {
  return jspb.Message.setProto3IntField(this, 17, value);
};


/**
 * optional int64 aggregate_type = 18;
 * @return {number}
 */
proto.core.eventstore.v1.Event.prototype.getAggregateType = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 18, 0));
};


/**
 * @param {number} value
 * @return {!proto.core.eventstore.v1.Event} returns this
 */
proto.core.eventstore.v1.Event.prototype.setAggregateType = function(value) {
  return jspb.Message.setProto3IntField(this, 18, value);
};


/**
 * optional string aggregate_id = 19;
 * @return {string}
 */
proto.core.eventstore.v1.Event.prototype.getAggregateId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 19, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.eventstore.v1.Event} returns this
 */
proto.core.eventstore.v1.Event.prototype.setAggregateId = function(value) {
  return jspb.Message.setProto3StringField(this, 19, value);
};


/**
 * optional string event_data = 20;
 * @return {string}
 */
proto.core.eventstore.v1.Event.prototype.getEventData = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 20, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.eventstore.v1.Event} returns this
 */
proto.core.eventstore.v1.Event.prototype.setEventData = function(value) {
  return jspb.Message.setProto3StringField(this, 20, value);
};


goog.object.extend(exports, proto.core.eventstore.v1);
