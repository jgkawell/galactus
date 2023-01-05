// source: core/aggregates/v1/registry.proto
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
var options_gorm_pb = require('../../../options/gorm_pb.js');
goog.object.extend(proto, options_gorm_pb);
var validate_validate_pb = require('../../../validate/validate_pb.js');
goog.object.extend(proto, validate_validate_pb);
goog.exportSymbol('proto.core.aggregates.v1.Consumer', null, global);
goog.exportSymbol('proto.core.aggregates.v1.ConsumerKind', null, global);
goog.exportSymbol('proto.core.aggregates.v1.ProtocolKind', null, global);
goog.exportSymbol('proto.core.aggregates.v1.Registration', null, global);
goog.exportSymbol('proto.core.aggregates.v1.Route', null, global);
goog.exportSymbol('proto.core.aggregates.v1.ServiceStatus', null, global);
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
proto.core.aggregates.v1.Registration = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.core.aggregates.v1.Registration.repeatedFields_, null);
};
goog.inherits(proto.core.aggregates.v1.Registration, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.core.aggregates.v1.Registration.displayName = 'proto.core.aggregates.v1.Registration';
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
proto.core.aggregates.v1.Route = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.core.aggregates.v1.Route, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.core.aggregates.v1.Route.displayName = 'proto.core.aggregates.v1.Route';
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
proto.core.aggregates.v1.Consumer = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.core.aggregates.v1.Consumer, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.core.aggregates.v1.Consumer.displayName = 'proto.core.aggregates.v1.Consumer';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.core.aggregates.v1.Registration.repeatedFields_ = [16,17];



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
proto.core.aggregates.v1.Registration.prototype.toObject = function(opt_includeInstance) {
  return proto.core.aggregates.v1.Registration.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.core.aggregates.v1.Registration} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.aggregates.v1.Registration.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    domain: jspb.Message.getFieldWithDefault(msg, 2, ""),
    name: jspb.Message.getFieldWithDefault(msg, 3, ""),
    version: jspb.Message.getFieldWithDefault(msg, 4, ""),
    description: jspb.Message.getFieldWithDefault(msg, 5, ""),
    status: jspb.Message.getFieldWithDefault(msg, 6, 0),
    routesList: jspb.Message.toObjectList(msg.getRoutesList(),
    proto.core.aggregates.v1.Route.toObject, includeInstance),
    consumersList: jspb.Message.toObjectList(msg.getConsumersList(),
    proto.core.aggregates.v1.Consumer.toObject, includeInstance)
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
 * @return {!proto.core.aggregates.v1.Registration}
 */
proto.core.aggregates.v1.Registration.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.core.aggregates.v1.Registration;
  return proto.core.aggregates.v1.Registration.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.core.aggregates.v1.Registration} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.core.aggregates.v1.Registration}
 */
proto.core.aggregates.v1.Registration.deserializeBinaryFromReader = function(msg, reader) {
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
      msg.setDomain(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setVersion(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setDescription(value);
      break;
    case 6:
      var value = /** @type {!proto.core.aggregates.v1.ServiceStatus} */ (reader.readEnum());
      msg.setStatus(value);
      break;
    case 16:
      var value = new proto.core.aggregates.v1.Route;
      reader.readMessage(value,proto.core.aggregates.v1.Route.deserializeBinaryFromReader);
      msg.addRoutes(value);
      break;
    case 17:
      var value = new proto.core.aggregates.v1.Consumer;
      reader.readMessage(value,proto.core.aggregates.v1.Consumer.deserializeBinaryFromReader);
      msg.addConsumers(value);
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
proto.core.aggregates.v1.Registration.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.core.aggregates.v1.Registration.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.core.aggregates.v1.Registration} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.aggregates.v1.Registration.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getDomain();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVersion();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getDescription();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getStatus();
  if (f !== 0.0) {
    writer.writeEnum(
      6,
      f
    );
  }
  f = message.getRoutesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      16,
      f,
      proto.core.aggregates.v1.Route.serializeBinaryToWriter
    );
  }
  f = message.getConsumersList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      17,
      f,
      proto.core.aggregates.v1.Consumer.serializeBinaryToWriter
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.core.aggregates.v1.Registration.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.aggregates.v1.Registration} returns this
 */
proto.core.aggregates.v1.Registration.prototype.setId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string domain = 2;
 * @return {string}
 */
proto.core.aggregates.v1.Registration.prototype.getDomain = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.aggregates.v1.Registration} returns this
 */
proto.core.aggregates.v1.Registration.prototype.setDomain = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string name = 3;
 * @return {string}
 */
proto.core.aggregates.v1.Registration.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.aggregates.v1.Registration} returns this
 */
proto.core.aggregates.v1.Registration.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string version = 4;
 * @return {string}
 */
proto.core.aggregates.v1.Registration.prototype.getVersion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.aggregates.v1.Registration} returns this
 */
proto.core.aggregates.v1.Registration.prototype.setVersion = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string description = 5;
 * @return {string}
 */
