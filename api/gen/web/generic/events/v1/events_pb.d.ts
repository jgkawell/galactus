import * as jspb from 'google-protobuf'

import * as generic_events_v1_notifications_pb from '../../../generic/events/v1/notifications_pb';
import * as generic_events_v1_system_pb from '../../../generic/events/v1/system_pb';
import * as generic_events_v1_todo_pb from '../../../generic/events/v1/todo_pb';


export class EventType extends jspb.Message {
  getSystemCode(): generic_events_v1_system_pb.SystemEventCode;
  setSystemCode(value: generic_events_v1_system_pb.SystemEventCode): EventType;

  getNotificationCode(): generic_events_v1_notifications_pb.NotificationEventCode;
  setNotificationCode(value: generic_events_v1_notifications_pb.NotificationEventCode): EventType;

  getTodoEventCode(): generic_events_v1_todo_pb.TodoEventCode;
  setTodoEventCode(value: generic_events_v1_todo_pb.TodoEventCode): EventType;

  getCodeCase(): EventType.CodeCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EventType.AsObject;
  static toObject(includeInstance: boolean, msg: EventType): EventType.AsObject;
  static serializeBinaryToWriter(message: EventType, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EventType;
  static deserializeBinaryFromReader(message: EventType, reader: jspb.BinaryReader): EventType;
}

export namespace EventType {
  export type AsObject = {
    systemCode: generic_events_v1_system_pb.SystemEventCode,
    notificationCode: generic_events_v1_notifications_pb.NotificationEventCode,
    todoEventCode: generic_events_v1_todo_pb.TodoEventCode,
  }

  export enum CodeCase { 
    CODE_NOT_SET = 0,
    SYSTEM_CODE = 1,
    NOTIFICATION_CODE = 2,
    TODO_EVENT_CODE = 3,
  }
}

export enum AggregateType { 
  AGGREGATE_TYPE_INVALID = 0,
  AGGREGATE_TYPE_SYSTEM = 1,
  AGGREGATE_TYPE_NOTIFICATION = 2,
  AGGREGATE_TYPE_TODO = 3,
}
