package model

import (
	"fmt"

	"github.com/rahimuj570/go_http_jwt_role_based_practice/enums"
)

type User struct {
	Id   int        `json:"id"`
	Role enums.Role `json:"role"`
}

func (u User) String() string {
	return fmt.Sprintf("id=%v\trole=%v", u.Id, u.Role)
}
