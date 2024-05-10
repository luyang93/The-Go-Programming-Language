// Package params provides a reflection-based parser for URL parameters.
package params

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func Pack(ptr interface{}) (url.URL, error) {
	v := reflect.ValueOf(ptr).Elem()
	if v.Type().Kind() != reflect.Struct {
		return url.URL{}, fmt.Errorf("Pack: %v is not a struct.", ptr)
	}

	vals := &url.Values{}
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		vals.Add(name, fmt.Sprintf("%v", v.Field(i)))
	}
	return url.URL{RawQuery: vals.Encode()}, nil
}

type Check func(v interface{}) error

// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			if err := checks(name, value); err != nil {
				return err
			}
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func checks(name string, v string) error {
	switch name {
	case "p":
		if err := lt100(v); err != nil {
			return err
		}
	default:
		panic(fmt.Sprintf("unsupported type: %s", name))
	}
	return nil
}

func lt100(v interface{}) error {
	i, err := strconv.ParseInt(v.(string), 10, 64)
	if err != nil {
		return fmt.Errorf("not an int: %v", v)
	}
	if i >= 100 {
		return fmt.Errorf("%d >= 100", i)
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}