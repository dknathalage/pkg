package app

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// LoadEnvWithPrefix loads environment variables with the given prefix into the provided config struct
func LoadEnvWithPrefix(prefix string, config interface{}) {
	rv := reflect.ValueOf(config).Elem()
	rt := rv.Type()

	for i := 0; i < rv.NumField(); i++ {
		field := rt.Field(i)
		envVar := prefix + "_" + strings.ToUpper(field.Name)

		value, exists := os.LookupEnv(envVar)
		if !exists {
			fmt.Printf("Skipping %s: environment variable not set\n", envVar)
			continue
		}

		switch field.Type.Kind() {
		case reflect.String:
			rv.Field(i).SetString(value)
		case reflect.Int, reflect.Int64:
			intValue, err := strconv.Atoi(value)
			if err != nil {
				fmt.Printf("Invalid int value for %s: %s\n", envVar, value)
				continue
			}
			rv.Field(i).SetInt(int64(intValue))
		case reflect.Bool:
			rv.Field(i).SetBool(strings.ToLower(value) == "true")
		default:
			fmt.Printf("Skipping unsupported type for %s\n", envVar)
		}
	}
}
