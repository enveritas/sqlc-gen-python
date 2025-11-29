package poet

import "slices"

// TODO(quentin@escape.tech): check if this is complete
var reservedKeywords = []string{
	"and",
	"as",
	"assert",
	"async",
	"await",
	"break",
	"class",
	"continue",
	"def",
	"del",
	"elif",
	"else",
	"except",
	"finally",
	"for",
	"from",
	"global",
	"if",
	"import",
	"in",
	"is",
	"lambda",
	"nonlocal",
	"not",
	"or",
	"pass",
	"raise",
	"return",
	"try",
	"while",
	"with",
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
