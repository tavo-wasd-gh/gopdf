# gopdf

## Sample hello world

```go
package main

import (
	"html/template"
	"log"

	"github.com/tavo-wasd-gh/gopdf"
)

func main() {
	tmpl, err := template.New("pdf-template").Parse(`<p>{{.Content}}</p>`)
	if err != nil {
		log.Fatalf("failed to parse template: %v", err)
	}

	data := struct {
		Content string
	}{
		Content: "Hello, World!",
	}

	pdfBytes, err := gopdf.PDF(tmpl, data)
	if err != nil {
		log.Fatalf("failed to generate PDF: %v", err)
	}

	log.Printf("PDF generated successfully: %d bytes", len(pdfBytes))
}
```
