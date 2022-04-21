package rawmessage

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalRawMessage(t json.RawMessage) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		s, _ := t.MarshalJSON()
		_, _ = io.WriteString(w, string(s))
	})
}

func UnmarshalRawMessage(v interface{}) (json.RawMessage, error) {
	msg, ok := v.(json.RawMessage)
	if !ok {
		return nil, fmt.Errorf("casting to raw message")
	}

	return msg, nil
}
