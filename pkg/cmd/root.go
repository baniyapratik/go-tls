package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-tls/pkg/cert"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type Config struct {
	CACert *cert.CACert          `yaml:"caCert"`
	Cert   map[string]*cert.Cert `yaml:"certs"`
}

var cfgFilePath string
var config Config
var rootCmd = &cobra.Command{
	Use:   "tls",
	Short: "tls isa command line tool for TLS",
	Long: `tls is a command line tool for TLS.
           Mainly used for generation of x509 certificates.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFilePath, "config", "c", "", "config file (default is tls.yaml")
}

func initConfig() {
	if cfgFilePath == "" {
		cfgFilePath = "tls.yaml"
	}
	cfgFileBytes, err := ioutil.ReadFile(cfgFilePath)
	if err != nil {
		fmt.Errorf("error while reading cofig file: %s", err)
		return
	}
	err = yaml.Unmarshal(cfgFileBytes, &config)
	if err != nil {
		fmt.Errorf("error while parsing cofig file: %s", err)
		return
	}
}
