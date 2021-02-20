package main

import (
	"flag"
	"fmt"
	"os"
)

const usage = `
iglist

List FQDNs of the instance group provided. The tool waits until all members
will have not empty FQDNs.

Arguments:
  -help bool
        print this message and exist`

type argT struct {
	Folder *string
	Group  *string
}

var args = argT{
	Folder: flag.String(
		"folder",
		"",
		"intance group parent folder id, required",
	),
	Group: flag.String(
		"group",
		"",
		"instance group name, required",
	),
}

func init() {
	flag.CommandLine.SetOutput(os.Stdout)
	flag.Usage = func() {
		fmt.Println(usage)
		flag.PrintDefaults()
	}
}
