# NOTE: call scripts from /scripts instead of writing lengthy targets
SHELL=/bin/bash # this forces usage of bash instead of sh (which is default and doesn't have a lot of features)
DOCKER_REGISTRY=REPLACE_ME
BUILD_ID=$(shell date +%Y%m%d%H%M)
NAMESPACE=$(DEVELOPER_NAMESPACE)
CONTAINER=$(DOCKER_REGISTRY)/$(SERVICE):$(NAMESPACE)-$(BUILD_ID)
ROOT_DIR=$(PWD)
# Enable or disable integration/functional testing using argo rollouts
#  true -> run integration/functional tests (slower)
#  false -> just deploy without tests (faster)
HAS_INTEGRATION_TESTS?=false
HAS_FUNCTIONAL_TESTS?=false
SVC_DIRS=\
	services/core/commandhandler \
	services/core/eventstore \
	services/core/notifier \
	services/core/queryhandler \
	services/core/registry
SVCS=$(SVC_DIRS:services/core/%=%)
PKG_DIRS=\
	pkg/azkeyvault \
	pkg/chassis \
	pkg/clientcredentials \
	pkg/logging \
	pkg/messagebus
TOOLS_DIRS=\
	tools/hammer
# Setting default values for key vault configuration. export env vars to overwrite.
KEYVAULT?=dev-keyvault
RESOURCEGROUP?=cluster-resources
OVERRIDE_VALUES?=developer-namespace
# for linux shenanigans
USER := $(shell whoami)
OS := $(shell uname)

.PHONY: all
all: $(SVCS)

.PHONY: $(SVCS)
$(SVCS):
# build the service image
	VERSION=$$(cat services/$@/VERSION); \
	docker build -f ./services/$@/Dockerfile -t $(DOCKER_REGISTRY)/$@:$$VERSION-$(NAMESPACE)-$(BUILD_ID) .; \
	docker push $(DOCKER_REGISTRY)/$@:$$VERSION-$(NAMESPACE)-$(BUILD_ID)

# only build integration image if enabled
ifeq ("$(HAS_INTEGRATION_TESTS)", "true")
# read the integration tests for service from values.yaml
	$(eval TESTS := $(shell yq eval '.rollout.testing.integration.tests' services/$@/values.yaml | sed 's/^- //'))
	@for TEST in $(TESTS); do \
		VERSION=$$(cat services/$@/VERSION); \
		docker build -f ./test/integration/$${TEST}/Dockerfile -t $(DOCKER_REGISTRY)/$@-integration-$${TEST}:$$VERSION-$(NAMESPACE)-$(BUILD_ID) .; \
		docker push $(DOCKER_REGISTRY)/$@-integration-$${TEST}:$$VERSION-$(NAMESPACE)-$(BUILD_ID); \
	done
endif # end HAS_INTEGRATION_TESTS

# only build functional image(s) if enabled
ifeq ("$(HAS_FUNCTIONAL_TESTS)", "true")
# read the functional tests for service from values.yaml
	$(eval TESTS := $(shell yq eval '.rollout.testing.functional.tests' services/$@/values.yaml | sed 's/^- //'))
	@for TEST in $(TESTS); do \
		VERSION=$$(cat services/$@/VERSION); \
		docker build -f ./test/functional/$${TEST}/Dockerfile -t $(DOCKER_REGISTRY)/$@-functional-$${TEST}:$$VERSION-$(NAMESPACE)-$(BUILD_ID) .; \
		docker push $(DOCKER_REGISTRY)/$@-functional-$${TEST}:$$VERSION-$(NAMESPACE)-$(BUILD_ID); \
	done
endif # end HAS_FUNCTIONAL_TESTS

