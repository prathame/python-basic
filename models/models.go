package models

type Document struct {
	Document_id                           string
	Document_path                         string
	Document_type                         string
	Name                                  string
	Version                               string
	Publish_dt                            string
	Effective_dt                          string
	Review_dt                             string
	Publisher                             string
	Reviewe_name                          string
	Domination_of_control_geography       string
	Domination_of_control_bussiness       string
	Domination_of_control_technology_func string
	Scope_deploy_id                       string
	Scope_technique_id                    string
	Appli_deploy_id                       string
	Appli_technique_id                    string
}
type Statement struct {
	Statement_type           string
	Document_type            string
	Document_section_level_1 string
	Document_section_level_2 string
	Statement_parent_id      string
	Statement_id             string
	Statement_words          string
	Fulfilled_practice       string
	Prac_id                  string
	Act_id                   string
	Scope_technique_id       string
	Scope_deploy_id          string
	Appli_technique_id       string
	Appli_deploy_id          string
}

//doc[i].Document_id, doc[i].Document_path, doc[i].Document_type, doc[i].Name, doc[i].Version, doc[i].Publish_dt, doc[i].Effective_dt, doc[i].Review_dt, doc[i].Publisher, doc[i].Reviewe_name, doc[i].Domination_of_control_geography, doc[i].Domination_of_control_bussiness, doc[i].Domination_of_control_technology_func, doc[i].Scope_deploy_id, doc[i].Scope_technique_id, doc[i].Appli_deploy_id, doc[i].Appli_technique_id
