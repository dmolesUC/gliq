package options

import (
	"os"
	"slices"
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

func Verbose() bool {
	return verboseOutput
}

func DryRun() bool {
	return dryRun
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
	repo      = "repo"
	open      = "open"
	closed    = "closed"
	labels    = "labels"
	milestone = "milestone"
	related   = "include-related"
	unrelated = "exclude-related"
	verbose   = "verbose"
	dry       = "dry-run"
)

var repository string
var accessToken string
var states State
var includeLabels []string
var includeMilestone string
var includeRelated []string
var excludeRelated []string
var verboseOutput bool
var dryRun bool

func defineFlags(cmd *cobra.Command) {
	flags := cmd.PersistentFlags()

	flags.String(repo, "", "GitLab repository, in the form <user or organization>/<repository name>")

	flags.BoolP(open, "o", true, "whether to include open issues (--open=false to exclude)")
	flags.BoolP(closed, "c", false, "whether to include closed issues")

	flags.StringSliceP(related, "r", []string{}, "include only if related to any of these issues (comma-delimited list)")
	flags.StringSliceP(unrelated, "x", []string{}, "exclude if related to any of these issues (comma-delimited list)")

	flags.StringSliceP(labels, "l", []string{}, "comma-delimited list of labels to include")

	flags.StringP(milestone, "m", "", "include only issues assigned to the specified milestone")

	flags.BoolP(verbose, "v", false, "verbose output")
	flags.BoolP(dry, "n", false, "dry run: validate parameters but do not make an API request")

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

	includeRelated, excludeRelated = relatedIssues()

	includeLabels = viper.GetStringSlice(labels)
	includeMilestone = strings.TrimSpace(viper.GetString(milestone))

	verboseOutput = viper.GetBool(verbose)
	dryRun = viper.GetBool(dry)

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

func relatedIssues() (incl []string, excl []string) {
	incl = viper.GetStringSlice(related)
	slices.Sort(incl)

	excl = viper.GetStringSlice(unrelated)
	slices.Sort(excl)

	overlap := util.Intersect(incl, excl)
	if len(overlap) > 0 {
		util.Fail("can't return issues that both are and are not related to " + strings.Join(overlap, ","))
	}

	return incl, excl
}
