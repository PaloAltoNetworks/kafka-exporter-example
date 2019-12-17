package configuration

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"go.aporeto.io/addedeffect/lombric"
	"go.aporeto.io/elemental/cmd/elegen/versions"
	"go.aporeto.io/tg/tglib"
)

// Configuration holds the service configuration
type Configuration struct {
	LogFormat                string `mapstructure:"log-format"                  desc:"Log format"                                               default:"json"`
	LogLevel                 string `mapstructure:"log-level"                   desc:"Log level"                                                default:"info"`
	CACertPath               string `mapstructure:"cacert"                      desc:"Path to the Aporeto CA"                                   required:"true"`
	ClientCertificatePath    string `mapstructure:"client-cert"                 desc:"Path to the client certificate"                           required:"true"`
	ClientCertificateKeyPath string `mapstructure:"client-cert-key"             desc:"Path to the client certificate key"                       required:"true"`
	ClientCertificateKeyPass string `mapstructure:"client-cert-key-pass"        desc:"Password of the client certificate key"                   required:"true"`
	NATSPassword             string `mapstructure:"pubsub-pass"                 desc:"Password to use to connect to Nats"                       secret:"true"`
	NATSUser                 string `mapstructure:"pubsub-user"                 desc:"User name to use to connect to Nats"                      secret:"true"`
	NATSClientID             string `mapstructure:"pubsub-client-id"            desc:"Nats client ID"                                           `
	NATSClusterID            string `mapstructure:"pubsub-cluster-id"           desc:"Nats cluster ID"                                          default:"test-cluster"`
	NATSAddress              string `mapstructure:"pubsub-address"              desc:"Nats Address"                                             required:"true"`

	CAPool               *x509.CertPool
	ClientCertificate    string
	ClientCertificateKey string
	ClientCertificates   []tls.Certificate
	NATSQueueName        string
}

// Prefix returns the configuration prefix.
func (c *Configuration) Prefix() string { return "kafka-exporter" }

// PrintVersion prints the current version.
func (c *Configuration) PrintVersion() {
	fmt.Printf("kafka-exporter - %s (%s)\n", versions.ProjectVersion, versions.ProjectSha)
}

// NewConfiguration returns a new Configuration.
func NewConfiguration() *Configuration {

	c := &Configuration{}
	lombric.Initialize(c)

	pool, err := x509.SystemCertPool()
	if err != nil {
		panic("Unable to get system cert pool: " + err.Error())
	}
	cadata, err := ioutil.ReadFile(c.CACertPath)
	if err != nil {
		panic("Unable to read CA file: " + err.Error())
	}
	pool.AppendCertsFromPEM(cadata)
	c.CAPool = pool

	x509Cert, pkey, err := tglib.ReadCertificatePEM(c.ClientCertificatePath, c.ClientCertificateKeyPath, c.ClientCertificateKeyPass)
	if err != nil {
		panic("Unable to read certificate: " + err.Error())
	}
	cert, err := tglib.ToTLSCertificate(x509Cert, pkey)
	if err != nil {
		panic("Unable to read certificate: " + err.Error())
	}
	c.ClientCertificates = []tls.Certificate{cert}

	// All instances must listen a given queue so only one of them is processing a given publication.
	c.NATSQueueName = "de839119-6dad-4ca9-8888-181e785f3b53"

	return c
}
