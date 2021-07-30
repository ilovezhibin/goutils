package main

import (
	"fmt"
	"go-utils/httputils"
)

func main() {
	fmt.Println(httputils.Get("http://s-dev-backend.sysop.yy.com/web/api/getAllWebAreaType"))
}
