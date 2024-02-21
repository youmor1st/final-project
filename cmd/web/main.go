package main

import (
	"github.com/yourusername/merit-demerit/internal/app/api"
)

func main() {
	r := api.SetupRouter()
	r.Run(":8080")
}
