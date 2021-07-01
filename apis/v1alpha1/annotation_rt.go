package v1alpha1

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

const (
	conversionAnnotationField = "ack.aws.dev/"
)

// returns annotationName, annotationValue and an error
// TODO(a-hilaly) care about annotation key size limit (maybe hash the keys? and store a hashmap as annotation value?)
// TODO(a-hilaly) move this utility to runtime
func AnnotateFieldData(fieldName string, data interface{}) (string, string, error) {
	annotationKey := conversionAnnotationField + fieldName
	annotationValue := ""
	switch data.(type) {
	case *string:
		annotationValue = "string=" + *data.(*string)
	case *int:
		annotationValue = "int=" + strconv.Itoa(*data.(*int))
	case *int8:
	case *int16:
	case *int32:
	case *int64:
	default:
		bytes, err := json.Marshal(data)
		if err != nil {
			return "", "", err
		}
		annotationValue = "json=" + string(bytes)
	}

	return annotationKey, annotationValue, nil
}

func DecodeFieldDataAnnotation(annotationValue string, unmarshallTo interface{}) error {
	/* 	if !strings.HasPrefix(annotationKey, conversionAnnotationField) {
		return fmt.Errorf("not a conversion annotation")
	} */

	parts := strings.Split(annotationValue, "=")
	vType := parts[0]
	value := parts[1]
	switch vType {
	case "string":
		unmarshallTo = value
	case "int":
		unmarshallTo = value
	case "int8":
	case "int16":
	case "int32":
	case "int64":
	case "json":
		valueBytes := []byte(value)
		err := json.Unmarshal(valueBytes, unmarshallTo)
		if err != nil {
			return fmt.Errorf("unmarshalling value of type %s: %v", vType, err)
		}
	default:
		return fmt.Errorf("unsupported type")
	}

	return nil
}
