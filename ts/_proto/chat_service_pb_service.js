// package: proto
// file: chat_service.proto

var chat_service_pb = require("./chat_service_pb");

var Chat = (function () {
  function Chat() {}
  Chat.serviceName = "proto.Chat";
  return Chat;
}());

Chat.connect = {
  methodName: "connect",
  service: Chat,
  requestStream: true,
  responseStream: true,
  requestType: chat_service_pb.WebMessage,
  responseType: chat_service_pb.WebMessage
};

exports.Chat = Chat;

