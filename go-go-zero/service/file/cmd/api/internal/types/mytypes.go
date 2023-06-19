package types

import "mime/multipart"

type PicReq struct {
	Pic   multipart.FileHeader `json:"pic"`
	Order int64                `json:"order"`
}

type AtclPicsReq struct {
	AtclId int64    `json:"cmdtyId"`
	Pics   []PicReq `json:"pics"`
}

type CmdtyPicsReq struct {
	CmdtyId int64    `json:"cmdtyId"`
	Pics    []PicReq `json:"pics"`
}
