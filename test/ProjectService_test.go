//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package test

import (
	"testing"

	"github.com/apache/cloudstack-go/v2/cloudstack"
)

func TestProjectService(t *testing.T) {
	service := "ProjectService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testactivateProject := func(t *testing.T) {
		if _, ok := response["activateProject"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Project.NewActivateProjectParams("id")
		_, err := client.Project.ActivateProject(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ActivateProject", testactivateProject)

	testcreateProject := func(t *testing.T) {
		if _, ok := response["createProject"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Project.NewCreateProjectParams("displaytext", "name")
		_, err := client.Project.CreateProject(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CreateProject", testcreateProject)

	testdeleteProject := func(t *testing.T) {
		if _, ok := response["deleteProject"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Project.NewDeleteProjectParams("id")
		_, err := client.Project.DeleteProject(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteProject", testdeleteProject)

	testdeleteProjectInvitation := func(t *testing.T) {
		if _, ok := response["deleteProjectInvitation"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Project.NewDeleteProjectInvitationParams("id")
		_, err := client.Project.DeleteProjectInvitation(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteProjectInvitation", testdeleteProjectInvitation)

	testlistProjectInvitations := func(t *testing.T) {
		if _, ok := response["listProjectInvitations"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Project.NewListProjectInvitationsParams()
		_, err := client.Project.ListProjectInvitations(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListProjectInvitations", testlistProjectInvitations)

	testlistProjects := func(t *testing.T) {
		if _, ok := response["listProjects"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Project.NewListProjectsParams()
		_, err := client.Project.ListProjects(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListProjects", testlistProjects)

	testsuspendProject := func(t *testing.T) {
		if _, ok := response["suspendProject"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Project.NewSuspendProjectParams("id")
		_, err := client.Project.SuspendProject(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("SuspendProject", testsuspendProject)

	testupdateProject := func(t *testing.T) {
		if _, ok := response["updateProject"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Project.NewUpdateProjectParams("id")
		_, err := client.Project.UpdateProject(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UpdateProject", testupdateProject)

	testupdateProjectInvitation := func(t *testing.T) {
		if _, ok := response["updateProjectInvitation"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Project.NewUpdateProjectInvitationParams("projectid")
		_, err := client.Project.UpdateProjectInvitation(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UpdateProjectInvitation", testupdateProjectInvitation)

	testlistProjectRolePermissions := func(t *testing.T) {
		if _, ok := response["listProjectRolePermissions"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Project.NewListProjectRolePermissionsParams("projectid")
		_, err := client.Project.ListProjectRolePermissions(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListProjectRolePermissions", testlistProjectRolePermissions)

	testcreateProjectRolePermission := func(t *testing.T) {
		if _, ok := response["createProjectRolePermission"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Project.NewCreateProjectRolePermissionParams("permission", "projectid", "projectroleid", "rule")
		_, err := client.Project.CreateProjectRolePermission(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CreateProjectRolePermission", testcreateProjectRolePermission)

	testupdateProjectRolePermission := func(t *testing.T) {
		if _, ok := response["updateProjectRolePermission"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Project.NewUpdateProjectRolePermissionParams("projectid", "projectroleid")
		_, err := client.Project.UpdateProjectRolePermission(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UpdateProjectRolePermission", testupdateProjectRolePermission)

	testdeleteProjectRolePermission := func(t *testing.T) {
		if _, ok := response["deleteProjectRolePermission"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Project.NewDeleteProjectRolePermissionParams("id", "projectid")
		_, err := client.Project.DeleteProjectRolePermission(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteProjectRolePermission", testdeleteProjectRolePermission)

}