// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package oracle

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "oracle", asset.ModuleFieldsPri, AssetOracle); err != nil {
		panic(err)
	}
}

// AssetOracle returns asset data.
// This is the base64 encoded gzipped contents of module/oracle.
func AssetOracle() string {
	return "eJzEVdtu4jwQvucpPvWm/y9RHiAXK7FbkJBokdqq0l5VA5mAVcfO+tCWPv3KJtA0cWj3INYXEPkw32E8ngs88jaDNrSSPACccJIznC3ixNkAyNmujKic0CrDbho5OVqSZZQ69/Gc3WjjHlZaFWKdoSBpw6xhyWQ5w5IdDYBCsMxtNgCACygquQEdhttWnGFttK/qmRT8ATWMZsxmXEdLybaiFR+WUvE7GMlzbSH70QZvEgi/7xb28I+8fdYmb629I3F3IPE+TBIoZOOhEDKN1hbbwbrcJzOEsBCq0KaksNY6lRLb5CHakt5ISK3WicU+zV6JH54hclZOFIJNL2bCZHxgdAd4KiSHONAF3GZ3u5Hw8w3Vitd+1JThHcxb8cpNq0FL7V2ETyL3u9/kVdLLaLl1bJO7gB1chmN7jiaso+OKXkTpy0g62gKhjoTfEz0tyWh2ndxI9FMcC8N8Yje/9GwC7ja1vU0Z9ERChqoJVOAtm3h3R3E3rZwn2T1UCuUt3EZYPJEMVWbD0RxOwzpt6m2GJTnOUbKjGLS/GBw5n9L/S0WYln4+jfcqImQY349n8/HX+QTaYHZ9P57PLvHf/qNkUkEYvZURlC+XbIJEpV1Iu7c8jG7xC5WV5CFofyWSBKjRD3axn8kiN7qqOP//vNcUraRQ/PB3vJmTdXhU+lnVcWtHOg/WCAsVs337/XYxnQ7D/93kaojFdDqfXU+GWFyH/+DfzeTb4n5yMzreW9qNEJ/vK433vH7VLa35D/rL0YL8qBh/qw857UhG3FrCoeKGhzekvzBCVZ2Wb6zj+PkZflHdvzA0vElNgj8DAAD//3pQwhs="
}
