package pkg

type Config struct {
	Project  string `yaml:"project"`
	Profiles struct {
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