import axios from 'axios'

export interface BaseResp {
    code: number
    msg: string
}

export const atclService = axios.create({
    baseURL: 'http://localhost:10001/atcl'
})

export const chatService = axios.create({
    baseURL: 'http://localhost:10002/chat'
})

export const cmdtyService = axios.create({
    baseURL: 'http://localhost:10003/cmdty'
})

