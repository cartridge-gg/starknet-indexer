package big

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"io"
	basebig "math/big"
	"strconv"
	"strings"

	"entgo.io/ent/dialect"
)

var IntSchemaType = map[string]string{
	dialect.Postgres: "numeric",
}

type Int struct {
	*basebig.Int
}

func NewInt(i int64) Int {
	return Int{Int: basebig.NewInt(i)}
}

func FromBase(i *basebig.Int) Int {
	return Int{Int: i}
}

func (b *Int) Scan(src interface{}) error {
	var i sql.NullString
	if err := i.Scan(src); err != nil {
		return err
	}
	if !i.Valid {
		return nil
	}
	if b.Int == nil {
		b.Int = basebig.NewInt(0)
	}
	// Value came in a floating point format.
	if strings.ContainsAny(i.String, ".+e") {
		f := basebig.NewFloat(0)
		if _, err := fmt.Sscan(i.String, f); err != nil {
			return err
		}
		b.Int, _ = f.Int(b.Int)
	} else if _, err := fmt.Sscan(i.String, b.Int); err != nil {
		return err
	}
	return nil
}

func (b Int) Value() (driver.Value, error) {
	if b.Int == nil {
		return "", nil
	}
	return b.String(), nil
}

func (b Int) Add(c Int) Int {
	b.Int = b.Int.Add(b.Int, c.Int)
	return b
}

func (b Int) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, b.String())), nil
}

func (b *Int) UnmarshalJSON(p []byte) error {
	if string(p) == "null" {
		return nil
	}

	b.Int = new(basebig.Int)

	// Ints are represented as strings in JSON. We
	// remove the enclosing quotes to provide a plain
	// string number to SetString.
	s := string(p[1 : len(p)-1])
	if i, _ := b.Int.SetString(s, 10); i == nil {
		return fmt.Errorf("unmarshalling big int: %s", string(p))
	}

	return nil
}

func (b Int) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(b.String()))
}

func (b *Int) UnmarshalGQL(v interface{}) error {
	if bi, ok := v.(string); ok {
		b.Int = new(basebig.Int)
		b.Int, ok = b.Int.SetString(bi, 10)
		if !ok {
			return fmt.Errorf("invalid big number: %s", bi)
		}

		return nil
	}

	return fmt.Errorf("invalid big number")
}

func (b Int) Neg() Int {
	return FromBase(new(basebig.Int).Neg(b.Int))
}
