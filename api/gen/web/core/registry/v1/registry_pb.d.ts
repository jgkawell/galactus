import * as jspb from 'google-protobuf'

import * as core_aggregates_v1_registry_pb from '../../../core/aggregates/v1/registry_pb';


export class RegisterRequest extends jspb.Message {
  getName(): string;
  setName(value: string): RegisterRequest;

  getDomain(): string;
  setDomain(value: string): RegisterRequest;

  getVersion(): string;
  setVersion(value: string): RegisterRequest;

  getDescription(): string;
  setDescription(value: string): RegisterRequest;

  getProtocolsList(): Array<Protocol>;
  setProtocolsList(value: Array<Protocol>): RegisterRequest;
  clearProtocolsList(): RegisterRequest;
  addProtocols(value?: Protocol, index?: number): Protocol;

  getProducersList(): Array<core_aggregates_v1_registry_pb.Producer>;
  setProducersList(value: Array<core_aggregates_v1_registry_pb.Producer>): RegisterRequest;
  clearProducersList(): RegisterRequest;
  addProducers(value?: core_aggregates_v1_registry_pb.Producer, index?: number): core_aggregates_v1_registry_pb.Producer;

  getConsumersList(): Array<core_aggregates_v1_registry_pb.Consumer>;
  setConsumersList(value: Array<core_aggregates_v1_registry_pb.Consumer>): RegisterRequest;
  clearConsumersList(): RegisterRequest;
  addConsumers(value?: core_aggregates_v1_registry_pb.Consumer, index?: number): core_aggregates_v1_registry_pb.Consumer;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterRequest): RegisterRequest.AsObject;
  static serializeBinaryToWriter(message: RegisterRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterRequest;
  static deserializeBinaryFromReader(message: RegisterRequest, reader: jspb.BinaryReader): RegisterRequest;
}

export namespace RegisterRequest {
  export type AsObject = {
    name: string,
    domain: string,
    version: string,
    description: string,
    protocolsList: Array<Protocol.AsObject>,
    producersList: Array<core_aggregates_v1_registry_pb.Producer.AsObject>,
    consumersList: Array<core_aggregates_v1_registry_pb.Consumer.AsObject>,
  }
}

export class Protocol extends jspb.Message {
  getKind(): core_aggregates_v1_registry_pb.ProtocolKind;
  setKind(value: core_aggregates_v1_registry_pb.ProtocolKind): Protocol;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Protocol.AsObject;
  static toObject(includeInstance: boolean, msg: Protocol): Protocol.AsObject;
  static serializeBinaryToWriter(message: Protocol, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Protocol;
  static deserializeBinaryFromReader(message: Protocol, reader: jspb.BinaryReader): Protocol;
}

export namespace Protocol {
  export type AsObject = {
    kind: core_aggregates_v1_registry_pb.ProtocolKind,
  }
}

export class RegisterResponse extends jspb.Message {
  getRegistration(): core_aggregates_v1_registry_pb.Registration | undefined;
  setRegistration(value?: core_aggregates_v1_registry_pb.Registration): RegisterResponse;
  hasRegistration(): boolean;
  clearRegistration(): RegisterResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterResponse): RegisterResponse.AsObject;
  static serializeBinaryToWriter(message: RegisterResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterResponse;
  static deserializeBinaryFromReader(message: RegisterResponse, reader: jspb.BinaryReader): RegisterResponse;
}

export namespace RegisterResponse {
  export type AsObject = {
    registration?: core_aggregates_v1_registry_pb.Registration.AsObject,
  }
}

export class ConnectionRequest extends jspb.Message {
  getName(): string;
  setName(value: string): ConnectionRequest;

  getVersion(): string;
  setVersion(value: string): ConnectionRequest;

  getType(): core_aggregates_v1_registry_pb.ProtocolKind;
  setType(value: core_aggregates_v1_registry_pb.ProtocolKind): ConnectionRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ConnectionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ConnectionRequest): ConnectionRequest.AsObject;
  static serializeBinaryToWriter(message: ConnectionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ConnectionRequest;
  static deserializeBinaryFromReader(message: ConnectionRequest, reader: jspb.BinaryReader): ConnectionRequest;
}

export namespace ConnectionRequest {
  export type AsObject = {
    name: string,
    version: string,
    type: core_aggregates_v1_registry_pb.ProtocolKind,
  }
}

export class ConnectionResponse extends jspb.Message {
  getAddress(): string;
  setAddress(value: string): ConnectionResponse;

  getPort(): number;
  setPort(value: number): ConnectionResponse;

  getStatus(): ServiceStatus;
  setStatus(value: ServiceStatus): ConnectionResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ConnectionResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ConnectionResponse): ConnectionResponse.AsObject;
  static serializeBinaryToWriter(message: ConnectionResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ConnectionResponse;
  static deserializeBinaryFromReader(message: ConnectionResponse, reader: jspb.BinaryReader): ConnectionResponse;
}

export namespace ConnectionResponse {
  export type AsObject = {
    address: string,
    port: number,
    status: ServiceStatus,
  }
}

export enum ServiceStatus { 
  SERVICE_STATUS_INVALID = 0,
  SERVICE_STATUS_HEALTHY = 1,
  SERVICE_STATUS_UNHEALTHY = 2,
}
