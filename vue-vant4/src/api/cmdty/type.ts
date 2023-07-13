import {cmdtyInfo} from "../../models/cmdty.ts";
import {BaseResp} from "../index.ts";

export interface newestResp extends BaseResp {
    data: {
        sell: cmdtyInfo[]
        want: cmdtyInfo[]
    }
}