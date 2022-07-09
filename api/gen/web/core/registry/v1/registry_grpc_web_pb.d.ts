import * as grpcWeb from 'grpc-web';

import * as core_registry_v1_registry_pb from '../../../core/registry/v1/registry_pb';


export class RegistryClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  register(
    request: core_registry_v1_registry_pb.RegisterRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: core_registry_v1_registry_pb.RegisterResponse) => void
  ): grpcWeb.ClientReadableStream<core_registry_v1_registry_pb.RegisterResponse>;

  connection(
    request: core_registry_v1_registry_pb.ConnectionRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: core_registry_v1_registry_pb.ConnectionResponse) => void
  ): grpcWeb.ClientReadableStream<core_registry_v1_registry_pb.ConnectionResponse>;

}

export class RegistryPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  register(
    request: core_registry_v1_registry_pb.RegisterRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<core_registry_v1_registry_pb.RegisterResponse>;

  connection(
    request: core_registry_v1_registry_pb.ConnectionRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<core_registry_v1_registry_pb.ConnectionResponse>;

}

