package main

import (
	"context"
	"fmt"

	"myzgw/gen/examples/biz"
	"myzgw/gen/examples/conf"
	"myzgw/gen/examples/dal"
	"myzgw/gen/examples/dal/query"
)

func init() {
	dal.DB = dal.ConnectDB(conf.MySQLDSN).Debug()
}

func main() {
	// start your project here
	fmt.Println("hello world")
	defer fmt.Println("bye~")

	query.SetDefault(dal.DB)
	biz.Query(context.Background())
}
