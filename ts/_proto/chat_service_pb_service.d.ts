// package: proto
// file: chat_service.proto

import * as chat_service_pb from "./chat_service_pb";

type Chatconnect = {
  readonly methodName: string;
  readonly service: typeof Chat;
  readonly requestStream: true;
  readonly responseStream: true;
  readonly requestType: typeof chat_service_pb.WebMessage;
  readonly responseType: typeof chat_service_pb.WebMessage;
};

export class Chat {
  static readonly serviceName: string;
  static readonly connect: Chatconnect;
}
