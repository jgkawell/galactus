/**
 * @fileoverview gRPC-Web generated client stub for core.eventstore.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');



var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')
const proto = {};
proto.core = {};
proto.core.eventstore = {};
proto.core.eventstore.v1 = require('./eventstore_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.core.eventstore.v1.EventStoreClient =
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
proto.core.eventstore.v1.EventStorePromiseClient =
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
 *   !proto.core.eventstore.v1.CreateRequest,
 *   !proto.core.eventstore.v1.CreateResponse>}
 */
const methodDescriptor_EventStore_Create = new grpc.web.MethodDescriptor(
  '/core.eventstore.v1.EventStore/Create',
  grpc.web.MethodType.UNARY,
  proto.core.eventstore.v1.CreateRequest,
  proto.core.eventstore.v1.CreateResponse,
  /**
   * @param {!proto.core.eventstore.v1.CreateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.core.eventstore.v1.CreateResponse.deserializeBinary
);


/**
 * @param {!proto.core.eventstore.v1.CreateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.core.eventstore.v1.CreateResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.core.eventstore.v1.CreateResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.core.eventstore.v1.EventStoreClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/core.eventstore.v1.EventStore/Create',
      request,
      metadata || {},
      methodDescriptor_EventStore_Create,
      callback);
};


/**
 * @param {!proto.core.eventstore.v1.CreateRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.core.eventstore.v1.CreateResponse>}
 *     Promise that resolves to the response
 */
proto.core.eventstore.v1.EventStorePromiseClient.prototype.create =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/core.eventstore.v1.EventStore/Create',
      request,
      metadata || {},
      methodDescriptor_EventStore_Create);
};


module.exports = proto.core.eventstore.v1;

