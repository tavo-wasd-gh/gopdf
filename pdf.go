package gopdf

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func PDF(tmpl *template.Template, data interface{}) ([]byte, error) {
	var filled bytes.Buffer
	if err := tmpl.Execute(&filled, data); err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	wk, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, fmt.Errorf("failed to create PDF generator: %w", err)
	}

	wk.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader(filled.Bytes())))

	if err := wk.Create(); err != nil {
		return nil, fmt.Errorf("failed to generate PDF: %w", err)
	}

	return wk.Bytes(), nil
}
