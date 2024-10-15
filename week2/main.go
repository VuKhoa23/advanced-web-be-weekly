package main

import (
	_ "github.com/VuKhoa23/advanced-web-be/docs"
	_ "github.com/VuKhoa23/advanced-web-be/internal/utils/validator"
	"github.com/VuKhoa23/advanced-web-be/startup"
)

// @title           API for Advanced Web
// @version         1.0
// @description     API for Advanced Web
// @BasePath  /api/v1
func main() {
	startup.Execute()
}
