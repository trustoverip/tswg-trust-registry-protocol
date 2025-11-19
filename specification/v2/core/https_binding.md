## HTTPS Binding 

This section defines the requirements for making TRQP authorization and recognition queries over HTTPS with JSON.

---

### Common Request Headers

All HTTPS calls **MUST** include:

* `Content-Type: application/json`
* `Authorization: Bearer <token>` *(as required by the service)*
* Optional tracing header: `X-Request-ID: <uuid>`

---

### Authorization over HTTPS

#### HTTPS Authorization Request

```http
POST /authorization HTTP/1.1
Host: registry.example.com
Content-Type: application/json
Authorization: Bearer eyJ...
X-Request-ID: d4f34c12-9b7a-4e3a-a5d1-7e4f8c2c9f10

{
  "entity_id":    "user-1234",
  "authority_id": "auth-service-A",
  "action":       "issue",
  "resource":     "engineer-license",
  "context": {
    "time": "2025-06-19T11:30:00Z"
  }
}
```

#### HTTPS Authorization Response

A successful authorization call returns HTTP 200 with the JSON body below:

```http
HTTP/1.1 200 OK
Content-Type: application/json
X-Request-ID: d4f34c12-9b7a-4e3a-a5d1-7e4f8c2c9f10

{
  "entity_id":          "user-1234",
  "authority_id":       "auth-service-A",
  "action":             "issue",
  "resource":           "engineer-license",
  "authorized":         true,
  "time_requested":     "2025-06-19T11:30:00Z",
  "time_evaluated":     "2025-06-19T11:30:00Z",
  "message":            "User-1234 holds the admin role.",
  "context": {
    "time": "2025-06-19T11:30:00Z"
  }
}
```

---

### Recognition over HTTPS

#### HTTPS Recognition Request

```http
POST /recognition HTTP/1.1
Host: registry.example.com
Content-Type: application/json
Authorization: Bearer eyJ...
X-Request-ID: bfe9eb29-ab87-4ca3-be83-a1d5d8305716

{
  "entity_id":    "service-42",
  "authority_id": "did:example",
  "action":       "govern",
  "resource":     "professional-engineers",
  "context": {
    "time": "2025-06-19T10:00:00Z"
  }
}
```

#### HTTPS Recognition Response

A successful recognition returns HTTP 200 with the JSON body below:

```http
HTTP/1.1 200 OK
Content-Type: application/json
X-Request-ID: bfe9eb29-ab87-4ca3-be83-a1d5d8305716

{
  "entity_id":      "service-42",
  "authority_id":   "did:example",
  "action":         "govern",
  "resource":       "professional-engineers",
  "recognized":     true,
  "time_requested": "2025-06-19T10:00:00Z",
  "time_evaluated": "2025-06-19T10:00:00Z",
  "message":        "Service-42 is recognized by auth-master.",
  "context": {
    "time": "2025-06-19T10:00:00Z"
  }
}
```

Error conditions (e.g. malformed JSON, unauthorized, not found) are signaled via standard HTTP 4xx/5xx status codes and a Problem Details JSON body.

---

### Error Handling

* **400 Bad Request** — invalid JSON or missing required fields
* **401 Unauthorized** — missing/invalid bearer token
* **404 Not Found** — entity, authority, action, or resource not recognized
* **500 Internal Server Error** — unexpected server failure

Error responses use the [Problem Details for HTTP APIs](https://datatracker.ietf.org/doc/html/rfc7807) format:

```json
{
  "type":   "https://example.com/problems/invalid-action",
  "title":  "action not found",
  "status": 404,
  "detail": "action \"issue\" is not defined for authority auth-service-A."
}
```

