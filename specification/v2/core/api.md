## Recognition API 

The Recognition API asks “Is **entity\_id** recognized by **authority\_id** for **assertion\_id** under **context**?”

### RecognitionRequest 

The RecognitionRequest JSON Schema file is located here: TODO

```json
[[insert: ./specification/v2/core/schema/trqp_recognition_request.jsonschema]]
```


**Example request:**

```http
POST /recognition
Content-Type: application/json

{
  "entity_id":    "service-42",
  "ecosystem_id": "did:example",
  "assertion_id": "peer-recognition",
  "context": {
    "time": "2025-06-19T10:00:00Z"
  }
}
```

### RecognitionResponse

```json
[[insert: ./specification/v2/core/schema/trqp_recognition_response.jsonschema]]
```


**Example response:**

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "entity_id":      "service-42",
  "ecosystem_id":   "did:example",
  "assertion_id":       "peer-recognition",
  "recognized":     true,
  "message":        "Service-42 is recognized by auth-master.",
}
```

---

# Authorization API 

The Authorization API asks “Does **entity\_id** hold **assertion\_id** according to **authority\_id** under **context**?”

### AuthorizationRequest 

```json
[[insert: ./specification/v2/core/schema/trqp_authorization_request.jsonschema]]
```


**Example request:**

```http
POST /authorization
Content-Type: application/json

{
  "entity_id":    "user-1234",
  "ecosystem_id": "auth-service-A",
  "assertion_id": "role-admin",
  "context": {
    "time": "2025-06-19T11:30:00Z"
  }
}
```

### AuthorizationResponse

```json
[[insert: ./specification/v2/core/schema/trqp_authorization_response.jsonschema]]
```


**Example response:**

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "entity_id":          "user-1234",
  "ecosystem_id":       "auth-service-A",
  "assertion_id":       "role-admin",
  "assertion_verified": true,
  "time":               "2025-06-19T11:30:00Z",
  "message":            "User-1234 holds the admin role.",
}
``` 

