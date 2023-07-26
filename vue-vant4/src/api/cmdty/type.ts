import {CmdtyInfo} from "../../models/cmdty.ts";
import {BaseResp} from "../index.ts";

export interface ListCisPageResp extends BaseResp {
    data: CmdtyInfo[]

}

export interface GetInfoResp extends BaseResp {
    data: {
        cmdtyInfo: CmdtyInfo
        isCollected: number

    }

}