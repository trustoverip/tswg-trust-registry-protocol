
---

## Authorization API 

The Authorization API asks “Does **entity\_id** hold **assertion\_id** according to **ecosystem\_id** (under **context**)?”

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
  "authority_id": "auth-service-A",
  "action":       "issue",
  "resource":     "country:state:driverlicense",
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
  "entity_id":    "did:user-1234",
  "authority_id": "auth-service-A",
  "action":       "issue",
  "resource":     "country:state:driverlicense",
  "authorized":   true,
  "time":         "2025-06-19T11:30:00Z",
  "message":      "did:user-1234 is authorized for issue+country:state:driverlicense (action+resource) by auth-service-A.",
}
``` 

## Recognition API 

The Recognition API asks “Is **entity\_id** recognized by **ecosystem\_id** for **assertion\_id** under **context**?”

### RecognitionRequest 

The RecognitionRequest JSON Schema file is located [here](./schema/trqp_recognition_request.jsonschema)

```json
[[insert: ./specification/v2/core/schema/trqp_recognition_request.jsonschema]]
```


**Example request:**

```http
POST /recognition
Content-Type: application/json

{
  "entity_id":    "service-42",
  "authority_id": "did:example",
  "action":       "recognize",
  "resource":     "listed-registry",
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
  "entity_id":    "service-42",
  "authority_id": "did:example",
  "action":       "recognize",
  "resource":     "listed-registry",
  "recognized":   true,
  "message":      "Service-42 is recognized by did:example.",
}
```
