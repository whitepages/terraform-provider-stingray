package stingray

import "encoding/json"

type resourceList struct {
	Children []resourceListItem `json:"children"`
}

type resourceListItem struct {
	Name string `json:"name"`
	HRef string `json:"href"`
}

func (rl *resourceList) decode(data []byte) error {
	return json.Unmarshal(data, &rl)
}

func (rl *resourceList) names() []string {
	names := []string{}
	for _, item := range rl.Children {
		names = append(names, item.Name)
	}

	return names
}
