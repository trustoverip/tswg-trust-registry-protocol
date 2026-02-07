
---

## Authorization Query and Response Schemas

*This section is normative.*

The purpose of a TRQP authorization query is to ask the question “Does `authority_id` authorize `entity_id` to take `action` on `resource` (with optional context conditions such as `time`)?" 

TRQP authorization queries and responses MUST conform to the JSON schemas defined in this section.

### Authorization Query Schema

```json
[[insert: ./specification/v2/core/schema/trqp_authorization_request.jsonschema]]
```

**Example authorization query:**

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

### Authorization Response Schema

```json
[[insert: ./specification/v2/core/schema/trqp_authorization_response.jsonschema]]
```

**Example authorization response:**

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

## Recognition Query and Response Schemas
 
*This section is normative.*

The purpose of a TRQP recognition query is to ask the question “Does `authority_id` recognize `entity_id` (another authority) to be authoritative for `action` on `resource`?"

TRQP recognition queries and responses MUST conform to the JSON schemas defined in this section.

### Recognition Query Schema

```json
[[insert: ./specification/v2/core/schema/trqp_recognition_request.jsonschema]]
```

**Example recognitiion query:**

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

### Recognition Response

```json
[[insert: ./specification/v2/core/schema/trqp_recognition_response.jsonschema]]
```

**Example recognition response:**

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
