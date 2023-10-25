package db_local

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/google/uuid"
	filehelpers "github.com/turbot/go-kit/files"
	"github.com/turbot/pipe-fittings/filepaths"
	"github.com/turbot/pipe-fittings/utils"
)

// Passwords :: structure for working with DB passwords
type Passwords struct {
	Root      string
	Steampipe string
}

func writePasswordFile(password string) error {
	return os.WriteFile(filepaths_steampipe.GetPasswordFileLocation(), []byte(password), 0600)
}

// readPasswordFile reads the password file and returns it contents.
// the the password file could not be found, then it generates a new
// password and writes it to the password file, before returning it
func readPasswordFile() (string, error) {
	if !filehelpers.FileExists(filepaths_steampipe.GetPasswordFileLocation()) {
		p := generatePassword()
		if err := writePasswordFile(p); err != nil {
			return "", err
		}
		return p, nil
	}
	contentBytes, err := os.ReadFile(filepaths_steampipe.GetPasswordFileLocation())
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(contentBytes)), nil
}

func generatePassword() string {
	// Create a simple, random password of the form f9fe-442f-90fb
	// Simple to read / write, and has a strength rating of 4 per https://lowe.github.io/tryzxcvbn/
	// Yes, this UUIDv4 does always include a 4, but good enough for our needs.
	u, err := uuid.NewRandom()
	if err != nil {
		// Should never happen?
		panic(err)
	}
	s := u.String()
	return strings.ReplaceAll(s[9:23], "-", "_")
}

func migrateLegacyPasswordFile() error {
	utils.LogTime("db_local.migrateLegacyPasswordFile start")
	defer utils.LogTime("db_local.migrateLegacyPasswordFile end")
	if filehelpers.FileExists(filepaths_steampipe.GetLegacyPasswordFileLocation()) {
		p, err := getLegacyPasswords()
		if err != nil {
			return err
		}
		os.Remove(filepaths_steampipe.GetLegacyPasswordFileLocation())
		return writePasswordFile(p.Steampipe)
	}
	return nil
}

func getLegacyPasswords() (*Passwords, error) {
	contentBytes, err := os.ReadFile(filepaths_steampipe.GetLegacyPasswordFileLocation())
	if err != nil {
		return nil, err
	}
	var passwords = new(Passwords)
	err = json.Unmarshal(contentBytes, passwords)
	if err != nil {
		return nil, err
	}
	return passwords, nil
}
