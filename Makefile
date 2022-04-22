all: bench

TAGS ?= -tags containers_image_openpgp
# Dumb hack to use rooted containerd socket.
ifneq ($(shell id -u),0)
$(info Executing 'sudo' for containerd socket access)
GAINROOTFN := sudo
endif

BENCH_TIME ?= 100x
BENCH_CPUS ?= 1,4,8

bench:
	$(if $(V),,@)CGO_ENABLED=0 go test -c -o ./bin/b.test $(TAGS)
	$(if $(V),,@)$(GAINROOTFN) ./bin/b.test -test.bench=. -test.benchmem -test.cpu $(BENCH_CPUS) -test.benchtime=$(BENCH_TIME) $(if $(V),-test.v,)
