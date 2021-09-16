NAMESPACE  := logicmonitor
REPOSITORY := collectorset-controller
VERSION    := 1.0.0

default: build

gofmt:
ifeq ($(shell uname -s), Darwin)
	find pkg/ -type f | grep go | egrep -v "mocks|gomock" | xargs gofmt -l -d -s -w; sync
	find pkg/ -type f | grep go | egrep -v "mocks|gomock" | xargs gofumpt -l -d -s -w; sync
	find pkg/ -type f | grep go | egrep -v "mocks|gomock" | xargs gci -w; sync
	find pkg/ -type f | grep go | egrep -v "mocks|gomock" | xargs goimports -l -d -w; sync
	find cmd/ -type f | grep go | egrep -v "mocks|gomock" | xargs gofmt -l -d -s -w; sync
	find cmd/ -type f | grep go | egrep -v "mocks|gomock" | xargs gofumpt -l -d -s -w; sync
	find cmd/ -type f | grep go | egrep -v "mocks|gomock" | xargs gci -w; sync
	find cmd/ -type f | grep go | egrep -v "mocks|gomock" | xargs goimports -l -d -w; sync
	gofmt -l -d -s -w main.go; sync
	gofumpt -l -d -s -w main.go; sync
	gci -w main.go; sync
	goimports -l -d -w main.go; sync
endif

build: gofmt
	docker build --build-arg VERSION=$(VERSION) -t $(NAMESPACE)/$(REPOSITORY):latest .
	docker run --rm -v "$(shell pwd)":/out --entrypoint=cp $(NAMESPACE)/$(REPOSITORY):latest /tmp/api.pb.go /out/api
	docker run --rm -v "$(shell pwd)":/out --entrypoint=cp $(NAMESPACE)/$(REPOSITORY):latest /tmp/zz_generated.deepcopy.go /out/pkg/apis/v1alpha2/
	docker tag $(NAMESPACE)/$(REPOSITORY):latest $(NAMESPACE)/$(REPOSITORY):$(VERSION)

dev: gofmt
	docker build --build-arg VERSION=$(VERSION) -t $(NAMESPACE)/$(REPOSITORY):latest -f Dockerfile.dev .
	docker run --rm -v "$(shell pwd)":/out --entrypoint=cp $(NAMESPACE)/$(REPOSITORY):latest /tmp/api.pb.go /out/api
	docker run --rm -v "$(shell pwd)":/out --entrypoint=cp $(NAMESPACE)/$(REPOSITORY):latest /tmp/zz_generated.deepcopy.go /out/pkg/apis/v1alpha2/
	docker tag $(NAMESPACE)/$(REPOSITORY):latest $(NAMESPACE)/$(REPOSITORY):$(VERSION)

	
