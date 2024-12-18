package pkg

type Config struct {
	Name       string `yaml:"name"`
	Entrypoint string `yaml:"entrypoint"`
	Version    int    `yaml:"version"`
	Profiles   struct {
		Debug struct {
			Flags []string `yaml:"flags"`
			Env   struct {
				GOENV string `yaml:"GO_ENV"`
			} `yaml:"env"`
			Output  string   `yaml:"output"`
			Tags    []string `yaml:"tags"`
			Ldflags string   `yaml:"ldflags"`
		} `yaml:"debug"`
		Release struct {
			Flags []string `yaml:"flags"`
			Env   struct {
				GOENV string `yaml:"GO_ENV"`
			} `yaml:"env"`
			Output  string   `yaml:"output"`
			Tags    []string `yaml:"tags"`
			Ldflags string   `yaml:"ldflags"`
		} `yaml:"release"`
	} `yaml:"profiles"`
}
