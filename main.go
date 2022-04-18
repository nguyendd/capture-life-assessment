package main

import (
	"capturelife.assessment.daniel/routers"
)

func main() {
	r := routers.InitRouter()

	r.Run("localhost:8080")
}
