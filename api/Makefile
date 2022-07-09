OUTPUT_ROOT = gen
#Additional Makefiles go here when they are written. The build target should be the default target and don't forget to
#define a clean.
MAKEFILES := $(shell ls -d ./${OUTPUT_ROOT}/*)
# make sure to run as user (not root which is default on linux)
USER := $(shell whoami)
OS := $(shell uname)
BUILDER_IMAGE = bufbuilder

.PHONY: build clean docker-build docker-clean
.SILENT: build clean docker-build docker-clean

# build all generated files (according to the build targets in the sub-folders)
build: run-bufbuilder
	$(foreach mf,$(MAKEFILES),$(MAKE) --no-print-directory -C $(mf);)

# clean all generated files (according to the clean targets in the sub-folders)
clean:
	@rm -rf buf.lock
	$(foreach mf,$(MAKEFILES),$(MAKE) --no-print-directory clean -C $(mf);)

# build and tag the protobuf builder image
docker-build:
	docker build -t $(BUILDER_IMAGE):v1 -f ./Dockerfile .;

# remove the built Dockerfile image for the protobuf builder
docker-clean:
	docker rmi `docker images '$(BUILDER_IMAGE)' -a -q`

run-bufbuilder:
	@echo "Running the protobuf builder"
	@docker run --rm --name bufbuilder --volume $(PWD):/workspace --workdir /workspace bufbuilder:v1 mod update
	@docker run --rm --name bufbuilder --volume $(PWD):/workspace --workdir /workspace bufbuilder:v1 generate
	@if [ "$(OS)" = "Linux" ]; then\
		sudo chown -R $(USER):$(USER) $(PWD);\
	fi
	@echo "Finished running the protobuf builder"
