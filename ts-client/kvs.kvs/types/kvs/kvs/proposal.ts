/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "kvs.kvs";

export interface Proposal {
  index: string;
  value: string;
  acknowledgments: string[];
}

function createBaseProposal(): Proposal {
  return { index: "", value: "", acknowledgments: [] };
}

export const Proposal = {
  encode(message: Proposal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    for (const v of message.acknowledgments) {
      writer.uint32(26).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Proposal {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProposal();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        case 3:
          message.acknowledgments.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Proposal {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      value: isSet(object.value) ? String(object.value) : "",
      acknowledgments: Array.isArray(object?.acknowledgments) ? object.acknowledgments.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: Proposal): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.value !== undefined && (obj.value = message.value);
    if (message.acknowledgments) {
      obj.acknowledgments = message.acknowledgments.map((e) => e);
    } else {
      obj.acknowledgments = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Proposal>, I>>(object: I): Proposal {
    const message = createBaseProposal();
    message.index = object.index ?? "";
    message.value = object.value ?? "";
    message.acknowledgments = object.acknowledgments?.map((e) => e) || [];
    return message;
  },
};

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
