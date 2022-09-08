entity=

default:
	@echo

ent.init:
	go run entgo.io/ent/cmd/ent init ${entity}

ent.generate:
	go generate ./ent
