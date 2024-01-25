package internal

import (
	"encoding/json"
	"os"
	"strconv"
)

func ReadConfig(Path string, cc interface{}) error {
	fb, err := os.ReadFile(os.ExpandEnv(Path))
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
