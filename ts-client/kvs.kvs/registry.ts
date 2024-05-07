import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgDataProposal } from "./types/kvs/kvs/tx";
import { MsgDataConfirmation } from "./types/kvs/kvs/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/kvs.kvs.MsgDataProposal", MsgDataProposal],
    ["/kvs.kvs.MsgDataConfirmation", MsgDataConfirmation],
    
];

export { msgTypes }