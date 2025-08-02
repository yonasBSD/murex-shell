package string

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func marshal(_ *lang.Process, iface any) (b []byte, err error) {
	qs := make(url.Values)

	switch v := iface.(type) {
	case []string:
		for i := range v {
			qs.Add(strconv.Itoa(i), v[i])
		}
		b = []byte(qs.Encode())

	case []any:
		for i := range v {
			t, err := types.ConvertGoType(v[i], types.String)
			if err != nil {
				t = fmt.Sprint(v[i])
			}
			qs.Add(strconv.Itoa(i), t.(string))
		}
		b = []byte(qs.Encode())

	case map[string]string:
		for s := range v {
			qs.Add(s, v[s])
		}
		b = []byte(qs.Encode())

	case map[string]any:
		for s := range v {
			t, err := types.ConvertGoType(v[s], types.String)
			if err != nil {
				t = fmt.Sprint(v[s])
			}
			qs.Add(s, t.(string))
		}
		b = []byte(qs.Encode())

	case map[any]any:
		for s := range v {
			t1, err := types.ConvertGoType(s, types.String)
			if err != nil {
				t1 = fmt.Sprint(s)
			}
			t2, err := types.ConvertGoType(v[s], types.String)
			if err != nil {
				t1 = fmt.Sprint(v[s])
			}
			qs.Add(t1.(string), t2.(string))
		}
		b = []byte(qs.Encode())

	case map[any]string:
		for s := range v {
			t, err := types.ConvertGoType(s, types.String)
			if err != nil {
				t = fmt.Sprint(s)
			}
			qs.Add(t.(string), v[s])
		}
		b = []byte(qs.Encode())

	case any:
		qs.Add(fmt.Sprint(v), "")

	default:
		err = errors.New("I don't know how to marshal that data into a `" + dataType + "`. Data possibly too complex?")
	}

	return
}

func unmarshal(p *lang.Process) (any, error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(b) == 0 {
		return nil, nil
	}

	if b[0] == '?' {
		if len(b) == 1 {
			return nil, nil
		}
		b = b[1:]
	}

	values, err := url.ParseQuery(string(b))
	if err != nil {
		return nil, err
	}

	qs := make(map[string]any)
	for s := range values {
		if len(values[s]) == 1 {
			float, tnErr := toNumber(values[s][0])
			if tnErr != nil {
				qs[s] = values[s][0]
				continue
			}
			qs[s] = float

		} else {
			qs[s] = values[s]
		}
	}

	return qs, nil
}

func toNumber(s string) (f float64, err error) {
	f, err = strconv.ParseFloat(s, 64)
	if err != nil {
		return
	}

	if s != strconv.FormatFloat(f, 'f', -1, 64) {
		err = errors.New("Input doesn't match converted output. Possibly due to padding?")
	}

	return
}
