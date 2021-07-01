package excel

var defaultOption = ExcelOption{}

type Option interface {
	apply(*ExcelOption)
}



type titleOpt struct {
	titles []string
}

func (t titleOpt) apply(eo *ExcelOption) {
	eo.titles = t.titles
}

type fileNameOpt struct {
	fileName string
}

func (t fileNameOpt) apply(eo *ExcelOption) {
	eo.fileName = t.fileName
}

type sheetNameOpt struct {
	sheetName string
}

func (t sheetNameOpt) apply(eo *ExcelOption) {
	eo.sheetName = t.sheetName
}

