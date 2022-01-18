package conf

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	err := ReadConf("../config/application.toml")
	fmt.Println(err)

}
