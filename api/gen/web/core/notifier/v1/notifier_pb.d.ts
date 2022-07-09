import * as jspb from 'google-protobuf'



export class ConnectionRequest extends jspb.Message {
  getUserId(): string;
  setUserId(value: string): ConnectionRequest;

  getClientId(): string;
  setClientId(value: string): ConnectionRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ConnectionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ConnectionRequest): ConnectionRequest.AsObject;
  static serializeBinaryToWriter(message: ConnectionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ConnectionRequest;
  static deserializeBinaryFromReader(message: ConnectionRequest, reader: jspb.BinaryReader): ConnectionRequest;
}

export namespace ConnectionRequest {
  export type AsObject = {
    userId: string,
    clientId: string,
  }
}

export class Notification extends jspb.Message {
  getNotificationType(): NotificationType;
  setNotificationType(value: NotificationType): Notification;

  getData(): string;
  setData(value: string): Notification;

  getTransactionId(): string;
  setTransactionId(value: string): Notification;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Notification.AsObject;
  static toObject(includeInstance: boolean, msg: Notification): Notification.AsObject;
  static serializeBinaryToWriter(message: Notification, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Notification;
  static deserializeBinaryFromReader(message: Notification, reader: jspb.BinaryReader): Notification;
}

export namespace Notification {
  export type AsObject = {
    notificationType: NotificationType,
    data: string,
    transactionId: string,
  }
}

export class Heartbeat extends jspb.Message {
  getSessionId(): string;
  setSessionId(value: string): Heartbeat;

  getExpirationDeadline(): number;
  setExpirationDeadline(value: number): Heartbeat;

  getClientId(): string;
  setClientId(value: string): Heartbeat;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Heartbeat.AsObject;
  static toObject(includeInstance: boolean, msg: Heartbeat): Heartbeat.AsObject;
  static serializeBinaryToWriter(message: Heartbeat, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Heartbeat;
  static deserializeBinaryFromReader(message: Heartbeat, reader: jspb.BinaryReader): Heartbeat;
}

export namespace Heartbeat {
  export type AsObject = {
    sessionId: string,
    expirationDeadline: number,
    clientId: string,
  }
}

export enum NotificationType { 
  INVALID = 0,
  HEARTBEAT = 1,
  TODO_CREATED = 2,
}
