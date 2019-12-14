package cli

import (
	"fmt"
	"reflect"
	"strings"
)

func ExecCommand(handler interface{}, cmd string) error {
	if cmd == "" {
		return nil
	}
	method := reflect.ValueOf(handler).MethodByName(fmt.Sprintf("Cmd%s", strings.ToUpper(cmd[:1])+cmd[1:]))
	if method.IsValid() {
		return method.Interface().(func() error)()
	}
	return fmt.Errorf("command not found.[cmd]=%s", cmd)
}
