package reflect

type structType struct {
	name   string
	fields map[string]Type
}

func (t *structType) String() string { return t.name }
func (t *structType) IsRef() bool    { return false }

func (t *structType) Size() uint {
	var size uint
	for _, t := range t.fields {
		size += t.Size()
	}
	return size
}

func Struct(name string, fields map[string]Type) Type {
	return &structType{
		name:   name,
		fields: fields,
	}
}
