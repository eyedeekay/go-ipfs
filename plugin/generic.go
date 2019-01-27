package plugin

// PluginGeneric is an interface that can be implemented to trigger functionality
// using the IPFS API or an external API like those provided by Tor or I2P
type PluginGeneric interface {
	Plugin

	Exists() error
}
