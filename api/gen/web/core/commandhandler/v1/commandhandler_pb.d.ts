import * as jspb from 'google-protobuf'


import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class ApplyCommandRequest extends jspb.Message {
  getAggregateType(): string;
  setAggregateType(value: string): ApplyCommandRequest;

  getEventType(): string;
  setEventType(value: string): ApplyCommandRequest;

  getEventCode(): string;
  setEventCode(value: string): ApplyCommandRequest;

  getAggregateId(): string;
  setAggregateId(value: string): ApplyCommandRequest;

  getCommandData(): string;
  setCommandData(value: string): ApplyCommandRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ApplyCommandRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ApplyCommandRequest): ApplyCommandRequest.AsObject;
  static serializeBinaryToWriter(message: ApplyCommandRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ApplyCommandRequest;
  static deserializeBinaryFromReader(message: ApplyCommandRequest, reader: jspb.BinaryReader): ApplyCommandRequest;
}

export namespace ApplyCommandRequest {
  export type AsObject = {
    aggregateType: string,
    eventType: string,
    eventCode: string,
    aggregateId: string,
    commandData: string,
  }
}

export class ApplyCommandResponse extends jspb.Message {
  getId(): string;
  setId(value: string): ApplyCommandResponse;

  getTransactionId(): string;
  setTransactionId(value: string): ApplyCommandResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ApplyCommandResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ApplyCommandResponse): ApplyCommandResponse.AsObject;
  static serializeBinaryToWriter(message: ApplyCommandResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ApplyCommandResponse;
  static deserializeBinaryFromReader(message: ApplyCommandResponse, reader: jspb.BinaryReader): ApplyCommandResponse;
}

export namespace ApplyCommandResponse {
  export type AsObject = {
    id: string,
    transactionId: string,
  }
}

