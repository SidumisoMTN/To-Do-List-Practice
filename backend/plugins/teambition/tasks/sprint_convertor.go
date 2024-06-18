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

package tasks

import (
	"fmt"
	"github.com/apache/incubator-devlake/core/dal"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/models/domainlayer"
	"github.com/apache/incubator-devlake/core/models/domainlayer/ticket"
	"github.com/apache/incubator-devlake/core/plugin"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/teambition/models"
	"reflect"
)

var ConvertSprintsMeta = plugin.SubTaskMeta{
	Name:             "convertSprints",
	EntryPoint:       ConvertSprints,
	EnabledByDefault: true,
	Description:      "convert teambition projects",
	DomainTypes:      []string{plugin.DOMAIN_TYPE_TICKET},
}

func ConvertSprints(taskCtx plugin.SubTaskContext) errors.Error {
	rawDataSubTaskArgs, data := CreateRawDataSubTaskArgs(taskCtx, RAW_SPRINT_TABLE)
	db := taskCtx.GetDal()
	clauses := []dal.Clause{
		dal.From(&models.TeambitionSprint{}),
		dal.Where("connection_id = ? AND project_id = ?", data.Options.ConnectionId, data.Options.ProjectId),
	}

	cursor, err := db.Cursor(clauses...)
	if err != nil {
		return err
	}
	defer cursor.Close()
	converter, err := helper.NewDataConverter(helper.DataConverterArgs{
		RawDataSubTaskArgs: *rawDataSubTaskArgs,
		InputRowType:       reflect.TypeOf(models.TeambitionSprint{}),
		Input:              cursor,
		Convert: func(inputRow interface{}) ([]interface{}, errors.Error) {
			userTool := inputRow.(*models.TeambitionSprint)
			sprint := &ticket.Sprint{
				DomainEntity: domainlayer.DomainEntity{
					Id: getSprintIdGen().Generate(data.Options.ConnectionId, userTool.Id),
				},
				Name:            userTool.Name,
				Url:             fmt.Sprintf("https://www.teambition.com/project/%s/sprint/section/%s", userTool.ProjectId, userTool.Id),
				Status:          userTool.Status,
				StartedDate:     userTool.StartDate.ToNullableTime(),
				EndedDate:       userTool.DueDate.ToNullableTime(),
				CompletedDate:   userTool.Accomplished.ToNullableTime(),
				OriginalBoardID: getProjectIdGen().Generate(data.Options.ConnectionId, userTool.ProjectId),
			}

			return []interface{}{
				sprint,
				&ticket.BoardSprint{
					BoardId:  getProjectIdGen().Generate(data.Options.ConnectionId, userTool.ProjectId),
					SprintId: getSprintIdGen().Generate(data.Options.ConnectionId, userTool.Id),
				},
			}, nil
		},
	})

	if err != nil {
		return err
	}

	return converter.Execute()
}
