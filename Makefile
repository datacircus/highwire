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

gen: buf-lint
	@buf generate

clean:
	rm -rf gen/
