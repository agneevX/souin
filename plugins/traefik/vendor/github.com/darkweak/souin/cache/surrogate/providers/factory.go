package providers

import (
	"github.com/darkweak/souin/configurationtypes"
)

func SurrogateFactory(config configurationtypes.AbstractConfigurationInterface) SurrogateInterface {
	cdn := config.GetDefaultCache().GetCDN()

	switch cdn.Provider {
	case "akamai":
		return generateAkamaiInstance(config)
	case "fastly":
		return generateFastlyInstance(config)
	default:
		return generateSouinInstance(config)
	}
}