package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// encode writes to buf an S-expression representation of v.
func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

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
		return encode(buf, v.Elem())

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(')')

	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}

			fieldInfo := v.Type().Field(i)
			tag := fieldInfo.Tag
			name := tag.Get("sexpr")
			if name == "" {
				name = fieldInfo.Name
			}

			if _, err := fmt.Fprintf(buf, "(%s ", name); err != nil {
				return err
			}
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	case reflect.Map: // ((key value) ...)
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteByte('(')
			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	case reflect.Float32:
		if _, err := fmt.Fprintf(buf, "%s", strconv.FormatFloat(v.Float(), 'f', -1, 32)); err != nil {
			return err
		}
	case reflect.Float64:
		if _, err := fmt.Fprintf(buf, "%s", strconv.FormatFloat(v.Float(), 'f', -1, 64)); err != nil {
			return err
		}

	case reflect.Complex64:
		if _, err := fmt.Fprintf(
			buf,
			"#C(%s %s)",
			strconv.FormatFloat(real(v.Complex()), 'f', -1, 32),
			strconv.FormatFloat(imag(v.Complex()), 'f', -1, 32),
		); err != nil {
			return err
		}

	case reflect.Complex128:
		if _, err := fmt.Fprintf(
			buf,
			"#C(%s %s)",
			strconv.FormatFloat(real(v.Complex()), 'f', -1, 64),
			strconv.FormatFloat(imag(v.Complex()), 'f', -1, 64),
		); err != nil {
			return err
		}

	case reflect.Bool:
		if v.Bool() {
			buf.WriteByte('t')
		} else {
			buf.WriteString("nil")
		}

	case reflect.Interface:
		if _, err := fmt.Fprintf(buf, "(%q ", reflect.Indirect(v).Type()); err != nil {
			return err
		}
		if err := encode(buf, reflect.Indirect(v).Elem()); err != nil {
			return err
		}
		buf.WriteByte(')')
	default: // chan, func
		return fmt.Errorf("unsupported type: %s", v.Type())
	}

	return nil
}
