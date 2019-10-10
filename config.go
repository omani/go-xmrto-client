package xmrto

// Config holds the configuration of a xmr.to API client.
type Config struct {
	APIBaseAddress         string
	APIVersion             string
	APIConversionDirection string
	Testnet                bool
}
