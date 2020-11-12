package model

type Period struct {
	Id         int64  `db:"id" json:"id"`
	HappenTime string `db:"happen_time" json:"happenTime"`
	Type       int    `db:"type" json:"type"`
	CreateTime string `db:"create_time" json:"createTime"`
	UpdateTime string `db:"update_time" json:"updateTime"`
}
