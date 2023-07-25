import {cmdtyService} from "../index.ts";
import {GetInfoResp, ListCisPageResp} from "./type.ts";

export const listCisPageReq = (type: number, page: number) => {
    return cmdtyService.get<any, ListCisPageResp>("/cache", {
        params: {
            type: type,
            page: page,
        }
    })
}

export const getInfoReq = (id: number) => {
    return cmdtyService.get<any, GetInfoResp>(`/${id}`)
}
