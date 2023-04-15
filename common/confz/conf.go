package confz

import (
	"log"
	"os"
	"reflect"
	"zero-template/common/utils/crypt"

	"github.com/zeromicro/go-zero/core/conf"
)

type SecurityConf struct {
	Encrypted bool   `json:",optional"`
	KeyEnv    string `json:",optional"`
}

func findSecurityConfInStruct(v interface{}) (SecurityConf, bool) {
	if reflect.ValueOf(v).Kind() == reflect.Ptr {
		v = reflect.ValueOf(v).Elem().Interface()
	}
	t := reflect.TypeOf(v)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type == reflect.TypeOf(SecurityConf{}) {
			return reflect.ValueOf(v).FieldByIndex(field.Index).Interface().(SecurityConf), true
		}
	}
	return SecurityConf{}, false
}

func MustLoad(path string, v interface{}, opts ...conf.Option) {
	if err := conf.Load(path, v, opts...); err != nil {
		log.Fatalf("error: config file %s, %s", path, err.Error())
	}
	c, ok := findSecurityConfInStruct(v)
	if ok && c.Encrypted {
		key := os.Getenv(c.KeyEnv)
		decoded, err := crypt.Decode(v, key)
		if err != nil {
			log.Fatalf("error: config file %s, %s", path, err.Error())
		}
		if reflect.TypeOf(v).Kind() == reflect.Ptr {
			reflect.ValueOf(v).Elem().Set(reflect.ValueOf(decoded).Elem())
			return
		}
		reflect.ValueOf(v).Set(reflect.ValueOf(decoded))
	}
}
