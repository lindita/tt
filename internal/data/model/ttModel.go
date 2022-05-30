package model

type TtModel struct {
	BaseModel

	UserId uint64 `gorm:"column:userId;type:bigint(20) unsigned;not null;default:0"` // 用户id
	name   string `gorm:"column:alipayUserId;type:varchar(45);not null;default:''"`  //

}

func (m *TtModel) TableName() string {
	return "stt_alipay_bind_log"
}
