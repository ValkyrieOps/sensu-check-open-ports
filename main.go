package main

import (
	"fmt"
        "os"
        "os/exec"
	"strconv"
	"strings"
	"io/ioutil"
	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	"github.com/sensu/sensu-go/types"
)

// Config represents the check plugin config.
type Config struct {
        sensu.PluginConfig
        User          string
        Warn          int
        Crit          int
}

var (
        plugin = Config{
                PluginConfig: sensu.PluginConfig{
                        Name:     "sensu-check-open-ports",
                        Short:    "Sensu Check for Open Ports",
                        Keyspace: "sensu.io/plugins/sensu-ofd-check/config",
                },
        }

        options = []*sensu.PluginConfigOption{
                &sensu.PluginConfigOption{
                        Path:      "user",
                        Env:       "CHECK_USER",
                        Argument:  "user",
                        Shorthand: "u",
                        Default:   "sensu",
                        Usage:     "User to query for open port count",
                        Value:     &plugin.User,
                },
                &sensu.PluginConfigOption{
                        Path:      "warn",
                        Env:       "CHECK_WARN",
                        Argument:  "warn",
                        Shorthand: "w",
                        Default:   nil,
                        Usage:     "Warning threshold - count of open ports required for warning state",
                        Value:     &plugin.Warn,
                },
                &sensu.PluginConfigOption{
                        Path:      "crit",
                        Env:       "CHECK_CRITICAL",
                        Argument:  "crit",
                        Shorthand: "c",
                        Default:   nil,
                        Usage:     "Critical threshold - count of open ports required for critical state",
                        Value:     &plugin.Crit,
                },
        }
)


func main() {
	check := sensu.NewGoCheck(&plugin.PluginConfig, options, checkArgs, executeCheck, false)
	check.Execute()
}

func checkArgs(event *types.Event) (int, error) {
	if len(plugin.User) == 0 {
		return sensu.CheckStateWarning, fmt.Errorf("--user or CHECK_USER environment variable is required")
	}
	return sensu.CheckStateOK, nil
}

func handleError(err error) {
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
}

func executeCheck(event *types.Event) (int, error) {
	get_lsof := exec.Command("lsof", "-i")
	get_user := exec.Command("grep", plugin.User)
	get_count := exec.Command("wc", "-l")

	get_user.Stdin, _ = get_lsof.StdoutPipe()
	get_count.Stdin, _ = get_user.StdoutPipe()

	stdout, err := get_count.StdoutPipe()
	handleError(err)

	get_count.Start()
	get_user.Start()
	get_lsof.Start()
	err = get_lsof.Start()


	defer get_count.Wait()
	defer get_user.Wait()

	ports, err := ioutil.ReadAll(stdout)

	int_ports, _ := strconv.Atoi(strings.TrimSuffix(string(ports), "\n"))

	if int_ports >= plugin.Warn && int_ports < plugin.Crit {
                        fmt.Println("WARNING\nOpen Ports for " + plugin.User + ":", int_ports)
                        return sensu.CheckStateWarning, nil
                } else if int_ports >= plugin.Crit {
                        fmt.Println("CRITICAL\nOpen Ports for " + plugin.User + ":", int_ports)
                        return sensu.CheckStateCritical, nil
                } else {
                        fmt.Println("OK\nOpen Ports for " + plugin.User + ":", int_ports)
                        return sensu.CheckStateOK, nil
                }

}
