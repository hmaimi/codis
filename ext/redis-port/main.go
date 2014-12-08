// Copyright 2014 Wandoujia Inc. All Rights Reserved.
// Licensed under the MIT (MIT-LICENSE.txt) license.

package main

import (
	"github.com/hmaimi/codis/ext/redis-port/args"
	"github.com/hmaimi/codis/ext/redis-port/cmd"
)

func main() {
	switch args.Code() {
	case "decode":
		cmd.Decode(args.NCPU(), args.Input(), args.Output())
	case "restore":
		cmd.Restore(args.NCPU(), args.Input(), args.Target())
	case "dump":
		cmd.Dump(args.NCPU(), args.From(), args.Output())
	case "sync":
		cmd.Sync(args.NCPU(), args.From(), args.Target())
	}
}
