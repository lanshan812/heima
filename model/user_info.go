package model

type UserInfo struct {
	Id             int    `gorm:"int(60) not null; primary_key;AUTO_INCREMENT" json:"id"`
	USerId         int    `gorm:"int(254) not null " json:"user_id"`
	ReadPatentList string `gorm:"varchar(254) not null " json:"read_patent_list"`
}
