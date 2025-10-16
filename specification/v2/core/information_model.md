## Information Model 

The TRQP Assertion API is based on four core entities—**authority\_id**,
**entity\_id**, **assertion\_id**, and **context**—each represented as JSON
objects or values.  All JSON objects follow [RFC 8259]({{RFC8259}}).

---

### authority\_id 

The **ecosystem\_id** identifies the party (service or system) asserting or evaluating the claim.

* **Type:** string
* **Required:** yes
* **Description:**
  A unique identifier for the authority making the assertion or performing the evaluation.
* **Example:**

  ```json
  "ecosystem_id": "auth-service-A"
  ```

---

### entity\_id 

The **entity\_id** identifies the subject (user, device, service, etc.) about whom the assertion is made.

* **Type:** string
* **Required:** yes
* **Description:**
  A unique identifier for the entity under evaluation, scoped by the authority’s domain.
* **Example:**

  ```json
  "entity_id": "user-1234"
  ```

---

### assertion\_id 

The **assertion\_id** specifies which claim, role, or right is being queried.

* **Type:** string
* **Required:** yes
* **Description:**
  A unique identifier for the specific claim or authorization type.
* **Example:**

  ```json
  "assertion_id": "role-admin"
  ```

---

### context 

The **context** object carries auxiliary parameters that influence evaluation, such as timestamps.

* **Type:** object
* **Required:** no
* **Description:**
  A JSON object whose members convey evaluation context.  At minimum, implementations SHOULD recognize a `time` field in [RFC 3339]({{RFC3339}}) format; additional keys MAY be defined by profiles or bindings.
* **Properties:**

  ```yaml
  time:
    type: string
    format: date-time
    description: |
      RFC 3339 timestamp (UTC, “Z” suffix).  
      If omitted, the server MUST use its current time.
  ```
* **Example:**

  ```json
  "context": {
    "time": "2025-06-01T12:00:00Z"

