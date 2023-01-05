import * as jspb from 'google-protobuf'

import * as validate_validate_pb from '../../../validate/validate_pb';
import * as options_gorm_pb from '../../../options/gorm_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class Registration extends jspb.Message {
  getId(): string;
  setId(value: string): Registration;

  getName(): string;
  setName(value: string): Registration;

  getVersion(): string;
  setVersion(value: string): Registration;

  getDomain(): string;
  setDomain(value: string): Registration;

  getDescription(): string;
  setDescription(value: string): Registration;

  getAddress(): string;
  setAddress(value: string): Registration;

  getStatus(): ServiceStatus;
  setStatus(value: ServiceStatus): Registration;

  getProtocolsList(): Array<Protocol>;
  setProtocolsList(value: Array<Protocol>): Registration;
  clearProtocolsList(): Registration;
  addProtocols(value?: Protocol, index?: number): Protocol;

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
    name: string,
    version: string,
    domain: string,
    description: string,
    address: string,
    status: ServiceStatus,
    protocolsList: Array<Protocol.AsObject>,
    consumersList: Array<Consumer.AsObject>,
  }
}

export class Protocol extends jspb.Message {
  getId(): string;
  setId(value: string): Protocol;

  getKind(): ProtocolKind;
  setKind(value: ProtocolKind): Protocol;

  getVersion(): string;
  setVersion(value: string): Protocol;

  getPort(): number;
  setPort(value: number): Protocol;

  getRoute(): string;
  setRoute(value: string): Protocol;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Protocol.AsObject;
  static toObject(includeInstance: boolean, msg: Protocol): Protocol.AsObject;
  static serializeBinaryToWriter(message: Protocol, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Protocol;
  static deserializeBinaryFromReader(message: Protocol, reader: jspb.BinaryReader): Protocol;
}

export namespace Protocol {
  export type AsObject = {
    id: string,
    kind: ProtocolKind,
    version: string,
    port: number,
    route: string,
  }
}

export class Consumer extends jspb.Message {
  getId(): string;
  setId(value: string): Consumer;

  getRoutingKey(): string;
  setRoutingKey(value: string): Consumer;

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
    routingKey: string,
    kind: ConsumerKind,
  }
}

export enum ServiceStatus { 
  SERVICE_STATUS_INVALID = 0,
  SERVICE_STATUS_REGISTERED = 1,
  SERVICE_STATUS_DEREGISTERED = 2,
  SERVICE_STATUS_HEALTHY = 3,
  SERVICE_STATUS_UNHEALTHY = 4,
}
export enum ProtocolKind { 
  PROTOCOL_KIND_INVALID = 0,
  PROTOCOL_KIND_GRPC = 1,
  PROTOCOL_KIND_HTTP = 2,
}
export enum ConsumerKind { 
  CONSUMER_KIND_INVALID = 0,
  CONSUMER_KIND_QUEUE = 1,
  CONSUMER_KIND_TOPIC = 2,
}