proto.core.aggregates.v1.Registration.prototype.getDescription = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.aggregates.v1.Registration} returns this
 */
proto.core.aggregates.v1.Registration.prototype.setDescription = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional ServiceStatus status = 6;
 * @return {!proto.core.aggregates.v1.ServiceStatus}
 */
proto.core.aggregates.v1.Registration.prototype.getStatus = function() {
  return /** @type {!proto.core.aggregates.v1.ServiceStatus} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/**
 * @param {!proto.core.aggregates.v1.ServiceStatus} value
 * @return {!proto.core.aggregates.v1.Registration} returns this
 */
proto.core.aggregates.v1.Registration.prototype.setStatus = function(value) {
  return jspb.Message.setProto3EnumField(this, 6, value);
};


/**
 * repeated Route routes = 16;
 * @return {!Array<!proto.core.aggregates.v1.Route>}
 */
proto.core.aggregates.v1.Registration.prototype.getRoutesList = function() {
  return /** @type{!Array<!proto.core.aggregates.v1.Route>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.core.aggregates.v1.Route, 16));
};


/**
 * @param {!Array<!proto.core.aggregates.v1.Route>} value
 * @return {!proto.core.aggregates.v1.Registration} returns this
*/
proto.core.aggregates.v1.Registration.prototype.setRoutesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 16, value);
};


/**
 * @param {!proto.core.aggregates.v1.Route=} opt_value
 * @param {number=} opt_index
 * @return {!proto.core.aggregates.v1.Route}
 */
proto.core.aggregates.v1.Registration.prototype.addRoutes = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 16, opt_value, proto.core.aggregates.v1.Route, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.core.aggregates.v1.Registration} returns this
 */
proto.core.aggregates.v1.Registration.prototype.clearRoutesList = function() {
  return this.setRoutesList([]);
};


/**
 * repeated Consumer consumers = 17;
 * @return {!Array<!proto.core.aggregates.v1.Consumer>}
 */
proto.core.aggregates.v1.Registration.prototype.getConsumersList = function() {
  return /** @type{!Array<!proto.core.aggregates.v1.Consumer>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.core.aggregates.v1.Consumer, 17));
};


/**
 * @param {!Array<!proto.core.aggregates.v1.Consumer>} value
 * @return {!proto.core.aggregates.v1.Registration} returns this
*/
proto.core.aggregates.v1.Registration.prototype.setConsumersList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 17, value);
};


/**
 * @param {!proto.core.aggregates.v1.Consumer=} opt_value
 * @param {number=} opt_index
 * @return {!proto.core.aggregates.v1.Consumer}
 */
proto.core.aggregates.v1.Registration.prototype.addConsumers = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 17, opt_value, proto.core.aggregates.v1.Consumer, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.core.aggregates.v1.Registration} returns this
 */
proto.core.aggregates.v1.Registration.prototype.clearConsumersList = function() {
  return this.setConsumersList([]);
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
proto.core.aggregates.v1.Route.prototype.toObject = function(opt_includeInstance) {
  return proto.core.aggregates.v1.Route.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.core.aggregates.v1.Route} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.aggregates.v1.Route.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    path: jspb.Message.getFieldWithDefault(msg, 2, ""),
    host: jspb.Message.getFieldWithDefault(msg, 3, ""),
    port: jspb.Message.getFieldWithDefault(msg, 4, 0),
    kind: jspb.Message.getFieldWithDefault(msg, 5, 0)
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
 * @return {!proto.core.aggregates.v1.Route}
 */
proto.core.aggregates.v1.Route.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.core.aggregates.v1.Route;
  return proto.core.aggregates.v1.Route.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.core.aggregates.v1.Route} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.core.aggregates.v1.Route}
 */
proto.core.aggregates.v1.Route.deserializeBinaryFromReader = function(msg, reader) {
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
      msg.setPath(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setHost(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPort(value);
      break;
    case 5:
      var value = /** @type {!proto.core.aggregates.v1.ProtocolKind} */ (reader.readEnum());
      msg.setKind(value);
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
proto.core.aggregates.v1.Route.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.core.aggregates.v1.Route.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.core.aggregates.v1.Route} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.aggregates.v1.Route.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getPath();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getHost();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getPort();
  if (f !== 0) {
    writer.writeInt32(
      4,
      f
    );
  }
  f = message.getKind();
  if (f !== 0.0) {
    writer.writeEnum(
      5,
      f
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.core.aggregates.v1.Route.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.aggregates.v1.Route} returns this
 */
proto.core.aggregates.v1.Route.prototype.setId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string path = 2;
 * @return {string}
 */
proto.core.aggregates.v1.Route.prototype.getPath = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.aggregates.v1.Route} returns this
 */
proto.core.aggregates.v1.Route.prototype.setPath = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string host = 3;
 * @return {string}
 */
proto.core.aggregates.v1.Route.prototype.getHost = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.aggregates.v1.Route} returns this
 */
proto.core.aggregates.v1.Route.prototype.setHost = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional int32 port = 4;
 * @return {number}
 */
proto.core.aggregates.v1.Route.prototype.getPort = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.core.aggregates.v1.Route} returns this
 */
