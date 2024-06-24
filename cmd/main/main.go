package main

import (
	"GoScan/internal/utils"
)

func main() {
	utils.ScanForDirectories("http://127.0.0.1:8080", "../../wordlists/common.txt")
}