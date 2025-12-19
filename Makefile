# binary name
BINARY_NAME=sfvbrr

# go related variables
GO=go
GOBIN=$(shell $(GO) env GOPATH)/bin

# build variables
BUILD_DIR=build
VERSION=$(shell git describe --tags 2>/dev/null || echo "dev")
BUILD_TIME=$(shell date +%FT%T%z)
LDFLAGS=-ldflags "-X github.com/autobrr/sfvbrr/cmd.version=${VERSION} -X github.com/autobrr/sfvbrr/cmd.buildTime=${BUILD_TIME}"

# race detector settings
GORACE=log_path=./race_report.log \
       history_size=2 \
       halt_on_error=1 \
       atexit_sleep_ms=2000

# make all builds and installs
.PHONY: all
all: clean build install

# build binary
.PHONY: build
build:
	@echo "Building ${BINARY_NAME}..."
	@mkdir -p ${BUILD_DIR}
	CGO_ENABLED=0 $(GO) build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY_NAME}

# build with PGO
.PHONY: build-pgo
build-pgo:
	@echo "Building ${BINARY_NAME} with PGO..."
	@if [ ! -f "default.pgo" ]; then \
		echo "No PGO profile found. Run 'make profile' first."; \
		exit 1; \
	fi
	@mkdir -p ${BUILD_DIR}
	CGO_ENABLED=0 $(GO) build -pgo=default.pgo ${LDFLAGS} -o ${BUILD_DIR}/${BINARY_NAME}

# generate PGO profile with various workloads
.PHONY: profile
profile:
	@echo "Generating PGO profile..."
	@go build -o ${BUILD_DIR}/${BINARY_NAME}

	@echo "Running profile workload 1: SFV..."
	@${BUILD_DIR}/${BINARY_NAME} sfv ./test/sfv/TEST.FOLDER.v0.1.X64-GRP --cpuprofile=./cpu1.pprof

	@echo "Running profile workload 2: Validate..."
	@${BUILD_DIR}/${BINARY_NAME} validate \
		./test/validate/01_1/App.Pro.v1.1.1.Linux.RPM.ARM64.Incl.Keymaker-GRP \
		./test/validate/01_2/Corporation.Program.v1.0.0.0.x64.Multilingual.Incl.Keymaker-GRP \
		./test/validate/02_1/First_Last_-_Title-AUDiOBOOK-WEB-EN-2025-GRP \
		./test/validate/02_2/Artist-Title-AUDiOBOOK-WEB-EN-2025-GRP \
		./test/validate/03_1/First.Last.The.Title.2025.RETAiL.ePub.eBook-GRP \
		./test/validate/04_1/Comics.-.Title.Vol.01.2025.Retail.Comic.eBook-GRP \
		./test/validate/05_1/Learning.-.Topic.UPDATED.November.2025.BOOKWARE-GRP \
		./test/validate/05_2/OREILLY_CLOUD-GRP \
		./test/validate/06_1/Show.S01E06.1080p.WEB.h264-GRP \
		./test/validate/06_2/TVShow.S10.E999.720p.BluRay.x264-GRP \
		./test/validate/07_1/Great.Game.Udate.v1.2.34.567890.incl.DLC-GRP \
		./test/validate/08_1/Title.No.100.2025.GERMAN.HYBRID.MAGAZINE.eBook-GRP \
		./test/validate/09_1/The.Movie.2025.1080P.BLURAY.X264-GRP \
		./test/validate/10_1/First_Last_-_Title-WEB-CZ-2025-GRP \
		./test/validate/10_2/Artist-Title-WEB-2025-GRP \
		./test/validate/11_1/Show.Season.S01.COMPLETE.HDTV.x264-GRP \
		./test/validate/11_2/Show.S01.1080p.WEB.H264-GRP \
		--cpuprofile=./cpu2.pprof

	@echo "Running profile workload 3: ZIP..."
	@${BUILD_DIR}/${BINARY_NAME} zip ./test/zip/TEST.FOLDER.v0.1.X64-GRP --cpuprofile=./cpu3.pprof

	@echo "Merging profiles..."
	@if [ -f "cpu1.pprof" ] && [ -f "cpu2.pprof" ] && [ -f "cpu3.pprof" ]; then \
		go tool pprof -proto cpu1.pprof cpu2.pprof cpu3.pprof> default.pgo; \
		rm cpu1.pprof cpu2.pprof cpu3.pprof; \
		echo "Profile generated at default.pgo"; \
	else \
		echo "Error: Profile files not generated correctly"; \
		exit 1; \
	fi

# install binary in system path
.PHONY: install
install: build
	@echo "Installing ${BINARY_NAME}..."
	@if [ "$$(id -u)" = "0" ]; then \
		install -m 755 ${BUILD_DIR}/${BINARY_NAME} /usr/local/bin/; \
	else \
		install -m 755 ${BUILD_DIR}/${BINARY_NAME} ${GOBIN}/; \
	fi

# install binary with PGO optimization
.PHONY: install-pgo
install-pgo:
	@echo "Installing ${BINARY_NAME} with PGO..."
	@if [ ! -f "default.pgo" ]; then \
		echo "No PGO profile found. Run 'make profile' first."; \
		exit 1; \
	fi
	@mkdir -p ${BUILD_DIR}
	CGO_ENABLED=0 $(GO) build -pgo=default.pgo ${LDFLAGS} -o ${BUILD_DIR}/${BINARY_NAME}
	@if [ "$$(id -u)" = "0" ]; then \
		install -m 755 ${BUILD_DIR}/${BINARY_NAME} /usr/local/bin/; \
	else \
		install -m 755 ${BUILD_DIR}/${BINARY_NAME} ${GOBIN}/; \
	fi

# run all tests (excluding large tests)
.PHONY: test
test:
	@echo "Running tests..."
	$(GO) test -v ./...

# run tests with coverage
.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	GORACE="$(GORACE)" $(GO) test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
	$(GO) tool cover -html=coverage.txt -o coverage.html
	@if [ -f "./race_report.log" ]; then \
		echo "Race conditions detected! Check race_report.log"; \
		cat "./race_report.log"; \
	fi

# run golangci-lint
.PHONY: lint
lint:
	@echo "Running linter..."
	@if ! command -v golangci-lint &> /dev/null; then \
		echo "Installing golangci-lint..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	golangci-lint run

# clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -rf ${BUILD_DIR}
	@rm -f coverage.txt coverage.html

# show help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  all            - Clean, build, and install the binary"
	@echo "  build          - Build the binary"
	@echo "  build-pgo      - Build the binary with PGO optimization"
	@echo "  install        - Install the binary in GOPATH"
	@echo "  install-pgo    - Install the binary with PGO optimization"
	@echo "  test           - Run tests (excluding large tests)"
	@echo "  test-coverage  - Run tests with coverage report"
	@echo "  lint           - Run golangci-lint"
	@echo "  clean          - Remove build artifacts"
	@echo "  help           - Show this help"
