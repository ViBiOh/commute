package templ

type Field struct {
	Err   error
	Value string
}

type Fields map[string]Field

func (f Fields) HasError() bool {
	for _, field := range f {
		if field.Err != nil {
			return true
		}
	}

	return false
}
