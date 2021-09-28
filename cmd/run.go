package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func run(c *cli.Context) error {
	if c.Args().Len() != 1 {
		logger.Fatalf("mlsql run script-path")
	}

	_mlsqlHome, err := os.Executable()
	if err != nil {
		logger.Fatalf("mlsql lang is not execute.")
	}
	mlsqlHome := filepath.Dir(filepath.Dir(_mlsqlHome))

	var javaHome = os.Getenv("JAVA_HOME")

	if javaHome == "" {
		if _, err := os.Stat(path.Join(mlsqlHome, "jdk8")); !os.IsNotExist(err) {
			javaHome = path.Join(mlsqlHome, "jdk8")
		}
	}

	var mlsqlConfig = make(map[string]string)
	mlsqlConfigStr, err := os.ReadFile(".mlsql.config")
	if err == nil {
		for _, _s := range strings.Split(string(mlsqlConfigStr), "\n") {
			s := strings.TrimSpace(_s)
			if strings.HasPrefix(s, "#") {
				continue
			}
			if s == "" {
				continue
			}
			kv := strings.SplitN(s, "=", 2)
			mlsqlConfig[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
	}

	var xmx = ""

	if value, ok := mlsqlConfig["engine.memory"]; ok {
		xmx = "-Xmx" + value
	}

	var owner = "admin"

	if value, ok := mlsqlConfig["user.owner"]; ok {
		owner = value
	}

	var executable = "java"

	if javaHome != "" {
		executable = path.Join(javaHome, "bin", "java")
	}

	mainLib := path.Join(mlsqlHome, "main", "*")
	libsLib := path.Join(mlsqlHome, "libs", "*")
	pluginLib := path.Join(mlsqlHome, "plugin", "*")
	sparkLib := path.Join(mlsqlHome, "spark", "*")

	const mainClass = "streaming.core.StreamingApp"

	var command = []string{
		"-cp",
		fmt.Sprintf("%s:%s:%s:%s", mainLib, libsLib, pluginLib, sparkLib),
		mainClass,
		"-streaming.master", "local[*]",
		"-streaming.name", "MLSQL-desktop",
		"-streaming.rest", "false",
		"-streaming.thrift", "false",
		"-streaming.platform", "spark",
		"-streaming.spark.service", "false",
		"-streaming.job.cancel", "true",
		"-streaming.datalake.path", path.Join(".", "data"),
		"-streaming.driver.port", "9003",
		"-streaming.plugin.clzznames", "tech.mlsql.plugins.ds.MLSQLExcelApp,tech.mlsql.plugins.shell.app.MLSQLShell,tech.mlsql.plugins.assert.app.MLSQLAssert",
		"-streaming.platform_hooks", "tech.mlsql.runtime.SparkSubmitMLSQLScriptRuntimeLifecycle",
		"-streaming.mlsql.script.path", c.Args().First(),
		"-streaming.mlsql.script.owner", owner,
		"-streaming.mlsql.sctipt.jobName", "mlsql-cli",
	}

	if xmx != "" {
		command = append([]string{xmx}, command...)
	}

	logger.Infof("%v", command)

	r := exec.Command(executable, command...)
	r.Stdout = os.Stdout
	r.Stderr = os.Stderr

	err = r.Start()
	if err != nil {
		panic(err)
	}

	r.Wait()
	return nil
}

func runFlags() *cli.Command {
	cmd := &cli.Command{
		Name:      "run",
		Usage:     "run mlsql lang script",
		ArgsUsage: "script path",
		Action:    run,
		Flags:     []cli.Flag{},
	}
	return cmd
}
