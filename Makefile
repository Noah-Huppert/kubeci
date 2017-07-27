.PHONY: cca pa tag

NO_COLOR=\033[0m

OK_TAG=OK   #
OK_COLOR=\033[32;01m

WARN_TAG=WARN #
WARN_COLOR=\033[33;01m

ERROR_TAG=ERROR
ERROR_COLOR=\033[31;01m

define log = # level, message
$(if $(findstring ok, $(1)), @printf "$(OK_COLOR)[$(OK_TAG)] $(2)$(NO_COLOR)\n")
$(if $(findstring warn, $(1)), @printf "$(WARN_COLOR)[$(WARN_TAG)] $(2)$(NO_COLOR)\n")
$(if $(findstring error, $(1)), @printf "$(ERROR_COLOR)[$(ERROR_TAG)] $(2)$(NO_COLOR)\n")
endef

THIS_FILE := $(lastword $(MAKEFILE_LIST))

# Executable extension, windows has .exe, all other platforms have none
EXT =
ifeq ("${OS}","windows")
EXT = .exe
endif

clean:
	$(call log,ok,Cleaning build and dist directory)
	rm -rf ./build
	rm -rf ./dist

# cross-compile: Build kubeci for specific platform
#
# Args:
#   - OS (string): Operating system to build for
#   - ARCH (string): System architecture to build for
cross-compile:
ifeq ($(OS),)
	$(call log,error,OS argument must be provided to build command)
	@exit 1
endif

ifeq ($(ARCH),)
	$(call log,error,ARCH argument must be provided to build command)
	@exit 1
endif

	$(call log,ok,Building for $(OS) $(ARCH))
	mkdir -p build/$(OS)/$(ARCH)
	GOOS=$(OS) GOARCH=$(ARCH) go build -o build/$(OS)/$(ARCH)/kubeci ./kubeci

ifeq ($(OS),windows)
	mv build/$(OS)/$(ARCH)/kubeci build/$(OS)/$(ARCH)/kubeci.exe
endif

	$(call log,ok,Build OK)

# cross-compile-all: Builds kubeci for all platform targets:
# 	- Linux
#       - 64-bit
#       - 32-bit
# 	- OSX
#       - 64-bit
#       - 32-bit
# 	- Windows
#       - 64-bit
#       - 32-bit
cross-compile-all:
	@OS=linux ARCH=amd64 ${MAKE} -f ${THIS_FILE} cross-compile
	@OS=linux ARCH=386 ${MAKE} -f ${THIS_FILE} cross-compile

	@OS=darwin ARCH=amd64 ${MAKE} -f ${THIS_FILE} cross-compile
	@OS=darwin ARCH=386 ${MAKE} -f ${THIS_FILE} cross-compile

	@OS=windows ARCH=amd64 ${MAKE} -f ${THIS_FILE} cross-compile
	@OS=windows ARCH=386 ${MAKE} -f ${THIS_FILE} cross-compile

# cca: Shorthand for cross-compile-all
cca:
	${MAKE} -f ${THIS_FILE} cross-compile-all

# package: Package kubeci build
#
# Args:
#   - OS (string): Operating system to build for
#   - ARCH (string): System architecture to build for
package:
ifeq ($(OS),)
	$(call log,error,OS argument must be provided to package command)
	@exit 1
endif

ifeq ($(ARCH),)
	$(call log,error,ARCH argument must be provided to package command)
	@exit 1
endif

ifeq ("$(wildcard build/${OS}/${ARCH}/kubeci${EXT})","")
	$(call log,error,kubeci not built for ${OS} ${ARCH})
	exit 1
endif

	$(call log,ok,Packaging kubeci build for ${OS} ${ARCH})
	mkdir -p dist
	tar -zcvf dist/${OS}-${ARCH}-kubeci.tar.gz build/${OS}/${ARCH}/kubeci${EXT}

	$(call log,ok,Package OK)

# package-all: Package kubeci builds for all platforms
package-all:
	@OS=linux ARCH=amd64 ${MAKE} -f ${THIS_FILE} package
	@OS=linux ARCH=386 ${MAKE} -f ${THIS_FILE} package

	@OS=darwin ARCH=amd64 ${MAKE} -f ${THIS_FILE} package
	@OS=darwin ARCH=386 ${MAKE} -f ${THIS_FILE} package

	@OS=windows ARCH=amd64 ${MAKE} -f ${THIS_FILE} package
	@OS=windows ARCH=386 ${MAKE} -f ${THIS_FILE} package

# pa: package-all short alias
pa:
	@${MAKE} -f ${THIS_FILE} package-all

# tag: Make a Git tag for version on specific branch
#
# Args:
#	- VERSION (string): SemVer 2.0 version to tag
tag:
ifeq ($(VERSION),)
	$(call log,error,VERSION argument must be provided to tag command)
	@exit 1
endif

	$(call log,ok,Tagging current commit as version ${VERSION})

	git tag -a v${VERSION} -m "kubeci version ${VERSION}"

	$(call log,ok,Tag OK)
