import * as jspb from 'google-protobuf'


import * as options_gorm_pb from '../../../options/gorm_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class Event extends jspb.Message {
  getId(): string;
  setId(value: string): Event;

  getReceivedTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setReceivedTime(value?: google_protobuf_timestamp_pb.Timestamp): Event;
  hasReceivedTime(): boolean;
  clearReceivedTime(): Event;

  getPublishedTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setPublishedTime(value?: google_protobuf_timestamp_pb.Timestamp): Event;
  hasPublishedTime(): boolean;
  clearPublishedTime(): Event;

  getTransactionId(): string;
  setTransactionId(value: string): Event;

  getAggregateType(): string;
  setAggregateType(value: string): Event;

  getEventType(): string;
  setEventType(value: string): Event;

  getEventCode(): string;
  setEventCode(value: string): Event;

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
    id: string,
    receivedTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    publishedTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    transactionId: string,
    aggregateType: string,
    eventType: string,
    eventCode: string,
    aggregateId: string,
    eventData: string,
  }
}

