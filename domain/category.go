package domain

type Category struct {
	Id     string
	Parent *Category
	Title  string
}
