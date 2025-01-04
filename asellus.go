package asellus

type Chunk struct {
	Header         string
	Content        string
	Level          int
	Prefix         string
	ParentLevel    int
	ParentHeader   string
	ParentPrefix   string
	Lineage        string
	SimpleMetaData string                 // Additional metadata if needed
	Metadata       map[string]interface{} // additional metadata
	KeyWords       []string               // Keywords that could be extracted from the content
}
