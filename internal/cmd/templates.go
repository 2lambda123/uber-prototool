// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/uber/prototool/internal/exec"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	allCmdTemplate = &cmdTemplate{
		Use:   "all dirOrProtoFiles...",
		Short: "Compile, then format and overwrite, then re-compile and generate, then lint, stopping if any step fails.",
		Run: func(runner exec.Runner, args []string, flags *flags) error {
			return runner.All(args, flags.disableFormat, flags.disableLint, !flags.noRewrite)
		},
		BindFlags: func(flagSet *pflag.FlagSet, flags *flags) {
			flags.bindDirMode(flagSet)
			flags.bindDisableFormat(flagSet)
			flags.bindDisableLint(flagSet)
			flags.bindNoRewrite(flagSet)
		},
	}

	binaryToJSONCmdTemplate = &cmdTemplate{
		Use:   "binary-to-json dirOrProtoFiles... messagePath data",
		Short: "Convert the data from json to binary for the message path and data.",
		Args:  cobra.MinimumNArgs(3),
		Run: func(runner exec.Runner, args []string, flags *flags) error {
			return runner.BinaryToJSON(args)
		},
		BindFlags: func(flagSet *pflag.FlagSet, flags *flags) {
			flags.bindDirMode(flagSet)
		},
	}

	cleanCmdTemplate = &cmdTemplate{
		Use:   "clean",
		Short: "Delete the cache.",
		Args:  cobra.NoArgs,
		Run: func(runner exec.Runner, args []string, flags *flags) error {
			return runner.Clean()
		},
	}

	compileCmdTemplate = &cmdTemplate{
		Use:   "compile dirOrProtoFiles...",
		Short: "Compile with protoc to check for failures.",
		Run: func(runner exec.Runner, args []string, flags *flags) error {
			return runner.Compile(args, flags.dryRun)
		},
		BindFlags: func(flagSet *pflag.FlagSet, flags *flags) {
			flags.bindDirMode(flagSet)
			flags.bindDryRun(flagSet)
		},
	}

	createCmdTemplate = &cmdTemplate{
		Use:   "create files...",
		Short: "Create the given Protobuf files according to a template that passes default prototool lint.",
		Run: func(runner exec.Runner, args []string, flags *flags) error {
			return runner.Create(args, flags.pkg)
		},
		BindFlags: func(flagSet *pflag.FlagSet, flags *flags) {
			flags.bindPackage(flagSet)
		},
	}

	descriptorProtoCmdTemplate = &cmdTemplate{
		Use:   "descriptor-proto dirOrProtoFiles... messagePath",
		Short: "Get the descriptor proto for the message path.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(runner exec.Runner, args []string, flags *flags) error {
			return runner.DescriptorProto(args)
		},
		BindFlags: func(flagSet *pflag.FlagSet, flags *flags) {
			flags.bindDirMode(flagSet)
		},
	}

	downloadCmdTemplate = &cmdTemplate{
		Use:   "download",
		Short: "Download the protobuf artifacts to a cache.",
		Args:  cobra.NoArgs,
		Run: func(runner exec.Runner, args []string, flags *flags) error {
			return runner.Download()
		},
	}

	fieldDescriptorProtoCmdTemplate = &cmdTemplate{
		Use:   "field-descriptor-proto dirOrProtoFiles... fieldPath",
		Short: "Get the field descriptor proto for the field path.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(runner exec.Runner, args []string, flags *flags) error {
			return runner.FieldDescriptorProto(args)
		},
		BindFlags: func(flagSet *pflag.FlagSet, flags *flags) {
			flags.bindDirMode(flagSet)
		},
	}

	filesCmdTemplate = &cmdTemplate{
		Use:   "files dirOrProtoFiles...",
		Short: "Print all files that match the input arguments.",
		Run: func(runner exec.Runner, args []string, flags *flags) error {
			return runner.Files(args)
		},
	}

	formatCmdTemplate = &cmdTemplate{
		Use:   "format dirOrProtoFiles...",
		Short: "Format a proto file and compile with protoc to check for failures.",
		Run: func(runner exec.Runner, args []string, flags *flags) error {
			return runner.Format(args, flags.overwrite, flags.diffMode, flags.lintMode, !flags.noRewrite)
		},
		BindFlags: func(flagSet *pflag.FlagSet, flags *flags) {
			flags.bindDiffMode(flagSet)
			flags.bindLintMode(flagSet)
			flags.bindOverwrite(flagSet)
			flags.bindNoRewrite(flagSet)
		},
	}

	genCmdTemplate = &cmdTemplate{
		Use:   "gen dirOrProtoFiles...",
		Short: "Generate with protoc.",
		Run: func(runner exec.Runner, args []string, flags *flags) error {
			return runner.Gen(args, flags.dryRun)
		},
		BindFlags: func(flagSet *pflag.FlagSet, flags *flags) {
			flags.bindDirMode(flagSet)
			flags.bindDryRun(flagSet)
		},
	}

	grpcCmdTemplate = &cmdTemplate{
		Use:   "grpc dirOrProtoFiles...",
		Short: "Call a gRPC endpoint. Be sure to set required flags address, method, and either data or stdin.",
		Run: func(runner exec.Runner, args []string, flags *flags) error {
			return runner.GRPC(args, flags.headers, flags.address, flags.method, flags.data, flags.callTimeout, flags.connectTimeout, flags.keepaliveTime, flags.stdin)
		},
		BindFlags: func(flagSet *pflag.FlagSet, flags *flags) {
			flags.bindAddress(flagSet)
			flags.bindCallTimeout(flagSet)
			flags.bindConnectTimeout(flagSet)
			flags.bindData(flagSet)
			flags.bindDirMode(flagSet)
			flags.bindHeaders(flagSet)
			flags.bindKeepaliveTime(flagSet)
			flags.bindMethod(flagSet)
			flags.bindStdin(flagSet)
		},
	}

	initCmdTemplate = &cmdTemplate{
		Use:   "init [dirPath]",
		Short: "Generate an initial config file in the current or given directory.",
		Args:  cobra.MaximumNArgs(1),
		Run: func(runner exec.Runner, args []string, flags *flags) error {
			return runner.Init(args, flags.uncomment)
		},
		BindFlags: func(flagSet *pflag.FlagSet, flags *flags) {
			flags.bindUncomment(flagSet)
		},
	}

	jsonToBinaryCmdTemplate = &cmdTemplate{
		Use:   "json-to-binary dirOrProtoFiles... messagePath data",
		Short: "Convert the data from json to binary for the message path and data.",
		Args:  cobra.MinimumNArgs(2),
		Run: func(runner exec.Runner, args []string, flags *flags) error {
			return runner.JSONToBinary(args)
		},
		BindFlags: func(flagSet *pflag.FlagSet, flags *flags) {
			flags.bindDirMode(flagSet)
		},
	}

	lintCmdTemplate = &cmdTemplate{
		Use:   "lint dirOrProtoFiles...",
		Short: "Lint proto files and compile with protoc to check for failures.",
		Run: func(runner exec.Runner, args []string, flags *flags) error {
			return runner.Lint(args, flags.listAllLinters, flags.listLinters)
		},
		BindFlags: func(flagSet *pflag.FlagSet, flags *flags) {
			flags.bindDirMode(flagSet)
			flags.bindListAllLinters(flagSet)
			flags.bindListLinters(flagSet)
		},
	}

	listAllLintGroupsCmdTemplate = &cmdTemplate{
		Use:   "list-all-lint-groups",
		Short: "List all the available lint groups.",
		Args:  cobra.NoArgs,
		Run: func(runner exec.Runner, args []string, flags *flags) error {
			return runner.ListAllLintGroups()
		},
	}

	listLintGroupCmdTemplate = &cmdTemplate{
		Use:   "list-lint-group group",
		Short: "List the linters in the given lint group.",
		Args:  cobra.ExactArgs(1),
		Run: func(runner exec.Runner, args []string, flags *flags) error {
			return runner.ListLintGroup(args[0])
		},
	}

	serviceDescriptorProtoCmdTemplate = &cmdTemplate{
		Use:   "service-descriptor-proto dirOrProtoFiles... servicePath",
		Short: "Get the service descriptor proto for the service path.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(runner exec.Runner, args []string, flags *flags) error {
			return runner.ServiceDescriptorProto(args)
		},
		BindFlags: func(flagSet *pflag.FlagSet, flags *flags) {
			flags.bindDirMode(flagSet)
		},
	}

	versionCmdTemplate = &cmdTemplate{
		Use:   "version",
		Short: "Print the version.",
		Args:  cobra.NoArgs,
		Run: func(runner exec.Runner, args []string, flags *flags) error {
			return runner.Version()
		},
	}
)

