# gopdf

## Sample hello world

```go
package main

import (
	"fmt"
	"os"

	"github.com/tavo-wasd-gh/gopdf"
)

func main() {
	htmlTemplate := `<h1>{{ .Message }}</h1>`

	data := struct {
		Message string
	}{
		Message: "Hello, world!",
	}

	pdfBytes, err := gopdf.PDF(htmlTemplate, data)
	if err != nil {
		fmt.Println("Error generating PDF:", err)
		return
	}

	err = os.WriteFile("hello_world.pdf", pdfBytes, 0644)
	if err != nil {
		fmt.Println("Error saving PDF:", err)
		return
	}

	fmt.Println("PDF generated successfully: hello_world.pdf")
}
```
