package export

type Export struct {
	Sheet string
}

func(t *Export) AddSheet(sheetName string)  {
	if sheetName == "" {
		t.Sheet = "sheet1"
	}else {
		t.Sheet = sheetName
	}
}

func(t *Export) AddRow()  {

}
