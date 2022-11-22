package provider

import (
	"fmt"
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
				Description: "URL of provider service API.",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap: map[string]*schema.Resource{
			"provider_query": resourceQuery(),
		},
	}
	return p
}

// newProviderClient is a factory for creating ProviderClient structs
func newProviderClient(url string) ProviderClient {
	p := ProviderClient{
		url: url,
	}
	privateKey, publicKey := pkidClient.GenerateKeyPair()
	p.Client = pkidClient.NewPkidClient(privateKey, publicKey, url, 5*time.Second)

	return p
}

// providerConfigure parses the config into the Terraform provider meta object
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	url := d.Get("url").(string)
	if url == "" {
		log.Println("Defaulting environment in URL config to use API default PROVIDER_URL...")
	}

	return newProviderClient(url), nil
}

// resourceQuery is where we define the schema of the Terraform data source
func resourceQuery() *schema.Resource {
	return &schema.Resource{
		Create: resourceQuerySet,
		Read:   resourceQueryGet,
		Update: resourceQuerySet,
		Delete: resourceQueryDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"project": {
				Type:     schema.TypeString,
				Required: true,
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
			"encrypt": {
				Type:     schema.TypeBool,
				Required: true,
			},
		},
	}
}

// resourceQuerySet tells Terraform how to contact our microservice and set the necessary data
func resourceQuerySet(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(ProviderClient)
	client := provider.Client

	project := d.Get("project").(string)
	key := d.Get("key").(string)
	value := d.Get("value").(string)
	encrypt := d.Get("encrypt").(bool)

	err := client.Set(project, key, value, encrypt)
	if err != nil {
		return fmt.Errorf("error creating a provider query: %v", err)
	}
	d.SetId(project + "_" + key)
	log.Printf("[INFO] provider query created")
	return resourceQueryGet(d, meta)
}

// resourceQueryGet tells Terraform how to contact our microservice and retrieve the necessary data
func resourceQueryGet(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(ProviderClient)
	client := provider.Client

	project := d.Get("project").(string)
	key := d.Get("key").(string)

	value, err := client.Get(project, key)
	if err != nil {
		return fmt.Errorf("error getting a provider query: %v", err)
	}
	d.SetId(project + "_" + key)
	log.Printf("[INFO] provider query read: %v", value)
	return nil
}

// resourceQueryDelete tells Terraform how to contact our microservice and deletes the necessary data
func resourceQueryDelete(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(ProviderClient)
	client := provider.Client

	project := d.Get("project").(string)
	key := d.Get("key").(string)

	err := client.Delete(project, key)
	if err != nil {
		return fmt.Errorf("error deleting a provider query: %v", err)
	}
	d.SetId(project + "_" + key)
	log.Printf("[INFO] provider query delete")
	return resourceQueryGet(d, meta)
}
