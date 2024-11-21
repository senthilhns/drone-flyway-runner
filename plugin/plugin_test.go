// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func IsTestFunctionality() bool {
	return os.Getenv("IS_TEST_FUNCTIONALITY") == "TRUE"
}

func TestFunctionalityCleanWithArgs(t *testing.T) {
	if !IsTestFunctionality() {
		return
	}
	args := GetArgsForFunctionalTesting(
		getDefaultPluginDriverPath(),
		"clean",
		getDefaultPluginLocations(),
		getDefaultPluginCommandLineArgs(),
		getDefaultPluginUrl(),
		getDefaultPluginUser(),
		getDefaultPluginPassword())

	fmt.Println(args.ToStr())

	fp, err := Exec(context.TODO(), args)
	if err != nil {
		fmt.Println("Error in Exec: " + err.Error())
		t.Fail()
	}
	_ = fp
}

func TestFunctionalityCleanWithConfigFile(t *testing.T) {
	if !IsTestFunctionality() {
		return
	}
	args := GetArgsForFunctionalTesting(
		"",                         // driver path
		"clean",                    // command
		"",                         // location
		ConfigFileOpt+"",           // command line args
		getDefaultPluginUrl(),      // url
		getDefaultPluginUser(),     //  user
		getDefaultPluginPassword()) // password

		args.IsDryRun = true
		fmt.Println(args.ToStr())

	fp, err := Exec(context.TODO(), args)
	if err != nil {
		fmt.Println("Error in Exec: " + err.Error())
		t.Fail()
	}
	_ = fp
}

func TestFunctionalityBaseline(t *testing.T) {
	if IsTestFunctionality() {
		return
	}

}

func TestFunctionalityMigrate(t *testing.T) {
	if IsTestFunctionality() {
		return
	}
}

func TestFunctionalityRepair(t *testing.T) {
	if IsTestFunctionality() {
		return
	}
}

func TestFunctionalityValidate(t *testing.T) {
	if IsTestFunctionality() {
		return
	}
}

func TestUnitTcClean(t *testing.T) {
}

func TestUnitTcBaseline(t *testing.T) {
}

func TestUnitTcMigrate(t *testing.T) {
}

func TestUnitTcRepair(t *testing.T) {
}

func TestUnitTcValidate(t *testing.T) {
}

func GetArgsForFunctionalTesting(pluginDriverPath, pluginFlywayCommand, pluginLocations,
	pluginCommandLineArgs, pluginUrl, pluginUser, pluginPassword string) Args {

	defaultArgs := Args{
		FlywayEnvPluginArgs: FlywayEnvPluginArgs{
			DriverPath:      pluginDriverPath,
			FlywayCommand:   pluginFlywayCommand,
			Locations:       pluginLocations,
			CommandLineArgs: pluginCommandLineArgs,
			Url:             pluginUrl,
			UserName:        pluginUser,
			Password:        pluginPassword,
		},
	}

	return defaultArgs
}

func getDefaultPluginDriverPath() string {
	return os.Getenv("PLUGIN_DRIVER_PATH")
}

func getDefaultPluginFlywayCommand() string {
	return os.Getenv("PLUGIN_FLYWAY_COMMAND")
}

func getDefaultPluginLocations() string {
	return os.Getenv("PLUGIN_LOCATIONS")
}

func getDefaultPluginCommandLineArgs() string {
	return os.Getenv("PLUGIN_COMMAND_LINE_ARGS")
}

func getDefaultPluginUrl() string {
	return os.Getenv("PLUGIN_URL")
}

func getDefaultPluginUser() string {
	return os.Getenv("PLUGIN_USERNAME")
}

func getDefaultPluginPassword() string {
	return os.Getenv("PLUGIN_PASSWORD")
}
