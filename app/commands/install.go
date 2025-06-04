package commands

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/lem3s/fg/app/cmd"
	"github.com/mholt/archiver/v3"
)

type InstallCmd struct {
	Ctx *cmd.AppContext
}

func init() {
	cmd.Register("install", func(ctx *cmd.AppContext) cmd.Command {
		return &InstallCmd{Ctx: ctx}
	})
}

func (h *InstallCmd) Run(args []string) error {
	// TODO[Lemes] : alterar quando tiverem as flags
	installVersion := "1.2.3"

	versionsData, err := cmd.GetVersionsData()
	if err != nil {
		return err
	}

	versionInfo := getVersion(versionsData, installVersion)
	if versionInfo == nil {
		return fmt.Errorf("FG %s not available in repository", installVersion)
	}

	// Start of JDK install
	fileExtension := getFileExtension(h.Ctx.OS)

	err = downloadCompressedJDK(
		versionInfo.JDK.PlatformsCompressedSourceUrl[h.Ctx.OS],
		filepath.Join(h.Ctx.FgHome, "temp"),
		"jdk."+fileExtension,
	)
	if err != nil {
		return err
	}

	if !folderExists(filepath.Join(h.Ctx.FgHome, installVersion, "jdk")) {
		err = extractCompressedJDK(
			filepath.Join(h.Ctx.FgHome, "temp", "jdk."+fileExtension),
			filepath.Join(h.Ctx.FgHome, installVersion, "jdk"),
		)
		if err != nil {
			return err
		}
	}

	// TODO [Lemes] : Adicionar arquivo de configuração do FG
	// e colocar path do executável no arquivo de configuração
	// para o run não precisar buscar o executável
	_, err = searchForJavaExecutable(
		filepath.Join(h.Ctx.FgHome, installVersion, "jdk"),
		h.Ctx.OS,
	)
	if err != nil {
		return err
	}

	// End of JDK install

	// Start of apps install

	for _, app := range versionInfo.Apps {
		err = createFolder(filepath.Join(h.Ctx.FgHome, installVersion, app.Name))
		if err != nil {
			return err
		}

		for _, dependecyUrl := range app.DependenciesUrl {
			err = downloadDependecy(dependecyUrl, filepath.Join(h.Ctx.FgHome, installVersion, app.Name))
			if err != nil {
				return err
			}
		}

		for _, fileUrl := range app.FilesUrl {
			err = downloadDependecy(fileUrl, filepath.Join(h.Ctx.FgHome, installVersion, app.Name))
			if err != nil {
				return err
			}
		}
	}

	// End of apps install

	return nil
}

func getVersion(versionsData []cmd.Version, installVersion string) *cmd.Version {
	for _, version := range versionsData {
		if version.Version == installVersion {
			return &version
		}
	}

	return nil
}

func getFileExtension(os string) string {
	if os == "windows" {
		return "zip"
	} else {
		return "tar.gz"
	}
}

func downloadCompressedJDK(url, destinationDirectory, fileNameDestination string) error {
	err := os.MkdirAll(destinationDirectory, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath.Join(destinationDirectory, fileNameDestination))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func extractCompressedJDK(sourceDirectory, destinationDirectory string) error {
	return archiver.Unarchive(sourceDirectory, destinationDirectory)
}

func searchForJavaExecutable(rootDir, system string) ([]string, error) {
	var targetFile string
	if system == "windows" {
		targetFile = "java.exe"
	} else {
		targetFile = "java"
	}

	var matches []string

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Warning: cannot access %s: %v\n", path, err)
			return filepath.SkipDir
		}

		if !info.IsDir() && strings.EqualFold(info.Name(), targetFile) {
			matches = append(matches, path)
		}

		return nil
	})

	return matches, err
}

func folderExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func createFolder(path string) error {
	return os.MkdirAll(path, 0755)
}

func getFileName(url string) string {
	return filepath.Base(url)
}

func downloadDependecy(url, destinationDirectory string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath.Join(destinationDirectory, getFileName(url)))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
