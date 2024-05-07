/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "kvs.kvs";

export interface Acl {
  adresses: string[];
}

function createBaseAcl(): Acl {
  return { adresses: [] };
}

export const Acl = {
  encode(message: Acl, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.adresses) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Acl {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAcl();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.adresses.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Acl {
    return { adresses: Array.isArray(object?.adresses) ? object.adresses.map((e: any) => String(e)) : [] };
  },

  toJSON(message: Acl): unknown {
    const obj: any = {};
    if (message.adresses) {
      obj.adresses = message.adresses.map((e) => e);
    } else {
      obj.adresses = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Acl>, I>>(object: I): Acl {
    const message = createBaseAcl();
    message.adresses = object.adresses?.map((e) => e) || [];
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
