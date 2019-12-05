//author: liu.yang02@ucarinc.com

package vos

import "fmt"

const (
	EnvValueSplit = ":"
)

func getUserEnvProfiles() []string {
	userName := GetCurrentUserName()

	files := []string{
		"/etc/bashrc",
		"/etc/profile",
	}

	if userName == "root" {
		return append(files, "/root/.bashrc", "/root/.bash_profile")
	}

	return append(files, fmt.Sprintf("/home/%s/.bashrc", userName),
		fmt.Sprintf("/home/%s/.bash_profile", userName))
}