proto.core.aggregates.v1.Route.prototype.setPort = function(value) {
  return jspb.Message.setProto3IntField(this, 4, value);
};


/**
 * optional ProtocolKind kind = 5;
 * @return {!proto.core.aggregates.v1.ProtocolKind}
 */
proto.core.aggregates.v1.Route.prototype.getKind = function() {
  return /** @type {!proto.core.aggregates.v1.ProtocolKind} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {!proto.core.aggregates.v1.ProtocolKind} value
 * @return {!proto.core.aggregates.v1.Route} returns this
 */
proto.core.aggregates.v1.Route.prototype.setKind = function(value) {
  return jspb.Message.setProto3EnumField(this, 5, value);
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
proto.core.aggregates.v1.Consumer.prototype.toObject = function(opt_includeInstance) {
  return proto.core.aggregates.v1.Consumer.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.core.aggregates.v1.Consumer} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.aggregates.v1.Consumer.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    aggregateType: jspb.Message.getFieldWithDefault(msg, 2, ""),
    eventType: jspb.Message.getFieldWithDefault(msg, 3, ""),
    eventCode: jspb.Message.getFieldWithDefault(msg, 4, ""),
    kind: jspb.Message.getFieldWithDefault(msg, 5, 0)
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
 * @return {!proto.core.aggregates.v1.Consumer}
 */
proto.core.aggregates.v1.Consumer.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.core.aggregates.v1.Consumer;
  return proto.core.aggregates.v1.Consumer.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.core.aggregates.v1.Consumer} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.core.aggregates.v1.Consumer}
 */
proto.core.aggregates.v1.Consumer.deserializeBinaryFromReader = function(msg, reader) {
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
      msg.setAggregateType(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setEventType(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setEventCode(value);
      break;
    case 5:
      var value = /** @type {!proto.core.aggregates.v1.ConsumerKind} */ (reader.readEnum());
      msg.setKind(value);
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
proto.core.aggregates.v1.Consumer.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.core.aggregates.v1.Consumer.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.core.aggregates.v1.Consumer} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.aggregates.v1.Consumer.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAggregateType();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getEventType();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getEventCode();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getKind();
  if (f !== 0.0) {
    writer.writeEnum(
      5,
      f
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.core.aggregates.v1.Consumer.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.aggregates.v1.Consumer} returns this
 */
proto.core.aggregates.v1.Consumer.prototype.setId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string aggregate_type = 2;
 * @return {string}
 */
proto.core.aggregates.v1.Consumer.prototype.getAggregateType = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.aggregates.v1.Consumer} returns this
 */
proto.core.aggregates.v1.Consumer.prototype.setAggregateType = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string event_type = 3;
 * @return {string}
 */
proto.core.aggregates.v1.Consumer.prototype.getEventType = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.aggregates.v1.Consumer} returns this
 */
proto.core.aggregates.v1.Consumer.prototype.setEventType = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string event_code = 4;
 * @return {string}
 */
proto.core.aggregates.v1.Consumer.prototype.getEventCode = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.core.aggregates.v1.Consumer} returns this
 */
proto.core.aggregates.v1.Consumer.prototype.setEventCode = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional ConsumerKind kind = 5;
 * @return {!proto.core.aggregates.v1.ConsumerKind}
 */
proto.core.aggregates.v1.Consumer.prototype.getKind = function() {
  return /** @type {!proto.core.aggregates.v1.ConsumerKind} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {!proto.core.aggregates.v1.ConsumerKind} value
 * @return {!proto.core.aggregates.v1.Consumer} returns this
 */
proto.core.aggregates.v1.Consumer.prototype.setKind = function(value) {
  return jspb.Message.setProto3EnumField(this, 5, value);
};


/**
 * @enum {number}
 */
proto.core.aggregates.v1.ServiceStatus = {
  SERVICE_STATUS_UNSPECIFIED: 0,
  SERVICE_STATUS_REGISTERED: 1,
  SERVICE_STATUS_DEREGISTERED: 2,
  SERVICE_STATUS_HEALTHY: 3,
  SERVICE_STATUS_UNHEALTHY: 4
};

/**
 * @enum {number}
 */
proto.core.aggregates.v1.ProtocolKind = {
  PROTOCOL_KIND_UNSPECIFIED: 0,
  PROTOCOL_KIND_GRPC: 1,
  PROTOCOL_KIND_HTTP: 2
};

/**
 * @enum {number}
 */
proto.core.aggregates.v1.ConsumerKind = {
  CONSUMER_KIND_UNSPECIFIED: 0,
  CONSUMER_KIND_QUEUE: 1,
  CONSUMER_KIND_TOPIC: 2
};

goog.object.extend(exports, proto.core.aggregates.v1);
