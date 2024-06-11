.PHONY: generate clean

generate:
	GO_POST_PROCESS_FILE="/usr/bin/sed -i -e s,place/holder,bsky.watch,g" \
	./openapi-generator-cli generate \
		--enable-post-process-file \
		-i _spec.yaml \
		-g go \
		--skip-validate-spec \
		-c generator-config.yml \
		--global-property modelDocs=false,apiDocs=false \
		--type-mappings=integer=int,string+byte=[]byte
	sed -i -e s,place/holder,bsky.watch,g go.mod
	go mod tidy
	go test ./...

clean:
	cat .openapi-generator/FILES | grep -v '^api/' | xargs rm || true
	rm -rf docs test || true
