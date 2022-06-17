package models

type Posts struct {
	ID          uint64 `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id" form:"id"`
	PostNo      string `gorm:"column:post_no;NOT NULL" json:"post_no" form:"post_no"`                                   // 唯一编码
	Title       string `gorm:"column:title;NOT NULL" json:"title" form:"title"`                                         // 文章标题
	Description string `gorm:"column:description;NOT NULL" json:"description" form:"description"`                       // 文章描述
	Author      string `gorm:"column:author;NOT NULL" json:"author" form:"author"`                                      // 作者
	SourceSite  string `gorm:"column:source_site;NOT NULL" json:"source_site" form:"source_site"`                       // 来源网站
	SourceUrl   string `gorm:"column:source_url;NOT NULL" json:"source_url" form:"source_url"`                          // 来源网址
	PublishTime int    `gorm:"column:publish_time;default:0;NOT NULL" json:"publish_time" form:"publish_time"`          // 前端展示时间
	Status      int    `gorm:"column:status;default:0;NOT NULL" json:"status" form:"status"`                            // 状态，1是正常，0删除
	CreatedAt   int64  `gorm:"autoCreateTime;column:created_at;default:0;NOT NULL" json:"created_at" form:"created_at"` // 时间
	UpdatedAt   int64  `gorm:"autoUpdateTime;column:updated_at;default:0;NOT NULL" json:"updated_at" form:"updated_at"` // 更新时间
}

func (m *Posts) TableName() string {
	return "posts"
}
