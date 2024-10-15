package main

import (
	_ "github.com/VuKhoa23/advanced-web-be/docs"
	"github.com/VuKhoa23/advanced-web-be/startup"
)

// @title           API for Advanced Web
// @version         1.0
// @description     API for Advanced Web
// @BasePath  /api/v1
func main() {
	startup.Execute()
}
