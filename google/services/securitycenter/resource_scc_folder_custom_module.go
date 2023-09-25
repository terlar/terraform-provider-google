// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package securitycenter

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceSecurityCenterFolderCustomModule() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityCenterFolderCustomModuleCreate,
		Read:   resourceSecurityCenterFolderCustomModuleRead,
		Update: resourceSecurityCenterFolderCustomModuleUpdate,
		Delete: resourceSecurityCenterFolderCustomModuleDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSecurityCenterFolderCustomModuleImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"custom_config": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `The user specified custom configuration for the module.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"predicate": {
							Type:     schema.TypeList,
							Required: true,
							Description: `The CEL expression to evaluate to produce findings. When the expression evaluates
to true against a resource, a finding is generated.`,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"expression": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `Textual representation of an expression in Common Expression Language syntax.`,
									},
									"description": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `Description of the expression. This is a longer text which describes the
expression, e.g. when hovered over it in a UI.`,
									},
									"location": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `String indicating the location of the expression for error reporting, e.g. a
file name and a position in the file.`,
									},
									"title": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `Title for the expression, i.e. a short string describing its purpose. This can
be used e.g. in UIs which allow to enter the expression.`,
									},
								},
							},
						},
						"recommendation": {
							Type:     schema.TypeString,
							Required: true,
							Description: `An explanation of the recommended steps that security teams can take to resolve
the detected issue. This explanation is returned with each finding generated by
this module in the nextSteps property of the finding JSON.`,
						},
						"resource_selector": {
							Type:     schema.TypeList,
							Required: true,
							Description: `The resource types that the custom module operates on. Each custom module
can specify up to 5 resource types.`,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"resource_types": {
										Type:        schema.TypeList,
										Required:    true,
										Description: `The resource types to run the detector on.`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"severity": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: verify.ValidateEnum([]string{"CRITICAL", "HIGH", "MEDIUM", "LOW"}),
							Description:  `The severity to assign to findings generated by the module. Possible values: ["CRITICAL", "HIGH", "MEDIUM", "LOW"]`,
						},
						"custom_output": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Custom output properties.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"properties": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: `A list of custom output properties to add to the finding.`,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: `Name of the property for the custom output.`,
												},
												"value_expression": {
													Type:     schema.TypeList,
													Optional: true,
													Description: `The CEL expression for the custom output. A resource property can be specified
to return the value of the property or a text string enclosed in quotation marks.`,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"expression": {
																Type:        schema.TypeString,
																Required:    true,
																Description: `Textual representation of an expression in Common Expression Language syntax.`,
															},
															"description": {
																Type:     schema.TypeString,
																Optional: true,
																Description: `Description of the expression. This is a longer text which describes the
expression, e.g. when hovered over it in a UI.`,
															},
															"location": {
																Type:     schema.TypeString,
																Optional: true,
																Description: `String indicating the location of the expression for error reporting, e.g. a
file name and a position in the file.`,
															},
															"title": {
																Type:     schema.TypeString,
																Optional: true,
																Description: `Title for the expression, i.e. a short string describing its purpose. This can
be used e.g. in UIs which allow to enter the expression.`,
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `Text that describes the vulnerability or misconfiguration that the custom
module detects. This explanation is returned with each finding instance to
help investigators understand the detected issue. The text must be enclosed in quotation marks.`,
						},
					},
				},
			},
			"display_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateRegexp(`^[a-z][\w_]{0,127}$`),
				Description: `The display name of the Security Health Analytics custom module. This
display name becomes the finding category for all findings that are
returned by this custom module. The display name must be between 1 and
128 characters, start with a lowercase letter, and contain alphanumeric
characters or underscores only.`,
			},
			"enablement_state": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: verify.ValidateEnum([]string{"ENABLED", "DISABLED"}),
				Description:  `The enablement state of the custom module. Possible values: ["ENABLED", "DISABLED"]`,
			},
			"folder": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Numerical ID of the parent folder.`,
			},
			"ancestor_module": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `If empty, indicates that the custom module was created in the organization, folder,
or project in which you are viewing the custom module. Otherwise, ancestor_module
specifies the organization or folder from which the custom module is inherited.`,
			},
			"last_editor": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The editor that last updated the custom module.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name of the custom module. Its format is "folders/{folder_id}/securityHealthAnalyticsSettings/customModules/{customModule}".
