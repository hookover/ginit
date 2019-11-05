package models

type User struct {
	Id        int64  `json:"id" xorm:"bigint autoincr pk comment('自增id')"`
	Email     string `json:"email" xorm:"varchar(64) not null unique('uniq_idx_email') comment('邮箱')"`
	Phone     string `json:"phone" xorm:"varchar(64) not null unique('uniq_idx_phone') comment('手机号')"`
	Name      string `json:"name" xorm:"varchar(64) not null default '' comment('姓名')"`
	Password  string `json:"password" xorm:"varchar(128) not null default '' comment('密码')"`
	CreatedAt int64  `json:"created_at" xorm:"created bigint not null default 0 comment('创建时间')"`
	UpdatedAt int64  `json:"updated_at" xorm:"updated bigint not null default 0 comment('更新时间')"`
}
