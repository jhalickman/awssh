package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/mitchellh/goamz/aws"
	"github.com/peterh/liner"
	"github.com/spf13/viper"
)

func setupConfig() {
	viper.SetDefault("key_folder", "$HOME/.awssh/keys")
	viper.SetDefault("login_name", "ubuntu")

	viper.SetConfigName("config")

	viper.AddConfigPath("$HOME/.awssh")
	viper.AddConfigPath(".")

	viper.ReadInConfig()
}

func runSSH(keypath string, hostpart string) error {

	binary, lookErr := exec.LookPath("ssh")
	if lookErr != nil {
		return lookErr
	}

	// `Exec` requires arguments in slice form (as
	// apposed to one big string). We'll give `ls` a few
	// common arguments. Note that the first argument should
	// be the program name.
	args := []string{"ssh", "-i", keypath, hostpart}

	// `Exec` also needs a set of [environment variables](environment-variables)
	// to use. Here we just provide our current
	// environment.
	env := os.Environ()

	// Here's the actual `os.Exec` call. If this call is
	// successful, the execution of our process will end
	// here and be replaced by the `/bin/ls -a -l -h`
	// process. If there is an error we'll get a return
	// value.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		return execErr
	}
	return nil
}

func promptForInstance(instances InstanceList) (string, error) {
	term := liner.NewLiner()
	defer term.Close()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for _ = range signalChan {
			//fmt.Println("\nReceived an interrupt, stopping services...\n")
			term.Close()
			os.Exit(1)
		}
	}()

	for i, instance := range instances {
		fmt.Printf("%d) %s\n", i, instance.getName())
	}

	term.SetCompleter(func(line string) (c []string) {
		for _, instance := range instances {
			name := instance.getName()
			if strings.HasPrefix(name, strings.ToLower(line)) {
				c = append(c, name)
			}
		}
		return
	})

	line, err := term.Prompt("Which instance do you want to login to? ")

	if err != nil {
		return "", err
	}
	return line, nil
}

func main() {

	setupConfig()

	instances, err := getRunningInstances(viper.GetString("access_token"), viper.GetString("access_secret"), aws.USWest2)
	if err != nil {
		fmt.Println(err)
		return
	}

	var line string
	if len(os.Args) == 2 {
		line = os.Args[1]
	} else {
		line, err = promptForInstance(instances)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	var instance *Instance
	server_index, err := strconv.ParseInt(line, 10, 64) //Ignore error because could be string
	if err == nil {
		instance = &instances[server_index]
	}

	if instance == nil {
		instance = instances.getInstance(line)
	}

	if instance == nil {
		fmt.Printf("No instance found with name: %v\n", line)
		return
	}

	keypath := fmt.Sprintf("%s/%s.pem", viper.GetString("key_folder"), instance.KeyName)
	hostpart := fmt.Sprintf("%s@%s", viper.GetString("login_name"), instance.PrivateIpAddress)
	keypath = os.ExpandEnv(keypath)

	err = runSSH(keypath, hostpart)
	if err != nil {
		fmt.Println(err)
	}
}
