// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"encoding/json"
	"log"
)

type Args struct {
	Pipeline
	FlywayEnvPluginArgs
	Level string `envconfig:"PLUGIN_LOG_LEVEL"`
}

type FlywayEnvPluginArgs struct {
	PluginDriverPath      string `envconfig:"PLUGIN_DRIVER_PATH"`
	PluginFlywayCommand   string `envconfig:"PLUGIN_FLYWAY_COMMAND"`
	PluginLocations       string `envconfig:"PLUGIN_LOCATIONS"`
	PluginCommandLineArgs string `envconfig:"PLUGIN_COMMAND_LINE_ARGS"`
	PluginUrl             string `envconfig:"PLUGIN_URL"`
	PluginUser            string `envconfig:"PLUGIN_USERNAME"`
	PluginPassword        string `envconfig:"PLUGIN_PASSWORD"`
}

type FlywayPlugin struct {
	InputArgs         *Args
	IsMultiFileUpload bool
	ProcessingInfo
}

func (a *Args) ToStr() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		log.Printf("Error marshalling Args to JSON: %v", err)
		return ""
	}
	return string(jsonData)
}

type ProcessingInfo struct {
	IsDryRun    bool
	ExecCommand string
}

func GetNewPlugin(ctx context.Context, args Args) (FlywayPlugin, error) {
	return FlywayPlugin{}, nil
}

func Exec(ctx context.Context, args Args) (FlywayPlugin, error) {
	plugin, err := GetNewPlugin(ctx, args)
	if err != nil {
		return plugin, err
	}

	err = plugin.Init(&args)
	if err != nil {
		return plugin, err
	}
	defer func(p FlywayPlugin) {
		err := p.DeInit()
		if err != nil {
			LogPrintln("Error in DeInit: " + err.Error())
		}
	}(plugin)

	err = plugin.ValidateAndProcessArgs(args)
	if err != nil {
		return plugin, err
	}

	err = plugin.DoPostArgsValidationSetup(args)
	if err != nil {
		return plugin, err
	}

	err = plugin.Run()
	if err != nil {
		return plugin, err
	}

	return plugin, nil
}
