
GOARCH ?= $(shell go env GOHOSTARCH)
PLATFORM ?= $(shell go env GOHOSTOS)/$(shell go env GOHOSTARCH)

USER_TKM ?= ${QUAY_USER}
TAG_TKM ?= latest
IMAGE_TKM_OPERATOR ?= quay.io/$(USER_TKM)/tkm-operator:$(TAG_TKM)
IMAGE_TKM_AGENT ?= quay.io/$(USER_TKM)/tkm-agent:$(TAG_TKM)

.PHONY: build
build: ## Build local tkm binary.
	go build -o bin/tkm-operator cmd/tkm-operator/main.go
	go build -o bin/tkm-agent cmd/tkm-agent/main.go

.PHONY: build-images
build-images: build ## Build tkm in a container images
	DOCKER_BUILDKIT=0 docker buildx build --progress=plain -t ${IMAGE_TKM_OPERATOR}  --platform ${PLATFORM} --load -f ./Containerfile.tkm-operator .
	DOCKER_BUILDKIT=0 docker buildx build --progress=plain -t ${IMAGE_TKM_AGENT}  --platform ${PLATFORM} --load -f ./Containerfile.tkm-agent .

.PHONY: deploy
deploy: ## Deploy all examples to the cluster specified in ~/.kube/config.
	kubectl apply -f config/tkm-operator/deployment.yaml
	kubectl apply -f config/tkm-agent/daemonset.yaml

.PHONY: undeploy
undeploy: ## Deploy all examples to the cluster specified in ~/.kube/config.
	kubectl delete -f config/tkm-agent/daemonset.yaml
	kubectl delete -f config/tkm-operator/deployment.yaml
