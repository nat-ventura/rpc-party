// package: proto
// file: chat_service.proto

import * as jspb from "google-protobuf";

export class WebMessage extends jspb.Message {
  getText(): string;
  setText(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): WebMessage.AsObject;
  static toObject(includeInstance: boolean, msg: WebMessage): WebMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: WebMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): WebMessage;
  static deserializeBinaryFromReader(message: WebMessage, reader: jspb.BinaryReader): WebMessage;
}

export namespace WebMessage {
  export type AsObject = {
    text: string,
  }
}

export class GetMessageRequest extends jspb.Message {
  getText(): string;
  setText(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetMessageRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetMessageRequest): GetMessageRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetMessageRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetMessageRequest;
  static deserializeBinaryFromReader(message: GetMessageRequest, reader: jspb.BinaryReader): GetMessageRequest;
}

export namespace GetMessageRequest {
  export type AsObject = {
    text: string,
  }
}

