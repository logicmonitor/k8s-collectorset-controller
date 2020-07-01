NAMESPACE  := logicmonitor
REPOSITORY := collectorset-controller
VERSION    := 1.0.0

all:
	docker build --build-arg VERSION=$(VERSION) -t $(NAMESPACE)/$(REPOSITORY):v2latest .
	docker run --rm -v "$(shell pwd)":/out --entrypoint=cp $(NAMESPACE)/$(REPOSITORY):v2latest /tmp/api.pb.go /out/api
	docker run --rm -v "$(shell pwd)":/out --entrypoint=cp $(NAMESPACE)/$(REPOSITORY):v2latest /tmp/zz_generated.deepcopy.go /out/pkg/apis/v1alpha1/
	docker tag $(NAMESPACE)/$(REPOSITORY):v2latest $(NAMESPACE)/$(REPOSITORY):$(VERSION)

dev:
	docker build --build-arg VERSION=$(VERSION) -t $(NAMESPACE)/$(REPOSITORY):v2latest -f Dockerfile.dev .
	docker run --rm -v "$(shell pwd)":/out --entrypoint=cp $(NAMESPACE)/$(REPOSITORY):v2latest /tmp/api.pb.go /out/api
	docker run --rm -v "$(shell pwd)":/out --entrypoint=cp $(NAMESPACE)/$(REPOSITORY):v2latest /tmp/zz_generated.deepcopy.go /out/pkg/apis/v1alpha1/
	docker tag $(NAMESPACE)/$(REPOSITORY):v2latest $(NAMESPACE)/$(REPOSITORY):$(VERSION)