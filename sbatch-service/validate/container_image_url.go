// Copyright (C) 2023 DeepSquare Association
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package validate

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	regexImage = `[[:lower:][:digit:]/._-]+`
	regexTag   = `[[:alnum:]._:-]+`
)

var (
	// regexContainerURL is a matcher from https://github.com/NVIDIA/enroot/blob/master/src/docker.sh
	regexContainerURL = regexp.MustCompilePOSIX(
		fmt.Sprintf(
			"^(%s)(:(%s))?$",
			regexImage,
			regexTag,
		),
	)
)

func CheckContainerImage(
	username string,
	password string,
	registry string,
	img string,
) (err error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if registry == "" {
		registry = "registry-1.docker.io"
	}

	img, tag, _ := strings.Cut(img, ":")
	if tag == "" {
		tag = "latest"
	}

	// Check library prefix
	if strings.Contains(registry, "docker.io") {
		_, _, ok := strings.Cut(img, "/")
		if !ok {
			img = "library/" + img
		}
	}

	url := fmt.Sprintf("https://%s/v2/%s/manifests/%s", registry, img, tag)

	// Authenticate
	token, err := firstRequestForAuth(ctx, url, username, password)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add(
		"Accept",
		"application/vnd.docker.distribution.manifest.v2+json, application/vnd.oci.image.manifest.v1+json",
	)
	req.Header.Add("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 && resp.StatusCode < 200 {
		return errors.New("failure to fetch manifest")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var jsonResp map[string]interface{}
	if err = json.Unmarshal(body, &jsonResp); err != nil {
		return err
	}
	if _, ok := jsonResp["errors"]; ok {
		return errors.New(string(body))
	}

	return nil
}

func firstRequestForAuth(
	ctx context.Context,
	url, user, pass string,
) (token string, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 300 && resp.StatusCode >= 200 {
		return
	}
	if resp.StatusCode != 401 {
		return "", err
	}
	wwwAuthenticate := resp.Header.Get("Www-Authenticate")
	if wwwAuthenticate == "" {
		return
	}

	// Parse token request
	prefix, value, _ := strings.Cut(wwwAuthenticate, " ")
	switch prefix {
	case "Bearer":
		// Request a new token.
		realm, params, err := parseBearerToken(value)
		if err != nil {
			return "", err
		}
		return requestAuthToken(ctx, realm, params, user, pass)
	case "Basic":
		// Check that we have valid credentials and save them if successful.
		return fmt.Sprintf(
			"Basic %s",
			base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user, pass))),
		), nil
	}

	return
}

func parseBearerToken(
	token string,
) (realm string, options map[string]string, err error) {
	options = make(map[string]string)
	keyValue := strings.Split(token, ",")
	for _, kv := range keyValue {
		splitted := strings.Split(kv, "=")
		if len(splitted) != 2 {
			err = fmt.Errorf("wrong formatting of the token")
			return
		}
		splitted[1] = strings.Trim(splitted[1], `"`)
		switch splitted[0] {
		case "realm":
			realm = splitted[1]
		default:
			options[splitted[0]] = splitted[1]
		}
	}
	return
}

func requestAuthToken(
	ctx context.Context,
	realm string,
	options map[string]string,
	user, pass string,
) (authToken string, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", realm, nil)
	if err != nil {
		return
	}

	query := req.URL.Query()
	for k, v := range options {
		query.Add(k, v)
	}
	if user != "" && pass != "" {
		query.Add("offline_token", "true")
		req.SetBasicAuth(user, pass)
	}
	req.URL.RawQuery = query.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		err = fmt.Errorf("error in getting the token, http request failed %s", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		err = fmt.Errorf("authorization error %s", resp.Status)
		return
	}

	var jsonResp map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&jsonResp); err != nil {
		return
	}
	authTokenInterface, ok := jsonResp["token"]
	if !ok {
		authTokenInterface, ok = jsonResp["access_token"]
	}
	if ok {
		authToken = "Bearer " + authTokenInterface.(string)
	} else {
		err = fmt.Errorf("didn't get the token key from the server")
		return
	}
	return
}

func ContainerURLValidator(url string) bool {
	return strings.HasPrefix(url, "/") || regexContainerURL.MatchString(url)
}
