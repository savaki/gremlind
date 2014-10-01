package manifest

type Program struct {
	Cmd []string
}

type Service struct {
	Port        int
	HealthCheck string
}

type Check struct {
	Script   string
	Url      string
	Interval string
}

type Manifest struct {
	Name    string
	Repo    string
	Notes   string
	Program map[string]Program
	Service map[string]Service
	Tags    map[string]string
	Check   map[string]Check
}
