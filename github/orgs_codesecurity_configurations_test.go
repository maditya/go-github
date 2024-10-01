package github

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestOrganizationsService_GetOrganizationCodeSecurityConfigurations(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/orgs/o/code-security/configurations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[
		{
			"id":1,
			"name":"config1",
			"code_scanning_default_setup": "enabled"
		},
		{
			"id":2,
			"name":"config2",
			"private_vulnerability_reporting": "enabled"
		}]`)
	})

	configurations, _, err := client.Organizations.GetOrganizationCodeSecurityConfigurations(context.Background(), "o")
	if err != nil {
		t.Errorf("Organizations.GetOrganizationCodeSecurityConfigurations returned error: %v", err)
	}

	want := []*CodeSecurityConfiguration{
		{ID: Int64(1), Name: String("config1"), CodeScanningDefaultSetup: String("enabled")},
		{ID: Int64(2), Name: String("config2"), PrivateVulnerabilityReporting: String("enabled")},
	}
	if !reflect.DeepEqual(configurations, want) {
		t.Errorf("Organizations.GetOrganizationCodeSecurityConfigurations returned %+v, want %+v", configurations, want)
	}
}
