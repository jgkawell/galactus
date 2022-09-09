import * as jspb from 'google-protobuf'


import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class CreateRequest extends jspb.Message {
  getAggregateType(): string;
  setAggregateType(value: string): CreateRequest;

  getEventType(): string;
  setEventType(value: string): CreateRequest;

  getEventCode(): string;
  setEventCode(value: string): CreateRequest;

  getAggregateId(): string;
  setAggregateId(value: string): CreateRequest;

  getEventData(): string;
  setEventData(value: string): CreateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateRequest): CreateRequest.AsObject;
  static serializeBinaryToWriter(message: CreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateRequest;
  static deserializeBinaryFromReader(message: CreateRequest, reader: jspb.BinaryReader): CreateRequest;
}

export namespace CreateRequest {
  export type AsObject = {
    aggregateType: string,
    eventType: string,
    eventCode: string,
    aggregateId: string,
    eventData: string,
  }
}

export class CreateResponse extends jspb.Message {
  getId(): string;
  setId(value: string): CreateResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateResponse): CreateResponse.AsObject;
  static serializeBinaryToWriter(message: CreateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateResponse;
  static deserializeBinaryFromReader(message: CreateResponse, reader: jspb.BinaryReader): CreateResponse;
}

export namespace CreateResponse {
  export type AsObject = {
    id: string,
  }
}

