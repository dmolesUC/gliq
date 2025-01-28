package options

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/dmolesUC/gliq/util"
)

// ------------------------------------------------------------
// Exported

func Repository() string {
	return repository
}

func AccessToken() string {
	return accessToken
}

func Milestone() string {
	return includeMilestone
}

func Labels() []string {
	return includeLabels
}

func StateFlags() State {
	return states
}

func InitOptions(cmd *cobra.Command) {
	defineFlags(cmd)
	cobra.OnInitialize(func() {
		configure(cmd.PersistentFlags())
	})
}

// ------------------------------------------------------------
// Package-local

const (
	token     = "token"
	repo      = "repository"
	open      = "open"
	closed    = "closed"
	labels    = "labels"
	milestone = "milestone"
)

var repository string
var accessToken string
var states State
var includeLabels []string
var includeMilestone string

func defineFlags(cmd *cobra.Command) {
	flags := cmd.PersistentFlags()

	flags.StringP(repo, "r", "", "GitLab repo, in the form <user or organization>/<repository name>")

	flags.BoolP(open, "o", true, "Whether to include open issues (--open=false to exclude)")
	flags.BoolP(closed, "c", false, "Whether to include closed issues")

	// TODO: should this be on the root command?
	flags.StringSliceP(labels, "l", []string{}, "comma-delimited list of labels to include")

	// TODO: should this be on the root command?
	flags.StringP(milestone, "m", "", "include only issues assigned to the specified milestone")

	flags.String(token, "", "GitLab personal access token")
}

func configure(flags *pflag.FlagSet) {
	readConfigFile()
	readFlags(flags)

	repository = ensureRepo()
	accessToken = strings.TrimSpace(viper.GetString(token))

	var includeOpen = viper.GetBool(open)
	var includeClosed = viper.GetBool(closed)
	states = toState(includeOpen, includeClosed)

	includeLabels = viper.GetStringSlice(labels)
	includeMilestone = strings.TrimSpace(viper.GetString(milestone))

	// TODO: use issue/:iid/links endpoint to filter on related/unrelated
}

func readConfigFile() {
	// Find home directory.
	home := util.Safely(os.UserHomeDir)

	viper.AddConfigPath(home)
	viper.SetConfigName(".gliq")

	err := viper.ReadInConfig()
	if err == nil {
		util.Log("Using configuration file:", viper.ConfigFileUsed())
	} else {
		util.Log("Error reading configuration file:", err)
	}
}

func readFlags(flags *pflag.FlagSet) {
	err := viper.BindPFlags(flags)
	util.QuietlyHandle(err)
}

func ensureRepo() string {
	r := viper.GetString(repo)
	if r == "" {
		util.Fail("no repository specified")
	}
	return r
}
