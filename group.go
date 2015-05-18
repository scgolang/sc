package sc

type node struct {
	id int32 `json:"id" xml:"id,attr"`
}

type group struct {
	node
	children []*node `json:"children" xml:children>child"`
}

