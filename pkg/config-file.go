package pkg

type Config struct {
	Name       string `yaml:"name"`
	Author     string `yaml:"author"`
	Entrypoint string `yaml:"entrypoint"`
	Version    string `yaml:"version"`
	Build      struct {
		Profiles struct {
			Debug struct {
				Flags []string `yaml:"flags"`
				Env   struct {
					GOENV string `yaml:"GO_ENV"`
				} `yaml:"env"`
				Output   string   `yaml:"output"`
				Tags     []string `yaml:"tags"`
				Ldflags  string   `yaml:"ldflags"`
				OsTarget string   `yaml:"os-target"`
				Arch     string   `yaml:"arch"`
			} `yaml:"debug"`
			Release struct {
				Flags []string `yaml:"flags"`
				Env   struct {
					GOENV string `yaml:"GO_ENV"`
				} `yaml:"env"`
				Output   string   `yaml:"output"`
				Tags     []string `yaml:"tags"`
				Ldflags  string   `yaml:"ldflags"`
				OsTarget string   `yaml:"os-target"`
				Arch     string   `yaml:"arch"`
			} `yaml:"release"`
		} `yaml:"profiles"`
	} `yaml:"build"`
}
