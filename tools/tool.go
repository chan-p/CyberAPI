package tools

type Res_Json struct {
  ProjectMemberCount int `json:projectmembercount`
}

type Res_Json_Domain struct {
  Domain_Type   map[string]int `json:domain_type`
}

type Res_Json_Member_List struct {
  Name          []string  `json:name`
}
type Error struct {
  Massege string `json:massege`
}
