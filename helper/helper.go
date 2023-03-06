package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/Masterminds/semver"
)

const (
	PREFIX   string = ".arduino15/packages/rp2040/hardware/rp2040"
	FILENAME string = "lib/platform_inc.txt"

	ADDITIONAL_HEADER_0 string = "cores/rp2040"
	ADDITIONAL_HEADER_1 string = "include"
	ADDITIONAL_HEADER_2 string = "variants/rpipico"
)

func main() {
	homeDir := findHomeDir()
	version := findVersion(homeDir)
	headers := readHeaders(homeDir, version)
	fmt.Println(headers)
}

func findHomeDir() string {
	// Get the current user info
	user, _ := user.Current()
	return user.HomeDir
}

func findVersion(homeDir string) string {
	// Read the prefix path
	files, err := os.ReadDir(homeDir + "/" + PREFIX)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	// The vscode-arduino extension allows only one version at a time.
	var version string
	for _, file := range files {
		_, err := semver.NewVersion(file.Name())
		if err != nil {
			continue
		} else {
			version = file.Name()
		}
	}

	// Couldn't find a directory with semver format? exit.
	if version == "" {
		log.Println("cannot find a directory from" + homeDir + "/" + PREFIX)
		os.Exit(-1)
	}

	return version
}

func readHeaders(homeDir, version string) string {
	var ret string

	// Open the given file
	f, _ := os.Open(homeDir + "/" + PREFIX + "/" + version + "/" + FILENAME)
	// Iterate over the file contents
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, "/")
		// TODO: maybe 2 back slashes for Windows?
		line = strings.Join(lineSplit[1:], "/")
		line = "\"" + homeDir + "/" + PREFIX + "/" + version + "/" + line + "\",\n"
		ret += line
	}

	// Need to add few header files other than the headers stated in the platform_inc.txt.
	additionalHeaders := []string{
		ADDITIONAL_HEADER_0,
		ADDITIONAL_HEADER_1,
		ADDITIONAL_HEADER_2,
	}
	for _, h := range additionalHeaders {
		ret += "\"" + homeDir + "/" + PREFIX + "/" + version + "/" + h + "\",\n"
	}
	return ret
}
