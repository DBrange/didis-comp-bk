package assets

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func StructToBsonMap(s any) (bson.M, error) {
	fmt.Printf("%v", s)
	result := bson.M{}
	v := reflect.Indirect(reflect.ValueOf(s))

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input is not a struct")
	}

	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		if !field.IsExported() || field.Name == "ID" {
			continue
		}

		key := field.Tag.Get("bson")
		if key == "" {
			key = field.Name
		} else {
			key = strings.SplitN(key, ",", 2)[0]
			if key == "-" {
				continue
			}
		}

		if fieldValue.Kind() == reflect.Ptr {
			if fieldValue.IsNil() {
				continue
			}
			fieldValue = fieldValue.Elem()
		}

		switch fieldValue.Kind() {
		case reflect.Struct:
			if fieldValue.Type() == reflect.TypeOf(time.Time{}) {
				result[key] = fieldValue.Interface()
			} else {
				nestedMap, err := StructToBsonMap(fieldValue.Interface())
				if err != nil {
					return nil, err
				}
				result[key] = nestedMap
			}
		case reflect.Slice:
			sliceLen := fieldValue.Len()
			sliceValue := make([]any, sliceLen)
			for j := 0; j < sliceLen; j++ {
				item := fieldValue.Index(j)
				if item.Kind() == reflect.Struct {
					nestedMap, err := StructToBsonMap(item.Interface())
					if err != nil {
						return nil, err
					}
					sliceValue[j] = nestedMap
				} else {
					sliceValue[j] = item.Interface()
				}
			}
			result[key] = sliceValue
		default:
			result[key] = fieldValue.Interface()
		}
	}

	return result, nil
}
