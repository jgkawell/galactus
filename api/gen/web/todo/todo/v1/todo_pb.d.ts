import * as jspb from 'google-protobuf'

import * as todo_aggregates_v1_todo_pb from '../../../todo/aggregates/v1/todo_pb';


export class CreateTodoRequest extends jspb.Message {
  getPayload(): todo_aggregates_v1_todo_pb.Todo | undefined;
  setPayload(value?: todo_aggregates_v1_todo_pb.Todo): CreateTodoRequest;
  hasPayload(): boolean;
  clearPayload(): CreateTodoRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateTodoRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateTodoRequest): CreateTodoRequest.AsObject;
  static serializeBinaryToWriter(message: CreateTodoRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateTodoRequest;
  static deserializeBinaryFromReader(message: CreateTodoRequest, reader: jspb.BinaryReader): CreateTodoRequest;
}

export namespace CreateTodoRequest {
  export type AsObject = {
    payload?: todo_aggregates_v1_todo_pb.Todo.AsObject,
  }
}

export class CreateTodoResponse extends jspb.Message {
  getResult(): todo_aggregates_v1_todo_pb.Todo | undefined;
  setResult(value?: todo_aggregates_v1_todo_pb.Todo): CreateTodoResponse;
  hasResult(): boolean;
  clearResult(): CreateTodoResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateTodoResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateTodoResponse): CreateTodoResponse.AsObject;
  static serializeBinaryToWriter(message: CreateTodoResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateTodoResponse;
  static deserializeBinaryFromReader(message: CreateTodoResponse, reader: jspb.BinaryReader): CreateTodoResponse;
}

export namespace CreateTodoResponse {
  export type AsObject = {
    result?: todo_aggregates_v1_todo_pb.Todo.AsObject,
  }
}

