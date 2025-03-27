package common

import (
	"encoding/json"
	"fmt"
	"os"
)

const emptyString = ""

func ReadAndParseConfig() (string, error) {
	data, err := os.ReadFile("config.json")
	if err != nil {
		return emptyString, fmt.Errorf("error reading config: %v", err)
	}

	var config map[string]interface{}
	if err := json.Unmarshal(data, &config); err != nil {
		return emptyString, fmt.Errorf("error parsing config: %v", err)
	}

	prettyJSON, err := json.MarshalIndent(config, emptyString, "    ")
	if err != nil {
		return emptyString, fmt.Errorf("error formatting config: %v", err)
	}

	return string(prettyJSON), nil
}
