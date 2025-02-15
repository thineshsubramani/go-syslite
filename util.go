package syslite

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RunCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func ReadFileContent(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	return string(content), err
}

type OSRelease map[string]string

func ReadOSRelease() OSRelease {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return nil
	}
	defer file.Close()

	osRelease := make(OSRelease)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := strings.Trim(parts[1], `"'`)
			osRelease[key] = value
		}
	}

	return osRelease
}

func IsDistro(targetIDs ...string) bool {
	osRelease := ReadOSRelease()
	if osRelease == nil {
		return false
	}

	id, exists := osRelease["ID"]
	if !exists {
		return false
	}

	for _, targetID := range targetIDs {
		if strings.ToUpper(id) == strings.ToUpper(targetID) {
			return true
		}
	}

	return false
}

func printHelp() {
	fmt.Println("Available checks:")
	fmt.Println("  isLinux")
	fmt.Println("  isWindows")
	fmt.Println("  isDarwin")
	fmt.Println("  isARM64")
	fmt.Println("  isAMD64")
	fmt.Println("  Isx86_64")
	fmt.Println("  is386")
	fmt.Println("  isARM")
	fmt.Println("  isPPC64")
	fmt.Println("  isMIPS")
	fmt.Println("  isDebian")
	fmt.Println("  isUbuntu")
	fmt.Println("  isFedora")
	fmt.Println("  isCentOS")
	fmt.Println("  isRHEL")
	fmt.Println("  isArch")
	fmt.Println("  isAlpine")
	fmt.Println("  isOpenSUSE")
}
