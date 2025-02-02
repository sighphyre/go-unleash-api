# go-unleash-api

This is a Go client library to interact with the [Unleash Admin API](https://docs.getunleash.io/api) - in early development.

**Requires `unleash-server` v4.13 or higher.**

## Supported Endpoints

- [x] Features
  - [x] GetFeature
  - [x] CreateFeature
  - [x] UpdateFeature
  - [x] ArchiveFeature
  - [x] GetAllFeaturesByProject
  - [x] AddStrategyToFeature
  - [x] UpdateFeatureStrategy
  - [x] DeleteStrategyFromFeature
  - [x] EnableFeatureOnEnvironment
  - [x] GetAllFeatureTags
  - [x] CreateFeatureTags
  - [x] UpdateFeatureTags
  - [x] DeleteFeatureTags
- [x] Strategies
  - [x] CreateStrategy
  - [x] UpdateStrategy
  - [x] DeprecateStrategy
  - [x] ReactivateStrategy
  - [x] GetAllStrategies
  - [x] GetStrategyByName
- [x] Projects
  - [x] GetProjectById
  - [x] CreateProject
  - [x] UpdateProject
  - [x] DeleteProject
- [x] Features v2
  - [x] AddUserProject
  - [x] UpdateUserProject
  - [x] DeleteUserProject

## How to use

Import the package into your project:

```
import (
  ...
  "github.com/sighphyre/go-unleash-api/api"
)
```

Create an API Client to your Unleash instance:
```
client, err := api.NewClient(&http.Client{}, "<unleash-instance-api-url>", "<unleash-api-token>")
```

### Feature Toggle Tags service

Get all the tags of a feature toggle:

```
allFeatureTags, response, err := client.FeatureTags.GetAllFeatureTags("<your-toggle-id>")

if err != nil {
  fmt.Printf("client: error making http request: %s\n", err)
  body, err := ioutil.ReadAll(response.Request.Body)
  if err != nil {
    panic(err)
  } else {
    fmt.Print(body)
  }
} else {
  fmt.Print(allFeatureTags)
  fmt.Print(response)
}
```

Create tags into a feature toggle:

```
createdFeatureTag, response, err := client.FeatureTags.CreateFeatureTags("<your-toggle-id>",
  api.FeatureTag{
    Type:  "simple",
    Value: "add-my-tag",
  },
)

if err != nil {
  fmt.Printf("client: error making http request: %s\n", err)
  body, err := ioutil.ReadAll(response.Request.Body)
  if err != nil {
    panic(err)
  } else {
    fmt.Print(body)
  }
} else {
  fmt.Print(createdFeatureTag)
  fmt.Print(response)
}
```

Update tags into a feature toggle:

```
updatedFeatureTag, response, err := client.FeatureTags.UpdateFeatureTags("teeste", []api.FeatureTag{}, []api.FeatureTag{
	{
    Type:  "simple",
    Value: "update-my-tag",
  },
})

if err != nil {
  fmt.Printf("client: error making http request: %s\n", err)
  body, err := ioutil.ReadAll(response.Request.Body)
  if err != nil {
    panic(err)
  } else {
    fmt.Print(body)
  }
} else {
  fmt.Print(updatedFeatureTag)
  fmt.Print(response)
}
```

Delete tags from a feature toggle:

```
response, err := client.FeatureTags.DeleteFeatureTags("teeste", api.FeatureTag{
  Type:  "simple",
  Value: "feature2",
})

if err != nil {
  fmt.Printf("client: error making http request: %s\n", err)
} else {
  fmt.Print(response)
}
```