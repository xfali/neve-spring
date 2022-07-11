// Copyright (C) 2019-2022, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package customargs

import (
	"fmt"
	"github.com/spf13/pflag"
	"k8s.io/gengo/args"
)

type NeveArgs struct {
	Annotation string
}

func NewDefault() (*args.GeneratorArgs, *NeveArgs) {
	args := args.Default().WithoutDefaultFlagParsing()
	cusArgs := &NeveArgs{
		Annotation: "+neve:",
	}
	args.CustomArgs = cusArgs
	args.OutputFileBaseName = "zz_generated"
	return args, cusArgs
}

func (arg *NeveArgs) String() string {
	return arg.Annotation
}

func (arg *NeveArgs) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&arg.Annotation, "annotation", arg.Annotation, "Annotation name")
}

func Validate(args *args.GeneratorArgs) error {
	_ = args.CustomArgs.(*NeveArgs)

	if len(args.OutputFileBaseName) == 0 {
		return fmt.Errorf("Output file base name cannot be empty. ")
	}
	if len(args.InputDirs) == 0 {
		return fmt.Errorf("Input directory cannot be empty. ")
	}
	return nil
}
