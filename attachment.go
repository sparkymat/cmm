package cmm

import "os"

type Attachment struct {
	Path string
	File *os.File
}
