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
  SYSTEM_EVENT_CODE_INVALID_UNSPECIFIED = 0,
  SYSTEM_EVENT_CODE_ERROR = 1,
}
export enum SystemErrorCode { 
  SYSTEM_ERROR_CODE_INVALID_UNSPECIFIED = 0,
  SYSTEM_ERROR_CODE_FAILED_EVENT_PUBLISH = 1,
  SYSTEM_ERROR_CODE_FAILED_EVENT_SAVED = 2,
  SYSTEM_ERROR_CODE_FAILED_EVENT_FORWARD = 3,
  SYSTEM_ERROR_CODE_MALFORMED_EVENT_DATA = 4,
}
