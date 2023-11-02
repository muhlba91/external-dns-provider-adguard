package adguard

// Configuration holds configuration from environmental variables
type Configuration struct {
	URL      string `env:"ADGUARD_URL,notEmpty"`
	User     string `env:"ADGUARD_USER"`
	Password string `env:"ADGUARD_PASSWORD"`
	DryRun   bool   `env:"DRY_RUN" envDefault:"false"`
}
