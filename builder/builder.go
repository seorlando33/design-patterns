package builder

import "time"

// Returning an interface you can chain the functions
// Like SetFormat().SetCreationDate().GetDocument() :)
type TextBuildProcess interface {
	SetFormat() TextBuildProcess
	SetData(txt string) TextBuildProcess
	SetCreationDate() TextBuildProcess
	GetDocument() Document
}

type Document struct {
	Format       string
	Data         string
	CreationDate time.Time
}

type TextDirector struct {
	builder	TextBuildProcess
}

func (f *TextDirector) Construct(txt string) {
	f.builder.SetFormat().SetData(txt).SetCreationDate()
}

func (f *TextDirector) SetBuilder(b TextBuildProcess)  {
	f.builder = b
}

type TXTBuilder struct {
	d Document
}

func (t *TXTBuilder) SetFormat() TextBuildProcess {
	t.d.Format = "txt"
	return t
}

func (t *TXTBuilder) SetData(txt string) TextBuildProcess {
	t.d.Data = txt
	return t
}

func (t *TXTBuilder) SetCreationDate() TextBuildProcess {
	t.d.CreationDate = time.Now()
	return t
}

func (t *TXTBuilder) GetDocument() Document {
	return t.d
}


type PDFBuilder struct {
	d Document
}

func (p *PDFBuilder) SetFormat() TextBuildProcess {
	p.d.Format = "pdf"
	return p
}

func (p *PDFBuilder) SetData(txt string) TextBuildProcess {
	p.d.Data = txt
	return p
}

func (p *PDFBuilder) SetCreationDate() TextBuildProcess {
	p.d.CreationDate = time.Now()
	return p
}

func (p *PDFBuilder) GetDocument() Document {
	return p.d
}