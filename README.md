[![Actions Status](https://github.com/Hellthrashers/pulumi-airflow/workflows/release/badge.svg)](https://github.com/Hellthrashers/pulumi-airflow/actions)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

# Apache Airflow Resource Provider

The Pulumi Apache Airflow Resource Provider enables you to manage [Apache Airflow](https://airflow.apache.org/) resources using infrastructure as code with Pulumi. This provider is built on top of the [Terraform Airflow Provider](https://github.com/drfaust92/terraform-provider-airflow) and allows you to declaratively manage Airflow connections, variables, pools, and other resources.

## Overview

Apache Airflow is an open-source workflow management platform for data engineering pipelines. This Pulumi provider enables you to:

- Manage Airflow connections programmatically
- Configure Airflow variables
- Set up and manage connection pools
- Integrate Airflow configuration with your infrastructure as code workflows
- Support multiple authentication methods (Basic Auth and OAuth2)

This provider works with any Airflow instance that exposes the Airflow REST API, including:
- Self-hosted Airflow instances
- Google Cloud Composer (both version 1 and 2)
- Amazon MWAA (Managed Workflows for Apache Airflow)
- Azure Data Factory Managed Airflow

## Supported Resources

This provider currently supports the following Airflow resources:

- **Connections**: Manage Airflow connections to external systems
- **Variables**: Configure Airflow variables for your DAGs
- **Pools**: Manage Airflow pools for task parallelism control

## Installation

The Pulumi Airflow provider is available as a package in multiple languages:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm` or `yarn`:

```bash
$ npm install @pulumi/airflow
```

or

```bash
$ yarn add @pulumi/airflow
```

### Python

To use from Python, install using `pip`:

```bash
$ pip install pulumi-airflow
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
$ go get github.com/Hellthrashers/pulumi-airflow/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
$ dotnet add package Pulumi.Airflow
```

## Configuration

The Provider resource accepts the following configuration arguments:

- **baseEndpoint** (Required): The base endpoint URL for the Airflow REST API (e.g., `https://my-airflow-instance/api/v1`)
- **oauth2Token** (Optional): OAuth2 access token for API authentication (recommended for Google Cloud Composer 2)
- **password** (Optional): Password for HTTP basic authentication
- **username** (Optional): Username for HTTP basic authentication

**Note**: You must provide either `oauth2Token` OR both `username` and `password` for authentication.

These arguments can be passed as an object of type `ProviderArgs` to the Provider constructor.

## Example Usage

### Provider Setup with Node.js (JavaScript/TypeScript)

#### Google Cloud Composer 2 (OAuth2)

Google Cloud Composer 2 uses OAuth2 for authentication. You can use the Composer Airflow web UI endpoint with an OAuth2 access token:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as airflow from "@pulumi/airflow";

const myProvider = new airflow.Provider("my-provider", {
    baseEndpoint: "https://my-airflow-instance/api/v1",
    oauth2Token: process.env.AIRFLOW_OAUTH2_TOKEN,
});

// Use `myProvider` to create resources in your Pulumi program.
```

#### Basic Authentication

For self-hosted Airflow or services using basic authentication:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as airflow from "@pulumi/airflow";

const myProvider = new airflow.Provider("my-provider", {
    baseEndpoint: "https://my-airflow-instance/api/v1",
    username: "your-username",
    password: pulumi.secret("your-secure-password"),
});
```

### Creating an Airflow Connection

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as airflow from "@pulumi/airflow";

const postgresConnection = new airflow.Connection("postgres-example", {
    connectionId: "postgres_default",
    connType: "postgres",
    host: "postgres.example.com",
    port: 5432,
    login: "airflow",
    password: pulumi.secret("db-password"),
    schema: "airflow_db",
});
```

### Python Example

```python
import pulumi
import pulumi_airflow as airflow

# Configure the provider
config = pulumi.Config()
provider = airflow.Provider("my-provider",
    base_endpoint="https://my-airflow-instance/api/v1",
    username=config.require("airflow-username"),
    password=config.require_secret("airflow-password")
)

# Create a connection
connection = airflow.Connection("example",
    connection_id="example",
    conn_type="http",
    host="api.example.com",
    opts=pulumi.ResourceOptions(provider=provider)
)
```

## Importing Existing Resources

You can import existing Airflow resources into your Pulumi state. For example, to import a connection:

```sh
$ pulumi import airflow:index/connection:Connection default example
```

Replace `example` with the connection ID of the resource you want to import.

## Development

### Prerequisites

- [Pulumi CLI](https://www.pulumi.com/docs/get-started/install/)
- [Go](https://golang.org/dl/) (1.18 or later)
- [Node.js](https://nodejs.org/en/) (for Node.js SDK)
- [Python 3](https://www.python.org/downloads/) (for Python SDK)
- [.NET Core SDK](https://dotnet.microsoft.com/download) (for .NET SDK)

### Building the Provider

To build the provider and all SDKs, run:

```bash
$ make development
```

This will:
1. Build the provider binary
2. Generate all language SDKs (Node.js, Python, Go, .NET)
3. Install the SDKs locally

### Running Tests

To run the integration tests:

```bash
$ make test
```

## Architecture

This provider is built using the [Pulumi Terraform Bridge](https://github.com/pulumi/pulumi-terraform-bridge), which enables Terraform providers to be used as Pulumi providers. The architecture consists of:

1. **Provider Core**: Go-based provider implementation that bridges the Terraform Airflow provider
2. **Language SDKs**: Auto-generated SDKs for Node.js, Python, Go, and .NET
3. **Resource Mapping**: Maps Terraform resources to Pulumi resources with appropriate naming conventions

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details on:
- Setting up your development environment
- Submitting pull requests
- Code review process
- Commit message conventions

Please also read our [Code of Conduct](CODE-OF-CONDUCT.md) before contributing.

## Support and Resources

- **Documentation**: [Pulumi Airflow Provider Docs](https://www.pulumi.com/registry/packages/airflow/)
- **Examples**: Check the [examples/](examples/) directory for complete examples
- **Issues**: Report bugs or request features via [GitHub Issues](https://github.com/Hellthrashers/pulumi-airflow/issues)
- **Discussions**: Ask questions in [GitHub Discussions](https://github.com/Hellthrashers/pulumi-airflow/discussions)
- **Pulumi Community**: Join the [Pulumi Community Slack](https://slack.pulumi.com/)

## Related Projects

- [Apache Airflow](https://airflow.apache.org/) - The workflow orchestration platform
- [Terraform Airflow Provider](https://github.com/drfaust92/terraform-provider-airflow) - The underlying Terraform provider
- [Pulumi](https://www.pulumi.com/) - Modern infrastructure as code platform
- [Pulumi Terraform Bridge](https://github.com/pulumi/pulumi-terraform-bridge) - Bridge for using Terraform providers with Pulumi

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.