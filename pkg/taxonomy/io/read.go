// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package taxonomyio

import (
	"io/ioutil"
	"path/filepath"

	"sigs.k8s.io/yaml"
	"xyz.fybrik.io/vanity/pkg/taxonomy/model"
)

// ReadDocumentFromFile loads a document model from a JSON or YAML file
func ReadDocumentFromFile(path string) (*model.Document, error) {
	data, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	return unmarshalToDocument(data)
}

func unmarshalToDocument(content []byte) (*model.Document, error) {
	result := &model.Document{}
	err := yaml.Unmarshal(content, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
