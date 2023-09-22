package utils

import (
	"errors"
	"reflect"
	"regexp"
	"strings"
)

type Rules map[string][]string

type RulesMap map[string]Rules

var CustomizeMap = make(map[string]Rules)

//@author: [123]()
//@function: RegisterRule
//@description: 注册自定义规则方案建议在路由初始化层即注册

//@author: []()
//@function: Verify
//@description: 检验方法
//@param: st interface{}, roleMap Rules(入参实例, 规则Map)
//@return: err error

func Verify(st interface{}, roleMap Rules) (err error) {
	compareMap := map[string]bool{
		"lt": true,
		"le": true,
		"eq": true,
		"ne": true,
		"ge": true,
		"gt": true,
	}

	typ := reflect.TypeOf(st)
	val := reflect.ValueOf(st) // 获取refect.Type 类型

	kd := val.Kind() // 获取到st对应的类型
	if kd != reflect.Struct {
		return errors.New("expect struct")
	}
	num := val.NumField()
	// 遍历结构体的所有字段
	for i := 0; i < num; i++ {
		tagVal := typ.Field(i)
		val := val.Field(i)
		if tagVal.Type.Kind() == reflect.Struct {
			if err = Verify(val.Interface(), roleMap); err != nil {
				return err
			}
		}
		if len(roleMap[tagVal.Name]) > 0 {
			for _, v := range roleMap[tagVal.Name] {
				switch {
				case v == "notEmpty":
					if isBlank(val) {
						return errors.New(tagVal.Name + "值不能为空")
					}
				case strings.Split(v, "=")[0] == "regexp":
					if !regexMatch(strings.Split(v, "=")[1], val.String()) {
						return errors.New(tagVal.Name + "校验格式不通过")
					}
				case compareMap[strings.Split(v, "=")[0]]:
					if !compareVerify(val, v) {
						return errors.New(tagVal.Name + "长度或值不在合法范围, " + v)
					}
				}
			}
		}
	}
}

//@author: [123]()
//@function: compareVerify
//@description: 长度和数字的检验方法 根据类型自动校验
//@param: value reflect.Value, VerifyStr string
//@return: bool

func compareVerify(value reflect.Value, VerifyStr string) bool {
	switch value.Kind() {
	case reflect.String:
		return compare(len())
	}
}

//@author: [123]()
//@function: isBlank
//@description: 非空校验
//@param: value relect.Value
//@return: bool

func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String, reflect.Slice:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

func regexMatch(rule, matchStr string) bool {
	return regexp.MustCompile(rule).MatchString(matchStr)
}
