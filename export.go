package shop

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"reflect"
)

func (accountMutex AccountMutex) ExportAccountsCSV() ([]byte, error) {
	export := make(map[interface{}]interface{})
	for acc := range accountMutex.Account {
		export[acc] = accountMutex.Account[acc]
	}
	buffer := bytes.Buffer{}
	writer := csv.NewWriter(&buffer)
	for acc := range export {
		var record []string
		value := reflect.ValueOf(export[acc])
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i).Interface()
			var rec string
			switch field.(type) {
			case float64, float32:
				rec = fmt.Sprintf("%f", field)
			default:
				rec = fmt.Sprintf("%v", field)
			}
			record = append(record, rec)
		}

		err := writer.Write(record)
		if err != nil {
			return nil, errors.New("not write")
		}
	}
	writer.Flush()
	return buffer.Bytes(), writer.Error()

}

func (productMutex ProductMutex) ExportProductsCSV() ([]byte, error) {
	export := make(map[interface{}]interface{})
	for prd := range productMutex.Product {
		export[prd] = productMutex.Product[prd]
	}
	buffer := bytes.Buffer{}
	writer := csv.NewWriter(&buffer)
	for acc := range export {
		var record []string
		value := reflect.ValueOf(export[acc])
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i).Interface()
			var rec string
			switch field.(type) {
			case float64, float32:
				rec = fmt.Sprintf("%f", field)
			default:
				rec = fmt.Sprintf("%v", field)
			}

			record = append(record, rec)
		}

		err := writer.Write(record)
		if err != nil {
			return nil, errors.New("not write")
		}
	}
	writer.Flush()
	return buffer.Bytes(), writer.Error()
}
