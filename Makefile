.PHONY: \
	buf-breaking \
	buf-lint \
	build-descriptor \
	clean \
	gen \
	sync

buf-breaking:
	@buf breaking --against .git#branch=origin/main

buf-lint:
	@buf lint

buf-format:
	@buf format -w

build-descriptor:
	@buf build -o gen/coffeeco.bin

sync:
	@buf export buf.build/sideshow/coffeeco -o proto/

gen: buf-lint
	@buf generate

clean:
	rm -rf gen/
