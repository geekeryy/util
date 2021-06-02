package excel

var defaultOption = excelOption{}

type Option interface {
	apply(excelOption)
}



type titleOpt struct {
	titles []string
}

func (t titleOpt) apply(eo excelOption) {
	eo.titles = t.titles
}

type fileNameOpt struct {
	fileName string
}

func (t fileNameOpt) apply(eo excelOption) {
	eo.fileName = t.fileName
}

type sheetNameOpt struct {
	sheetName string
}

func (t sheetNameOpt) apply(eo excelOption) {
	eo.sheetName = t.sheetName
}

