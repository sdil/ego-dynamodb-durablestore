# eGo DynamoDB Durable Store Plugin

This plugin provides a DynamoDB-based durable state store for the eGo framework. It allows you to persist the state of your eGo actors in an AWS DynamoDB table.

## Features

- **Durable State Storage**: Persist the full state of your eGo actors in DynamoDB.
- **Stateless Client**: The DynamoDB client used in this plugin is stateless, meaning no connection management is required.
- **Versioning**: Each state is versioned to ensure consistency.

## Installation

To install the plugin, run:

```bash
go get github.com/sdil/ego-dynamodb-durablestore
```

## Usage

Configuration

First, ensure that you have states_store DynamoDB table in your AWS account.

And then, you can initialize DynamoDB durable store like below:

```go
import (
    "context"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/tochemey/ego/v3"
    "github.com/sdil/ego-dynamodb-durablestore/dynamodb"
)

func main() {
    // Create the DynamoDB durable store
    durableStore := dynamodb.NewStateStore()

    // Create the eGo engine
    engine := ego.NewEngine("Sample", nil, ego.WithStateStore(durableStore))

    // Start the eGo engine
    if err := engine.Start(context.Background()); err != nil {
        log.Fatalf("failed to start eGo engine: %v", err)
    }
}
```

## Implementing Durable State Behavior

Define your actor's state, commands, and behavior using Google Protocol Buffers. Implement the DurableStateBehavior interface for your actor:

```go
type AccountBehavior struct {
    id string
}

func NewAccountBehavior(id string) *AccountBehavior {
    return &AccountBehavior{id: id}
}

func (a *AccountBehavior) ID() string {
    return a.id
}

func (a *AccountBehavior) InitialState() ego.State {
    return ego.State(new(samplepb.Account))
}

func (a *AccountBehavior) HandleCommand(_ context.Context, command ego.Command, _ ego.State) (ego.State, error) {
    switch cmd := command.(type) {
    case *samplepb.CreateAccount:
        return &samplepb.Account{
            AccountId:      cmd.GetAccountId(),
            AccountBalance: cmd.GetAccountBalance(),
        }, nil
    default:
        return nil, errors.New("unhandled command")
    }
}
```

## Running the Example

To run the example, ensure you have a DynamoDB table named states_store with the appropriate schema. Then, execute your Go application:

```bash
go run .
```

## Schema

The DynamoDB table schema should be as follows:

- Partition Key: PersistenceID (String)
- Attributes:
  - VersionNumber (Number)
  - StatePayload (Binary)
  - StateManifest (String)
  - Timestamp (Number)
  - ShardNumber (Number)

## Contributing

Contributions are welcome! Please read the contributing guidelines for more information.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

