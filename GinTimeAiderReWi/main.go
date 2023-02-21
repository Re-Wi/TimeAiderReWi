package main

import (
	"fmt"
	"os"
)

var (
	configFile string
)

func init() {
	configFile = "GinTimeAiderReWi_CONFIG"
}

func getEnvStr(env string, defaultValue string) string {
	v := os.Getenv(env)
	if v == "" {
		return defaultValue
	}
	return v
}

func main() {
	fmt.Printf("GinTimeAiderReWi 's GinTimeAiderReWi_CONFIG --> %s", configFile)
}
