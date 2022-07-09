GO=go
GENERATED := $(shell ls -d ./*/)
OS := $(shell uname)

.PHONY: build clean

build:
	@echo "Processing extra build setps for 'go/'"
	@# https://github.com/infobloxopen/protoc-gen-gorm/issues/224
	@# On OSX, replace -i '' with -i '.bak' to generate before and after if the sed is misbehaving.
	@if [ "$(OS)" = "Linux" ]; then \
		find . -name "*.pb.gorm.go" | xargs sed -i 's/_ "google.golang.org\/protobuf\/types\/known\/timestamppb"/"google.golang.org\/protobuf\/types\/known\/timestamppb"/g'; \
	else \
		find . -name "*.pb.gorm.go" | xargs sed -i '' 's/_ "google.golang.org\/protobuf\/types\/known\/timestamppb"/"google.golang.org\/protobuf\/types\/known\/timestamppb"/g'; \
	fi
	$(GO) mod tidy && $(GO) build ./...
	@echo "Finished with 'go/'"

clean:
	@echo "Removing generated code for 'go/'"
	@rm -rf $(GENERATED)