The id {customModule} is server-generated and is not user settable. It will be a numeric id containing 1-20 digits.`,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The time at which the custom module was last updated.

A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and
up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceSecurityCenterFolderCustomModuleCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandSecurityCenterFolderCustomModuleDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	enablementStateProp, err := expandSecurityCenterFolderCustomModuleEnablementState(d.Get("enablement_state"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enablement_state"); !tpgresource.IsEmptyValue(reflect.ValueOf(enablementStateProp)) && (ok || !reflect.DeepEqual(v, enablementStateProp)) {
		obj["enablementState"] = enablementStateProp
	}
	customConfigProp, err := expandSecurityCenterFolderCustomModuleCustomConfig(d.Get("custom_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("custom_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(customConfigProp)) && (ok || !reflect.DeepEqual(v, customConfigProp)) {
		obj["customConfig"] = customConfigProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "folders/{{folder}}/securityHealthAnalyticsSettings/customModules")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterBasePath}}folders/{{folder}}/securityHealthAnalyticsSettings/customModules")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new FolderCustomModule: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating FolderCustomModule: %s", err)
	}
	if err := d.Set("name", flattenSecurityCenterFolderCustomModuleName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "folders/{{folder}}/securityHealthAnalyticsSettings/customModules/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating FolderCustomModule %q: %#v", d.Id(), res)

	return resourceSecurityCenterFolderCustomModuleRead(d, meta)
}

func resourceSecurityCenterFolderCustomModuleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterBasePath}}folders/{{folder}}/securityHealthAnalyticsSettings/customModules/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("SecurityCenterFolderCustomModule %q", d.Id()))
	}

	if err := d.Set("name", flattenSecurityCenterFolderCustomModuleName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderCustomModule: %s", err)
	}
	if err := d.Set("display_name", flattenSecurityCenterFolderCustomModuleDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderCustomModule: %s", err)
	}
	if err := d.Set("enablement_state", flattenSecurityCenterFolderCustomModuleEnablementState(res["enablementState"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderCustomModule: %s", err)
	}
	if err := d.Set("update_time", flattenSecurityCenterFolderCustomModuleUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderCustomModule: %s", err)
	}
	if err := d.Set("last_editor", flattenSecurityCenterFolderCustomModuleLastEditor(res["lastEditor"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderCustomModule: %s", err)
	}
	if err := d.Set("ancestor_module", flattenSecurityCenterFolderCustomModuleAncestorModule(res["ancestorModule"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderCustomModule: %s", err)
	}
	if err := d.Set("custom_config", flattenSecurityCenterFolderCustomModuleCustomConfig(res["customConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderCustomModule: %s", err)
	}

	return nil
}

func resourceSecurityCenterFolderCustomModuleUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	enablementStateProp, err := expandSecurityCenterFolderCustomModuleEnablementState(d.Get("enablement_state"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enablement_state"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enablementStateProp)) {
		obj["enablementState"] = enablementStateProp
	}
	customConfigProp, err := expandSecurityCenterFolderCustomModuleCustomConfig(d.Get("custom_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("custom_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, customConfigProp)) {
		obj["customConfig"] = customConfigProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "folders/{{folder}}/securityHealthAnalyticsSettings/customModules")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterBasePath}}folders/{{folder}}/securityHealthAnalyticsSettings/customModules/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating FolderCustomModule %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("enablement_state") {
		updateMask = append(updateMask, "enablementState")
	}

	if d.HasChange("custom_config") {
		updateMask = append(updateMask, "customConfig")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating FolderCustomModule %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating FolderCustomModule %q: %#v", d.Id(), res)
	}

	return resourceSecurityCenterFolderCustomModuleRead(d, meta)
}

func resourceSecurityCenterFolderCustomModuleDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	lockName, err := tpgresource.ReplaceVars(d, config, "folders/{{folder}}/securityHealthAnalyticsSettings/customModules")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterBasePath}}folders/{{folder}}/securityHealthAnalyticsSettings/customModules/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting FolderCustomModule %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "FolderCustomModule")
	}

	log.Printf("[DEBUG] Finished deleting FolderCustomModule %q: %#v", d.Id(), res)
	return nil
}

func resourceSecurityCenterFolderCustomModuleImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"folders/(?P<folder>[^/]+)/securityHealthAnalyticsSettings/customModules/(?P<name>[^/]+)",
		"(?P<folder>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "folders/{{folder}}/securityHealthAnalyticsSettings/customModules/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenSecurityCenterFolderCustomModuleName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenSecurityCenterFolderCustomModuleDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterFolderCustomModuleEnablementState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterFolderCustomModuleUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterFolderCustomModuleLastEditor(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterFolderCustomModuleAncestorModule(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterFolderCustomModuleCustomConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["predicate"] =
		flattenSecurityCenterFolderCustomModuleCustomConfigPredicate(original["predicate"], d, config)
	transformed["custom_output"] =
		flattenSecurityCenterFolderCustomModuleCustomConfigCustomOutput(original["customOutput"], d, config)
	transformed["resource_selector"] =
		flattenSecurityCenterFolderCustomModuleCustomConfigResourceSelector(original["resourceSelector"], d, config)
	transformed["severity"] =
		flattenSecurityCenterFolderCustomModuleCustomConfigSeverity(original["severity"], d, config)
	transformed["description"] =
		flattenSecurityCenterFolderCustomModuleCustomConfigDescription(original["description"], d, config)
	transformed["recommendation"] =
		flattenSecurityCenterFolderCustomModuleCustomConfigRecommendation(original["recommendation"], d, config)
	return []interface{}{transformed}
}
func flattenSecurityCenterFolderCustomModuleCustomConfigPredicate(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["expression"] =
		flattenSecurityCenterFolderCustomModuleCustomConfigPredicateExpression(original["expression"], d, config)
	transformed["title"] =
		flattenSecurityCenterFolderCustomModuleCustomConfigPredicateTitle(original["title"], d, config)
	transformed["description"] =
		flattenSecurityCenterFolderCustomModuleCustomConfigPredicateDescription(original["description"], d, config)
	transformed["location"] =
		flattenSecurityCenterFolderCustomModuleCustomConfigPredicateLocation(original["location"], d, config)
	return []interface{}{transformed}
}
func flattenSecurityCenterFolderCustomModuleCustomConfigPredicateExpression(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterFolderCustomModuleCustomConfigPredicateTitle(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterFolderCustomModuleCustomConfigPredicateDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterFolderCustomModuleCustomConfigPredicateLocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterFolderCustomModuleCustomConfigCustomOutput(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["properties"] =
		flattenSecurityCenterFolderCustomModuleCustomConfigCustomOutputProperties(original["properties"], d, config)
	return []interface{}{transformed}
}
func flattenSecurityCenterFolderCustomModuleCustomConfigCustomOutputProperties(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"name":             flattenSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesName(original["name"], d, config),
			"value_expression": flattenSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpression(original["valueExpression"], d, config),
		})
	}
	return transformed
}
func flattenSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpression(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["expression"] =
		flattenSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpressionExpression(original["expression"], d, config)
	transformed["title"] =
		flattenSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpressionTitle(original["title"], d, config)
	transformed["description"] =
		flattenSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpressionDescription(original["description"], d, config)
	transformed["location"] =
		flattenSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpressionLocation(original["location"], d, config)
	return []interface{}{transformed}
}
func flattenSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpressionExpression(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpressionTitle(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpressionDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpressionLocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterFolderCustomModuleCustomConfigResourceSelector(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["resource_types"] =
		flattenSecurityCenterFolderCustomModuleCustomConfigResourceSelectorResourceTypes(original["resourceTypes"], d, config)
	return []interface{}{transformed}
}
func flattenSecurityCenterFolderCustomModuleCustomConfigResourceSelectorResourceTypes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterFolderCustomModuleCustomConfigSeverity(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterFolderCustomModuleCustomConfigDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterFolderCustomModuleCustomConfigRecommendation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandSecurityCenterFolderCustomModuleDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleEnablementState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPredicate, err := expandSecurityCenterFolderCustomModuleCustomConfigPredicate(original["predicate"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPredicate); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["predicate"] = transformedPredicate
	}

	transformedCustomOutput, err := expandSecurityCenterFolderCustomModuleCustomConfigCustomOutput(original["custom_output"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCustomOutput); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["customOutput"] = transformedCustomOutput
	}

	transformedResourceSelector, err := expandSecurityCenterFolderCustomModuleCustomConfigResourceSelector(original["resource_selector"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResourceSelector); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["resourceSelector"] = transformedResourceSelector
	}

	transformedSeverity, err := expandSecurityCenterFolderCustomModuleCustomConfigSeverity(original["severity"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSeverity); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["severity"] = transformedSeverity
	}

	transformedDescription, err := expandSecurityCenterFolderCustomModuleCustomConfigDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedRecommendation, err := expandSecurityCenterFolderCustomModuleCustomConfigRecommendation(original["recommendation"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRecommendation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["recommendation"] = transformedRecommendation
	}

	return transformed, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigPredicate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpression, err := expandSecurityCenterFolderCustomModuleCustomConfigPredicateExpression(original["expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["expression"] = transformedExpression
	}

	transformedTitle, err := expandSecurityCenterFolderCustomModuleCustomConfigPredicateTitle(original["title"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTitle); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["title"] = transformedTitle
	}

	transformedDescription, err := expandSecurityCenterFolderCustomModuleCustomConfigPredicateDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLocation, err := expandSecurityCenterFolderCustomModuleCustomConfigPredicateLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigPredicateExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigPredicateTitle(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigPredicateDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigPredicateLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigCustomOutput(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProperties, err := expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputProperties(original["properties"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProperties); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["properties"] = transformedProperties
	}

	return transformed, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputProperties(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedValueExpression, err := expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpression(original["value_expression"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValueExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["valueExpression"] = transformedValueExpression
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpression, err := expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpressionExpression(original["expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["expression"] = transformedExpression
	}

	transformedTitle, err := expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpressionTitle(original["title"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTitle); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["title"] = transformedTitle
	}

	transformedDescription, err := expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpressionDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLocation, err := expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpressionLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpressionExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpressionTitle(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpressionDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpressionLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigResourceSelector(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedResourceTypes, err := expandSecurityCenterFolderCustomModuleCustomConfigResourceSelectorResourceTypes(original["resource_types"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResourceTypes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["resourceTypes"] = transformedResourceTypes
	}

	return transformed, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigResourceSelectorResourceTypes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigSeverity(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigRecommendation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
