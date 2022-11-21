import * as jspb from 'google-protobuf'


import * as options_gorm_pb from '../../../options/gorm_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class Todo extends jspb.Message {
  getId(): string;
  setId(value: string): Todo;

  getTitle(): string;
  setTitle(value: string): Todo;

  getDescription(): string;
  setDescription(value: string): Todo;

  getStatus(): TodoStatus;
  setStatus(value: TodoStatus): Todo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Todo.AsObject;
  static toObject(includeInstance: boolean, msg: Todo): Todo.AsObject;
  static serializeBinaryToWriter(message: Todo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Todo;
  static deserializeBinaryFromReader(message: Todo, reader: jspb.BinaryReader): Todo;
}

export namespace Todo {
  export type AsObject = {
    id: string,
    title: string,
    description: string,
    status: TodoStatus,
  }
}

export enum TodoStatus { 
  TODO_STATUS_INVALID = 0,
  COMPLETE = 1,
  INCOMPLETE = 2,
}
