package main

import (
	"fmt"
	"light-ms/order/internal/config"
)

func main() {
	conf := config.LoadConf()

	fmt.Println(conf.DbUser)
}
