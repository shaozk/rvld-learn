VERSION = 0.1.0
COMMIT_ID = $(shell git rev-list -1 HEAD)
TESTS := $(wildcard tests/*.sh) 

build:
	@go build -ldflags "-X main.version=${VERSION}-${COMMIT_ID}"
	@ln -sf rvld-learn ld

clean:
	go clean
	rm -rf out/

test: build
	@CC="riscv64-linux-gnu-gcc" \
	$(MAKE) $(TESTS)
	@printf '\e[32mPassed all tests\e[0m\n'

$(TESTS):
	@echo 'Testing' $@
	@./$@
	@printf '\e[32mOk\e[0m\n'

.PHONY: build clean test $(TESTS)

