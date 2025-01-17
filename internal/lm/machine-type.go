/**
# Copyright (c) 2022, NVIDIA CORPORATION.  All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
**/

package lm

import (
	"fmt"
	"os"
	"strings"
)

func newMachineTypeLabeler(machineTypePath string) (Labeler, error) {
	data, err := os.ReadFile(machineTypePath)
	if err != nil {
		return nil, fmt.Errorf("error getting machine type: %v", err)
	}
	machineType := strings.TrimSpace(string(data))

	l := Labels{
		"nvidia.com/gpu.machine": strings.Replace(machineType, " ", "-", -1),
	}
	return l, nil
}
