/**
 * @fileoverview gRPC-Web generated client stub for todo.todo.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var todo_aggregates_v1_todo_pb = require('../../../todo/aggregates/v1/todo_pb.js')
const proto = {};
proto.todo = {};
proto.todo.todo = {};
proto.todo.todo.v1 = require('./todo_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.todo.todo.v1.TodoClient =
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
proto.todo.todo.v1.TodoPromiseClient =
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
 *   !proto.todo.todo.v1.CreateTodoRequest,
 *   !proto.todo.todo.v1.CreateTodoResponse>}
 */
const methodDescriptor_Todo_Create = new grpc.web.MethodDescriptor(
  '/todo.todo.v1.Todo/Create',
  grpc.web.MethodType.UNARY,
  proto.todo.todo.v1.CreateTodoRequest,
  proto.todo.todo.v1.CreateTodoResponse,
  /**
   * @param {!proto.todo.todo.v1.CreateTodoRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.todo.todo.v1.CreateTodoResponse.deserializeBinary
);


/**
 * @param {!proto.todo.todo.v1.CreateTodoRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.todo.todo.v1.CreateTodoResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.todo.todo.v1.CreateTodoResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.todo.todo.v1.TodoClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/todo.todo.v1.Todo/Create',
      request,
      metadata || {},
      methodDescriptor_Todo_Create,
      callback);
};


/**
 * @param {!proto.todo.todo.v1.CreateTodoRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.todo.todo.v1.CreateTodoResponse>}
 *     Promise that resolves to the response
 */
proto.todo.todo.v1.TodoPromiseClient.prototype.create =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/todo.todo.v1.Todo/Create',
      request,
      metadata || {},
      methodDescriptor_Todo_Create);
};


module.exports = proto.todo.todo.v1;

