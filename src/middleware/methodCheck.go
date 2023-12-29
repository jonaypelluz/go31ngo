package middleware

import (
	"encoding/json"
	"fmt"
	"go31ngo/src/utils"
	"net/http"
	"reflect"
)

func RequestCheck[T any](method string, next func(http.ResponseWriter, *http.Request, T, utils.DBService), dbService utils.DBService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != method {
			utils.SendApiResponse(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var requestData T
		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			utils.SendApiResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := validateStruct(requestData); err != nil {
			utils.SendApiResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		next(w, r, requestData, dbService)
	})
}

func validateStruct[T any](s T) error {
	val := reflect.ValueOf(s)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if isEmptyValue(field) {
			return fmt.Errorf("missing or invalid field: %s", val.Type().Field(i).Name)
		}
	}
	return nil
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String, reflect.Array, reflect.Map, reflect.Slice:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
