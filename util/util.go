package util

import (
  "reflect"
  "runtime"
  "strings"
)

/*
public function for stringifying a function name. In this repo it is used as to get the name of the function passed
as the argument "statistic" to the two Boot functions (BootVector, BootMatrix) defined in the boot package

params:
function - an interace is used to genericize the function, but the value should be a pointer to a function

returns: the name of the function as a string
*/
func GetFunctionName(function interface{}) string {
  strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()), ".")
  return strs[len(strs)-1]
}
