import {CmdtyInfo, CmdtyPics} from "../../models/cmdty.ts";
import {BaseResp} from "../index.ts";
import {UserInfo} from "../../models/user.ts";

export interface ListCisPageResp extends BaseResp {
    data: CmdtyInfo[]

}

export interface GetInfoResp extends BaseResp {
    data: {
        cmdtyInfo: CmdtyInfo
        cmdtyPics: CmdtyPics[]
        userInfo: UserInfo
        isCollected: number
        isMine: number
    }
}