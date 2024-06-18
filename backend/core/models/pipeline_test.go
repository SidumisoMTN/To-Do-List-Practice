/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipelinePlan_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		plan PipelinePlan
		want bool
	}{
		{
			name: "empty",
			plan: PipelinePlan{},
			want: true,
		},
		{
			name: "empty",
			plan: []PipelineStage{{}, {}},
			want: true,
		},
		{
			name: "empty",
			plan: []PipelineStage{{}, {&PipelineTask{}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("length of plan is", len(tt.plan))
			assert.Equalf(t, tt.want, tt.plan.IsEmpty(), "IsEmpty()")
		})
	}
}
