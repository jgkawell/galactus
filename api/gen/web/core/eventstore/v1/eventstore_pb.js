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
goog.exportSymbol('proto.core.eventstore.v1.CreateRequest', null, global);
goog.exportSymbol('proto.core.eventstore.v1.CreateResponse', null, global);
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
proto.core.eventstore.v1.CreateRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.core.eventstore.v1.CreateRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.core.eventstore.v1.CreateRequest.displayName = 'proto.core.eventstore.v1.CreateRequest';
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
proto.core.eventstore.v1.CreateResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.core.eventstore.v1.CreateResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.core.eventstore.v1.CreateResponse.displayName = 'proto.core.eventstore.v1.CreateResponse';
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
proto.core.eventstore.v1.CreateRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.core.eventstore.v1.CreateRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.core.eventstore.v1.CreateRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.eventstore.v1.CreateRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    aggregateType: jspb.Message.getFieldWithDefault(msg, 1, ""),
    eventType: jspb.Message.getFieldWithDefault(msg, 2, ""),
    eventCode: jspb.Message.getFieldWithDefault(msg, 3, ""),
    aggregateId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    eventData: jspb.Message.getFieldWithDefault(msg, 5, "")
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
 * @return {!proto.core.eventstore.v1.CreateRequest}
 */
proto.core.eventstore.v1.CreateRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.core.eventstore.v1.CreateRequest;
  return proto.core.eventstore.v1.CreateRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.core.eventstore.v1.CreateRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.core.eventstore.v1.CreateRequest}
 */
proto.core.eventstore.v1.CreateRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setAggregateType(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setEventType(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setEventCode(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setAggregateId(value);
      break;
    case 5:
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
proto.core.eventstore.v1.CreateRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.core.eventstore.v1.CreateRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.core.eventstore.v1.CreateRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.eventstore.v1.CreateRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAggregateType();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getEventType();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getEventCode();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getAggregateId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getEventData();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
};


/**
 * optional string aggregate_type = 1;
 * @return {string}
 */
proto.core.eventstore.v1.CreateRequest.prototype.getAggregateType = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.eventstore.v1.CreateRequest} returns this
 */
proto.core.eventstore.v1.CreateRequest.prototype.setAggregateType = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string event_type = 2;
 * @return {string}
 */
proto.core.eventstore.v1.CreateRequest.prototype.getEventType = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.eventstore.v1.CreateRequest} returns this
 */
proto.core.eventstore.v1.CreateRequest.prototype.setEventType = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string event_code = 3;
 * @return {string}
 */
proto.core.eventstore.v1.CreateRequest.prototype.getEventCode = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.eventstore.v1.CreateRequest} returns this
 */
proto.core.eventstore.v1.CreateRequest.prototype.setEventCode = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string aggregate_id = 4;
 * @return {string}
 */
proto.core.eventstore.v1.CreateRequest.prototype.getAggregateId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.eventstore.v1.CreateRequest} returns this
 */
proto.core.eventstore.v1.CreateRequest.prototype.setAggregateId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string event_data = 5;
 * @return {string}
 */
proto.core.eventstore.v1.CreateRequest.prototype.getEventData = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.eventstore.v1.CreateRequest} returns this
 */
proto.core.eventstore.v1.CreateRequest.prototype.setEventData = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
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
proto.core.eventstore.v1.CreateResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.core.eventstore.v1.CreateResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.core.eventstore.v1.CreateResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.eventstore.v1.CreateResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, "")
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
 * @return {!proto.core.eventstore.v1.CreateResponse}
 */
proto.core.eventstore.v1.CreateResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.core.eventstore.v1.CreateResponse;
  return proto.core.eventstore.v1.CreateResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.core.eventstore.v1.CreateResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.core.eventstore.v1.CreateResponse}
 */
proto.core.eventstore.v1.CreateResponse.deserializeBinaryFromReader = function(msg, reader) {
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
proto.core.eventstore.v1.CreateResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.core.eventstore.v1.CreateResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.core.eventstore.v1.CreateResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.eventstore.v1.CreateResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.core.eventstore.v1.CreateResponse.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.eventstore.v1.CreateResponse} returns this
 */
proto.core.eventstore.v1.CreateResponse.prototype.setId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


goog.object.extend(exports, proto.core.eventstore.v1);
