/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "kvs.kvs";

export interface MsgDataProposal {
  creator: string;
  key: string;
  value: string;
}

export interface MsgDataProposalResponse {
}

export interface MsgDataConfirmation {
  creator: string;
  key: string;
}

export interface MsgDataConfirmationResponse {
}

export interface MsgAddressRegistration {
  creator: string;
  addresses: string[];
}

export interface MsgAddressRegistrationResponse {
}

function createBaseMsgDataProposal(): MsgDataProposal {
  return { creator: "", key: "", value: "" };
}

export const MsgDataProposal = {
  encode(message: MsgDataProposal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.key !== "") {
      writer.uint32(18).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(26).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDataProposal {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDataProposal();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.key = reader.string();
          break;
        case 3:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDataProposal {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: MsgDataProposal): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDataProposal>, I>>(object: I): MsgDataProposal {
    const message = createBaseMsgDataProposal();
    message.creator = object.creator ?? "";
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseMsgDataProposalResponse(): MsgDataProposalResponse {
  return {};
}

export const MsgDataProposalResponse = {
  encode(_: MsgDataProposalResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDataProposalResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDataProposalResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDataProposalResponse {
    return {};
  },

  toJSON(_: MsgDataProposalResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDataProposalResponse>, I>>(_: I): MsgDataProposalResponse {
    const message = createBaseMsgDataProposalResponse();
    return message;
  },
};

function createBaseMsgDataConfirmation(): MsgDataConfirmation {
  return { creator: "", key: "" };
}

export const MsgDataConfirmation = {
  encode(message: MsgDataConfirmation, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.key !== "") {
      writer.uint32(18).string(message.key);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDataConfirmation {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDataConfirmation();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.key = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDataConfirmation {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      key: isSet(object.key) ? String(object.key) : "",
    };
  },

  toJSON(message: MsgDataConfirmation): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.key !== undefined && (obj.key = message.key);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDataConfirmation>, I>>(object: I): MsgDataConfirmation {
    const message = createBaseMsgDataConfirmation();
    message.creator = object.creator ?? "";
    message.key = object.key ?? "";
    return message;
  },
};

function createBaseMsgDataConfirmationResponse(): MsgDataConfirmationResponse {
  return {};
}

export const MsgDataConfirmationResponse = {
  encode(_: MsgDataConfirmationResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDataConfirmationResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDataConfirmationResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDataConfirmationResponse {
    return {};
  },

  toJSON(_: MsgDataConfirmationResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDataConfirmationResponse>, I>>(_: I): MsgDataConfirmationResponse {
    const message = createBaseMsgDataConfirmationResponse();
    return message;
  },
};

function createBaseMsgAddressRegistration(): MsgAddressRegistration {
  return { creator: "", addresses: [] };
}

export const MsgAddressRegistration = {
  encode(message: MsgAddressRegistration, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    for (const v of message.addresses) {
      writer.uint32(18).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddressRegistration {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddressRegistration();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.addresses.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddressRegistration {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      addresses: Array.isArray(object?.addresses) ? object.addresses.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: MsgAddressRegistration): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    if (message.addresses) {
      obj.addresses = message.addresses.map((e) => e);
    } else {
      obj.addresses = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddressRegistration>, I>>(object: I): MsgAddressRegistration {
    const message = createBaseMsgAddressRegistration();
    message.creator = object.creator ?? "";
    message.addresses = object.addresses?.map((e) => e) || [];
    return message;
  },
};

function createBaseMsgAddressRegistrationResponse(): MsgAddressRegistrationResponse {
  return {};
}

export const MsgAddressRegistrationResponse = {
  encode(_: MsgAddressRegistrationResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddressRegistrationResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddressRegistrationResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgAddressRegistrationResponse {
    return {};
  },

  toJSON(_: MsgAddressRegistrationResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddressRegistrationResponse>, I>>(_: I): MsgAddressRegistrationResponse {
    const message = createBaseMsgAddressRegistrationResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  DataProposal(request: MsgDataProposal): Promise<MsgDataProposalResponse>;
  DataConfirmation(request: MsgDataConfirmation): Promise<MsgDataConfirmationResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  AddressRegistration(request: MsgAddressRegistration): Promise<MsgAddressRegistrationResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.DataProposal = this.DataProposal.bind(this);
    this.DataConfirmation = this.DataConfirmation.bind(this);
    this.AddressRegistration = this.AddressRegistration.bind(this);
  }
  DataProposal(request: MsgDataProposal): Promise<MsgDataProposalResponse> {
    const data = MsgDataProposal.encode(request).finish();
    const promise = this.rpc.request("kvs.kvs.Msg", "DataProposal", data);
    return promise.then((data) => MsgDataProposalResponse.decode(new _m0.Reader(data)));
  }

  DataConfirmation(request: MsgDataConfirmation): Promise<MsgDataConfirmationResponse> {
    const data = MsgDataConfirmation.encode(request).finish();
    const promise = this.rpc.request("kvs.kvs.Msg", "DataConfirmation", data);
    return promise.then((data) => MsgDataConfirmationResponse.decode(new _m0.Reader(data)));
  }

  AddressRegistration(request: MsgAddressRegistration): Promise<MsgAddressRegistrationResponse> {
    const data = MsgAddressRegistration.encode(request).finish();
    const promise = this.rpc.request("kvs.kvs.Msg", "AddressRegistration", data);
    return promise.then((data) => MsgAddressRegistrationResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
