import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as options_gorm_pb from '../../../options/gorm_pb';
import * as validate_validate_pb from '../../../validate/validate_pb';


export class Registration extends jspb.Message {
  getId(): string;
  setId(value: string): Registration;

  getDomain(): string;
  setDomain(value: string): Registration;

  getName(): string;
  setName(value: string): Registration;

  getVersion(): string;
  setVersion(value: string): Registration;

  getDescription(): string;
  setDescription(value: string): Registration;

  getStatus(): ServiceStatus;
  setStatus(value: ServiceStatus): Registration;

  getRoutesList(): Array<Route>;
  setRoutesList(value: Array<Route>): Registration;
  clearRoutesList(): Registration;
  addRoutes(value?: Route, index?: number): Route;

  getConsumersList(): Array<Consumer>;
  setConsumersList(value: Array<Consumer>): Registration;
  clearConsumersList(): Registration;
  addConsumers(value?: Consumer, index?: number): Consumer;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Registration.AsObject;
  static toObject(includeInstance: boolean, msg: Registration): Registration.AsObject;
  static serializeBinaryToWriter(message: Registration, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Registration;
  static deserializeBinaryFromReader(message: Registration, reader: jspb.BinaryReader): Registration;
}

export namespace Registration {
  export type AsObject = {
    id: string,
    domain: string,
    name: string,
    version: string,
    description: string,
    status: ServiceStatus,
    routesList: Array<Route.AsObject>,
    consumersList: Array<Consumer.AsObject>,
  }
}

export class Route extends jspb.Message {
  getId(): string;
  setId(value: string): Route;

  getPath(): string;
  setPath(value: string): Route;

  getHost(): string;
  setHost(value: string): Route;

  getPort(): number;
  setPort(value: number): Route;

  getKind(): ProtocolKind;
  setKind(value: ProtocolKind): Route;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Route.AsObject;
  static toObject(includeInstance: boolean, msg: Route): Route.AsObject;
  static serializeBinaryToWriter(message: Route, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Route;
  static deserializeBinaryFromReader(message: Route, reader: jspb.BinaryReader): Route;
}

export namespace Route {
  export type AsObject = {
    id: string,
    path: string,
    host: string,
    port: number,
    kind: ProtocolKind,
  }
}

export class Consumer extends jspb.Message {
  getId(): string;
  setId(value: string): Consumer;

  getAggregateType(): string;
  setAggregateType(value: string): Consumer;

  getEventType(): string;
  setEventType(value: string): Consumer;

  getEventCode(): string;
  setEventCode(value: string): Consumer;

  getKind(): ConsumerKind;
  setKind(value: ConsumerKind): Consumer;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Consumer.AsObject;
  static toObject(includeInstance: boolean, msg: Consumer): Consumer.AsObject;
  static serializeBinaryToWriter(message: Consumer, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Consumer;
  static deserializeBinaryFromReader(message: Consumer, reader: jspb.BinaryReader): Consumer;
}

export namespace Consumer {
  export type AsObject = {
    id: string,
    aggregateType: string,
    eventType: string,
    eventCode: string,
    kind: ConsumerKind,
  }
}

export enum ServiceStatus { 
  SERVICE_STATUS_UNSPECIFIED = 0,
  SERVICE_STATUS_REGISTERED = 1,
  SERVICE_STATUS_DEREGISTERED = 2,
  SERVICE_STATUS_HEALTHY = 3,
  SERVICE_STATUS_UNHEALTHY = 4,
}
export enum ProtocolKind { 
  PROTOCOL_KIND_UNSPECIFIED = 0,
  PROTOCOL_KIND_GRPC = 1,
  PROTOCOL_KIND_HTTP = 2,
}
export enum ConsumerKind { 
  CONSUMER_KIND_UNSPECIFIED = 0,
  CONSUMER_KIND_QUEUE = 1,
  CONSUMER_KIND_TOPIC = 2,
}
