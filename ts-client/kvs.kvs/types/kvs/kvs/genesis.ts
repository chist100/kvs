/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Data } from "./data";
import { Params } from "./params";
import { Proposal } from "./proposal";

export const protobufPackage = "kvs.kvs";

/** GenesisState defines the kvs module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  dataList: Data[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  proposalList: Proposal[];
}

function createBaseGenesisState(): GenesisState {
  return { params: undefined, dataList: [], proposalList: [] };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.dataList) {
      Data.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.proposalList) {
      Proposal.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.dataList.push(Data.decode(reader, reader.uint32()));
          break;
        case 3:
          message.proposalList.push(Proposal.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      dataList: Array.isArray(object?.dataList) ? object.dataList.map((e: any) => Data.fromJSON(e)) : [],
      proposalList: Array.isArray(object?.proposalList)
        ? object.proposalList.map((e: any) => Proposal.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.dataList) {
      obj.dataList = message.dataList.map((e) => e ? Data.toJSON(e) : undefined);
    } else {
      obj.dataList = [];
    }
    if (message.proposalList) {
      obj.proposalList = message.proposalList.map((e) => e ? Proposal.toJSON(e) : undefined);
    } else {
      obj.proposalList = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.dataList = object.dataList?.map((e) => Data.fromPartial(e)) || [];
    message.proposalList = object.proposalList?.map((e) => Proposal.fromPartial(e)) || [];
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
