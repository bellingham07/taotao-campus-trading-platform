export const SET_TOKEN = (token: string) => {
    localStorage.setItem('Authorization', token)
}
//本地存储获取数据
export const GET_TOKEN = () => {
    return localStorage.getItem('Authorization')
}
//本地存储删除数据方法
export const REMOVE_TOKEN = () => {
    localStorage.removeItem('Authorization')
}
