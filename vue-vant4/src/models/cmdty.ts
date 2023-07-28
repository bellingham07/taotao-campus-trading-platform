export interface CmdtyInfo {
    id: number
    userId: number
    briefIntro: string
    cover: string
    tag: string
    price: number
    brand: string
    model: string
    intro: string
    old: string
    status: number
    createAt: string
    publishAt: string
    view: number
    collect: number
    type: number
    like: number
}

export interface CmdtyPics {
    id: number;
    cmdty_id: number;
    objectname: string;
    order: number;
    url: string;
    user_id: number;
}