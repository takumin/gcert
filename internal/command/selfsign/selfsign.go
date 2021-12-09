package selfsign

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"math"
	"math/big"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/takumin/gcert/internal/config"
)

func NewCommands(c *config.Config, f []cli.Flag) []*cli.Command {
	flags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "rootca",
			Aliases: []string{"ca"},
			Usage:   "root ca selfsign",
			EnvVars: []string{"ROOTCA"},
		},
	}
	return []*cli.Command{
		{
			Name:    "selfsign",
			Aliases: []string{"self"},
			Usage:   "generate selfsign certificate",
			Flags:   append(flags, f...),
			Action:  action(c),
		},
	}
}

func action(c *config.Config) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		pub, priv, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			return err
		}

		serial, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
		if err != nil {
			return err
		}

		tmpl := &x509.Certificate{
			SerialNumber: serial,
		}

		crt, err := x509.CreateCertificate(rand.Reader, tmpl, tmpl, pub, priv)
		if err != nil {
			return err
		}

		if err := pem.Encode(os.Stdout, &pem.Block{
			Type:  "CERTIFICATE",
			Bytes: crt,
		}); err != nil {
			return err
		}

		return nil
	}
}
