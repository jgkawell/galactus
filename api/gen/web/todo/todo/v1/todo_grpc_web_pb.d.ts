import * as grpcWeb from 'grpc-web';

import * as todo_todo_v1_todo_pb from '../../../todo/todo/v1/todo_pb';


export class TodoClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  create(
    request: todo_todo_v1_todo_pb.CreateTodoRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: todo_todo_v1_todo_pb.CreateTodoResponse) => void
  ): grpcWeb.ClientReadableStream<todo_todo_v1_todo_pb.CreateTodoResponse>;

}

export class TodoPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  create(
    request: todo_todo_v1_todo_pb.CreateTodoRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<todo_todo_v1_todo_pb.CreateTodoResponse>;

}

