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
	labels = "labels"
)

// Repo GitLab repo, in the form <user or organization>/<repository name>
var Repo string

// Token GitLab personal access token
var Token string

// IncludeOpen whether to include open issues
var IncludeOpen bool

// IncludeClosed whether to include closed issues
var IncludeClosed bool

// Labels include only issues with the specified labels
var Labels []string // TODO: should this be in here? on the root command?

func DefineFlags(cmd *cobra.Command) {
	flags := cmd.PersistentFlags()
	flags.String(repo, "", "GitLab repo, in the form <user or organization>/<repository name>")
	flags.String(token, "", "GitLab personal access token")
	flags.BoolP(open, "o", true, "Whether to include open issues (--open=false to exclude)")
	flags.BoolP(closed, "c", false, "Whether to include closed issues")
	flags.StringSliceP(labels, "l", []string{}, "comma-delimited list of labels to include")
}

func Configure(flags *pflag.FlagSet) {
	readConfigFile()
	readFlags(flags)

	Repo = ensureRepo()
	Token = viper.GetString(token)
	IncludeOpen = viper.GetBool(open)
	IncludeClosed = viper.GetBool(closed)
	Labels = viper.GetStringSlice(labels)
}

func readConfigFile() {
	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.SetConfigName(".gliq")

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

func ensureRepo() string {
	r := viper.GetString(repo)
	if r == "" {
		util.Fail("no repository specified")
	}
	return r
}
