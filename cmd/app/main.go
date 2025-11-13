package main

import (
	"dang.z.v.task/internal/config"
)

func main() {
	cfg := config.MustLoad()
	_ = cfg
}
