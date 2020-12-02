package main

import (
	"fmt"
	"user_service/service"
)

func main() {
	userInfo, err := service.GetUserInfo(0)
	if err != nil {
		panic(err)
	}

	fmt.Printf("User ID: %d \n", userInfo.ID)
}
