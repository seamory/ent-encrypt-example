//go:build ignore

package main

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"log"
	"main/internal/tools/array"
	"os"
)

func main() {
	remove(".")
	codegen("./schema")
}

func remove(dir string) {
	if entries, err := os.ReadDir(dir); err != nil {
		log.Fatalln(err)
	} else {
		for _, entry := range entries {
			list := array.Strings{"entc.go", "generate.go", "schema"}
			if !list.Includes(entry.Name()) {
				if entry.IsDir() {
					os.RemoveAll(entry.Name())
				} else {
					os.Remove(entry.Name())
				}
			}
		}
	}
	log.Println("remove ent history files")
}

func codegen(schemaPath string) {
	err := entc.Generate(schemaPath, &gen.Config{
		Features: []gen.Feature{
			gen.FeatureSchemaConfig,
			gen.FeatureModifier,
			gen.FeatureExecQuery,
		},
	})
	if err != nil {
		log.Fatalln("running ent codegen:", err)
	}
	log.Println("ent code generator done")
}
