package cobra

// Serpent is a simple API that makes working with Cobra a bit easier. The API
// allows you to get straight to work by skipping over much of the complexity
// required to set up Cobra and Viper. Serpent's downside is that it makes a
// bunch of assumptions about how you want Cobra and Viper configured. If
// those assumptions are incorrect for your project, all the original bits of
// Cobra are still present so you can always drop back into snake charming the
// fearsome Cobra.

import (
	"fmt"
	"os"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var Serpent struct {
	Root    *Command
	LogFile os.File
}

type SerpentError struct {
	msg string
}

func (e SerpentError) Error() string {
	return e.msg
}

type App struct {
	Command
}

func NewApp(n string) (a *App) {
	a = &App{}
	a.Use = n
	return
}

type Config struct {
	RequireConfigFile       bool   // if true, execution stops if a config file is not found
	ReportMissingConfigFile bool   // logs an alert if the config file is optional and can't be found
	RequireLogFile          bool   // if true, execution stops if the log file cannot be created
	ReportMissingLogFile    bool   // logs an alert if the logfile is optional and cannot be created
	ConfigFile              string // the name of the config file without the extension
	SearchLocalConfig       bool   // if true, viper looks in the working directory for a config file
	UserConfigPath          string // a path relative the user's home directory that viper searches
	GlobalConfigPath        string // an arbitrary path to look for config files
	WatchConfig             bool   // if true, viper watches for changes in the config file
	UseEnvVariables         bool   // if true, viper looks for configs in environment variables
	EnvVarPrefix            string // viper only looks at environment variables with the prefix
}

func NewTestingConfig(logtags []string) *Config {
	c := &Config{}
	c.SetDefault("verbose", true)
	c.SetDefault("logtags", logtags)
	m := map[string]interface{}{}
	viper.Set("serpenttags", m)

	stdout := new(SerpentFormatter)
	stdout.Debug = true
	log.SetFormatter(stdout)
	log.SetLevel(log.DebugLevel)

	// Set up log tags.
	tags := GetStringSlice("logtags")
	if len(tags) > 0 {
		dt := viper.GetStringMap("serpenttags")
		for _, t := range tags {
			dt[t] = true
		}
	}

	return c
}

func NewConfig(n string) *Config {
	c := &Config{}
	c.ConfigFile = "config"
	c.SearchLocalConfig = true

	// Configure logging
	c.RequireLogFile = true
	c.ReportMissingLogFile = true
	m := map[string]interface{}{}
	viper.Set("serpenttags", m)

	// Configure stdout logging
	stdout := new(SerpentFormatter)
	log.SetFormatter(stdout)
	log.SetLevel(log.DebugLevel)

	return c
}

func (c *Config) SetDefault(k string, v interface{}) {
	viper.SetDefault(k, v)
}

func (c *Config) Set(k string, v interface{}) {
	viper.Set(k, v)
}

func Set(k string, v interface{}) {
	viper.Set(k, v)
}

func (c *Config) LogPanicOnly() {
	log.SetLevel(log.PanicLevel)
}

func (c *Config) LogNormal() {
	log.SetLevel(log.DebugLevel)
}

func NewCommand(n string) (c *Command) {
	c = &Command{
		Use: n,
		PreRun: func(cmd *Command, args []string) {
			// These two VisitAll calls are to bind the flags as late as
			// possible in order to support flags on different commands with
			// the same name but with different configurations, such as
			// different defult values.
			cmd.pflags.VisitAll(
				func(f *pflag.Flag) { viper.BindPFlag(f.Name, cmd.PersistentFlags().Lookup(f.Name)) },
			)
			cmd.flags.VisitAll(
				func(f *pflag.Flag) { viper.BindPFlag(f.Name, c.Flags().Lookup(f.Name)) },
			)
		},
	}
	return
}

// Init initializes Serpent and returns the root Command.
func Init(app *App, cfg *Config) (cmd *Command) {
	cmd = &app.Command
	cmd.AddFlags(NewStringFlag("config", Opts().Ubiq(true)))
	cmd.AddFlags(NewBoolFlag("verbose", Opts().Ubiq(true).Abbr("v")))
	cmd.AddFlags(NewStringFlag("log", Opts().Desc("path to the log file").Ubiq(true)))
	cmd.AddFlags(NewStringSliceFlag("logtags", Opts().Ubiq(true)))
	cmd.AddFlags(NewBoolFlag("logalltags", Opts().Ubiq(true)))
	OnInitialize(func() { loadConfigs(app.Use, cfg) })
	Serpent.Root = cmd
	return
}

// Cobra executes this function after the command line has been parsed, but
// before actual execution of the command.
func loadConfigs(n string, cfg *Config) {
	// Do we want to read environment variables, maybe with a prefix?
	if cfg.UseEnvVariables {
		if len(cfg.EnvVarPrefix) > 0 {
			viper.SetEnvPrefix(cfg.EnvVarPrefix)
		}
		viper.AutomaticEnv()
	}

	// Set config file name.
	if c, ok := CheckString("config"); ok {
		viper.SetConfigName(c)
	} else {
		viper.SetConfigName(cfg.ConfigFile)
	}

	// Set paths to check for config file.
	cfg.UserConfigPath = fmt.Sprintf(".%s", n)
	viper.AddConfigPath(".")
	viper.AddConfigPath(getHomedir(cfg.UserConfigPath))
	viper.AddConfigPath(fmt.Sprintf("%s/", cfg.GlobalConfigPath))

	// Load the config file.
	if err := viper.ReadInConfig(); err != nil {
		switch {
		case cfg.RequireConfigFile:
			Out("exiting: can't read config file: ", err)
			os.Exit(1)
		case cfg.ReportMissingConfigFile:
			Out("continuing without a config file")
		}
	}

	// Watch the config file for changes if requested.
	if cfg.WatchConfig {
		viper.WatchConfig()
	}

	// Set up log tags.
	tags := GetStringSlice("logtags")
	if len(tags) > 0 {
		dt := viper.GetStringMap("serpenttags")
		for _, t := range tags {
			dt[t] = true
		}
	}

	cfg.SetDefault("log", n+".log")
	LoadLogFile(cfg)

	// // Configure file logging via a logrus hook.
	// f, err := os.Create(GetString("log"))
	// switch {
	// case err != nil && cfg.RequireLogFile:
	// 	Out("exiting: can't access log file: ", err)
	// 	os.Exit(1)
	// case err != nil && cfg.ReportMissingLogFile:
	// 	Out("continuing without a log file")
	// default:
	// 	tf := new(SerpentFormatter)
	// 	tf.Debug = true
	// 	log.AddHook(lfshook.NewHook(f, tf))
	// }
}

func LoadLogFile(cfg *Config) {
	// Configure file logging via a logrus hook.
	f, err := os.Create(GetString("log"))
	switch {
	case err != nil && cfg.RequireLogFile:
		Out("exiting: can't access log file: ", err)
		os.Exit(1)
	case err != nil && cfg.ReportMissingLogFile:
		Out("continuing without a log file")
	default:
		tf := new(SerpentFormatter)
		tf.Debug = true
		log.AddHook(lfshook.NewHook(f, tf))
	}
}

func (c *Command) SubCmds(subc ...*Command) *Command {
	if c.subCmds == nil {
		c.subCmds = []*Command{}
	}
	for _, sub := range subc {
		c.subCmds = append(c.subCmds, sub)
	}
	return c
}

func getHomedir(p string) string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err) // TODO: Log!!!!!!!!!!!!
		return ""
	}
	return fmt.Sprintf("%s/%s/", home, p)
}

