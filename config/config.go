package config

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/dmolesUC/gliq/util"
)

const (
	token  = "token"
	repo   = "repo"
	open   = "open"
	closed = "closed"
)

// Token GitLab personal access token
var Token string

// Repo GitLab repo, in the form <user or organization>/<repository name>
var Repo string

// IncludeOpen whether to include open issues
var IncludeOpen bool

// IncludeClosed whether to include closed issues
var IncludeClosed bool

func DefineFlags(cmd *cobra.Command) {
	flags := cmd.PersistentFlags()
	flags.String(token, "", "GitLab personal access token")
	flags.String(repo, "", "GitLab repo, in the form <user or organization>/<repository name>")
	flags.BoolP(open, "o", true, "Whether to include open issues (--open=false to exclude)")
	flags.BoolP(closed, "c", false, "Whether to include closed issues")
}

func Configure(flags *pflag.FlagSet) {
	readConfigFile()
	readFlags(flags)

	Token = viper.GetString(token)
	Repo = viper.GetString(repo)
	IncludeOpen = viper.GetBool(open)
	IncludeClosed = viper.GetBool(closed)
}

func readConfigFile() {
	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.SetConfigName(".gliq")

	// If a config file is found, read it in.
	err = viper.ReadInConfig()
	if err == nil {
		util.Log("Using configuration file:", viper.ConfigFileUsed())
	} else {
		util.Log("Error reading configuration file:", err)
	}
}

func readFlags(flags *pflag.FlagSet) {
	err := viper.BindPFlags(flags)
	cobra.CheckErr(err)
}
