package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

type credentials struct {
	RoleArn         string `json:"RoleArn"`
	AccessKeyID     string `json:"AccessKeyId"`
	SecretAccessKey string `json:"SecretAccessKey"`
	Token           string `json:"Token"`
	Expiration      string `json:"Expiration"`
}

const (
	AWS_URL = "http://169.254.170.2"
)

var s = flag.String("s", "default", "section name")

func main() {
	flag.Parse()
	section := *s
	credentialsRelativeURI := os.Getenv("AWS_CONTAINER_CREDENTIALS_RELATIVE_URI")
	url := AWS_URL + credentialsRelativeURI
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	defer resp.Body.Close()

	cred := &credentials{}
	err = json.NewDecoder(resp.Body).Decode(&cred)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Fprintf(os.Stdout, "[%s]\naws_access_key_id=%s\naws_secret_access_key=%s\naws_session_token=%s\n", section, cred.AccessKeyID, cred.SecretAccessKey, cred.Token)
}
