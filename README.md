**Directory Structure**
```
fivetran-webhook-bridge/
├── go.mod
├── main.go
├── handler/
│   └── webhook.go         # Handles incoming /webhook endpoint
├── internal/
│   ├── config/
│   │   └── config.go       # Environment variable loading
│   ├── event/
│   │   ├── event.go        # Defines FivetranEvent struct (outer envelope)
│   │   ├── event_data.go   # Defines specific event Data structs + Parse methods
│   │   └── transformer.go  # Applies business rules to transform events
│   ├── eventgrid/
│   │   └── client.go       # Sends events to Azure Event Grid
│   └── logger/
│       └── logger.go       # Custom logger wrapper (optional, but clean)
├── models/
│   └── eventgrid.go        # Defines EventGridEvent struct (public model)
└── README.md
```
**Visual Flow for Event Grid**
```
flowchart TD
  A[Fivetran Webhook (POST)] --> B[handler/webhook.go]
  B --> C[internal/event/event.go (unmarshal Event)]
  C --> D[internal/event/event_data.go (parse inner Data)]
  D --> E[internal/event/transformer.go (apply rules)]
  E --> F[internal/eventgrid/client.go (send to Event Grid)]
```