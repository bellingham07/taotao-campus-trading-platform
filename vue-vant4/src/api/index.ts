import axios, {AxiosError, AxiosResponse} from 'axios'

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

const respInterceptor = (config: AxiosResponse) => {
    const code = config.data['code'] || 200;
    if (code == 200) {
        return Promise.resolve(config.data)
    }
    return Promise.reject(config.data)
}

const error = (error: AxiosError) => {
    console.log(error)
}

cmdtyService.interceptors.response.use(respInterceptor, error)

