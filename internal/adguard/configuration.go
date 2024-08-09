package adguard

// Configuration holds configuration from environmental variables
type Configuration struct {
	URL              string `env:"ADGUARD_URL,notEmpty"`
	User             string `env:"ADGUARD_USER"`
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
