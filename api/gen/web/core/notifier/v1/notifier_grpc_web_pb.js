/**
 * @fileoverview gRPC-Web generated client stub for core.notifier.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.core = {};
proto.core.notifier = {};
proto.core.notifier.v1 = require('./notifier_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.core.notifier.v1.NotifierClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.core.notifier.v1.NotifierPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.core.notifier.v1.ConnectionRequest,
 *   !proto.core.notifier.v1.Notification>}
 */
const methodDescriptor_Notifier_Connect = new grpc.web.MethodDescriptor(
  '/core.notifier.v1.Notifier/Connect',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.core.notifier.v1.ConnectionRequest,
  proto.core.notifier.v1.Notification,
  /**
   * @param {!proto.core.notifier.v1.ConnectionRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.core.notifier.v1.Notification.deserializeBinary
);


/**
 * @param {!proto.core.notifier.v1.ConnectionRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.core.notifier.v1.Notification>}
 *     The XHR Node Readable Stream
 */
proto.core.notifier.v1.NotifierClient.prototype.connect =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/core.notifier.v1.Notifier/Connect',
      request,
      metadata || {},
      methodDescriptor_Notifier_Connect);
};


/**
 * @param {!proto.core.notifier.v1.ConnectionRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.core.notifier.v1.Notification>}
 *     The XHR Node Readable Stream
 */
proto.core.notifier.v1.NotifierPromiseClient.prototype.connect =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/core.notifier.v1.Notifier/Connect',
      request,
      metadata || {},
      methodDescriptor_Notifier_Connect);
};


module.exports = proto.core.notifier.v1;

