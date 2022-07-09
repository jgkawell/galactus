import * as grpcWeb from 'grpc-web';

import * as core_commandhandler_v1_commandhandler_pb from '../../../core/commandhandler/v1/commandhandler_pb';


export class CommandHandlerClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  apply(
    request: core_commandhandler_v1_commandhandler_pb.ApplyCommandRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: core_commandhandler_v1_commandhandler_pb.ApplyCommandResponse) => void
  ): grpcWeb.ClientReadableStream<core_commandhandler_v1_commandhandler_pb.ApplyCommandResponse>;

}

export class CommandHandlerPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  apply(
    request: core_commandhandler_v1_commandhandler_pb.ApplyCommandRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<core_commandhandler_v1_commandhandler_pb.ApplyCommandResponse>;

}

