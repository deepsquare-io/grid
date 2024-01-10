// Copyright (C) 2024 DeepSquare Association
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
	"strings"
	"time"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

const (
	mediaTypeDockerSchema2Manifest     = "application/vnd.docker.distribution.manifest.v2+json"
	mediaTypeDockerSchema2ManifestList = "application/vnd.docker.distribution.manifest.list.v2+json"
	mediaTypeDockerTarGzipLayer        = "application/vnd.docker.image.rootfs.diff.tar.gzip"
)

type ImageFetcher interface {
	FetchContainerImage(
		username string,
		password string,
		registry string,
		img string,
	) (image ocispec.Image, err error)
}

var DefaultImageFetcher ImageFetcher = &imageFetcher{}

func OverrideImageFetcher(fetcher ImageFetcher) {
	DefaultImageFetcher = fetcher
}

type imageFetcher struct{}

func (*imageFetcher) FetchContainerImage(
	username string,
	password string,
	registry string,
	img string,
) (image ocispec.Image, err error) {
	if strings.HasPrefix(img, "/") {
		return image, nil
	}
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

	manifestBaseURL := fmt.Sprintf("https://%s/v2/%s/manifests/", registry, img)
	blobsBaseURL := fmt.Sprintf("https://%s/v2/%s/blobs/", registry, img)

	// Authenticate
	token, err := firstRequestForAuth(ctx, manifestBaseURL+tag, username, password)
	if err != nil {
		return image, err
	}

	manifest, err := fetchImageManifest(ctx, manifestBaseURL, tag, token)
	if err != nil {
		return image, err
	}

	image, err = fetchImage(ctx, blobsBaseURL, string(manifest.Config.Digest), token)
	if err != nil {
		return image, err
	}

	return image, nil
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

func fetchImageManifest(
	ctx context.Context,
	manifestBaseURL string,
	tag string,
	token string,
) (manifest ocispec.Manifest, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", manifestBaseURL+tag, nil)
	if err != nil {
		return manifest, err
	}

	req.Header.Add(
		"Accept",
		strings.Join(
			[]string{
				mediaTypeDockerSchema2Manifest,
				mediaTypeDockerSchema2ManifestList,
				ocispec.MediaTypeImageManifest,
				ocispec.MediaTypeImageIndex,
			},
			", ",
		),
	)
	req.Header.Add("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return manifest, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 && resp.StatusCode < 200 {
		return manifest, errors.New("failure to fetch manifest")
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return manifest, err
	}

	var jsonResp map[string]interface{}
	if err = json.Unmarshal(raw, &jsonResp); err != nil {
		return manifest, err
	}
	if _, ok := jsonResp["errors"]; ok {
		return manifest, errors.New(string(raw))
	}

	switch jsonResp["mediaType"] {
	case mediaTypeDockerSchema2Manifest, ocispec.MediaTypeImageManifest:
		if err = json.Unmarshal(raw, &manifest); err != nil {
			return manifest, err
		}
		return manifest, nil
	case mediaTypeDockerSchema2ManifestList, ocispec.MediaTypeImageIndex:
		var index ocispec.Index
		if err = json.Unmarshal(raw, &index); err != nil {
			return manifest, err
		}

		// TODO: fetch arm64 too
		for _, desc := range index.Manifests {
			if desc.Platform != nil && desc.Platform.OS == "linux" &&
				desc.Platform.Architecture == "amd64" {
				manifest, err = fetchImageManifest(
					ctx,
					manifestBaseURL,
					string(desc.Digest),
					token,
				)
				if err != nil {
					continue
				}
				return manifest, nil
			}
		}
	default:
		return manifest, fmt.Errorf("unknown mediaType: %s", jsonResp["mediaType"])
	}
	return manifest, errors.New("no manifest")
}

func fetchImage(
	ctx context.Context,
	blobsBaseURL string,
	digest string,
	token string,
) (image ocispec.Image, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", blobsBaseURL+digest, nil)
	if err != nil {
		return image, err
	}

	req.Header.Add(
		"Accept",
		strings.Join(
			[]string{
				mediaTypeDockerSchema2Manifest,
				ocispec.MediaTypeImageManifest,
			},
			", ",
		),
	)
	req.Header.Add("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return image, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 && resp.StatusCode < 200 {
		return image, errors.New("failure to fetch manifest")
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return image, err
	}

	var jsonResp map[string]interface{}
	if err = json.Unmarshal(raw, &jsonResp); err != nil {
		return image, err
	}
	if _, ok := jsonResp["errors"]; ok {
		return image, errors.New(string(raw))
	}

	if err = json.Unmarshal(raw, &image); err != nil {
		return image, err
	}
	return image, nil
}
