Use this plugin for deploying an application to [ContainerShip](https://containership.io).
You can override the default configuration with the following parameters:

- `organization` - Your ContainerShip organization
- `api_key` - Your ContainerShip organization API Key
- `cluster_id` - ID of your ContainerShip cluster
- `application` - Name of the application, defaults to repo name
- `docker_image` - Docker image to use, including tag (`MyOrg/MyImage:latest`)

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
deploy:
  containership:
    organization: MyContainerShipOrganization
    api_key: abcdef1234567890
    cluster_id: abcdef1234567890
    application: my-app-name
    docker_image: MyOrg/MyImage:latest
```
