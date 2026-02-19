package poet

import "slices"

// TODO(quentin@escape.tech): check if this is complete
var reservedKeywords = []string{
	"class",
	"if",
	"else",
	"elif",
	"not",
	"for",
	"and",
	"in",
	"is",
	"or",
	"with",
	"as",
	"assert",
	"break",
	"except",
	"finally",
	"try",
	"raise",
	"return",
	"yield",
}

func IsReserved(name string) bool {
	return slices.Contains(reservedKeywords, name)
}

func FieldName(name string) string {
	if IsReserved(name) {
		return name + "_"
	}
	return name
}