// cmdTemplate contains the static parts of a cobra.Command such as
// documentation that we want to store outside of runtime creation.
//
// We do not just store cobra.Commands as in theory they have fields
// with types such as slices that if we were to return a blind copy,
// would mean that both the global cmdTemplate and the runtime
// cobra.Command would point to the same location. By making a new
// struct, we can also do more fancy templating things like prepending
// the Short description to the Long description for consistency, and
// have our own abstractions for the Run command.
type cmdTemplate struct {
	// Use is the one-line usage message.
	// This field is required.
	Use string
	// Short is the short description shown in the 'help' output.
	// This field is required.
	Short string
	// Long is the long message shown in the 'help <this-command>' output.
	// The Short field will be prepended to the Long field with a newline
	// when applied to a *cobra.Command.
	// This field is optional.
	Long string
	// Expected arguments.
	// This field is optional.
	Args cobra.PositionalArgs
	// Run is the command to run given an exec.Runner, args, and flags.
	// This field is required.
	Run func(exec.Runner, []string, *flags) error
	// BindFlags binds flags to the *pflag.FlagSet on Build.
	// There is no corollary to this on *cobra.Command.
	// This field is optional, although usually will be set.
	// We need to do this before run as the flags are populated
	// before Run is called.
	BindFlags func(*pflag.FlagSet, *flags)
}

