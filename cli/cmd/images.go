// Copyright 2021 NetApp, Inc. All Rights Reserved.

package cmd

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	k8sclient "github.com/netapp/trident/cli/k8s_client"
	tridentconfig "github.com/netapp/trident/config"
	"github.com/netapp/trident/utils"
)

var K8sVersion string
var versionNotSupported = "The provided Kubernetes version is not supported: "
var versionNotSemantic = "The provided Kubernetes version was not given in semantic versioning: "

type ImageSet struct {
	Images []string `json:"images"`
	K8sVersion string `json:"k8sVersion"`
}

type ImageList struct {
	ImageSets []ImageSet `json:"imageSets"`
}

func init() {
	getImageCmd.Flags().StringVarP(&K8sVersion, "k8s-version", "v", "all",
		"Semantic version of Kubernetes cluster")
	RootCmd.AddCommand(getImageCmd)
}

var getImageCmd = &cobra.Command{
	Use:     "images",
	Short:   "Print a table of the container images Trident needs",
	Aliases: []string{},
	RunE: func(cmd *cobra.Command, args []string) error {
		if OperatingMode == ModeTunnel {
			command := []string{"images"}
			TunnelCommand(append(command, args...))
			return nil
		} else {
			return printImages()
		}
	},
}

func printImages() error {
	err := listImages()
	if err != nil {
		if err.Error() == versionNotSupported {
			log.Fatal(fmt.Sprintln(versionNotSupported, K8sVersion))
		} else {
			log.Fatal(fmt.Sprintln(versionNotSemantic, K8sVersion))
		}
	}

	return nil
}

func listImages() error {
	var err error
	var semVersion *utils.Version
	k8sVersions := make([]string, 0)
	imageMap := make(map[string][]string)
	var installYaml string
	var yamlErr error

	if K8sVersion == "all" {
		minMinorVersion := utils.MustParseSemantic(tridentconfig.KubernetesVersionMin).MinorVersion()
		maxMinorVersion := utils.MustParseSemantic(tridentconfig.KubernetesVersionMax).MinorVersion()
		for i := minMinorVersion; i <= maxMinorVersion; i++ {
			semVersion := fmt.Sprintf("v1.%d.0", i)
			k8sVersions = append(k8sVersions, semVersion)
			installYaml, yamlErr = getInstallYaml(utils.MustParseSemantic(semVersion))
			if yamlErr != nil {
				return yamlErr
			}
			imageMap[semVersion] = getImageNames(installYaml)
		}
	} else {
		semVersion, err = utils.ParseSemantic(K8sVersion)
		if err != nil {
			return err
		}
		k8sVersions = append(k8sVersions, K8sVersion)
		installYaml, yamlErr = getInstallYaml(semVersion)
		if yamlErr != nil {
			return yamlErr
		}
		imageMap[K8sVersion] = getImageNames(installYaml)
	}
	writeImages(k8sVersions, imageMap)
	return nil
}

func getInstallYaml(semVersion *utils.Version) (string, error) {
	var yaml string
	minorVersion := semVersion.ToMajorMinorVersion().MinorVersion()
	isSupportedVersion := minorVersion <= utils.MustParseSemantic(tridentconfig.KubernetesVersionMax).MinorVersion() &&
		minorVersion >= utils.MustParseSemantic(tridentconfig.KubernetesVersionMin).MinorVersion()

	if isSupportedVersion {
		if semVersion.LessThan(utils.MustParseSemantic(tridentconfig.KubernetesCSIVersionMinForced)) {
			csi = false
		} else {
			csi = true
		}
	} else {
		return "", errors.New(versionNotSupported)
	}

	// Get Deployment and Daemonset YAML and collect the names of the container images Trident needs to run.
	if csi {
		yaml = k8sclient.GetCSIDeploymentYAML(getDeploymentName(true),
			tridentconfig.BuildImage, tridentconfig.DefaultAutosupportImage, "", "",
			"", "", "", "", []string{}, nil,
			nil, false, false, true, semVersion, false)
		// trident image here is an empty string because we are already going to get it from the deployment yaml
		yaml += k8sclient.GetCSIDaemonSetYAML("", "", "", "",
			"", []string{}, nil, nil, false, false, semVersion)
	} else {
		yaml = k8sclient.GetDeploymentYAML("", tridentconfig.BuildImage, "", []string{}, nil,
			nil, false)
	}
	return yaml, nil
}

func getImageNames(yaml string) []string {
	var images []string
	lines := strings.Split(yaml, "\n")
	// don't get images that are empty strings
	imageRegex := regexp.MustCompile(`\s*image:\s*(.*)`)
	for i := 0; i < len(lines); i++ {
		imageLine := lines[i]
		image := strings.TrimSpace(imageRegex.ReplaceAllString(imageLine, "$1"))
		if matches := imageRegex.MatchString(imageLine); matches {
			// strip out "    image:" prefix
			if image != "" {
				images = append(images, image)
			}
		}
	}
	return images
}

func writeImages(k8sVersions []string, imageMap map[string][]string) {

	var imageSets []ImageSet
	for _, k8sVersion := range k8sVersions {
		var images []string
		for _, image := range imageMap[k8sVersion] {
			images = append(images, image)
		}
		imageSets = append(imageSets, ImageSet{Images: images, K8sVersion: k8sVersion})
	}
	switch OutputFormat {
	case FormatJSON:
		WriteJSON(ImageList{ImageSets: imageSets})
	case FormatYAML:
		WriteYAML(ImageList{ImageSets: imageSets})
	default:
		writeImageTable(k8sVersions, imageMap)
	}
}

func writeImageTable(k8sVersions []string, imageMap map[string][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Kubernetes Version", "Container Image"})
	if OutputFormat == FormatMarkdown {
		// print in markdown table format
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		table.SetCenterSeparator("|")
	}
	table.SetRowLine(true)
	for _, k8sVersion := range k8sVersions {
		table.Append([]string{k8sVersion, strings.Join(imageMap[k8sVersion], "\n")})
	}
	table.Render()
}
