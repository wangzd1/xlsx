package types_test

import (
	"encoding/xml"
	"github.com/plandem/xlsx/types"
	"github.com/stretchr/testify/require"
	"testing"
	"fmt"
)

func TestCellType(t *testing.T) {
	type Entity struct {
		Attribute types.CellType `xml:"attribute,attr"`
	}


	list := map[string] types.CellType{
		"b": types.CellTypeBool,
		"d": types.CellTypeDate,
		"n": types.CellTypeNumber,
		"e": types.CellTypeError,
		"s": types.CellTypeSharedString,
		"str": types.CellTypeFormula,
		"inlineStr": types.CellTypeInlineString,
	}

	for s, v := range list {
		t.Run(s, func(tt *testing.T){
			entity := Entity{Attribute: v}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			require.Equal(tt, fmt.Sprintf(`<Entity attribute="%s"></Entity>`, s), string(encoded))

			var decoded Entity
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
		})
	}
}