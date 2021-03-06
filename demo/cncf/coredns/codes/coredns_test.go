/*
Copyright 2016 - 2017 Huawei Technologies Co., Ltd. All rights reserved.

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

package main

import (
	"os"
	"testing"
)

//Testing git clone with a small repository.
func Test_git_clone(t *testing.T) {
	type repository struct {
		r    string
		dest string
	}

	tests := []struct {
		name    string
		r       repository
		wantErr bool
	}{
		{
			name: "Dockyard",
			r: repository{
				r:    "git@github.com:Huawei/dockyard.git",
				dest: "/tmp/tests",
			},
			wantErr: false,
		},
	}

	for _, v := range tests {
		os.MkdirAll(v.r.dest, 0777)
		err := gitClone(v.r.r, v.r.dest)

		if (err != nil) != v.wantErr {
			t.Errorf("GitClone() error = %v, and want error %v", err, v.wantErr)
			return
		}

		os.RemoveAll(v.r.dest)

		t.Log("GitClone() function test OK")
	}
}
