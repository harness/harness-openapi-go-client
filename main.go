package main

import (
	"context"
	"fmt"

	ngmanager "github.com/ashishsanodia-harness/harness-swagger-go-client-poc/ngmanager"

	"github.com/antihax/optional"
)

func main() {
}

func swaggerCodeGenGetOrgs() {

	cfg := ngmanager.Configuration{
		BasePath:      "https://qa.harness.io/gateway/ng/api",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
	}

	cfg.AddDefaultHeader("x-api-key", "<x-api-key-value>")

	swaggerClient := ngmanager.NewAPIClient(&cfg)

	organizationApiGetOrganizationsOpts := ngmanager.OrganizationApiGetOrganizationsOpts{
		Account: optional.NewString("<account>")}

	orgzz, _, orgzzErr := swaggerClient.OrganizationApi.GetOrganizations(context.Background(), &organizationApiGetOrganizationsOpts)

	responsez := orgzz[0]
	if orgzzErr == nil {
		fmt.Println("name            : ", responsez.Org.Name)
		fmt.Println("slug            : ", responsez.Org.Slug)
		fmt.Println("description     : ", responsez.Org.Description)
		fmt.Println("tags            : ", responsez.Org.Tags)
		fmt.Println("created         : ", responsez.Created)
		fmt.Println("updated         : ", responsez.Updated)
		fmt.Println("harness_managed : ", responsez.HarnessManaged)
		fmt.Println("-------------------------------")
	}
}

func swaggerCodeGenGetAccountScopedSecret() {

	cfg := ngmanager.Configuration{
		BasePath:      "<host>",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
	}

	cfg.AddDefaultHeader("x-api-key", "<x-api-key-value>")

	swaggerClient := ngmanager.NewAPIClient(&cfg)

	secretSlug := "samplegithubconnector1credential"
	accountSecretApiGetAccountScopedSecretOpts := ngmanager.AccountSecretApiGetAccountScopedSecretOpts{
		Account: optional.NewString("<account>")}

	secretz, _, secretzErr := swaggerClient.AccountSecretApi.GetAccountScopedSecret(context.Background(), secretSlug, &accountSecretApiGetAccountScopedSecretOpts)

	if secretzErr != nil {
		panic(secretzErr)
	}

	fmt.Println("name 							: ", secretz.Secret.Name)
	fmt.Println("slug 							: ", secretz.Secret.Slug)
	fmt.Println("org 								: ", secretz.Secret.Org)
	fmt.Println("project 						: ", secretz.Secret.Project)
	fmt.Println("description 				: ", secretz.Secret.Description)
	fmt.Println("tags 							: ", secretz.Secret.Tags)
	fmt.Println("created 						: ", secretz.Created)
	fmt.Println("updated						: ", secretz.Updated)
	fmt.Println("secret_type 				: ", secretz.Secret.SecretSpec.Type_)
	fmt.Println("spec.type 					: ", secretz.Secret.SshKerberosTgtPasswordSpec.Type_)
	fmt.Println("spec.port 					: ", secretz.Secret.SshKerberosTgtPasswordSpec.Port)
	fmt.Println("spec.principal 		: ", secretz.Secret.SshKerberosTgtPasswordSpec.Principal)
	fmt.Println("spec.realm 				: ", secretz.Secret.SshKerberosTgtPasswordSpec.Realm)
	fmt.Println("spec.password 			: ", secretz.Secret.SshKerberosTgtPasswordSpec.Password)
	fmt.Println("draft 							: ", secretz.Draft)
	fmt.Println("governance_metadata: ", secretz.GovernanceMetadata)
}

func swaggerCodeGenGetAccountScopedSecrets() {

	cfg := ngmanager.Configuration{
		BasePath:      "<host>",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
	}

	cfg.AddDefaultHeader("x-api-key", "<x-api-key-value>")
	swaggerClient := ngmanager.NewAPIClient(&cfg)

	accountSecretApiGetAccountScopedSecretsOpts := ngmanager.AccountSecretApiGetAccountScopedSecretsOpts{
		Account: optional.NewString("<account>")}
	secretzz, _, secretzzErr := swaggerClient.AccountSecretApi.GetAccountScopedSecrets(context.Background(), &accountSecretApiGetAccountScopedSecretsOpts)
	if secretzzErr == nil {
		for _, element := range secretzz {
			fmt.Println("name 				: ", element.Secret.Name)
			fmt.Println("slug 				:", element.Secret.Slug)
			fmt.Println("org 					:", element.Secret.Org)
			fmt.Println("project			:", element.Secret.Project)
			fmt.Println("description 	:", element.Secret.Description)
			fmt.Println("tags 				:", element.Secret.Tags)
			fmt.Println("created 			:", element.Created)
			fmt.Println("updated 			:", element.Updated)
			fmt.Println("secret_type 	:", element.Secret.SecretSpec.Type_)
			fmt.Println("spec 				:")
			if element.Secret.SecretSpec.Type_ == "SecretText" {
				fmt.Println("						", element.Secret.SecretTextSpec)
				fmt.Println("						", element.Secret.SecretTextSpec.Type_)
				fmt.Println("						", element.Secret.SecretTextSpec.Value)
				fmt.Println("						", element.Secret.SecretTextSpec.ValueType)
				fmt.Println("						", element.Secret.SecretTextSpec.SecretManagerSlug)
			}
		}
	}
}
