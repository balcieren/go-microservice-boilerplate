package main

import (
	"github.com/balcieren/go-microservice-boilerplate/pkg/entity"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.ApplyBasic(entity.Pet{}, entity.Owner{})

	g.Execute()
}
