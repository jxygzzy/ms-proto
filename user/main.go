package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"ms-proto/service"
)

func main() {
	user := &service.User{
		Username: "zzy",
		Age:      18,
	}
	marshal, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}

	newUser := &service.User{}
	err = proto.Unmarshal(marshal, newUser)
	if err != nil {
		panic(err)
	}

	fmt.Println(newUser.String())

}