# generate and apply helm template
	VERSION=$$(cat services/core/$@/VERSION) ; \
	helm template -n $@ --namespace $(NAMESPACE) \
		-f ./services/$@/values.yaml \
		-f ./deployments/values/all.yaml \
		-f ./deployments/values/developer-namespace.yaml \
		--set host.prefix=$(NAMESPACE) \
		--set image.tag=$$VERSION-$(NAMESPACE)-$(BUILD_ID) \
		--set namespace=$(NAMESPACE) \
		--set configMap.namespace=$(NAMESPACE) \
		--set configMap.keyVault=$(KEYVAULT) \
		--set configMap.resourceGroup=$(RESOURCEGROUP) \
		--set configMap.version=$$VERSION-$(NAMESPACE)-$(BUILD_ID) \
		--set rollout.testing.integration.isEnabled=$(HAS_INTEGRATION_TESTS) \
		--set rollout.testing.functional.isEnabled=$(HAS_FUNCTIONAL_TESTS) \
		./deployments/service/ | kubectl apply -f - -n $(NAMESPACE)

# helm template - render the deployment manifest for a specific service
# $ make template SVC="core/eventstore" NAMESPACE="dev" OVERRIDE_VALUES="dev"
.PHONY: template
template:
	VERSION=$$(cat services/$(SVC)/VERSION) ; \
	helm template -n $(SVC) --namespace $(NAMESPACE) \
		-f ./services/$(SVC)/values.yaml \
		-f ./deployments/values/all.yaml \
		-f ./deployments/values/$(OVERRIDE_VALUES).yaml \
		--set host.prefix=$(NAMESPACE) \
		--set image.tag=$$VERSION-$(NAMESPACE)-$(BUILD_ID) \
		--set namespace=$(NAMESPACE) \
		--set configMap.namespace=$(NAMESPACE) \
		--set configMap.keyVault=$(KEYVAULT) \
		--set configMap.resourceGroup=$(RESOURCEGROUP) \
		--set configMap.version=$$VERSION \
		--set rollout.testing.integration.isEnabled=$(HAS_INTEGRATION_TESTS) \
		--set rollout.testing.functional.isEnabled=$(HAS_FUNCTIONAL_TESTS) \
		./deployments/service/

# Uses prebuilt images from the registry with the versions defined in the service directory
# To deploy all services: make remote
# To deploy a specific service: make remote SVCS=eventstore
.PHONY: remote
remote:
	for service in $(SVCS) ; do \
		VERSION=$$(cat services/$$service/VERSION) ; \
		helm template -n $$service --namespace $(NAMESPACE) \
			-f ./services/$$service/values.yaml \
			-f ./deployments/values/all.yaml \
			-f ./deployments/values/developer-namespace.yaml \
			--set host.prefix=$(NAMESPACE) \
			--set image.tag=$$VERSION \
			--set image.repository=SOME_URL/$$service \
			--set namespace=$(NAMESPACE) \
			--set configMap.namespace=$(NAMESPACE) \
			--set configMap.keyVault=$(KEYVAULT) \
			--set configMap.resourceGroup=$(RESOURCEGROUP) \
			--set rollout.testing.integration.isEnabled=$(HAS_INTEGRATION_TESTS) \
			--set rollout.testing.functional.isEnabled=$(HAS_FUNCTIONAL_TESTS) \
			./deployments/service/ | kubectl apply -f - -n $(NAMESPACE) ; \
    done

###################
# BUILD NEW STUFF'S
###################

# create a new consumer service (ASYNC)
# $ make consumer SERVICE=GameOfficiating AGGREGATE=Football COMMAND=PauseGame
consumer:
	cd tools/hammer && go run main.go consumer --service_name="$(SERVICE)" --aggregate_name="$(AGGREGATE)" --command_name="$(COMMAND)" --output_path="../../services"

########################
# BLUE/GREEN DEVELOPMENT
########################
shawarma:
	docker build -f ./third_party/$@/Dockerfile -t $(DOCKER_REGISTRY)/$@:0.0.1-$(NAMESPACE)-$(BUILD_ID) ./third_party/$@; \
	docker push $(DOCKER_REGISTRY)/$@:0.0.1-$(NAMESPACE)-$(BUILD_ID)

shawarma-webhook:
	docker build -f ./third_party/$@/Dockerfile -t $(DOCKER_REGISTRY)/$@:0.0.1-$(NAMESPACE)-$(BUILD_ID) ./third_party/$@; \
	docker push $(DOCKER_REGISTRY)/$@:0.0.1-$(NAMESPACE)-$(BUILD_ID)