type adder interface {
	AddTo(*Command)
}

func (c *Command) AddFlags(flags ...adder) {
	for _, f := range flags {
		f.AddTo(c)
	}
}

func buildCommands(c *Command) *Command {
	for _, sub := range c.subCmds {
		buildCommands(sub)
	}
	c.AddCommand(c.subCmds...)
	return c
}

// parses the command line and executes the command
func Execute() {
	root := buildCommands(Serpent.Root)
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ShutDown() {
	Serpent.LogFile.Close()
}

// Serpent flag system --------------------------------------------------------

type SerpentFlag struct {
	Name    string      // the flag's name as used on the command line
	Abbr    string      // one letter appreviation of the name
	Desc    string      // the flag's description
	Req     bool        // true if the flag must be set on the command line
	Ubiq    bool        // true if the flag is inherited by subcommands
	Default interface{} // the default value of the flag when it is omitted
	Implied string      // the value of the flag when the flag has no argument
	Hide    bool        // if true, the flag doesn't show up in help
}

func (f *SerpentFlag) populate(n string, d interface{}, opts ...*FlagOptList) {
	f.Name = n
	for _, optlist := range opts {
		for _, opt := range *optlist {
			switch opt.FlagOptType {
			case cAbbreviation:
				f.Abbr = opt.Value.(string)
			case cDefault:
				if fmt.Sprintf("%T", opt.Value) != fmt.Sprintf("%T", d) {
					panic(SerpentError{msg: fmt.Sprintf("default value for flag %q is wrong type: %T instead of %T", n, opt.Value, d)})
				}
				f.Default = opt.Value
			case cImplied:
				if fmt.Sprintf("%T", opt.Value) != fmt.Sprintf("%T", d) {
					panic(SerpentError{msg: fmt.Sprintf("implied value for flag %q is wrong type: %T instead of %T", n, opt.Value, d)})
				}
				f.Implied = opt.Value.(string)
			case cDescription:
				f.Desc = opt.Value.(string)
			case cRequired:
				f.Req = true
			case cNotRequired:
				f.Req = false
			case cUbiquitous:
				f.Ubiq = true
			case cNotUbiquitous:
				f.Ubiq = false
			case cHide:
				f.Hide = true
			}
		}
	}
}

func (f *SerpentFlag) postAdd(c *Command) {
	switch {
	case f.Req && f.Ubiq:
		c.MarkPersistentFlagRequired(f.Name)
		// viper.BindPFlag(f.Name, c.PersistentFlags().Lookup(f.Name))
	case f.Req && !f.Ubiq:
		c.MarkFlagRequired(f.Name)
		// viper.BindPFlag(f.Name, c.Flags().Lookup(f.Name))
	case !f.Req && f.Ubiq:
		// viper.BindPFlag(f.Name, c.PersistentFlags().Lookup(f.Name))
	default:
		// viper.BindPFlag(f.Name, c.Flags().Lookup(f.Name))
	}

	switch {
	case len(f.Implied) > 0 && f.Ubiq:
		c.PersistentFlags().Lookup(f.Name).NoOptDefVal = f.Implied
	case len(f.Implied) > 0 && !f.Ubiq:
		c.Flags().Lookup(f.Name).NoOptDefVal = f.Implied
	}

	switch {
	case f.Hide && f.Ubiq:
		c.PersistentFlags().MarkHidden(f.Name)
	case f.Hide && !f.Ubiq:
		c.Flags().MarkHidden(f.Name)
	}
}

// func bindFlag(c *Command, f *pflag.Flag) {
// 	switch {
// 	case f.Req && f.Ubiq:
// 		// c.MarkPersistentFlagRequired(f.Name)
// 		viper.BindPFlag(f.Name, c.PersistentFlags().Lookup(f.Name))
// 	case f.Req && !f.Ubiq:
// 		// c.MarkFlagRequired(f.Name)
// 		viper.BindPFlag(f.Name, c.Flags().Lookup(f.Name))
// 	case !f.Req && f.Ubiq:
// 		viper.BindPFlag(f.Name, c.PersistentFlags().Lookup(f.Name))
// 	default:
// 		viper.BindPFlag(f.Name, c.Flags().Lookup(f.Name))
// 	}
// }

type FlagOpt struct {
	FlagOptType
	Value interface{}
}

type FlagOptType int

const (
	cAbbreviation FlagOptType = iota
	cDefault
	cDescription
	cEmpty
	cHide
	cImplied
	cNotRequired
	cNotUbiquitous
	cRequired
	cUbiquitous
)

type FlagOptList []FlagOpt

func Opts() *FlagOptList {
	f := FlagOptList{FlagOpt{cEmpty, nil}}
	return &f
}

func Reset() *FlagOptList {
	f := FlagOptList{FlagOpt{cEmpty, nil}}
	return &f
}

func (f *FlagOptList) Abbr(v string) *FlagOptList {
	*f = append(*f, FlagOpt{cAbbreviation, v})
	return f
}

func (f *FlagOptList) Hide() *FlagOptList {
	*f = append(*f, FlagOpt{cHide, true})
	return f
}

func (f *FlagOptList) Implied(v string) *FlagOptList {
	*f = append(*f, FlagOpt{cImplied, v})
	return f
}

func (f *FlagOptList) Default(d interface{}) *FlagOptList {
	*f = append(*f, FlagOpt{cDefault, d})
	return f
}

func (f *FlagOptList) Desc(v string) *FlagOptList {
	*f = append(*f, FlagOpt{cDescription, v})
	return f
}

func (f *FlagOptList) Req(b bool) *FlagOptList {
	*f = append(*f, FlagOpt{cRequired, b})
	return f
}

func (f *FlagOptList) Ubiq(b bool) *FlagOptList {
	*f = append(*f, FlagOpt{cUbiquitous, b})
	return f
}

// Access Viper values --------------------------------------------------------

// CheckString returns the value associated with the key as a string and a
// bool set to true if the key was found.
func CheckString(key string) (string, bool) { return viper.GetString(key), viper.IsSet(key) }

// GetString returns just the value associated with the key as a string.
func GetString(key string) string { return viper.GetString(key) }

// CheckBool returns the value associated with the key as a boolean and a bool
// set to true if the key was found.
func CheckBool(key string) (bool, bool) { return viper.GetBool(key), viper.IsSet(key) }

// GetBool just returns the value associated with the key as a boolean.
func GetBool(key string) bool { return viper.GetBool(key) }

// CheckInt returns the value associated with the key as an integer and a bool
// set to true if the key was found.
func CheckInt(key string) (int, bool) { return viper.GetInt(key), viper.IsSet(key) }

// GetInt just returns the value associated with the key as an integer.
func GetInt(key string) int { return viper.GetInt(key) }

// CheckInt32 returns the value associated with the key as an integer and a
// bool set to true if the key was found.
func CheckInt32(key string) (int32, bool) { return viper.GetInt32(key), viper.IsSet(key) }

// GetInt32 just returns the value associated with the key as an integer.
func GetInt32(key string) int32 { return viper.GetInt32(key) }

// CheckInt64 returns the value associated with the key as an integer and a
// bool set to true if the key was found.
func CheckInt64(key string) (int64, bool) { return viper.GetInt64(key), viper.IsSet(key) }

// GetInt64 just returns the value associated with the key as an integer.
func GetInt64(key string) int64 { return viper.GetInt64(key) }

// CheckFloat64 returns the value associated with the key as a float64 and a
// bool set to true if the key was found.
func CheckFloat64(key string) (float64, bool) { return viper.GetFloat64(key), viper.IsSet(key) }

// GetFloat64 just returns the value associated with the key as a float64.
func GetFloat64(key string) float64 { return viper.GetFloat64(key) }

// CheckTime returns the value associated with the key as time and a bool set
// to true if the key was found.
func CheckTime(key string) (time.Time, bool) { return viper.GetTime(key), viper.IsSet(key) }

// GetTime just returns the value associated with the key as time.
func GetTime(key string) time.Time { return viper.GetTime(key) }

// CheckDuration returns the value associated with the key as a duration and a
// bool set to true if the key was found.
func CheckDuration(key string) (time.Duration, bool) { return viper.GetDuration(key), viper.IsSet(key) }

// GetDuration just returns the value associated with the key as a duration.
func GetDuration(key string) time.Duration { return viper.GetDuration(key) }

// CheckStringSlice returns the value associated with the key as a slice of
// strings and a bool set to true if the key was found.
func CheckStringSlice(key string) ([]string, bool) { return viper.GetStringSlice(key), viper.IsSet(key) }

// GetStringSlice just returns the value associated with the key as a slice of
// strings.
func GetStringSlice(key string) []string { return viper.GetStringSlice(key) }

// CheckStringMap returns the value associated with the key as a map of
// interfaces and a bool set to true if the key was found.
func CheckStringMap(key string) (map[string]interface{}, bool) {
	return viper.GetStringMap(key), viper.IsSet(key)
}

// GetStringMap just returns the value associated with the key as a map of
// interfaces.
func GetStringMap(key string) map[string]interface{} { return viper.GetStringMap(key) }

// CheckStringMapString returns the value associated with the key as a map of
// strings and a bool set to true if the key was found.
func CheckStringMapString(key string) (map[string]string, bool) {
	return viper.GetStringMapString(key), viper.IsSet(key)
}

// GetStringMapString just returns the value associated with the key as a map
// of strings.
func GetStringMapString(key string) map[string]string { return viper.GetStringMapString(key) }

// CheckStringMapStringSlice returns the value associated with the key as a
// map to a slice of strings and a bool set to true if the key was found.
func CheckStringMapStringSlice(key string) (map[string][]string, bool) {
	return viper.GetStringMapStringSlice(key), viper.IsSet(key)
}

// GetStringMapStringSlice just returns the value associated with the key as a
// map to a slice of strings.
func GetStringMapStringSlice(key string) map[string][]string {
	return viper.GetStringMapStringSlice(key)
}

// CheckSizeInBytes returns the size of the value associated with the given
// key in bytes and a bool set to true if the key was found.
func CheckSizeInBytes(key string) (uint, bool) { return viper.GetSizeInBytes(key), viper.IsSet(key) }

// GetSizeInBytes just returns the size of the value associated with the given
// key in bytes.
func GetSizeInBytes(key string) uint { return viper.GetSizeInBytes(key) }
