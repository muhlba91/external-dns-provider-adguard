package adguard

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestDNSEntryFlags(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	cases := []struct {
		name          string
		config        Configuration
		expectedFlags string
	}{
		{
			name:          "minimal config for adguard provider",
			config:        Configuration{},
			expectedFlags: "",
		},
		{
			name: "enable important flag",
			config: Configuration{
				SetImportantFlag: true,
			},
			expectedFlags: ",important",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			flags := tc.config.DNSEntryFlags()
			assert.Equal(t, tc.expectedFlags, flags)
		})
	}
}
