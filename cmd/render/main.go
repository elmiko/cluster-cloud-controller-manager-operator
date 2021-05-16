package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/openshift/cluster-cloud-controller-manager-operator/pkg/render"
	"github.com/spf13/cobra"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"

	"k8s.io/klog/v2"
)

var (
	renderCmd = &cobra.Command{
		Use:   "run",
		Short: "Starts Cluster Cloud Controller Manager in render mode",
		Long:  "",
		RunE:  runRenderCmd,
	}

	renderOpts struct {
		destinationDir        string
		clusterInfrastructure string
		imagesFile            string
	}
)

func init() {
	renderCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	renderCmd.PersistentFlags().StringVar(&renderOpts.destinationDir, "dest-dir", "", "The destination dir where CCCMO writes the generated static pods for CCM.")
	renderCmd.PersistentFlags().StringVar(&renderOpts.clusterInfrastructure, "cluster-infrastructure-file", "", "Input path for the cluster infrastructure file.")
	renderCmd.PersistentFlags().StringVar(&renderOpts.imagesFile, "images-file", "", "Input path for the images config map file.")
	renderCmd.MarkFlagRequired("dest-dir")
	renderCmd.MarkFlagRequired("cluster-infrastructure-file")
	renderCmd.MarkFlagRequired("images-file")
}

func runRenderCmd(cmd *cobra.Command, args []string) error {
	flag.Set("logtostderr", "true")
	flag.Parse()

	if err := validate(
		renderOpts.destinationDir,
		renderOpts.clusterInfrastructure,
		renderOpts.imagesFile); err != nil {
		return err
	}

	if err := render.New(renderOpts.clusterInfrastructure, renderOpts.imagesFile).Run(renderOpts.destinationDir); err != nil {
		return err
	}

	return nil
}

// validate verifies all file and dirs exist
func validate(destinationDir, clusterInfrastructure, imagesFile string) error {
	errs := []error{}
	if err := isFile(clusterInfrastructure); err != nil {
		klog.Errorf("error reading --cluster-infrastructure-file=%q: %s", clusterInfrastructure, err)
		errs = append(errs, fmt.Errorf("error reading --cluster-infrastructure-file: %s", err))
	}
	if err := isFile(imagesFile); err != nil {
		klog.Errorf("error reading --images-file=%q: %s", imagesFile, err)
		errs = append(errs, fmt.Errorf("error reading --images-file: %s", err))
	}

	if len(errs) > 0 {
		return utilerrors.NewAggregate(errs)
	}

	return nil
}

func isFile(path string) error {
	st, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !st.Mode().IsRegular() {
		return fmt.Errorf("%q is not a regular file", path)
	}
	if st.Size() <= 0 {
		return fmt.Errorf("%q is empty", path)
	}

	return nil
}

func main() {
	if err := renderCmd.Execute(); err != nil {
		klog.Exitf("Error executing render: %v", err)
	}
}
