package domain

type Department struct {
	Id   int
	Name string
}

func (d Department) IsValid() bool {
	return d.Name != ""
}
