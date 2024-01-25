package cli

import (
	"bytes"
	"fmt"
	"strings"
)

func commandCreate(cfg *ApiConfig, s string) error {
	splitString := strings.Split(s, " ")
	if splitString[0] == "user" {
		if len(splitString) != 3 {

		}

		jsonBody := []byte(fmt.Sprintf(`{"name": "%s", "password": "%s"}`, splitString[1], splitString[2]))
		bodyReader := bytes.NewReader(jsonBody)
		cfg.ApiClient.HttpClient.Post(cfg.ApiClient.BaseURL, "application/json", bodyReader)
	}

	return nil
}
