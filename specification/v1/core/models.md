## TRQP Bindings
_This section is normative_

TRQP Bindings are technical specifications that define how to implement the TRQP
Core protocol over a specific transport protocol. Currently, only the RESTful
binding is available.

To be a compatible binding, the following requirements must be met:

* All compliant [[ref:TRQP Binding]]s MUST support the required interfaces described in the Required Interfaces (Section 9) section.
* A compliant [[ref:TRQP Binding]] MUST adhere to the [[ref:TRQP Core]] requirements.
* A compliant [[ref:TRQP Binding]] MUST support versioning using Semantic Versioning 2.0.

## Error Handling 

### Error Response Considerations

_this section is normative_

#### Query Error Handling Guidelines
_this section is informative_

This document outlines general guidelines for handling errors in responses to
queries within the Trust Registry Query Protocol. The approach described here is
abstracted from any specific transport or protocol (such as HTTP) to offer
guidance applicable across various implementations.

### Error Codes
_This section is normative_

Status Codes

TRQP uses a structured range of status codes to indicate the outcome of a query.
These codes are grouped to provide clarity and support future extensibility
without breaking compatibility:

* `0–99:` Success Codes — The operation completed successfully.
* `100–199:` General Errors — Unspecified or system-level errors.
* `200–299:` Resource Errors — Issues related to resource availability or access.
* `300–399:` Authentication Errors — Issues related to authenticating the querier.
* `400–499:` Validation Errors — Problems validating query inputs.

| Status Code | Return Message | Details                      |
|-------------|----------------|------------------------------|
| TRQP-0      | success        | query completed successfully |
| TRQP-100    | error          | error                        |
| TRQP-200    | notfound       | not found                    |
| TRQP-300    | unauthorized   | authorization error          |
| TRQP-400    | invalidrequest | invalid request              |

## Queries

This section defines the query types available under the TRQP Binding. Each query type is processed against a TRQP‐compliant trust registry endpoint and **MUST** adhere to the TRQP Core and TRQP Binding requirements. The following query types are defined:

1. **Description (Metadata) Query**  
2. **Authorization Query**  
3. **Ecosystem Recognition Query**  
4. **Delegation Query**

### Metadata (Description) Query

The metadata query supports descriptive information about an ecosystem or a
registry. If no ecosystem is provided, it's default behavior is to provide
information about the registry itself. It is expected for the [[ref:TRQP Profile]] to define the payloads in more detail.

#### Request Parameters Table

| Parameter    | Type   | Required? | Description                                                          | Example       |
|--------------|--------|-----------|----------------------------------------------------------------------|---------------|
| ecosystem_id | string | Optional  | Identifier for scoping the metadata request to a specific ecosystem. | "ecosystem A" |

*Example Request:*

```json
{ "ecosystem_id": "ecosystem A" }
```

#### Response

Response model is left to the [[ref:TRQP Profile]] and [[TRQP:Binding]] to define. The RESTful binding allows an arbitrary JSON response.

#### Metadata Query Errors

| Error Name                     | When                                              | Description                                                                    | Status Code |
|--------------------------------|---------------------------------------------------|--------------------------------------------------------------------------------|-------------|
| Ecosystem Identifier Not Found | The provided ecosystem identifier does not exist. | Indicates that the ecosystem identifier specified was not found.               | TRQP-201    |
| Malformed Request	Request   | parameters are missing or incorrectly formatted.  | Indicates that the request lacks required parameters or contains invalid data. | TRQP-400    |

### Authorization Query

In an authorization statement, an authority grants an authorization to an entity
under its authority. In the ToIP governance model, this entity is called a
governed party. This query serves the authorization statements of the ecosystem. 

#### Request Parameters Table

| Parameter        | Type   | Required? | Description                                                              | Example                |
|------------------|--------|-----------|--------------------------------------------------------------------------|------------------------|
| ecosystem_id     | string | Yes       | An ecosystem identifier as defined in the TRQP Binding.                  | “ecosystem A”          |
| authorization_id | string | Yes       | MUST match one of the defined authorization types in the TRQP Binding.   | “credential-A-issuer”  |
| entity_id        | string | Yes       | Identifies the entity for which the authorization is being queried.      | “random-id-1234”       |
| time             | string | Optional  | A timestamp in RFC3339 UTC format indicating when to evaluate the query. | “2025-04-01T00:00:00Z” |

Example Request:

```json
{
  "ecosystem_id": "ecosystem A",
  "authorization_id": "credential-A-issuer",
  "entity_id": "random-id-1234",
  "time": "2025-04-01T00:00:00Z"
}
```

#### **Response**

* TRQP bindings MUST return the Authorization Status of the entry. 

The Status Table below describes possible statuses. The response **MUST** have one of the following statuses:

**Authorization Status Table**

| Authorization  Status | Description                                                    |
|-----------------------|----------------------------------------------------------------|
| authorized            | The entity holds the requested authorization.                  |
| not-authorized        | The entity does not hold the requested authorization.          |
| revoked               | The authorization was previously granted but has been revoked. |
| unknown-subject       | The entity is not recognized or does not exist.                |
| error                 | An error occurred while evaluating the authorization query.    |


Additional details, such as validity information or supporting proof references,
MAY be included in the response as per the binding and profile requirements.

**Required Behavior**

The system MUST clearly indicate whether the subject holds the specified
authorization at the evaluated time. If no `time` is provided, `time` SHOULD be
evaluated as the time the request was received by the registry.

