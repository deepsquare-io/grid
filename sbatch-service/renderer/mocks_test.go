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

package renderer_test

import (
	"fmt"

	"github.com/deepsquare-io/grid/sbatch-service/utils"
	"github.com/deepsquare-io/grid/sbatch-service/validate"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

type mockImageFetcher struct{}

func (*mockImageFetcher) FetchContainerImage(
	_ string,
	_ string,
	_ string,
	_ string,
) (image ocispec.Image, err error) {
	return ocispec.Image{
		Config: ocispec.ImageConfig{
			WorkingDir: "/",
		},
	}, nil
}

func init() {
	fmt.Println("mocked ImageFetcher")
	validate.OverrideImageFetcher(&mockImageFetcher{})
	utils.MockRandomString("[random_string]")
}
