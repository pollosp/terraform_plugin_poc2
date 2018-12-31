package artifact

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"gopkg.in/resty.v1"
)

type artifactsjson struct {
	Type string   `json:"type"`
	List []string `json:"list"`
}

func dataSource() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRead,

		Schema: map[string]*schema.Schema{
			"artifacttype": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"artifactlist": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
		},
	}
}

func artifactType() (artifactsjson, error) {
	resp, _ := resty.R().Get("http://localhost:8080/test")

	var artifacts artifactsjson
	json.Unmarshal(resp.Body(), &artifacts)
	return artifacts, nil

}

func dataSourceRead(d *schema.ResourceData, meta interface{}) error {

	artifact, err := artifactType()

	d.Set("artifactlist", artifact.List)

	if err == nil {
		d.Set("artifacttype", artifact.Type)
		d.SetId(time.Now().UTC().String())
	} else {
		return fmt.Errorf("Error requesting external Artifact Type: %d", err)
	}

	return nil

}
