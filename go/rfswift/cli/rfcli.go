/* This code is part of RF Swift by @Penthertz
*  Author(s): Sébastien Dudek (@FlUxIuS)
 */

package cli

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
	rfdock "penthertz/rfswift/dock"
	rfutils "penthertz/rfswift/rfutils"
)

var DImage string
var ContID string
var ExecCmd string
var FilterLast string
var ExtraBind string
var XDisplay string
var SInstall string
var ImageRef string
var ImageTag string
var ExtraHost string

func forceXhost() {
	os := runtime.GOOS
	if os == "windows" {
		rfdock.DockerSetx11("/run/desktop/mnt/host/wslg/.X11-unix:/tmp/.X11-unix,/run/desktop/mnt/host/wslg:/mnt/wslg")
	} else {
		rfutils.XHostEnable() // force xhost to add local connections ALCs, TODO: to optimize later
	}
}

var rootCmd = &cobra.Command{
	Use:   "rfswift",
	Short: "rfswift - a simple CLI to transform and inspect strings",
	Long: `rfswift is a super fancy CLI (kidding)

One can use stringer to modify or inspect strings straight from the terminal`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use '-h' for help")
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "create and run a program",
	Long:  `Create a container and run a program inside the docker container`,
	Run: func(cmd *cobra.Command, args []string) {
		forceXhost()

		rfdock.DockerSetShell(ExecCmd)
		rfdock.DockerAddBiding(ExtraBind)
		rfdock.DockerSetImage(DImage)
		rfdock.DockerSetExtraHosts(ExtraHost)
		rfdock.DockerRun()
	},
}

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "exec a command",
	Long:  `Exec a program on a created docker container, even not started`,
	Run: func(cmd *cobra.Command, args []string) {
		forceXhost()

		rfdock.DockerSetShell(ExecCmd)
		rfdock.DockerExec(ContID, "/root")
	},
}

var lastCmd = &cobra.Command{
	Use:   "last",
	Short: "last container run",
	Long:  `Display the latest container that was run`,
	Run: func(cmd *cobra.Command, args []string) {
		rfdock.DockerLast(FilterLast)
	},
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install function script",
	Long:  `Install function script inside the container`,
	Run: func(cmd *cobra.Command, args []string) {
		rfdock.DockerSetShell(ExecCmd)
		rfdock.DockerInstallFromScript(ContID)
	},
}

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "commit a container",
	Long:  `Commit a container with change we have made`,
	Run: func(cmd *cobra.Command, args []string) {
		rfdock.DockerSetImage(DImage)
		rfdock.DockerCommit(ContID)
	},
}

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "pull a container",
	Long:  `Pull a container from internet`,
	Run: func(cmd *cobra.Command, args []string) {
		rfdock.DockerPull(ImageRef, ImageTag)
	},
}

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "rename an image",
	Long:  `Rename an image with another tag`,
	Run: func(cmd *cobra.Command, args []string) {
		rfdock.DockerRename(ImageRef, ImageTag)
	},
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove a container",
	Long:  `Remore an existing container`,
	Run: func(cmd *cobra.Command, args []string) {
		rfdock.DockerRemove(ContID)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(lastCmd)
	rootCmd.AddCommand(execCmd)
	rootCmd.AddCommand(commitCmd)
	rootCmd.AddCommand(pullCmd)
	rootCmd.AddCommand(renameCmd)
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().StringVarP(&ContID, "container", "c", "", "container to remove")

	installCmd.Flags().StringVarP(&ExecCmd, "install", "i", "", "function for installation")
	installCmd.Flags().StringVarP(&ContID, "container", "c", "", "container to run")

	pullCmd.Flags().StringVarP(&ImageRef, "image", "i", "", "image reference")
	pullCmd.Flags().StringVarP(&ImageTag, "tag", "t", "", "rename to target tag")
	pullCmd.MarkFlagRequired("image")
	pullCmd.MarkFlagRequired("tag")

	renameCmd.Flags().StringVarP(&ImageRef, "image", "i", "", "image reference")
	renameCmd.Flags().StringVarP(&ImageTag, "tag", "t", "", "rename to target tag")

	commitCmd.Flags().StringVarP(&ContID, "container", "c", "", "container to run")
	commitCmd.Flags().StringVarP(&DImage, "image", "i", "", "image (by default: 'myrfswift:latest')")
	commitCmd.MarkFlagRequired("container")
	commitCmd.MarkFlagRequired("image")

	execCmd.Flags().StringVarP(&ContID, "container", "c", "", "container to run")
	execCmd.Flags().StringVarP(&ExecCmd, "command", "e", "", "command to exec (required!)")
	execCmd.Flags().StringVarP(&SInstall, "install", "i", "", "install from function script (e.g: 'sdrpp_soft_install')")
	execCmd.MarkFlagRequired("command")

	runCmd.Flags().StringVarP(&ExtraHost, "extrahosts", "x", "", "set extra hosts (by default: 'pluto.local:192.168.1.2' separated by commas)")
	runCmd.Flags().StringVarP(&XDisplay, "display", "d", "", "set X Display (by default: 'DISPLAY=:0' separated by commas)")
	runCmd.Flags().StringVarP(&ExecCmd, "command", "e", "", "command to exec (by default: '/bin/bash')")
	runCmd.Flags().StringVarP(&ExtraBind, "bind", "b", "", "extra bindings (separed with commas)")
	runCmd.Flags().StringVarP(&DImage, "image", "i", "", "image (by default: 'myrfswift:latest')")

	lastCmd.Flags().StringVarP(&FilterLast, "filter", "f", "", "filter by image name")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
