import {CmdtyInfo} from "../../models/cmdty.ts";
import {BaseResp} from "../index.ts";

export interface ListCisPageResp extends BaseResp {
    data: {
        cis: CmdtyInfo[]
    }
}