## Recognition API 

The Recognition API asks “Is **entity\_id** recognized by **authority\_id** for **assertion\_id** under **context**?”

### RecognitionRequest 

```json
{
  "$id": "trqp-recognition-request",
  "title": "RecognitionRequest",
  "type": "object",
  "required": ["entity_id","authority_id"],
  "properties": {
    "entity_id": {
      "type": "string",
      "description": "The entity being recognized."
    },
    "authority_id": {
      "type": "string",
      "description": "The authority asserting recognition."
    },
    "assertion_id": {
      "type": "string",
      "description": "The specific recognition relationship or claim."
    },
    "context": {
      "type": "object",
      "description": "Optional parameters influencing evaluation.",
      "properties": {
        "time": {
          "type": "string",
          "format": "date-time",
          "description": "RFC3339 timestamp; defaults to server time."
        }
      },
      "additionalProperties": {
        "type": "string"
      }
    }
  }
}
```

**Example request:**

```http
POST /v1/recognition
Content-Type: application/json

{
  "entity_id":    "service-42",
  "authority_id": "auth-master",
  "assertion_id": "peer-recognition",
  "context": {
    "time": "2025-06-19T10:00:00Z"
  }
}
```

### RecognitionResponse

```json
{
  "$id": "trqp-recognition-response",
  "title": "RecognitionResponse",
  "type": "object",
  "required": [
    "entity_id",
    "authority_id",
    "recognized",
  ],
  "properties": {
    "entity_id":      { "type":"string", "description":"Queried entity." },
    "authority_id":   { "type":"string", "description":"Queried authority." },
    "scope_id":       { "type":"string", "description":"Scope of the recognition" },
    "recognized":     { "type":"boolean", "description":"True if recognized." },
    "message": {
      "type":"string",
      "description":"Optional human-readable details."
    }
  }
}
```

**Example response:**

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "entity_id":      "service-42",
  "authority_id":   "auth-master",
  "scope_id":       "peer-recognition",
  "recognized":     true,
  "time_requested": "2025-06-19T10:00:00Z",
  "time_evaluated": "2025-06-19T10:00:00Z",
  "message":        "Service-42 is recognized by auth-master.",
}
```

---

# Authorization API 

The Authorization API asks “Does **entity\_id** hold **assertion\_id** according to **authority\_id** under **context**?”

### AuthorizationRequest 

```json
{
  "$id": "trqp-authorization-request",
  "title": "AuthorizationRequest",
  "type": "object",
  "required": ["entity_id","authority_id","assertion_id"],
  "properties": {
    "entity_id": {
      "type": "string",
      "description": "The entity being queried."
    },
    "authority_id": {
      "type": "string",
      "description": "The authority making the claim."
    },
    "assertion_id": {
      "type": "string",
      "description": "The specific claim or right to evaluate."
    },
    "context": {
      "type": "object",
      "description": "Optional parameters influencing evaluation.",
      "properties": {
        "time": {
          "type": "string",
          "format": "date-time",
          "description": "RFC3339 timestamp; defaults to server time."
        }
      },
      "additionalProperties": {
        "type": "string"
      }
    }
  }
}
```

**Example request:**

```http
POST /v1/authorization
Content-Type: application/json

{
  "entity_id":    "user-1234",
  "authority_id": "auth-service-A",
  "assertion_id": "role-admin",
  "context": {
    "time": "2025-06-19T11:30:00Z"
  }
}
```

### AuthorizationResponse

```json
{
  "$id": "trqp-authorization-response",
  "title": "AuthorizationResponse",
  "type": "object",
  "required": [
    "entity_id",
    "authority_id",
    "assertion_id",
    "assertion_verified",
  ],
  "properties": {
    "entity_id": {
      "type":"string",
      "description":"Queried entity."
    },
    "authority_id": {
      "type":"string",
      "description":"Queried authority."
    },
    "assertion_id": {
      "type":"string",
      "description":"Queried claim."
    },
    "assertion_verified": {
      "type":"boolean",
      "description":"True if the claim holds."
    },
    "time": {
      "type":"string","format":"date-time",
      "description":"Client time, if supplied."
    },
    "message": {
      "type":"string",
      "description":"Optional human-readable details."
    }
  }
}
```

**Example response:**

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "entity_id":          "user-1234",
  "authority_id":       "auth-service-A",
  "assertion_id":       "role-admin",
  "assertion_verified": true,
  "time":               "2025-06-19T11:30:00Z",
  "message":            "User-1234 holds the admin role.",
}


