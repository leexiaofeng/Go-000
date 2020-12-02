package controller

import (
	"fmt"

	"github.com/GO-000/week02/dao"
)

// GetAllUser is get all user from db
func GetAllUser() {
	fmt.Println("ok")

	u, err := dao.GetUser()
	if err != nil {
		fmt.Printf("GetAllUser %+v\n", err)
		return
	}

	fmt.Println(u)
}
