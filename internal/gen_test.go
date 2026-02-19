package python

import (
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sqlc-dev/sqlc-gen-python/internal/printer"
)

func Test_fieldNode(t *testing.T) {
	cases := []struct {
		Expected           string
		EmitPydanticModels bool
		Field              Field
	}{
		{
			Expected: "class_: int",
			Field: Field{
				Name: "class",
				Type: pyType{
					InnerType: "int",
					IsNull:    false,
					IsArray:   false,
				},
			},
			EmitPydanticModels: false,
		},
		{
			Expected: "class_: int = pydantic.Field(\n    alias=\"class\",\n)",
			Field: Field{
				Name: "class",
				Type: pyType{
					InnerType: "int",
					IsNull:    false,
					IsArray:   false,
				},
			},
			EmitPydanticModels: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Expected+" "+strconv.FormatBool(tc.EmitPydanticModels), func(t *testing.T) {
			res := fieldNode(tc.Field, tc.EmitPydanticModels)
			result := printer.Print(res, printer.Options{})
			if diff := cmp.Diff(strings.TrimSpace(tc.Expected), strings.TrimSpace(string(result.Python))); diff != "" {
				t.Errorf("node to python code mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
