package entity

type UserEntity struct {
	ID        int64  `json:"id"`
	Phone     string `json:"phone"`
	CreatedAt int64  `json:"created_at"`
}

func (UserEntity) TableName() string {
	return "users"
}