// Build builds a *cobra.Command from the cmdTemplate.
func (c *cmdTemplate) Build(exitCodeAddr *int, stdin io.Reader, stdout io.Writer, stderr io.Writer, flags *flags) *cobra.Command {
	command := &cobra.Command{}
	command.Use = c.Use
	command.Short = c.Short
	if c.Long != "" {
		command.Long = fmt.Sprintf("%s\n%s", c.Short, c.Long)
	}
	command.Args = c.Args
	command.Run = func(_ *cobra.Command, args []string) {
		checkCmd(exitCodeAddr, stdin, stdout, stderr, args, flags, c.Run)
	}
	if c.BindFlags != nil {
		c.BindFlags(command.PersistentFlags(), flags)
	}
	return command
}

func checkCmd(exitCodeAddr *int, stdin io.Reader, stdout io.Writer, stderr io.Writer, args []string, flags *flags, f func(exec.Runner, []string, *flags) error) {
	runner, err := getRunner(stdin, stdout, stderr, flags)
	if err != nil {
		*exitCodeAddr = printAndGetErrorExitCode(err, stdout)
		return
	}
	if err := f(runner, args, flags); err != nil {
		*exitCodeAddr = printAndGetErrorExitCode(err, stdout)
	}
}

func getRunner(stdin io.Reader, stdout io.Writer, stderr io.Writer, flags *flags) (exec.Runner, error) {
	logger, err := getLogger(stderr, flags.debug)
	if err != nil {
		return nil, err
	}
	runnerOptions := []exec.RunnerOption{
		exec.RunnerWithLogger(logger),
	}
	if flags.cachePath != "" {
		runnerOptions = append(
			runnerOptions,
			exec.RunnerWithCachePath(flags.cachePath),
		)
	}
	if flags.dirMode {
		runnerOptions = append(
			runnerOptions,
			exec.RunnerWithDirMode(),
		)
	}
	if flags.harbormaster {
		runnerOptions = append(
			runnerOptions,
			exec.RunnerWithHarbormaster(),
		)
	}
	if flags.printFields != "" {
		runnerOptions = append(
			runnerOptions,
			exec.RunnerWithPrintFields(flags.printFields),
		)
	}
	if flags.protocURL != "" {
		runnerOptions = append(
			runnerOptions,
			exec.RunnerWithProtocURL(flags.protocURL),
		)
	}
	workDirPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return exec.NewRunner(workDirPath, stdin, stdout, runnerOptions...), nil
}

func getLogger(stderr io.Writer, debug bool) (*zap.Logger, error) {
	level := zapcore.InfoLevel
	if debug {
		level = zapcore.DebugLevel
	}
	return zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(
				zap.NewDevelopmentEncoderConfig(),
			),
			zapcore.Lock(zapcore.AddSync(stderr)),
			zap.NewAtomicLevelAt(level),
		),
	), nil
}
