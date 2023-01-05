// source: core/commandhandler/v1/commandhandler.proto
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
var validate_validate_pb = require('../../../validate/validate_pb.js');
goog.object.extend(proto, validate_validate_pb);
goog.exportSymbol('proto.core.commandhandler.v1.ApplyCommandRequest', null, global);
goog.exportSymbol('proto.core.commandhandler.v1.ApplyCommandResponse', null, global);
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
proto.core.commandhandler.v1.ApplyCommandRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.core.commandhandler.v1.ApplyCommandRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.core.commandhandler.v1.ApplyCommandRequest.displayName = 'proto.core.commandhandler.v1.ApplyCommandRequest';
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
proto.core.commandhandler.v1.ApplyCommandResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.core.commandhandler.v1.ApplyCommandResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.core.commandhandler.v1.ApplyCommandResponse.displayName = 'proto.core.commandhandler.v1.ApplyCommandResponse';
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
proto.core.commandhandler.v1.ApplyCommandRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.core.commandhandler.v1.ApplyCommandRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.core.commandhandler.v1.ApplyCommandRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.commandhandler.v1.ApplyCommandRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    aggregateType: jspb.Message.getFieldWithDefault(msg, 1, ""),
    eventType: jspb.Message.getFieldWithDefault(msg, 2, ""),
    eventCode: jspb.Message.getFieldWithDefault(msg, 3, ""),
    aggregateId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    commandData: jspb.Message.getFieldWithDefault(msg, 5, "")
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
 * @return {!proto.core.commandhandler.v1.ApplyCommandRequest}
 */
proto.core.commandhandler.v1.ApplyCommandRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.core.commandhandler.v1.ApplyCommandRequest;
  return proto.core.commandhandler.v1.ApplyCommandRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.core.commandhandler.v1.ApplyCommandRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.core.commandhandler.v1.ApplyCommandRequest}
 */
proto.core.commandhandler.v1.ApplyCommandRequest.deserializeBinaryFromReader = function(msg, reader) {
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
      msg.setCommandData(value);
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
proto.core.commandhandler.v1.ApplyCommandRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.core.commandhandler.v1.ApplyCommandRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.core.commandhandler.v1.ApplyCommandRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.commandhandler.v1.ApplyCommandRequest.serializeBinaryToWriter = function(message, writer) {
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
  f = message.getCommandData();
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
proto.core.commandhandler.v1.ApplyCommandRequest.prototype.getAggregateType = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.commandhandler.v1.ApplyCommandRequest} returns this
 */
proto.core.commandhandler.v1.ApplyCommandRequest.prototype.setAggregateType = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string event_type = 2;
 * @return {string}
 */
proto.core.commandhandler.v1.ApplyCommandRequest.prototype.getEventType = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.commandhandler.v1.ApplyCommandRequest} returns this
 */
proto.core.commandhandler.v1.ApplyCommandRequest.prototype.setEventType = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string event_code = 3;
 * @return {string}
 */
proto.core.commandhandler.v1.ApplyCommandRequest.prototype.getEventCode = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.commandhandler.v1.ApplyCommandRequest} returns this
 */
proto.core.commandhandler.v1.ApplyCommandRequest.prototype.setEventCode = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string aggregate_id = 4;
 * @return {string}
 */
proto.core.commandhandler.v1.ApplyCommandRequest.prototype.getAggregateId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.commandhandler.v1.ApplyCommandRequest} returns this
 */
proto.core.commandhandler.v1.ApplyCommandRequest.prototype.setAggregateId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string command_data = 5;
 * @return {string}
 */
proto.core.commandhandler.v1.ApplyCommandRequest.prototype.getCommandData = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.commandhandler.v1.ApplyCommandRequest} returns this
 */
proto.core.commandhandler.v1.ApplyCommandRequest.prototype.setCommandData = function(value) {
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
proto.core.commandhandler.v1.ApplyCommandResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.core.commandhandler.v1.ApplyCommandResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.core.commandhandler.v1.ApplyCommandResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.commandhandler.v1.ApplyCommandResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    transactionId: jspb.Message.getFieldWithDefault(msg, 2, "")
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
 * @return {!proto.core.commandhandler.v1.ApplyCommandResponse}
 */
proto.core.commandhandler.v1.ApplyCommandResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.core.commandhandler.v1.ApplyCommandResponse;
  return proto.core.commandhandler.v1.ApplyCommandResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.core.commandhandler.v1.ApplyCommandResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.core.commandhandler.v1.ApplyCommandResponse}
 */
proto.core.commandhandler.v1.ApplyCommandResponse.deserializeBinaryFromReader = function(msg, reader) {
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
      var value = /** @type {string} */ (reader.readString());
      msg.setTransactionId(value);
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
proto.core.commandhandler.v1.ApplyCommandResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.core.commandhandler.v1.ApplyCommandResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.core.commandhandler.v1.ApplyCommandResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.commandhandler.v1.ApplyCommandResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getTransactionId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.core.commandhandler.v1.ApplyCommandResponse.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.commandhandler.v1.ApplyCommandResponse} returns this
 */
proto.core.commandhandler.v1.ApplyCommandResponse.prototype.setId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string transaction_id = 2;
 * @return {string}
 */
proto.core.commandhandler.v1.ApplyCommandResponse.prototype.getTransactionId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.commandhandler.v1.ApplyCommandResponse} returns this
 */
proto.core.commandhandler.v1.ApplyCommandResponse.prototype.setTransactionId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


goog.object.extend(exports, proto.core.commandhandler.v1);
