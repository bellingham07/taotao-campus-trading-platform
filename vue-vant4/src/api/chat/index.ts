import {chatService} from "../index.ts";
import {CreateRoomReq, CreateRoomResp} from "./type.ts";

export const createRoom = (req: CreateRoomReq) => {
    return chatService.post<any, CreateRoomResp>('/room', req)
}