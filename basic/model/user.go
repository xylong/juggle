package model

// User 用户
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	NickName string `json:"nick_name" binding:"required,nefield=Name"`
	Password string `json:"password" binding:"required"`
	Birthday string `json:"birthday" binding:"omitempty,birthday"`
	Gender   int    `json:"gender" binding:"omitempty,min=0,max=1"`
	Phone    string `json:"phone"`
	Email    string `json:"email" binding:"omitempty,email"`
}

type Users struct {
	List []User `json:"list" binding:"gt=0,lt=100,size,dive"` // 切片长度>0，<100,dive验证每一项
	Size int    `json:"size"`
}

func NewUser(ID int, name string) *User {
	return &User{ID: ID, Name: name}
}

type UserQuery struct {
	Name     string `json:"name" form:"name"`
	Page     int    `json:"page" form:"page" binding:"required"`
	PageSize int    `json:"page_size" form:"page_size" binding:"required"`
}
