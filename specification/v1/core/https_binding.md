## HTTPS Binding {#https-binding}

This section shows how to carry the Recognition and Authorization calls over HTTPS with JSON.

---

### Common Request Headers

All HTTPS calls **MUST** include:

* `Content-Type: application/json`
* `Authorization: Bearer <token>` *(as required by the service)*
* Optional tracing header: `X-Request-ID: <uuid>`

---

## Recognition over HTTPS

### HTTPS Recognition Request

```http
POST /v1/recognition HTTP/1.1
Host: registry.example.com
Content-Type: application/json
Authorization: Bearer eyJ...
X-Request-ID: bfe9eb29-ab87-4ca3-be83-a1d5d8305716

{
  "entity_id":    "service-42",
  "authority_id": "auth-master",
  "assertion_id": "peer-recognition",
  "context": {
    "time": "2025-06-19T10:00:00Z"
  }
}
```

### HTTPS Recognition Response

A successful recognition returns HTTP 200 with the JSON body below:

```http
HTTP/1.1 200 OK
Content-Type: application/json
X-Request-ID: bfe9eb29-ab87-4ca3-be83-a1d5d8305716

{
  "entity_id":      "service-42",
  "authority_id":   "auth-master",
  "assertion_id":   "peer-recognition",
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

## Authorization over HTTPS

### HTTPS Authorization Request

```http
POST /v1/authorization HTTP/1.1
Host: registry.example.com
Content-Type: application/json
Authorization: Bearer eyJ...
X-Request-ID: d4f34c12-9b7a-4e3a-a5d1-7e4f8c2c9f10

{
  "entity_id":    "user-1234",
  "authority_id": "auth-service-A",
  "assertion_id": "role-admin",
  "context": {
    "time": "2025-06-19T11:30:00Z"
  }
}
```

### HTTPS Authorization Response

A successful authorization call returns HTTP 200 with:

```http
HTTP/1.1 200 OK
Content-Type: application/json
X-Request-ID: d4f34c12-9b7a-4e3a-a5d1-7e4f8c2c9f10

{
  "entity_id":          "user-1234",
  "authority_id":       "auth-service-A",
  "assertion_id":       "role-admin",
  "assertion_verified": true,
  "time_requested":     "2025-06-19T11:30:00Z",
  "time_evaluated":     "2025-06-19T11:30:00Z",
  "message":            "User-1234 holds the admin role.",
  "context": {
    "time": "2025-06-19T11:30:00Z"
  }
}
```

---

### Error Handling

* **400 Bad Request** — invalid JSON or missing required fields
* **401 Unauthorized** — missing/invalid bearer token
* **404 Not Found** — entity, authority, or assertion not recognized
* **500 Internal Server Error** — unexpected server failure

Error responses use the [Problem Details for HTTP APIs](https://datatracker.ietf.org/doc/html/rfc7807) format:

```json
{
  "type":   "https://example.com/problems/invalid-assertion",
  "title":  "Assertion not found",
  "status": 404,
  "detail": "Assertion \"role-admin\" is not defined for authority auth-service-A."
}
```

