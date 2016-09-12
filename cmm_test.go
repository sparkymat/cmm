package cmm

import (
	"fmt"
	"os"
	"testing"
)

func TestListFiles(t *testing.T) {
	if _, err := os.Stat("sample.cmm"); os.IsNotExist(err) {
		t.Errorf("File not found")
	}

	cmmFile := NewCmmFile("sample.cmm")

	err = cmmFile.OpenForRead()
	if err != nil {
		t.Errorf(err.Error())
	}

	defer cmmFile.CloseForRead()

	files, err := cmmFile.ListFiles()
	if err != nil {
		t.Errorf(err.Error())
	}

	for _, f := range files {
		fmt.Printf("Filename: %v\n", f.Name)
	}
}
