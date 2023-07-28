export interface CreateRoomReq {
    cmdtyId: number
    sellerId: number
    seller: string
    buyerId: number
    buyer: string
    cover: string
}

export interface CreateRoomResp {
    cmdtyId: number
    ownerId: number
    ownerName: string
    cover: string
    type: number
}