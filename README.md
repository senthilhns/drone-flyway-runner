A plugin to Drone Flyway runner DB migration support plugin.

# Usage

The following settings changes this plugin's behavior.

* param1 (optional) does something.
* param2 (optional) does something different.

Below is an example `.drone.yml` that uses this plugin.

```yaml
kind: pipeline
name: default

steps:
- name: run harnesscommunity/drone-flyway-runner plugin
  image: harnesscommunity/drone-flyway-runner
  pull: if-not-exists
  settings:
    param1: foo
    param2: bar
```

# Building

Build the plugin binary:

```text
scripts/build.sh
```

Build the plugin image:

```text
docker build -t harnesscommunity/drone-flyway-runner -f docker/Dockerfile .
```

# Testing

Execute the plugin from your current working directory:

### Migrate command example

```bash
docker run --rm -e \  
  -e PLUGIN_FLYWAY_COMMAND=migrate \
  -e PLUGIN_USERNAME=hns_usr \
  -e PLUGIN_PASSWORD=sksi$k89 \
  -e PLUGIN_URL=jdbc:mysql://23.12.7.8:3306/flyway_db \
  -e PLUGIN_LOCATIONS=/harness/db-migrations  
  harnesscommunity/drone-flyway-runner
```
### Migrate command with config file example

```bash
docker run --rm -e \  
  -e PLUGIN_FLYWAY_COMMAND=migrate \
  -e PLUGIN_COMMAND_LINE_ARGS='-configFiles=/harness/hns/test-resources/flyway/config1/flyway.conf'

```

### Migrate command with JDBC driver example

```bash
docker run --rm -e \  
  -e PLUGIN_FLYWAY_COMMAND=migrate \
  -e PLUGIN_USERNAME=hns_usr \
  -e PLUGIN_PASSWORD=sksi$k89 \
  -e PLUGIN_URL=jdbc:mysql://23.12.7.8:3306/flyway_db \
  -e PLUGIN_LOCATIONS=/harness/db-migrations \
  -e PLUGIN_DRIVER_PATH=/harness/drivers/mysql-connector-java-8.0.23.jar \  
  harnesscommunity/drone-flyway-runner
```

### Repair command with command line args example

```bash
docker run --rm -e \  
  -e PLUGIN_FLYWAY_COMMAND=repair \
  -e PLUGIN_USERNAME=hnstest03 \
  -e PLUGIN_PASSWORD=3cbc98835323 \
  -e PLUGIN_URL=jdbc:mysql://23.12.7.8:3306/flyway_db \
  -e PLUGIN_LOCATIONS=/harness/db-migrations \
  -e PLUGIN_COMMAND_LINE_ARGS='-connectRetries=3 -cleanDisabled=false -validateOnMigrate=true -baselineOnMigrate=true'
```
