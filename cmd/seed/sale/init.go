package sale

import (
	"os"
	"strconv"
)

var interval int

func Init() {
	v := os.Getenv("SEED_INTERVAL_MS")
	cv, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	interval = cv
	_ = interval
}
