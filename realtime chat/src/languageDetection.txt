/*
https://github.com/pemistahl/lingua-go#library-dependency

To Use:
1. git clone https://github.com/pemistahl/lingua-go.git
2. cd lingua-go
3. go build
*/

package main

import (
	"fmt"

	"github.com/pemistahl/lingua-go"
)

func languageDetect() {
	languages := []lingua.Language{
		lingua.English,
		lingua.Japanese,
		lingua.German,
		lingua.Spanish,
		lingua.Chinese,
	}

	detector := lingua.NewLanguageDetectorBuilder().
		FromLanguages(languages...).
		Build()

	if language, exists := detector.DetectLanguageOf("languages are awesome"); exists {
		fmt.Println(language)
	}

	// Output: English
}
