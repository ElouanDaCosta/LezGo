package templates

import "fmt"

func RenderLezgoConfigTemplate(name string) string {
	const configTemplate = `name: %v

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
    ldflags: "-X main.version=1.0.0"`
	return fmt.Sprintf(configTemplate, name)
}
