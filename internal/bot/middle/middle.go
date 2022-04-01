package middle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Dict is a struct that holds the data from the API
type Dicts []Dict

type Dict struct {
	// Word is the word that was searched
	Word string `json:"word"`

	// Phonetics is the phonetic spelling of the word
	Phonetics []struct {
		Audio string `json:"audio"`
		// SourceURL string `json:"sourceUrl,omitempty"`
		Text string `json:"text,omitempty"`
	} `json:"phonetics"`

	// Meanings is a slice of structs that hold the meanings of the word
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`

		Definitions []struct {
			Definition string `json:"definition"`
			// Synonyms   []interface{} `json:"synonyms"`
			// Antonyms   []interface{} `json:"antonyms"`
		} `json:"definitions"`

		// Synonyms []string      `json:"synonyms"`
		// Antonyms []interface{} `json:"antonyms"`
	} `json:"meanings"`

	// SourceUrls []string `json:"sourceUrls"`
}

// GetDicts is a function that gets the data from the API
func GetDicts(word string) (Dicts, error) {
	url := "https://api.dictionaryapi.dev/api/v2/entries/en/" + word
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var dicts Dicts
	////////////// i'm not sure about error handling here
	_ = json.Unmarshal(body, &dicts)
	return dicts, nil
}

// Valuable is a function that returns the most valuable Dict from Dicts struct
func (ds Dicts) Valuable() Dict {
	var d Dict
	for _, dict := range ds {
		for _, ph := range dict.Phonetics {
			if ph.Text != "" {
				d = dict
				d.Phonetics[0] = ph
				return d
			}
		}
	}
	return ds[0]
}

func GenerateMessage(word string) string {
	ds, _ := GetDicts(word)

	dict := ds.Valuable()

	return fmt.Sprintf("%s - %s (%s)\n\n%s: %s", dict.Word, dict.Phonetics[0].Text, dict.Phonetics[0].Audio, dict.Meanings[0].PartOfSpeech, dict.Meanings[0].Definitions[0].Definition)
}
