import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgDataConfirmation } from "./types/kvs/kvs/tx";
import { MsgAddressRegistration } from "./types/kvs/kvs/tx";
import { MsgDataProposal } from "./types/kvs/kvs/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/kvs.kvs.MsgDataConfirmation", MsgDataConfirmation],
    ["/kvs.kvs.MsgAddressRegistration", MsgAddressRegistration],
    ["/kvs.kvs.MsgDataProposal", MsgDataProposal],
    
];

export { msgTypes }