all:
	c-for-go tcml.yml

clean:
	rm -f tcml/cgo_helpers.go tcml/cgo_helpers.h tcml/cgo_helpers.c
	rm -f tcml/const.go tcml/doc.go tcml/types.go
	rm -f tcml/tcml.go

test:
	cd tcml && go build
	