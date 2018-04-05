import {grpc} from "grpc-web-client";
import {Chat} from "../_proto/chat_service_pb_service";
import {WebMessage} from "../_proto/chat_service_pb";

declare const USE_TLS: boolean;
const host = USE_TLS ? "https://localhost:9091" : "http://localhost:9090";

function chat() {
  grpc.invoke(Chat.connect, {
    host: host,
    onMessage: ((message: WebMessage) => {
      console.log("onMessage", message);
    }),
      onEnd: ((status: grpc.Code, statusMessage: string, trailers: grpc.Metadata) => {
        console.log("onEnd", status, statusMessage, trailers);
      }),
  });
}
