package cmm

type Node struct {
	Id          uint32
	Name        string
	Description string
	Connections []*Connection
	Attachments []*Attachment
}
