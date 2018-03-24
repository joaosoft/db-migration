package gomapper

import (
	"fmt"
	"reflect"
	"strings"
)

// Map ...
func (mapper *GoMapper) Map(value interface{}) (map[string]interface{}, error) {
	mapping := make(map[string]interface{})

	if err := convertToMap(value, "", mapping); err != nil {
		return nil, err
	}

	return mapping, nil
}

func convertToMap(obj interface{}, path string, mapping map[string]interface{}) error {
	types := reflect.TypeOf(obj)
	value := reflect.ValueOf(obj)

	if value.Kind() == reflect.Ptr {
		value = reflect.ValueOf(value).Elem()

		if value.IsValid() {
			types = value.Type()
		} else {
			return nil
		}
	}

	switch value.Kind() {
	case reflect.Struct:
		path = addPoint(path)
		for i := 0; i < types.NumField(); i++ {
			nextValue := value.Field(i)
			newPath := fmt.Sprintf("%s%s", path, strings.ToLower(types.Field(i).Name))
			convertToMap(nextValue.Interface(), newPath, mapping)
		}

	case reflect.Array, reflect.Slice:
		path = addPoint(path)
		for i := 0; i < value.Len(); i++ {
			nextValue := value.Index(i)
			newPath := fmt.Sprintf("%s[%d]", path, i)
			convertToMap(nextValue.Interface(), newPath, mapping)
		}

	case reflect.Map:
		path = addPoint(path)
		for _, key := range value.MapKeys() {
			nextValue := value.MapIndex(key)
			newPath := fmt.Sprintf("%s{%+v}", path, key)
			convertToMap(nextValue.Interface(), newPath, mapping)
		}

	default:
		if value.CanInterface() {
			mapping[path] = value.Interface()
			log.Debugf(fmt.Sprintf("%s=%+v", path, value.Interface()))
		} else {
			mapping[path] = value
			log.Debugf(fmt.Sprintf("%s=%+v", path, value))
		}

	}
	return nil
}

func addPoint(path string) string {
	if path != "" {
		path += "."
	}
	return path
}
