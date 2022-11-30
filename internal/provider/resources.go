package provider

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceKeyQuery is where we define the schema of the Terraform resource (for key)
func resourceKeyQuery() *schema.Resource {
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

// resourceKeyQuery is where we define the schema of the Terraform resource (for project)
func resourceProjectQuery() *schema.Resource {
	return &schema.Resource{
		Create: resourceQuerySet,
		Read:   resourceQueryListProject,
		Update: resourceQuerySet,
		Delete: resourceQueryDeleteProject,
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

// resourceQueryListProject tells Terraform how to contact our microservice and retrieve the necessary data
func resourceQueryListProject(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(ProviderClient)
	client := provider.Client

	project := d.Get("project").(string)

	value, err := client.List(project)
	if err != nil {
		return fmt.Errorf("error listing a provider query: %v", err)
	}
	d.SetId(project)
	log.Printf("[INFO] provider query list read: %v", value)
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
	return nil
}

// resourceQueryDeleteProject tells Terraform how to contact our microservice and deletes the necessary data
func resourceQueryDeleteProject(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(ProviderClient)
	client := provider.Client

	project := d.Get("project").(string)

	err := client.DeleteProject(project)
	if err != nil {
		return fmt.Errorf("error deleting project a provider query: %v", err)
	}
	d.SetId(project)
	log.Printf("[INFO] provider query delete")
	return nil
}
