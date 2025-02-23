package app

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// LoadEnvWithPrefix loads environment variables with the given prefix into the provided config struct.
func LoadEnvWithPrefix(prefix string, config interface{}) error {
	rv := reflect.ValueOf(config)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("config must be a pointer to a struct")
	}

	rv = rv.Elem()
	rt := rv.Type()

	for i := 0; i < rv.NumField(); i++ {
		field := rt.Field(i)
		if !rv.Field(i).CanSet() {
			continue // Skip unexported fields
		}

		envVar := prefix + "_" + strings.ToUpper(field.Name)
		value, exists := os.LookupEnv(envVar)
		if !exists {
			continue // Skip if the environment variable is not set
		}

		switch field.Type.Kind() {
		case reflect.String:
			rv.Field(i).SetString(value)
		case reflect.Int, reflect.Int64:
			intValue, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid int value for %s: %s", envVar, value)
			}
			rv.Field(i).SetInt(intValue)
		case reflect.Bool:
			boolValue := strings.EqualFold(value, "true")
			rv.Field(i).SetBool(boolValue)
		default:
			return fmt.Errorf("unsupported field type for %s", envVar)
		}
	}

	return nil
}
