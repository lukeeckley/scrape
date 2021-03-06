package main

import (
	"regexp"
	"strings"
)

var reCreds = regexp.MustCompile("(?m)^([a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+):([^ ~/$].*$)")
var reEmail = regexp.MustCompile("[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]+")
var rePrivKey = regexp.MustCompile("(?s)BEGIN (RSA|DSA|) PRIVATE KEY.*END (RSA|DSA|) PRIVATE KEY")
var reAwsKey = regexp.MustCompile("(?is).*(AKIA[A-Z0-9]{16}).*([A-Za-z0-9+/]{40})")

// Find AWS access keys and secrets
func processAWSKeys(contents, key string) bool {
	awsKeys := reAwsKey.FindAllStringSubmatch(contents, -1)

	// No keys found.
	if awsKeys == nil {
		return false
	}

	for _, awsKey := range awsKeys {
		conf.ds.Put("awskeys", strings.Join(awsKey[1:], ":"), key)
	}

	return true
}

// Look for email addresses and save them to a file.
func processEmails(contents, key string) bool {
	emails := reEmail.FindAllString(contents, -1)

	// No emails found.
	if emails == nil {
		return false
	}

	for _, email := range emails {
		conf.ds.Put("emails", strings.ToLower(email), key)
	}

	return true
}

// Look for credentials in the format of email:password and save them to a file.
func processCredentials(contents, key string) bool {
	creds := reCreds.FindAllString(contents, -1)

	// No creds found.
	if creds == nil {
		return false
	}

	for _, cred := range creds {
		conf.ds.Put("creds", cred, key)
	}

	return true
}

// Look for private keys.
func processPrivKey(contents, key string) bool {
	privKeys := rePrivKey.FindAllString(contents, -1)

	// No keys found.
	if privKeys == nil {
		return false
	}

	for _, privKey := range privKeys {
		conf.ds.Put("privkeys", privKey, key)
	}

	return true
}
