# Instructions

To add this generated code into the IBM Terraform Provider:

### Changes to `provider.go`

- Add the following entries to `DataSourcesMap`:
```
    "ibm_cm_catalog": dataSourceIBMCmCatalog(),
    "ibm_cm_offering": dataSourceIBMCmOffering(),
    "ibm_cm_version": dataSourceIBMCmVersion(),
    "ibm_cm_offering_instance": dataSourceIBMCmOfferingInstance(),
```

- Add the following entries to `ResourcesMap`:
```
    "ibm_cm_catalog": resourceIBMCmCatalog(),
    "ibm_cm_offering": resourceIBMCmOffering(),
    "ibm_cm_version": resourceIBMCmVersion(),
    "ibm_cm_offering_instance": resourceIBMCmOfferingInstance(),
```

### Changes to `config.go`

- Add an import for the generated Go SDK:
```
    "github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
```

- Add a method to the `ClientSession interface`:
```
    CatalogManagementV1()   (*catalogmanagementv1.CatalogManagementV1, error)
```

- Add two fields to the `clientSession struct`:
```
    catalogManagementClient     *catalogmanagementv1.CatalogManagementV1
    catalogManagementClientErr  error
```

- Implement a new method on the `clientSession struct`:
```
    func (session clientSession) CatalogManagementV1() (*catalogmanagementv1.CatalogManagementV1, error) {
        return session.catalogManagementClient, session.catalogManagementClientErr
    }
```

- In the `ClientSession()` method of `Config`, below the existing line that creates an authenticator:
```
    var authenticator *core.BearerTokenAuthenticator
```
  add the code to initialize the service client
```
    // Construct an "options" struct for creating the service client.
    catalogManagementClientOptions := &catalogmanagementv1.CatalogManagementV1Options{
        Authenticator: authenticator,
    }

    // Construct the service client.
    session.catalogManagementClient, err = catalogmanagementv1.NewCatalogManagementV1(catalogManagementClientOptions)
    if err == nil {
        // Enable retries for API calls
        session.catalogManagementClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
        // Add custom header for analytics
        session.catalogManagementClient.SetDefaultHeaders(gohttp.Header{
            "X-Original-User-Agent": { fmt.Sprintf("terraform-provider-ibm/%s", version.Version) },
        })
    } else {
        session.catalogManagementClientErr = fmt.Errorf("Error occurred while configuring Catalog Management API service: %q", err)
    }
```