#### Authorization Query Errors

| Error Name                   | When                                                        | Description                                                                        | Status Code |
|------------------------------|-------------------------------------------------------------|------------------------------------------------------------------------------------|-------------|
| Ecosystem ID Not Found       | The specified ecosystem identifier is not recognized.       | Indicates that the ecosystem identifier does not exist.                            | TRQP-201         |
| Invalid Authorization Type   | The provided authorization type does not match known types. | Indicates that the authorization type specified is invalid.                        | TRQP-200         |
| Authorization Type Not Found | The provided authorization type is not available.           | Indicates that the authorization type specified is not found.                      | TRQP-200         |
| Unknown Entity ID            | The provided entity identifier does not exist in records.   | Indicates that the entity ID is unknown.                                           | TRQP-200         |
| Invalid Time Requested       | The time parameter is invalid or incorrectly formatted.     | Indicates that the requested time does not conform to the expected RFC3339 format. | TRQP-400         |


### Ecosystem Recognition Query

In a recognition statement, one ecosystem governing authority recognizes another
ecosystem governing authority as a peer. The following query shares the recognition status. 

#### Request Parameters Table

| Parameter    | Type   | Required? | Description                                                                                             | Example                |
|--------------|--------|-----------|---------------------------------------------------------------------------------------------------------|------------------------|
| authority_id | string | Yes       | The identifier for the requesting ecosystem as defined in the TRQP Binding.                             | “ecosystem A”          |
| entity_id    | string | Optional  | Another ecosystem identifier against which recognition is being evaluated.                              | “ecosystem B”          |
| scope        | string | Optional  | A filter or context to narrow the recognition query; specific structure defined by individual profiles. | “financial-services”   |
| time         | string | Optional  | A timestamp in RFC3339 UTC format indicating when to evaluate the recognition query.                    | “2025-04-01T00:00:00Z” |

**Example Request:**

```json
{
  "authority_id": "ecosystem A",
  "entity_id": "ecosystem B",
  "scope": "financial-services",
  "time": "2025-04-01T00:00:00Z"
}
```

#### Response

Recognition Status: MUST be one of the following:

**Recognition Status Table**

| Recognition Status | Description                                |
|--------------------|--------------------------------------------|
| recognized         | The recognition relationship is confirmed. |
| not-recognized     | The recognition relationship is denied.    |

* Optional Fields: Additional supporting details such as proof references or log entries MAY be included.

**Behavior:**

The system MUST return a clear yes/no answer regarding ecosystem recognition and MAY provide further explanatory details as specified in the TRQP Binding.

#### Ecosystem Recognition Query Errors

| Error Name                    | When                                                        | Description                                                                        | Status Code |
|-------------------------------|-------------------------------------------------------------|------------------------------------------------------------------------------------|-------------|
| Authority ID Not Found        | The requesting authority identifier is not recognized.      | Indicates that the authority id is not registered.                                 | TRQP-201    |
| Entity ID Not Found           | The entity id  is unknown or unrecognized.                  | Indicates that the entity id  does not exist.                                      | TRQP-200    |
| Scope Not Found               | The specified scope does not match any known context.       | Indicates that the target ecosystem or scope is not found.                         | TRQP-200    |
| Malformed Recognition Request | Request parameters are incomplete or incorrectly formatted. | Indicates that essential elements of the recognition query are missing or invalid. | TRQP-400    |

### Delegation Query

:::note
The specifics of the Delegation Query model are pending further details.
The following serves as a placeholder specification and should be expanded as
additional requirements become available.
:::

#### Request Parameters Table

| Parameter	Type | Required? | Description | Example                                                                       |                        |
|-------------------|-----------|-------------|-------------------------------------------------------------------------------|------------------------|
| delegator_id      | string    | Yes         | The identifier for the authority delegating its authority.                    | “authority-123”        |
| delegatee_id      | string    | Yes         | The identifier for the entity or authority receiving the delegated authority. | “entity-456”           |
| scope             | string    | Optional    | The scope within which the delegation applies.                                | “limited-access”       |
| time              | string    | Optional    | A timestamp in RFC3339 UTC format indicating when to evaluate the delegation. | “2025-04-01T00:00:00Z” |

Example Request:

```json
{
  "delegator_id": "authority-123",
  "delegatee_id": "entity-456",
  "scope": "limited-access",
  "time": "2025-04-01T00:00:00Z"
}
```

#### Response
* Delegation Status: MUST be one of the following:

| Delegation Status | Description                               |
|-------------------|-------------------------------------------|
| delegated         | The delegation relationship is confirmed. |
| not-delegated     | The is no delegation relationship         |
| revoked           | The delegation relationship was revoked   |

Additional details or supporting information regarding the delegation MAY be included per the binding.

#### Delegation Query Errors

| Error Name                   | When                                                     | Description                                                                                    | Status Code |
|------------------------------|----------------------------------------------------------|------------------------------------------------------------------------------------------------|-------------|
| Delegator ID Not Found       | The specified delegator identifier is not recognized.    | Indicates that the delegator ID does not exist.                                                | TBD         |
| Delegatee ID Not Found       | The specified delegatee identifier is not recognized.    | Indicates that the delegatee ID does not exist.                                                | TBD         |
| Malformed Delegation Request | Request parameters are missing or incorrectly formatted. | Indicates that essential elements of the delegation query are missing or contain invalid data. | TBD         |
