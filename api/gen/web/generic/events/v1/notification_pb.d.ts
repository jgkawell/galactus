import * as jspb from 'google-protobuf'

import * as core_notifier_v1_notifier_pb from '../../../core/notifier/v1/notifier_pb';


export class NotificationDeliveryRequested extends jspb.Message {
  getActorId(): string;
  setActorId(value: string): NotificationDeliveryRequested;

  getClientId(): string;
  setClientId(value: string): NotificationDeliveryRequested;

  getNotification(): core_notifier_v1_notifier_pb.Notification | undefined;
  setNotification(value?: core_notifier_v1_notifier_pb.Notification): NotificationDeliveryRequested;
  hasNotification(): boolean;
  clearNotification(): NotificationDeliveryRequested;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NotificationDeliveryRequested.AsObject;
  static toObject(includeInstance: boolean, msg: NotificationDeliveryRequested): NotificationDeliveryRequested.AsObject;
  static serializeBinaryToWriter(message: NotificationDeliveryRequested, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NotificationDeliveryRequested;
  static deserializeBinaryFromReader(message: NotificationDeliveryRequested, reader: jspb.BinaryReader): NotificationDeliveryRequested;
}

export namespace NotificationDeliveryRequested {
  export type AsObject = {
    actorId: string,
    clientId: string,
    notification?: core_notifier_v1_notifier_pb.Notification.AsObject,
  }
}

export class NotificationDelivered extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NotificationDelivered.AsObject;
  static toObject(includeInstance: boolean, msg: NotificationDelivered): NotificationDelivered.AsObject;
  static serializeBinaryToWriter(message: NotificationDelivered, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NotificationDelivered;
  static deserializeBinaryFromReader(message: NotificationDelivered, reader: jspb.BinaryReader): NotificationDelivered;
}

export namespace NotificationDelivered {
  export type AsObject = {
  }
}

export enum NotificationEventCode { 
  INVALID_NOTIFICATION_EVENT_CODE = 0,
  NOTIFICATION_DELIVERY_REQUESTED = 1,
  NOTIFICATION_DELIVERED = 2,
}
