package logistic

import "fmt"

type Package struct {
	Name string
}

func (p *Package) String() string {
	return fmt.Sprintf("Package name: %s", p.Name)
}

var AllEntities = []*Package{
	{ Name: "Коробка"},
	{ Name: "Ящик"},
	{ Name: "Конверт"},
	{ Name: "Пачка"},
	{ Name: "Банка"},
	{ Name: "Коробка для конфет"},
	{ Name: "Ящик для конфет"},
	{ Name: "Конверт для тонких конфет"},
	{ Name: "Пачка для конфет"},
	{ Name: "Банка для конфет"},
}