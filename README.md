# kubeci
Command line tool which performs Kubernetes operations commonly required during continuous integration.

# Table Of Contents
- [Usage](#usage)
    - [Run](#run)
    - [List](#list)
    - [Help](#help)
- [Development](#development)
    - [Dependency Vendoring](#dependency-vendoring)
    - [Build](#build)
    - [Development Build](#development-build)
    - [Test](#test)
- [Release](#release)
    - [Pre Release Checklist](#pre-release-checklist)
    - [Release Build](#release-build)
    - [Uploading](#uploading)

# Usage
Kubeci is a binary which provides a command line interface. Subcommands are described in the following section.

## Run
Kubeci provides numerous operations. These operations run common Kubernetes actions.  

You can run these kubeci operations with the `run` command.  

Usage: `$ kubeci run [operation name] --os <operating system>`

Options:

- `operation name` (string): Name of operation to run
    - **required**
- `operating system` (string): Name of operating system kubeci is running commands on
    - **required**
    - Used to run different variations of certain commands (ex., package management commands)
    - Allowed values:
        - `alpine`: Alpine Linux

## List
List all available kubeci operations.

Usage: `$ kubeci list`

## Help
Provide general kubeci help or specific operation help.  

Usage: `$ kubeci help [operation name]`

Options:

- `operation name` (string): Name of operation to provide help about
    - *optional*
    - If no operation name is provided general kubeci help will be displayed
    
# Development
This section describes how to perform common development tasks for the kubeci tool.

## Dependency Vendoring
Kubeci uses [dep](https://github.com/golang/dep), Go's official dependency vendoring tool. To install it run:  

```bash
$ go get -u github.com/golang/dep/cmd/dep
```

Then to install kubeci's dependencies:

```bash
$ dep ensure
```

## Build
To build kubeci run:

```bash
$ go build -o build/kubeci ./kubeci
```

The `kubeci` binary will then be available in the `build` directory.

## Development Build
To run kubeci locally execute:

```bash
$ go run kubeci/kubeci.go
```

## Test
To test kubeci run:

```bash
$ go test -cover ./kubeci/...
```

# Release
This section describes the kubeci release process.  

Kubeci uses GNU Make in the release process, please ensure it is installed.

## Pre Release Checklist
Before a release can be built the following steps must occur:

1. Run tests
    - Run `$ go test ./kubeci/...`
1. Increment version number according to [SemVer 2.0](http://semver.org/)
    - In `kubeci/kubeci.go`
2. Tag Git commit with incremented version number
    - Run `$ VERSION=x.y.z make tag`

## Release Build
Kubeci is built for:

- Linux
    - 64-bit
    - 32-bit
- OSX
    - 64-bit
    - 32-bit
- Windows
    - 64-bit
    - 32-bit
    
To build all targets run:

```bash
$ make cross-compile-all
```

Or `make cca` for short. The kubeci binaries will be output to the `build` directory. Separated by directory for OS and 
system architecture.  

Next package all the builds up by running:

```bash
$ make package-all
```

Or `make pa` for short. The kubeci binary archives will be in the `dist` directory. Named in the form: 
`$OS-$ARCH-kubeci.tar.gz`.

Then push to the remote Git server with:  

```bash
$ git push origin master
```

# Uploading
To upload a release navigate to the [GitHub Release Creation Page](https://github.com/Noah-Huppert/kubeci/releases/new).  

Next select your newly created tag. Enter a release title and write a release description. Attatch the contents of the 
`dist` directory to the release with the upload feature.

Click the "Publish release" button to finalize the release.
