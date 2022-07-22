import * as jspb from 'google-protobuf'

import * as todo_aggregates_v1_todo_pb from '../../../todo/aggregates/v1/todo_pb';


export class TodoCreated extends jspb.Message {
  getTodo(): todo_aggregates_v1_todo_pb.Todo | undefined;
  setTodo(value?: todo_aggregates_v1_todo_pb.Todo): TodoCreated;
  hasTodo(): boolean;
  clearTodo(): TodoCreated;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TodoCreated.AsObject;
  static toObject(includeInstance: boolean, msg: TodoCreated): TodoCreated.AsObject;
  static serializeBinaryToWriter(message: TodoCreated, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TodoCreated;
  static deserializeBinaryFromReader(message: TodoCreated, reader: jspb.BinaryReader): TodoCreated;
}

export namespace TodoCreated {
  export type AsObject = {
    todo?: todo_aggregates_v1_todo_pb.Todo.AsObject,
  }
}

export class TodoCreationFailed extends jspb.Message {
  getTodo(): todo_aggregates_v1_todo_pb.Todo | undefined;
  setTodo(value?: todo_aggregates_v1_todo_pb.Todo): TodoCreationFailed;
  hasTodo(): boolean;
  clearTodo(): TodoCreationFailed;

  getError(): string;
  setError(value: string): TodoCreationFailed;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TodoCreationFailed.AsObject;
  static toObject(includeInstance: boolean, msg: TodoCreationFailed): TodoCreationFailed.AsObject;
  static serializeBinaryToWriter(message: TodoCreationFailed, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TodoCreationFailed;
  static deserializeBinaryFromReader(message: TodoCreationFailed, reader: jspb.BinaryReader): TodoCreationFailed;
}

export namespace TodoCreationFailed {
  export type AsObject = {
    todo?: todo_aggregates_v1_todo_pb.Todo.AsObject,
    error: string,
  }
}

export enum TodoEventCode { 
  TODO_EVENT_CODE_INVALID = 0,
  TODO_EVENT_CODE_CREATED = 1,
  TODO_EVENT_CODE_DELETED = 2,
}
