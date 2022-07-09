import * as jspb from 'google-protobuf'



export class SystemError extends jspb.Message {
  getCode(): SystemErrorCode;
  setCode(value: SystemErrorCode): SystemError;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SystemError.AsObject;
  static toObject(includeInstance: boolean, msg: SystemError): SystemError.AsObject;
  static serializeBinaryToWriter(message: SystemError, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SystemError;
  static deserializeBinaryFromReader(message: SystemError, reader: jspb.BinaryReader): SystemError;
}

export namespace SystemError {
  export type AsObject = {
    code: SystemErrorCode,
  }
}

export enum SystemEventCode { 
  UNKOWN_SYSTEM_EVENT = 0,
  SYSTEM_ERROR = 1,
}
export enum SystemErrorCode { 
  INVALID_SYSTEM_MESSAGE_DATA = 0,
  FAILED_EVENT_PUBLISH = 1,
  FAILED_EVENT_SAVED = 2,
  FAILED_EVENT_FORWARD = 3,
}
