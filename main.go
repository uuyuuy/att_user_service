package main

import (
	"github.com/uuyuuy/att_user_service.git/handler"
)

func main() {
	r := handler.Handle()
	r.Run()
}
