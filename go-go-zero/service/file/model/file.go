package model

import "time"

type (
	FileAtcl struct {
		Id         int64     `xorm:"id" json:"id"` // bigint自增
		AtclId     int64     `xorm:"atcl_id" json:"atclId"`
		UserId     int64     `xorm:"user_id" json:"userId"`
		Url        string    `xorm:"url" json:"url"`
		Objectname string    `xorm:"objectname" json:"objectname"`
		UploadAt   time.Time `xorm:"upload_at" json:"uploadAt"`
		IsCover    int64     `xorm:"is_cover" json:"isCover"` // 默认为0，封面为1
		Order      int64     `xorm:"order" json:"order"`
	}

	FileAvatar struct {
		Id         int64     `xorm:"id" json:"id"`
		UserId     int64     `xorm:"user_id" json:"userId"`
		Url        string    `xorm:"url" json:"url"`
		Objectname string    `xorm:"objectname" json:"objectname"`
		UploadAt   time.Time `xorm:"upload_at" json:"uploadAt"`
	}

	FileCmdty struct {
		Id         int64     `xorm:"id" json:"id"` // bigint自增
		CmdtyId    int64     `xorm:"cmdty_id" json:"cmdtyId"`
		UserId     int64     `xorm:"user_id" json:"userId"`
		Url        string    `xorm:"url" json:"url"`
		Objectname string    `xorm:"objectname" json:"objectname"`
		UploadAt   time.Time `xorm:"upload_at" json:"uploadAt"`
		IsCover    int64     `xorm:"is_cover" json:"isCover"` // 默认为0，封面为1
		Order      int64     `xorm:"order" json:"order"`
	}
)
