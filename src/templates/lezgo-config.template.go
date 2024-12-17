package templates

import "fmt"

func RenderLezgoConfigTemplate(name string) string {
	const configTemplate = `project: %v

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

# Cross-compilation
targets:
  - os: linux
    arch: amd64
    output: ./bin/linux_amd64/
  - os: windows
    arch: amd64
    output: ./bin/windows_amd64.exe

# custom scripts
scripts:
  clean: 
    command: "rm -rf ./bin/"
	`
	return fmt.Sprintf(configTemplate, name)
}
