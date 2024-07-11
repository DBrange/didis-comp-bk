package assets

import (
	"fmt"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func StructToBsonMap(s interface{}) (bson.M, error) {
    result := bson.M{}
    v := reflect.ValueOf(s)

    // Si es un puntero, obtenemos el valor al que apunta
    if v.Kind() == reflect.Ptr {
        v = v.Elem()
    }

    // Verifica que la entrada sea una estructura
    if v.Kind() != reflect.Struct {
        return nil, fmt.Errorf("input is not a struct")
    }

    t := v.Type()

    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        fieldValue := v.Field(i)

        // Omite campos no exportados
        if !field.IsExported() {
            continue
        }

        // Obtiene el nombre del campo de la etiqueta json
        key := field.Tag.Get("json")
        if key == "" {
            key = field.Name // Usa el nombre del campo si no hay etiqueta json
        } else {
            // Toma solo el nombre del campo, ignorando otras opciones como omitempty
            key = strings.SplitN(key, ",", 2)[0]
            if key == "-" {
                continue // Omite este campo si la etiqueta es "-"
            }
        }

        // Omite el campo "ID"
        if field.Name == "ID" {
            continue
        }

        // Manejo de punteros
        if fieldValue.Kind() == reflect.Ptr {
            if fieldValue.IsNil() {
                continue
            }
            fieldValue = fieldValue.Elem()
        }

        // Manejo recursivo para estructuras anidadas
        if fieldValue.Kind() == reflect.Struct {
            nestedMap, err := StructToBsonMap(fieldValue.Interface())
            if err != nil {
                return nil, err
            }
            result[key] = nestedMap
        } else if fieldValue.Kind() == reflect.Slice {
            // Manejo de slices
            sliceLen := fieldValue.Len()
            sliceValue := make([]interface{}, sliceLen)
            for j := 0; j < sliceLen; j++ {
                sliceValue[j] = fieldValue.Index(j).Interface()
            }
            result[key] = sliceValue
        } else {
            // Para otros tipos, simplemente asigna el valor
            result[key] = fieldValue.Interface()
        }
    }

    return result, nil
}