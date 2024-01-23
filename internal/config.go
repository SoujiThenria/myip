package internal

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

func ReadConfig(Path string, cc interface{}) error {
	envPath := os.Getenv("HOME")
	Path = strings.ReplaceAll(Path, "$HOME", envPath)
	fb, err := os.ReadFile(Path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(fb, cc); err != nil {
		return err
	}
	return nil
}

func BuildAddress(server string, port uint16) string {
	return server + ":" + strconv.Itoa(int(port))
}
