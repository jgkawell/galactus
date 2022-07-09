import * as jspb from 'google-protobuf'


import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class CreateEventRequest extends jspb.Message {
  getEvent(): Event | undefined;
  setEvent(value?: Event): CreateEventRequest;
  hasEvent(): boolean;
  clearEvent(): CreateEventRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateEventRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateEventRequest): CreateEventRequest.AsObject;
  static serializeBinaryToWriter(message: CreateEventRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateEventRequest;
  static deserializeBinaryFromReader(message: CreateEventRequest, reader: jspb.BinaryReader): CreateEventRequest;
}

export namespace CreateEventRequest {
  export type AsObject = {
    event?: Event.AsObject,
  }
}

export class CreateEventResponse extends jspb.Message {
  getId(): string;
  setId(value: string): CreateEventResponse;

  getIsPublished(): boolean;
  setIsPublished(value: boolean): CreateEventResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateEventResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateEventResponse): CreateEventResponse.AsObject;
  static serializeBinaryToWriter(message: CreateEventResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateEventResponse;
  static deserializeBinaryFromReader(message: CreateEventResponse, reader: jspb.BinaryReader): CreateEventResponse;
}

export namespace CreateEventResponse {
  export type AsObject = {
    id: string,
    isPublished: boolean,
  }
}

export class Event extends jspb.Message {
  getEventId(): string;
  setEventId(value: string): Event;

  getReceivedDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setReceivedDate(value?: google_protobuf_timestamp_pb.Timestamp): Event;
  hasReceivedDate(): boolean;
  clearReceivedDate(): Event;

  getPublishedDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setPublishedDate(value?: google_protobuf_timestamp_pb.Timestamp): Event;
  hasPublishedDate(): boolean;
  clearPublishedDate(): Event;

  getTransactionId(): string;
  setTransactionId(value: string): Event;

  getPublish(): boolean;
  setPublish(value: boolean): Event;

  getEventType(): number;
  setEventType(value: number): Event;

  getAggregateType(): number;
  setAggregateType(value: number): Event;

  getAggregateId(): string;
  setAggregateId(value: string): Event;

  getEventData(): string;
  setEventData(value: string): Event;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Event.AsObject;
  static toObject(includeInstance: boolean, msg: Event): Event.AsObject;
  static serializeBinaryToWriter(message: Event, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Event;
  static deserializeBinaryFromReader(message: Event, reader: jspb.BinaryReader): Event;
}

export namespace Event {
  export type AsObject = {
    eventId: string,
    receivedDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    publishedDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    transactionId: string,
    publish: boolean,
    eventType: number,
    aggregateType: number,
    aggregateId: string,
    eventData: string,
  }
}

