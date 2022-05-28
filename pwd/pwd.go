package pwd

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

const HashFile string = "hash.json"

func HashPassword(pwd string) string {
	pw := []byte(pwd)
	res, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Couldn't encrypt password: %s", err.Error())
	}
	log.Info("Hashed password")
	return string(res)
}

func ValidatePassword(hash, pwd string) bool {
	pw, hs := []byte(pwd), []byte(hash)
	err := bcrypt.CompareHashAndPassword(hs, pw)
	v := err == nil
	log.Infof("Validation result: %t", v)
	return v
}

func LoadHash(fn string) (PasswordHash, error) {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatalf("Couldn't open hash file (%s): %s", fn, err.Error())
		return PasswordHash{}, err
	}
	defer f.Close()

	var ph PasswordHash
	if err := json.NewDecoder(f).Decode(&ph); err != nil {
		log.Fatalf("Couldn't decode password hash: %s", err.Error())
		return PasswordHash{}, err
	}

	log.Info("Loaded password hash")
	return ph, nil
}

func ValidatePasswordWithStored(pwd, fn string) bool {
	ph, err := LoadHash(fn)
	if err != nil {
		return false
	}

	return ValidatePassword(ph.Hash, pwd)
}
