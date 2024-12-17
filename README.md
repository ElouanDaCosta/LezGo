# LezGo
A simple build and dependency management tool for Go, offering multiple build profiles and enhanced workflows for developers

## Installation
```bash
go get -u github .com/LezGo/lezgo
```

## Usage
```bash
lezgo [command]
```

## Configuration File
LezGo uses a configuration file named `lezgo.yml` to manage build profiles and dependencies. The configuration file should be placed in the root directory of the project.

### Example Configuration File
```yaml
project: your-project-name

# define build profile
profiles:
  debug:
    flags: ["-gcflags", "all=-N -l"]
    env:
      GO_ENV: debug
    output: ./bin/debug/
    tags: ["debug"]
    ldflags: ""

  release:
    flags: ["-ldflags", "-s -w"]
    env:
      GO_ENV: production
    output: ./bin/release/
    tags: ["release"]
    ldflags: "-X main.version=1.0.0"
```

### Cross-Compilation
LezGo supports cross-compilation by specifying the target platform and architecture in the configuration file.

```yaml
# Cross-compilation
targets:
  - os: linux
    arch: amd64
    output: ./bin/linux_amd64/
  - os: windows
    arch: amd64
    output: ./bin/windows_amd64.exe
	- os: arm64
    arch: arm64
    output: ./bin/arm64/
```
