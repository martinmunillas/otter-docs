package translations

import (
	"bytes"
	_ "embed"

	"github.com/martinmunillas/otter/i18n"
)

//go:embed en.json
var enJson []byte

//go:embed es.json
var esJson []byte

func Setup() {
	i18n.AddLocale("en", bytes.NewReader(enJson))
	i18n.AddLocale("es", bytes.NewReader(esJson))
}
