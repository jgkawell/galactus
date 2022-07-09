import * as grpcWeb from 'grpc-web';

import * as core_notifier_v1_notifier_pb from '../../../core/notifier/v1/notifier_pb';


export class NotifierClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  connect(
    request: core_notifier_v1_notifier_pb.ConnectionRequest,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<core_notifier_v1_notifier_pb.Notification>;

}

export class NotifierPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  connect(
    request: core_notifier_v1_notifier_pb.ConnectionRequest,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<core_notifier_v1_notifier_pb.Notification>;

}

