package adguard

// Configuration holds configuration from environmental variables
type Configuration struct {
	URL  string `env:"ADGUARD_URL,notEmpty"`
	User string `env:"ADGUARD_USER"`
	//nolint:gosec // this is not a secret, but a password for the AdGuard Home API, which is usually not exposed to the internet
	Password         string `env:"ADGUARD_PASSWORD"`
	DryRun           bool   `env:"DRY_RUN" envDefault:"false"`
	SetImportantFlag bool   `env:"ADGUARD_SET_IMPORTANT_FLAG" envDefault:"true"`
}

// DNSEntryFlags returns additional flags set for DNS entries
func (c *Configuration) DNSEntryFlags() string {
	flags := ""
	if c.SetImportantFlag {
		flags += ",important"
	}
	return flags
}
