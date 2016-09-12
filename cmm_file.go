package cmm

import (
	"archive/zip"
	"errors"
)

type CmmFile struct {
	FilePath   string
	ReadCloser *zip.ReadCloser
	Writer     *zip.Writer
}

func NewCmmFile(file string) CmmFile {
	return CmmFile{FilePath: file}
}

func (c *CmmFile) OpenForRead() error {
	var err error
	c.ReadCloser, err = zip.OpenReader(c.FilePath)

	if err != nil {
		return err
	}

	return nil
}

func (c *CmmFile) CloseForRead() error {
	if c.ReadCloser == nil {
		return errors.New("File not opened for read")
	}

	c.ReadCloser.Close()

	c.ReadCloser = nil

	return nil
}

func (c *CmmFile) ListFiles() ([]*File, error) {
	if c.ReaderCloser == nil {
		return errors.New("File not opened for read")
	}

	return c.ReaderClose.File, nil
}

func (c *CmmFile) GetFile(path string) (*File, error) {
	if c.ReaderCloser == nil {
		return errors.New("File not opened for read")
	}

	for _, f := range reader.File {
		if f.Name == path {
			return f, nil
		}
	}

	return nil, errors.New("File not found")
}

func (c CmmFile) GetRoot() (*CmmNode, error) {
}
