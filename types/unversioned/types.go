// Copyright 2017 Google Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package unversioned

type EnvVar struct {
	Key   string
	Value string
}

type Config struct {
	Env          map[string]string
	Entrypoint   []string
	Cmd          []string
	Volumes      []string
	Workdir      string
	ExposedPorts []string
}

type FlattenedMetadata struct {
	// Config Config `json:"config"`
	Config map[string]interface{} `json:"config"`
}

// type FlattenedConfig struct {
// 	Env []string `json:"Env"`
// 	Cmd []string `json:"Cmd"`
// }

type Command []string
