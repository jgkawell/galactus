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

  getProtocolsList(): Array<ProtocolRequest>;
  setProtocolsList(value: Array<ProtocolRequest>): RegisterRequest;
  clearProtocolsList(): RegisterRequest;
  addProtocols(value?: ProtocolRequest, index?: number): ProtocolRequest;

  getConsumersList(): Array<ConsumerRequest>;
  setConsumersList(value: Array<ConsumerRequest>): RegisterRequest;
  clearConsumersList(): RegisterRequest;
  addConsumers(value?: ConsumerRequest, index?: number): ConsumerRequest;

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
    protocolsList: Array<ProtocolRequest.AsObject>,
    consumersList: Array<ConsumerRequest.AsObject>,
  }
}

export class ProtocolRequest extends jspb.Message {
  getKind(): core_aggregates_v1_registry_pb.ProtocolKind;
  setKind(value: core_aggregates_v1_registry_pb.ProtocolKind): ProtocolRequest;

  getRoute(): string;
  setRoute(value: string): ProtocolRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ProtocolRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ProtocolRequest): ProtocolRequest.AsObject;
  static serializeBinaryToWriter(message: ProtocolRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ProtocolRequest;
  static deserializeBinaryFromReader(message: ProtocolRequest, reader: jspb.BinaryReader): ProtocolRequest;
}

export namespace ProtocolRequest {
  export type AsObject = {
    kind: core_aggregates_v1_registry_pb.ProtocolKind,
    route: string,
  }
}

export class ConsumerRequest extends jspb.Message {
  getKind(): core_aggregates_v1_registry_pb.ConsumerKind;
  setKind(value: core_aggregates_v1_registry_pb.ConsumerKind): ConsumerRequest;

  getOrder(): number;
  setOrder(value: number): ConsumerRequest;

  getAggregateType(): string;
  setAggregateType(value: string): ConsumerRequest;

  getEventType(): string;
  setEventType(value: string): ConsumerRequest;

  getEventCode(): string;
  setEventCode(value: string): ConsumerRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ConsumerRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ConsumerRequest): ConsumerRequest.AsObject;
  static serializeBinaryToWriter(message: ConsumerRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ConsumerRequest;
  static deserializeBinaryFromReader(message: ConsumerRequest, reader: jspb.BinaryReader): ConsumerRequest;
}

export namespace ConsumerRequest {
  export type AsObject = {
    kind: core_aggregates_v1_registry_pb.ConsumerKind,
    order: number,
    aggregateType: string,
    eventType: string,
    eventCode: string,
  }
}

export class RegisterResponse extends jspb.Message {
  getProtocolsList(): Array<ProtocolResponse>;
  setProtocolsList(value: Array<ProtocolResponse>): RegisterResponse;
  clearProtocolsList(): RegisterResponse;
  addProtocols(value?: ProtocolResponse, index?: number): ProtocolResponse;

  getConsumersList(): Array<ConsumerResponse>;
  setConsumersList(value: Array<ConsumerResponse>): RegisterResponse;
  clearConsumersList(): RegisterResponse;
  addConsumers(value?: ConsumerResponse, index?: number): ConsumerResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterResponse): RegisterResponse.AsObject;
  static serializeBinaryToWriter(message: RegisterResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterResponse;
  static deserializeBinaryFromReader(message: RegisterResponse, reader: jspb.BinaryReader): RegisterResponse;
}

export namespace RegisterResponse {
  export type AsObject = {
    protocolsList: Array<ProtocolResponse.AsObject>,
    consumersList: Array<ConsumerResponse.AsObject>,
  }
}

export class ProtocolResponse extends jspb.Message {
  getKind(): core_aggregates_v1_registry_pb.ProtocolKind;
  setKind(value: core_aggregates_v1_registry_pb.ProtocolKind): ProtocolResponse;

  getPort(): number;
  setPort(value: number): ProtocolResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ProtocolResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ProtocolResponse): ProtocolResponse.AsObject;
  static serializeBinaryToWriter(message: ProtocolResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ProtocolResponse;
  static deserializeBinaryFromReader(message: ProtocolResponse, reader: jspb.BinaryReader): ProtocolResponse;
}

export namespace ProtocolResponse {
  export type AsObject = {
    kind: core_aggregates_v1_registry_pb.ProtocolKind,
    port: number,
  }
}

export class ConsumerResponse extends jspb.Message {
  getKind(): core_aggregates_v1_registry_pb.ConsumerKind;
  setKind(value: core_aggregates_v1_registry_pb.ConsumerKind): ConsumerResponse;

  getOrder(): number;
  setOrder(value: number): ConsumerResponse;

  getRoutingKey(): string;
  setRoutingKey(value: string): ConsumerResponse;

  getExchange(): string;
  setExchange(value: string): ConsumerResponse;

  getQueueName(): string;
  setQueueName(value: string): ConsumerResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ConsumerResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ConsumerResponse): ConsumerResponse.AsObject;
  static serializeBinaryToWriter(message: ConsumerResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ConsumerResponse;
  static deserializeBinaryFromReader(message: ConsumerResponse, reader: jspb.BinaryReader): ConsumerResponse;
}

export namespace ConsumerResponse {
  export type AsObject = {
    kind: core_aggregates_v1_registry_pb.ConsumerKind,
    order: number,
    routingKey: string,
    exchange: string,
    queueName: string,
  }
}

export class ConnectionRequest extends jspb.Message {
  getPath(): string;
  setPath(value: string): ConnectionRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ConnectionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ConnectionRequest): ConnectionRequest.AsObject;
  static serializeBinaryToWriter(message: ConnectionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ConnectionRequest;
  static deserializeBinaryFromReader(message: ConnectionRequest, reader: jspb.BinaryReader): ConnectionRequest;
}

export namespace ConnectionRequest {
  export type AsObject = {
    path: string,
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
  SERVICE_STATUS_INVALID_UNSPECIFIED = 0,
  SERVICE_STATUS_HEALTHY = 1,
  SERVICE_STATUS_UNHEALTHY = 2,
}
