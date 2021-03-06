/*
Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"reflect"

	"github.com/golang/glog"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/grafeas/kritis/pkg/kritis/constants"
)

// RemoveGloballyWhitelistedImages returns all images that aren't globally whitelisted
func RemoveGloballyWhitelistedImages(images []string) []string {
	notWhitelisted := []string{}
	for _, image := range images {
		whitelisted, err := imageInWhitelist(image)
		if err != nil {
			glog.Errorf("couldn't check if %s is in global whitelist: %v", image, err)
		}
		if !whitelisted {
			notWhitelisted = append(notWhitelisted, image)
		}
	}
	return notWhitelisted
}

func imageInWhitelist(image string) (bool, error) {
	for _, w := range constants.GlobalImageWhitelist {
		whitelistRef, err := name.ParseReference(w, name.WeakValidation)
		if err != nil {
			return false, err
		}
		imageRef, err := name.ParseReference(image, name.WeakValidation)
		if err != nil {
			return false, err
		}
		// Make sure images have the same context
		if reflect.DeepEqual(whitelistRef.Context(), imageRef.Context()) {
			return true, nil
		}
	}
	return false, nil
}
