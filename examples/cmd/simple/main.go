package main

import (
	"fmt"

	"github.com/henomis/lingopipes/prompt/template"
)

func main() {

	promptTemplate := template.New(nil, nil, "Hello World", nil)
	output, _ := promptTemplate.Format(nil)
	fmt.Println(output)

}
