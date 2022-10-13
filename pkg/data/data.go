package data

import (
	"encoding/json"
	"path/filepath"
)

func Load() ([]*LegalAct, error) {
	result := make([]*LegalAct, 0)
	files, err := data.ReadDir(".")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) != "json" {
			continue
		}

		fileData, err := data.ReadFile(file.Name())
		if err != nil {
			return nil, err
		}

		output := &LegalAct{}
		err := json.Unmarshal(output, fileData)
		if err != nil {
			return nil, err
		}

		result = append(result, output)
	}

	return result, nil
}

// LegalAct represents a set of Law rules (e.g. Konstytucja RP)
type LegalAct struct {
	// ActName is a name of the Legal act (e.g. Konstytucja RP)
	ActName string

	// Rules is a set of rules
	Rules []Rule
}

type Rule struct {
	// Identifier is a "index" of article/paragraph
	// e.g.: "Rozdział 4, Artykuł 3, paragraf 7"
	Identifier string

	// Text is a text of the rule
	Text string

	// Links is a list to external resources. e.g. YouTube desctiption e.t.c.
	Links []string
}
