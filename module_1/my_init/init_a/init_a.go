package init_a

import (
	"fmt"

	_ "github.com/liubo0127/CloudNativeLearn/module_1/my_init/init_b"
)

func init() {
	fmt.Println("Init from a")
}
