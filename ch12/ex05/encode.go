package json

import (
	"bytes"
	"fmt"
	"reflect"
)

// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// encode writes to buf an S-expression representation of v.
func encode(buf *bytes.Buffer, v reflect.Value, indent int) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("null")

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if _, err := fmt.Fprintf(buf, "%d", v.Int()); err != nil {
			return err
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if _, err := fmt.Fprintf(buf, "%d", v.Uint()); err != nil {
			return err
		}

	case reflect.String:
		if _, err := fmt.Fprintf(buf, "%q", v.String()); err != nil {
			return err
		}

	case reflect.Ptr:
		return encode(buf, v.Elem(), indent)

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('[')
		indent++
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				if _, err := fmt.Fprintf(buf, "\n%*s", indent, ""); err != nil {
					return err
				}
			}
			if err := encode(buf, v.Index(i), indent); err != nil {
				return err
			}
			if i != v.Len()-1 {
				buf.WriteByte(',')
			}
		}
		buf.WriteByte(']')

	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('{')
		indent++
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				if _, err := fmt.Fprintf(buf, "\n%*s", indent, ""); err != nil {
					return err
				}
			}
			start := buf.Len()
			if _, err := fmt.Fprintf(buf, "%q: ", v.Type().Field(i).Name); err != nil {
				return err
			}
			if err := encode(buf, v.Field(i), indent+buf.Len()-start); err != nil {
				return err
			}
			if i != v.NumField()-1 {
				buf.WriteByte(',')
			}
		}
		buf.WriteByte('}')

	case reflect.Map: // ((key value) ...)
		buf.WriteByte('{')
		indent++
		for i, key := range v.MapKeys() {
			if i > 0 {
				if _, err := fmt.Fprintf(buf, "\n%*s", indent, ""); err != nil {
					return err
				}
			}
			start := buf.Len()
			if err := encode(buf, key, 0); err != nil {
				return err
			}
			buf.WriteString(": ")
			if err := encode(buf, v.MapIndex(key), indent+buf.Len()-start); err != nil {
				return err
			}
			if i != len(v.MapKeys())-1 {
				buf.WriteByte(',')
			}
		}
		buf.WriteByte('}')

	case reflect.Float32, reflect.Float64:
		if _, err := fmt.Fprintf(buf, "%g", v.Float()); err != nil {
			return err
		}

	case reflect.Bool:
		if v.Bool() {
			buf.WriteString("true")
		} else {
			buf.WriteString("false")
		}

	default: // chan, func, interface, complex
		return fmt.Errorf("unsupported type: %s", v.Type())
	}

	return nil
}
