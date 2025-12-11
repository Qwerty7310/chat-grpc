import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb'; // proto import: "google/protobuf/timestamp.proto"
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb'; // proto import: "google/protobuf/empty.proto"


export class RegisterRequest extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): RegisterRequest;

  getPassword(): string;
  setPassword(value: string): RegisterRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterRequest): RegisterRequest.AsObject;
  static serializeBinaryToWriter(message: RegisterRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterRequest;
  static deserializeBinaryFromReader(message: RegisterRequest, reader: jspb.BinaryReader): RegisterRequest;
}

export namespace RegisterRequest {
  export type AsObject = {
    username: string,
    password: string,
  }
}

export class LoginRequest extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): LoginRequest;

  getPassword(): string;
  setPassword(value: string): LoginRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LoginRequest): LoginRequest.AsObject;
  static serializeBinaryToWriter(message: LoginRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginRequest;
  static deserializeBinaryFromReader(message: LoginRequest, reader: jspb.BinaryReader): LoginRequest;
}

export namespace LoginRequest {
  export type AsObject = {
    username: string,
    password: string,
  }
}

export class AuthResponse extends jspb.Message {
  getOk(): boolean;
  setOk(value: boolean): AuthResponse;

  getToken(): string;
  setToken(value: string): AuthResponse;

  getUserId(): string;
  setUserId(value: string): AuthResponse;

  getError(): string;
  setError(value: string): AuthResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AuthResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AuthResponse): AuthResponse.AsObject;
  static serializeBinaryToWriter(message: AuthResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AuthResponse;
  static deserializeBinaryFromReader(message: AuthResponse, reader: jspb.BinaryReader): AuthResponse;
}

export namespace AuthResponse {
  export type AsObject = {
    ok: boolean,
    token: string,
    userId: string,
    error: string,
  }
}

export class User extends jspb.Message {
  getId(): string;
  setId(value: string): User;

  getUsername(): string;
  setUsername(value: string): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    id: string,
    username: string,
  }
}

export class Message extends jspb.Message {
  getId(): string;
  setId(value: string): Message;

  getUserId(): string;
  setUserId(value: string): Message;

  getUsername(): string;
  setUsername(value: string): Message;

  getText(): string;
  setText(value: string): Message;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Message;
  hasCreatedAt(): boolean;
  clearCreatedAt(): Message;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Message.AsObject;
  static toObject(includeInstance: boolean, msg: Message): Message.AsObject;
  static serializeBinaryToWriter(message: Message, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Message;
  static deserializeBinaryFromReader(message: Message, reader: jspb.BinaryReader): Message;
}

export namespace Message {
  export type AsObject = {
    id: string,
    userId: string,
    username: string,
    text: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class MessageToServer extends jspb.Message {
  getText(): string;
  setText(value: string): MessageToServer;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MessageToServer.AsObject;
  static toObject(includeInstance: boolean, msg: MessageToServer): MessageToServer.AsObject;
  static serializeBinaryToWriter(message: MessageToServer, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MessageToServer;
  static deserializeBinaryFromReader(message: MessageToServer, reader: jspb.BinaryReader): MessageToServer;
}

export namespace MessageToServer {
  export type AsObject = {
    text: string,
  }
}

export class SendAck extends jspb.Message {
  getOk(): boolean;
  setOk(value: boolean): SendAck;

  getId(): string;
  setId(value: string): SendAck;

  getError(): string;
  setError(value: string): SendAck;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SendAck.AsObject;
  static toObject(includeInstance: boolean, msg: SendAck): SendAck.AsObject;
  static serializeBinaryToWriter(message: SendAck, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SendAck;
  static deserializeBinaryFromReader(message: SendAck, reader: jspb.BinaryReader): SendAck;
}

export namespace SendAck {
  export type AsObject = {
    ok: boolean,
    id: string,
    error: string,
  }
}

export class HistoryRequest extends jspb.Message {
  getLimit(): number;
  setLimit(value: number): HistoryRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HistoryRequest.AsObject;
  static toObject(includeInstance: boolean, msg: HistoryRequest): HistoryRequest.AsObject;
  static serializeBinaryToWriter(message: HistoryRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HistoryRequest;
  static deserializeBinaryFromReader(message: HistoryRequest, reader: jspb.BinaryReader): HistoryRequest;
}

export namespace HistoryRequest {
  export type AsObject = {
    limit: number,
  }
}

export class GetHistoryResponse extends jspb.Message {
  getMessageList(): Array<Message>;
  setMessageList(value: Array<Message>): GetHistoryResponse;
  clearMessageList(): GetHistoryResponse;
  addMessage(value?: Message, index?: number): Message;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetHistoryResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetHistoryResponse): GetHistoryResponse.AsObject;
  static serializeBinaryToWriter(message: GetHistoryResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetHistoryResponse;
  static deserializeBinaryFromReader(message: GetHistoryResponse, reader: jspb.BinaryReader): GetHistoryResponse;
}

export namespace GetHistoryResponse {
  export type AsObject = {
    messageList: Array<Message.AsObject>,
  }
}

