/*
 * Minio Lambda (C) 2016 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/hashicorp/go-version"
)

// check if minimum Go version is met.
func checkGoVersion() {
	runtimeVersion := runtime.Version()

	// Checking version is always successful with go tip
	if strings.HasPrefix(runtimeVersion, "devel") {
		return
	}

	// Parsing golang version
	curVersion, err := version.NewVersion(runtimeVersion[2:])
	if err != nil {
		fmt.Println("Unable to determine current go version.", err)
		os.Exit(1)
	}

	// Prepare version constraint.
	constraints, err := version.NewConstraint(minGoVersion)
	if err != nil {
		fmt.Println("Unable to check go version.", err)
		os.Exit(1)
	}

	// Check for minimum version.
	if !constraints.Check(curVersion) {
		fmt.Println(fmt.Sprintf("Please recompile Minio with Golang version %s.", minGoVersion))
		os.Exit(1)
	}
}
