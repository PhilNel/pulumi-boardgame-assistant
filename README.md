# Board Game Assistant Citations

Pulumi infrastructure configuration for managing citation references used by the Board Game Assistant knowledge base.

## Prerequisites

- Pulumi CLI installed
- Go 1.23+
- AWS credentials configured

## Development

Below is an example of a citation used for the Nemesis boardgame:

```go
{
    GameId:        "nemesis",
    ReferenceId:   "R1-BASIC-ACTION",
    Type:          "rulebook",
    Title:         "Nemesis Core Rulebook",
    Section:       "Basic Actions",
    PageReference: "p.13",
}
```

Citations in the knowledge base use the format `[[R1-BASIC-ACTION]]` where `R1` refers to the Nemesis Core Rulebook and `BASIC-ACTION` identifies the specific rule section (Basic Actions in this case). 

The references deployment creates a DynamoDB table and seeds it with citation data that maps these identifiers to their source details.

### Deploy Components

```bash
# Deploy games configuration
make games

# Deploy references configuration
make references

# Format Go code
make fmt
```

## Related Repositories

- [`go-boardgame-assistant`](https://github.com/PhilNel/go-boardgame-assistant) - Collection of Lambdas used to process the knowledge base and provide an API to the Board Game Assistant project.

- [`infra-boardgame-assistant`](https://github.com/PhilNel/infra-boardgame-assistant) - Terraform configuration for deploying the infrastructure and managing Lambda permissions, S3 buckets, etc.

- [`knowledge-boardgame-assistant`](https://github.com/PhilNel/knowledge-boardgame-assistant) - Collection of structured board game rules in markdown format that forms the knowledge base for this project.

- [`vue-boardgame-assistant`](https://github.com/PhilNel/vue-boardgame-assistant) - The frontend Vue website that is used to interact with the Board Game Assistant functionality.
