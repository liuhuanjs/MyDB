package printer

import "fmt"

func GetMyDBInfo() string {
	return fmt.Sprintf("version : %s", "v0.0.1-alpha")
}
