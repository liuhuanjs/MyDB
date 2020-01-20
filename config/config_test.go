package config

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestStructTag(t *testing.T) {
	config := Config{"1.1.1.1", "2.2.2.2", 3306}
	jsonBytes, _ := json.Marshal(config)
	fmt.Println(string(jsonBytes))
}
