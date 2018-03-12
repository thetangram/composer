version=$(shell cat VERSION)
build=$(shell git rev-parse --short HEAD)
build-date=$(shell date --rfc-3339=seconds)
build-target-dir=dist
docker-image=jomoespe/tangram


compile:
	@go build ./...

run:
	@go run cmd/tangramd/main.go

test:
	@echo "Testing..." 
	@go test -cover ./...


clean: 
	@echo "Cleaning..." 
	@rm --force --recursive dist 


build: clean test
	@echo "Building Tangramd service..." 
	@CGO_ENABLED=0 GOOS=linux go build -a \
									   -installsuffix cgo \
	                                   -ldflags "-s -w \
	                                             -X 'main.version=$(version)' \
	                                             -X 'main.build=$(build)' \
	                                             -X 'main.buildDate=$(build-date)'"  \
	                                   -o $(build-target-dir)/tangramd \
									   cmd/tangramd/main.go

	@echo "Building Tangramd AWS Lambda..." 
	@go build -o $(build-target-dir)/tangram-aws cmd/lambda/main.go


package: build
	@echo "Packaging Docker container..."
	@docker build --rm \
	              --build-arg version=$(version) \
				  --file build/package/Dockerfile \
				  --tag "$(docker-image):$(version)" .
	@docker tag "$(docker-image):$(version)" "$(docker-image):latest"

	@echo "Packaging AWS Lamdba function..."
	@zip -r -q $(build-target-dir)/deployment.zip $(build-target-dir)/tangram-aws config/
