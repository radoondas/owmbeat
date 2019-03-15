// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package include

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("owmbeat", "fields.yml", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
