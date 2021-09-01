package gerrit

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type Project struct {
	Name              string `json:"-"`
	Parent            string `json:"parent,omitempty"`
	Description       string `json:"description,omitempty"`
	PermissionsOnly   bool   `json:"permissions_only"`
	CreateEmptyCommit bool   `json:"create_empty_commit"`
	SubmitType        string `json:"submit_type,omitempty"`
	Branches          string `json:"branches,omitempty"`
	Owners            string `json:"owners,omitempty"`
	RejectEmptyCommit string `json:"reject_empty_commit,omitempty"`
}

type Branch struct {
	Ref       string    `json:"ref"`
	Revision  string    `json:"revision"`
	CanDelete bool      `json:"can_delete"`
	WebLinks  []WebLink `json:"web_link"`
}

type WebLink struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
	Target   string `json:"target"`
}

func (gc Client) CreateProject(prj *Project) error {
	rsp, err := gc.resty.R().SetBody(prj).SetHeader("Content-Type", "application/json").
		Put(fmt.Sprintf("/projects/%s", prj.Name))

	return parseRestyResponse(rsp, err)
}

func (gc Client) GetProject(name string) (*Project, error) {
	rsp, err := gc.resty.R().
		SetHeader("accept", "application/json").
		Get(fmt.Sprintf("/projects/%s", name))

	if err != nil {
		return nil, errors.Wrapf(err, "Unable to get Gerrit project")
	}

	if rsp.StatusCode() != http.StatusOK {
		if rsp.StatusCode() == 404 {
			return nil, ErrDoesNotExist("does not exists")
		}

		return nil, errors.Errorf("wrong response code: %d, body: %s", rsp.StatusCode(), rsp.String())
	}

	var prj Project
	body := rsp.String()[5:]
	if err := json.Unmarshal([]byte(body), &prj); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal project response")
	}

	return &prj, nil
}

func (gc Client) UpdateProject(prj *Project) error {
	rsp, err := gc.resty.R().SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{
			"description":    prj.Description,
			"commit_message": "Update the project description",
		}).Put(fmt.Sprintf("/projects/%s/description", prj.Name))

	if err := parseRestyResponse(rsp, err); err != nil {
		return errors.Wrap(err, "unable to update project description")
	}

	rsp, err = gc.resty.R().SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{
			"parent":         prj.Parent,
			"commit_message": "Update the project parent",
		}).Put(fmt.Sprintf("/projects/%s/parent", prj.Name))

	return parseRestyResponse(rsp, err)
}

func (gc Client) DeleteProject(name string) error {
	rsp, err := gc.resty.R().SetHeader("Content-Type", "application/json").
		SetBody(map[string]bool{
			"force":    false,
			"preserve": false,
		}).Post(fmt.Sprintf("/projects/%s/delete-project~delete", name))

	return parseRestyResponse(rsp, err)
}

func (gc Client) ListProjects(_type string) ([]Project, error) {
	rsp, err := gc.resty.R().SetHeader("accept", "application/json").
		Get(fmt.Sprintf("/projects/?type=%s&d=1&t=1", _type))

	if err != nil {
		return nil, errors.Wrapf(err, "Unable to get Gerrit project")
	}

	if rsp.IsError() {
		return nil, errors.Errorf("wrong response code: %d, body: %s", rsp.StatusCode(), rsp.String())
	}

	var preProjects map[string]Project
	body := rsp.String()[5:]
	if err := json.Unmarshal([]byte(body), &preProjects); err != nil {
		return nil, errors.Wrapf(err, "unable to unmarshal project response, body: %s", rsp.String())
	}

	projects := make([]Project, 0, len(preProjects))
	for k, v := range preProjects {
		v.Name = k
		projects = append(projects, v)
	}

	return projects, nil
}

func (gc Client) ListProjectBranches(projectName string) ([]Branch, error) {
	rsp, err := gc.resty.R().SetHeader("accept", "application/json").
		Get(fmt.Sprintf("/projects/%s/branches/", projectName))

	if err != nil {
		return nil, errors.Wrapf(err, "Unable to get Gerrit project branches")
	}

	if rsp.IsError() {
		return nil, errors.Errorf("wrong response code: %d, body: %s", rsp.StatusCode(), rsp.String())
	}

	var branches []Branch
	body := rsp.String()[5:]
	if err := json.Unmarshal([]byte(body), &branches); err != nil {
		return nil, errors.Wrapf(err, "unable to unmarshal project response, body: %s", rsp.String())
	}

	return branches, nil
}