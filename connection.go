package cmm

type Connection struct {
	Id              uint32
	Name            string
	Description     string
	SourceNode      *Node
	DestinationNode *Node
	Attachments     []*Attachment
}
