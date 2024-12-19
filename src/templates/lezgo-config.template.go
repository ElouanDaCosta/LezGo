package templates

import "fmt"

func RenderLezgoConfigTemplate(name string) string {
	const configTemplate = `name: %v
author: <your_name>
entrypoint: main.go
version: 0.1.0
# define build profiles
build:
  profiles:
    debug:
      flags: ["-gcflags", "all=-N -l"]
      env:
        GO_ENV: debug
      output: ./bin/debug/
      tags: ["debug"]

    release:
      flags: ["-ldflags", "-s -w"]
      env:
        GO_ENV: production
      output: ./bin/release/
      tags: ["release"]`
	return fmt.Sprintf(configTemplate, name)
}
