Use the ContainerShip plugin to create or update an application on [ContainerShip](https://containership.io).

The following parameters are used to configure this plugin:

- `organization` - your ContainerShip organization
- `api_key` - your ContainerShip organization API Key
- `cluster_id` - ID of your ContainerShip cluster
- `application` - name of the application you wish to create or update, defaults to `$$DRONE_REPO`
- `image` - docker image to use for this application, including tag (`MyOrg/MyImage:latest`)

The following is a sample ContainerShip configuration in your `.drone.yml` file:

```yaml
deploy:
  containership:
    organization: MyContainerShipOrganization
    api_key: abcdef1234567890
    cluster_id: abcdef1234567890
    application: my-app-name
    image: MyOrg/MyImage:latest
```
