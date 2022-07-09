/**
 * @fileoverview gRPC-Web generated client stub for core.registry.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var core_aggregates_v1_registry_pb = require('../../../core/aggregates/v1/registry_pb.js')
const proto = {};
proto.core = {};
proto.core.registry = {};
proto.core.registry.v1 = require('./registry_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.core.registry.v1.RegistryClient =
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
proto.core.registry.v1.RegistryPromiseClient =
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
 *   !proto.core.registry.v1.RegisterRequest,
 *   !proto.core.registry.v1.RegisterResponse>}
 */
const methodDescriptor_Registry_Register = new grpc.web.MethodDescriptor(
  '/core.registry.v1.Registry/Register',
  grpc.web.MethodType.UNARY,
  proto.core.registry.v1.RegisterRequest,
  proto.core.registry.v1.RegisterResponse,
  /**
   * @param {!proto.core.registry.v1.RegisterRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.core.registry.v1.RegisterResponse.deserializeBinary
);


/**
 * @param {!proto.core.registry.v1.RegisterRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.core.registry.v1.RegisterResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.core.registry.v1.RegisterResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.core.registry.v1.RegistryClient.prototype.register =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/core.registry.v1.Registry/Register',
      request,
      metadata || {},
      methodDescriptor_Registry_Register,
      callback);
};


/**
 * @param {!proto.core.registry.v1.RegisterRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.core.registry.v1.RegisterResponse>}
 *     Promise that resolves to the response
 */
proto.core.registry.v1.RegistryPromiseClient.prototype.register =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/core.registry.v1.Registry/Register',
      request,
      metadata || {},
      methodDescriptor_Registry_Register);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.core.registry.v1.ConnectionRequest,
 *   !proto.core.registry.v1.ConnectionResponse>}
 */
const methodDescriptor_Registry_Connection = new grpc.web.MethodDescriptor(
  '/core.registry.v1.Registry/Connection',
  grpc.web.MethodType.UNARY,
  proto.core.registry.v1.ConnectionRequest,
  proto.core.registry.v1.ConnectionResponse,
  /**
   * @param {!proto.core.registry.v1.ConnectionRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.core.registry.v1.ConnectionResponse.deserializeBinary
);


/**
 * @param {!proto.core.registry.v1.ConnectionRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.core.registry.v1.ConnectionResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.core.registry.v1.ConnectionResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.core.registry.v1.RegistryClient.prototype.connection =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/core.registry.v1.Registry/Connection',
      request,
      metadata || {},
      methodDescriptor_Registry_Connection,
      callback);
};


/**
 * @param {!proto.core.registry.v1.ConnectionRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.core.registry.v1.ConnectionResponse>}
 *     Promise that resolves to the response
 */
proto.core.registry.v1.RegistryPromiseClient.prototype.connection =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/core.registry.v1.Registry/Connection',
      request,
      metadata || {},
      methodDescriptor_Registry_Connection);
};


module.exports = proto.core.registry.v1;

