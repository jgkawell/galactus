GENERATED := $(shell ls -d ./*/)
OS := $(shell uname)

.PHONY: build clean

build:
	@echo "Processing extra build steps for 'web/'"
	@# https://github.com/infobloxopen/protoc-gen-gorm/issues/224
	@# On OSX, replace -i '' with -i '.d.ts' to generate before and after if the sed is misbehaving.
	@if [ "$(OS)" = "Linux" ]; then \
		find . -name "*.d.ts" | xargs sed -i 's/import \* as validate_validate_pb from '\''\.\.\/\.\.\/\.\.\/validate\/validate_pb'\''\;//g'; \
		find . -name "*_pb.js" | xargs sed -i '/validate/d'; \
	else \
		find . -name "*.d.ts" | xargs sed -i '' 's/import \* as validate_validate_pb from '\''\.\.\/\.\.\/\.\.\/validate\/validate_pb'\''\;//g'; \
		find . -name "*_pb.js" | xargs sed -i '' '/validate/d'; \
	fi
	@echo "Finished with 'web/'"

clean:
	@echo "Removing generated code for 'web/'"
	@rm -rf $(GENERATED)
