package builder

import (
	"testing"
)

func TestBuilderPattern(t *testing.T) {

	manufacturingComplex := TextDirector{}

	pdfBuilder := &PDFBuilder{}

	manufacturingComplex.SetBuilder(pdfBuilder)

	manufacturingComplex.Construct("Testing PDF builder")

	pdf := pdfBuilder.GetDocument()

	if pdf.Format != "pdf" {
		t.Errorf("Format on a pdf must be 'pdf' and was %s\n", pdf.Format)
	}
	if pdf.Data != "Testing PDF builder" {
		t.Errorf("Expected data was 'Testing PDF builder' but was %s\n",
			pdf.Data)
	}

	if pdf.CreationDate.IsZero() {
		t.Error("Time was not set")
	}

	txtBuilder := &TXTBuilder{}

	manufacturingComplex.SetBuilder(txtBuilder)

	manufacturingComplex.Construct("Testing TXT builder")

	txt := txtBuilder.GetDocument()

	if txt.Format != "txt" {
		t.Errorf("Format on a txt must be 'txt' and was %s\n",
			txt.Format)
	}

	if txt.Data != "Testing TXT builder" {
		t.Errorf("Expected data was 'Testing PDF builder' but was %s\n",
			txt.Data)
	}
}
