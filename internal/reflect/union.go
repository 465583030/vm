package reflect

type unionType struct {
	name   string
	fields map[string]Type
}

func (t *unionType) String() string { return t.name }
func (t *unionType) IsRef() bool    { return false }

func (t *unionType) Size() uint {
	var max uint
	for _, t := range t.fields {
		if t.Size() > max {
			max = t.Size()
		}
	}
	return max
}

func Union(name string, fields map[string]Type) Type {
	return &unionType{
		name:   name,
		fields: fields,
	}
}
