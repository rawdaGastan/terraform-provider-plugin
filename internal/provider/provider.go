package provider

import (
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	pkidClient "github.com/rawdaGastan/pkid/client"
)

// ProviderClient holds metadata / config for use by Terraform resources
type ProviderClient struct {
	url    string
	Client pkidClient.PkidClient
}

// Provider returns a schema.Provider for my provider
func Provider() terraform.ResourceProvider {
	p := &schema.Provider{
		ConfigureFunc: providerConfigure,
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("url", nil),
				Description: "server url of provider service API.",
			},
			"seed": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("seed", nil),
				Description: "seed for generating keys of provider service API.",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap: map[string]*schema.Resource{
			"plugin_pkid_key_query":     resourceKeyQuery(),
			"plugin_pkid_project_query": resourceProjectQuery(),
		},
	}
	return p
}

// newProviderClient is a factory for creating ProviderClient structs
func newProviderClient(seed string, url string) ProviderClient {
	p := ProviderClient{
		url: url,
	}
	privateKey, publicKey, err := pkidClient.GenerateKeyPairUsingSeed(seed)
	if err != nil {
		log.Printf("An error happened while generating keys for pkid: %v\n", err)
	}
	p.Client = pkidClient.NewPkidClient(privateKey, publicKey, url, 5*time.Second)

	return p
}

// providerConfigure parses the config into the Terraform provider meta object
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	url := d.Get("url").(string)
	if url == "" {
		log.Println("Defaulting environment in URL config to use API default PROVIDER_URL...")
	}

	seed := d.Get("seed").(string)
	if seed == "" {
		log.Println("Defaulting environment in SEED config to use API default PROVIDER_SEED...")
	}

	return newProviderClient(seed, url), nil
}
