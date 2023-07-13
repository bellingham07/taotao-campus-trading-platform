import axios from 'axios'

export const atclService = axios.create({
    baseURL: 'http://localhost:11001'
})

export const chatService = axios.create({
    baseURL: 'http://localhost:11002'
})

export const cmdtyService = axios.create({
    baseURL: 'http://localhost:11003'
})

