import {grpc} from "grpc-web-client";
import {Chat} from "../_proto/chat_service_pb_service";
import {WebMessage, GetMessageRequest} from "../_proto/chat_service_pb";

declare const USE_TLS: boolean;
const host = USE_TLS ? "https://localhost:10000" : "https://localhost:10000";

function chat() {
  console.log("here")
  // const client = grpc.client(Chat.GetMessage, {
  //   host: "https://localhost:10000",
  // });

  // client.start();
  // const message = new WebMessage();
  // message.setText("hello");

  // client.send(message);

grpc.unary(Chat.GetMessage, {
  request: new GetMessageRequest(),
  host: host,
  onEnd: res => {
    console.log("res", res)
  }
})  
//   grpc.unary(WebMessage, {
//     request: new WebMessage(),
//     host: host
//     // onEnd:  => {
//       // const { status, statusMessage, headers, message, trailers } = res;
//       // console.log("getBook.onEnd.status", status, statusMessage);
//     // }
// });
  // grpc.invoke(Chat.connect, {
  //   host: host,
  //   onMessage: ((message: WebMessage) => {
  //     console.log("onMessage", message);
  //   }),
  //     onEnd: ((status: grpc.Code, statusMessage: string, trailers: grpc.Metadata) => {
  //       console.log("onEnd", status, statusMessage, trailers);
  //     }),
  // });
}

chat();
