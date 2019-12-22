package version

// -X 'github.com/fox-one/pkg/version.Version=$(VERSION)'
// -X 'github.com/fox-one/pkg/version.Commit=$(COMMIT)'
//

var (
	Version string
	Commit  string
)

func Print() string {
	if Commit != "" {
		return Version + "-" + Commit
	}

	return Version
}
