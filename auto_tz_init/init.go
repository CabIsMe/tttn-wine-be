package tzinit

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("Init timezone")
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")
}
