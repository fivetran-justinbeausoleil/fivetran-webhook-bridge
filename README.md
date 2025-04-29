# Fivetran Webhook Bridge

A lightweight Go server that receives Fivetran webhook events, transforms them into Azure Event Grid schema, and publishes them to an Event Grid custom topic.

✅ Built with Go, `net/http`, and clean interfaces  
✅ Can be exposed using ngrok for local development  
✅ Supports easy extension to other systems (ServiceNow, Splunk, etc.)

---

## 📚 Project Architecture

```text
Fivetran Webhook → Webhook Bridge Server → Azure Event Grid
(optional) via ngrok for HTTPS
```

- Receives POST events from Fivetran
- Validates and parses the webhook payload
- Transforms into Azure Event Grid event schema
- Posts to Azure Event Grid Topic

---

## 🛠 Tech Stack

- Go 1.21+
- `net/http`
- `encoding/json`
- Azure Event Grid (Destination)
- ngrok (for local HTTPS exposure)

---

## 🚀 Local Development Setup

### 1. Clone the Repository

```bash
git clone https://github.com/fivetran-justinbeausoleil/fivetran-webhook-bridge.git
cd fivetran-webhook-bridge
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Set Environment Variables

```bash
export EVENT_GRID_TOPIC_URL=https://<your-transformers-grid-endpoint>.eventgrid.azure.net/api/events
export EVENT_GRID_SAS_KEY=<your-transformers-grid-sas-key>
```

### 4. Run Locally

```bash
make run
# or
go run main.go
```

Server will start at: [http://localhost:8080](http://localhost:8080)

---

## 🌍 Expose Server Using ngrok (for Fivetran Integration)

Since Fivetran requires an HTTPS endpoint:

### 1. Install ngrok

```bash
brew install ngrok/ngrok/ngrok
```
Or download manually from [ngrok.com/download](https://ngrok.com/download)

### 2. Authenticate ngrok

```bash
ngrok config add-authtoken <your-auth-token>
```

### 3. Start the Tunnel

```bash
ngrok http 8080
```

You will see a forwarding URL like:  
`https://your-ngrok-id.ngrok-free.app -> http://localhost:8080`

### 4. Configure the Fivetran Webhook Connector

- **Payload URL:** `https://your-ngrok-id.ngrok-free.app/webhook/eventgrid`
- **Events:** Select desired events (e.g., `sync_start`, `sync_end`)
- **Secret:** (Optional)

Save and test the webhook!

---

## 📦 Example `curl` Test

```bash
curl -X POST http://localhost:8080/webhook/eventgrid   -H "Content-Type: application/json"   -d '{
        "transformers": "test_event",
        "created": "2025-04-27T00:00:00Z",
        "connector_type": "test",
        "connector_id": "connector_123",
        "connector_name": "Test Connector",
        "sync_id": "sync_123",
        "destination_group_id": "group_123",
        "data": {
            "status": "SUCCESSFUL"
        }
      }'
```

---

## 👨‍💻 Maintainers

- Justin Beausoleil | [GitHub](https://github.com/fivetran-justinbeausoleil)

---
