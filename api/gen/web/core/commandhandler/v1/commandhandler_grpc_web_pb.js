/**
 * @fileoverview gRPC-Web generated client stub for core.commandhandler.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')

var validate_validate_pb = require('../../../validate/validate_pb.js')
const proto = {};
proto.core = {};
proto.core.commandhandler = {};
proto.core.commandhandler.v1 = require('./commandhandler_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.core.commandhandler.v1.CommandHandlerClient =
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
proto.core.commandhandler.v1.CommandHandlerPromiseClient =
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
 *   !proto.core.commandhandler.v1.ApplyCommandRequest,
 *   !proto.core.commandhandler.v1.ApplyCommandResponse>}
 */
const methodDescriptor_CommandHandler_Apply = new grpc.web.MethodDescriptor(
  '/core.commandhandler.v1.CommandHandler/Apply',
  grpc.web.MethodType.UNARY,
  proto.core.commandhandler.v1.ApplyCommandRequest,
  proto.core.commandhandler.v1.ApplyCommandResponse,
  /**
   * @param {!proto.core.commandhandler.v1.ApplyCommandRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.core.commandhandler.v1.ApplyCommandResponse.deserializeBinary
);


/**
 * @param {!proto.core.commandhandler.v1.ApplyCommandRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.core.commandhandler.v1.ApplyCommandResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.core.commandhandler.v1.ApplyCommandResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.core.commandhandler.v1.CommandHandlerClient.prototype.apply =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/core.commandhandler.v1.CommandHandler/Apply',
      request,
      metadata || {},
      methodDescriptor_CommandHandler_Apply,
      callback);
};


/**
 * @param {!proto.core.commandhandler.v1.ApplyCommandRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.core.commandhandler.v1.ApplyCommandResponse>}
 *     Promise that resolves to the response
 */
proto.core.commandhandler.v1.CommandHandlerPromiseClient.prototype.apply =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/core.commandhandler.v1.CommandHandler/Apply',
      request,
      metadata || {},
      methodDescriptor_CommandHandler_Apply);
};


module.exports = proto.core.commandhandler.v1;

