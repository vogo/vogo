//author: liu.yang02@ucarinc.com

package vos

const (
	EnvValueSplit = ":"
)

// getUserEnvProfiles mac can't get user env from user file .bash_profile
func getUserEnvProfiles() []string {
	return []string{
		"/etc/bashrc",
		"/etc/profile",
	}
}
