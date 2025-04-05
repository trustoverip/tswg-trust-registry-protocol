## TRQP Bindings
_This section is normative_

TRQP Bindings are technical specifications that define how to implement the TRQP
Core protocol over a specific transport protocol. Currently, only the RESTful
binding is available.

To be a compatible binding, the following requirements must be met:

* All compliant [[ref:TRQP Binding]]s MUST support the required interfaces described in the Required Interfaces (Section 9) section.
* A compliant [[ref:TRQP Binding]] MUST adhere to the [[ref:TRQP Core]] requirements.
* A compliant [[ref:TRQP Binding]] MUST support versioning using Semantic Versioning 2.0.

### RESTful Binding

The RESTful binding can be seen in the [api docs](https://trustoverip.github.io/tswg-trust-registry-protocol/api-docs/) or described in the swagger document below: 

<details>
<summary>Click To View OpenAPI 3.0 Specification</summary>

```yaml
openapi: 3.0.1
info:
  title: TRQP Restful Binding
  version: 1.0.0
  description: |
    This specification defines a RESTful TRQP Binding.
    It includes endpoints for retrieving Trust Registry metadata,
    authorization data, verifying entity authorization status,
    and checking ecosystem recognition.
servers:
  - url: https://example-trust-registry.com
    description: Production server (example)

tags:
  - name: trqp
    description: TRQP Compliant Queries

paths:
  /metadata:
    get:
      summary: Retrieve Trust Registry Metadata
      tags:
        - trqp
      description: |
        Returns Trust Registry Metadata as a JSON object.
      operationId: getTrustRegistryMetadata
      parameters:
        - name: egf_id
          in: query
          required: false
          description: An optional identifier specifying which ecosystem's metadata should be retrieved.
          schema:
            type: string
      responses:
        "200":
          description: Successfully retrieved Trust Registry Metadata.
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/TrustRegistryMetadata"
        "404":
          description: Metadata not found.
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/ProblemDetails"
        "401":
          description: Unauthorized request.
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/ProblemDetails"

  /registries/{entity_id}/recognition:
    get:
      summary: Check Ecosystem Recognition
      tags:
        - trqp
      description: Verifies if the ecosystem governing authority identified by `entity_id` is recognized by the ecosystem governing authority identified by `authority_id`
      operationId: checkEcosystemRecognition
      parameters:
        - name: entity_id
          in: path
          required: true
          description: Unique identifier of the ecosystem governing authority being recognized.
          schema:
            type: string
        - name: authority_id
          in: query
          required: true
          description: Unique identifier of the ecosystem governing authority asserting recognition. Defaults to the ecosystem governing authority of the trust registry (but only if the trust registry serves only a single ecosystem governing authority).
          schema:
            type: string
        - name: time
          in: query
          required: false
          description: RFC3339 timestamp indicating when recognition is checked. Defaults to "now" on system being queried.
          schema:
            type: string
            format: date-time
      responses:
        "200":
          description: Ecosystem recognition successfully verified.
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/RecognitionResponse"
        "401":
          description: Unauthorized request.
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/ProblemDetails"
        "404":
          description: Ecosystem not recognized or not found.
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/ProblemDetails"

  /entities/{entity_id}/authorization:
    get:
      summary: Check Entity Authorization Status
      tags:
        - trqp
      description: |
        Determines if the specified entity (`entity_id`) is authorized under the given authorization identifier (`authorization_id`)
        within the specified governance framework (`egf_id`). Optionally, returns a list of authorizations if `all` is true.
      operationId: checkAuthorizationStatus
      parameters:
        - name: entity_id
          in: path
          required: true
          description: Unique identifier of the entity.
          schema:
            type: string
        - name: authorization_id
          in: query
          required: true
          description: Authorization identifier to evaluate.
          schema:
            type: string
        - name: authority_id
          in: query
          required: true
          description: Unique identifier of the ecosystem governing authority granting authorization.
          schema:
            type: string
        - name: time
          in: query
          required: false
          description: |
            ISO8601/RFC3339 timestamp for authorization status evaluation.
            Defaults to the current time if omitted.
          schema:
            type: string
            format: date-time
      responses:
        "200":
          description: Authorization status successfully retrieved.
          content:
            application/json:
              schema:
                oneOf:
                  - $ref: "#/components/schemas/AuthorizationResponse"
                  - type: array
                    items:
                      $ref: "#/components/schemas/AuthorizationResponse"
        "404":
          description: Entity not found.
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/ProblemDetails"
        "401":
          description: Unauthorized request.
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/ProblemDetails"

components:
  schemas:
    ProblemDetails:
      type: object
      description: |
        A Problem Details object as defined in [RFC 7807](https://datatracker.ietf.org/doc/html/rfc7807).
      properties:
        type:
          type: string
          format: uri
          description: A URI reference that identifies the problem type.
        title:
          type: string
          description: A short, human-readable summary of the problem.
        status:
          type: integer
          description: The HTTP status code (e.g., 404 for "Not Found").
        detail:
          type: string
          description: A human-readable explanation specific to this occurrence of the problem.
        instance:
          type: string
          format: uri
          description: A URI reference that identifies the specific occurrence of the problem.
      additionalProperties: true

    TrustRegistryMetadata:
      type: object
      properties:
        id:
          type: string
          description: Unique identifier of the Trust Registry.
        default_egf_id:
          type: string
          description: Default EGF, identified by DID, that will be used if none is supplied in various queries.
          #TODO: review thinking on defaultEGF_DID
        description:
          type: string
          maxLength: 4096
          description: A description of the Trust Registry.
        name:
          type: string
          description: Human-readable name of the Trust Registry.
        controllers:
          type: array
          description: List of unique identifiers representing the controllers of the Trust Registry.
          items:
            type: string
          minItems: 1
      required:
        - id
        - description
        - name
        - controllers

    AuthorizationResponse:
      type: object
      properties:
        egf_id:
          type: string
          description: EGF DID this authorization response relates to.
        recognized:
          type: boolean
          description: Indicates whether the entity is recognized by the Trust Registry.
        authorized:
          type: boolean
          description: Specifies whether the entity is authorized under the provided authorization ID.
        message:
          type: string
          description: Additional context or information regarding the authorization status.
        evaluated_at:
          type: string
          format: date-time
          description: Timestamp when the authorization status was evaluated.
        response_time:
          type: string
          format: date-time
          description: Timestamp when the response was generated.
        expiry_time:
          type: string
          format: date-time
          description: Timestamp when the authorization status expires (if applicable).
        jws:
          type: string
          description: Signed response object as specified in [RFC 7515](https://datatracker.ietf.org/doc/html/rfc7515) from the controller of the Trust Registry.
      required:
        - recognized
        - authorized
        - message
        - evaluated_at
        - response_time

    RecognitionResponse:
      type: object
      properties:
        recognized:
          type: boolean
          description: Indicates whether the ecosystem ID is recognized by the Trust Registry.
        message:
          type: string
          description: Additional information regarding the recognition status.
        egf_id:
          type: string
          description: EGF DID this recognition applies to.
        evaluated_at:
          type: string
          format: date-time
          description: Timestamp when the recognition status was evaluated.
        response_time:
          type: string
          format: date-time
          description: Timestamp when the response was generated.
        expiry_time:
          type: string
          format: date-time
          description: Timestamp when the recognition status expires (if applicable).
        jws:
          type: string
          description: Signed response object as specified in [RFC 7515](https://datatracker.ietf.org/doc/html/rfc7515) from the controller of the Trust Registry.
      required:
        - recognized
        - message
        - evaluated_at
        - response_time
```
</details>

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
| TRQP-0      | success        | Query completed successfully |
| TRQP-100    | error          | error                        |
| TRQP-200    | notfound       | not found                    |
| TRQP-300    | unauthorized   | Authorization error          |
| TRQP-400    | invalidrequest | Invalid request              |


Below is the complete specification in Markdown format. You can copy and paste this into your Markdown file.

## Queries

This section defines the query types available under the TRQP Binding. Each query type is processed against a TRQP‐compliant trust registry endpoint and **MUST** adhere to the TRQP Core and TRQP Binding requirements. The following query types are defined:

1. **Description (Metadata) Query**  
2. **Authorization Query**  
3. **Ecosystem Recognition Query**  
4. **Delegation Query**

### Description (Metadata) Query

#### Metadata Query Models

#### Request Parameters Table

| Parameter    | Type   | Required? | Description                                                          | Example       |
|--------------|--------|-----------|----------------------------------------------------------------------|---------------|
| ecosystem_id | string | Optional  | Identifier for scoping the metadata request to a specific ecosystem. | "ecosystem A" |

*Example Request:*

```json
{ "ecosystem_id": "ecosystem A" }
```

##### Response

* Fields:
  * id (string): Uniquely identifies the registry. If an ecosystem_id is provided, the response MUST clearly reflect that the returned metadata is scoped to the specified ecosystem (e.g., by including an explicit reference such as “ecosystem A”).


##### Metadata Query Errors

| Error Name                     | When                                              | Description	Status Code                                                     |     |
|--------------------------------|---------------------------------------------------|--------------------------------------------------------------------------------|-----|
| Ecosystem Identifier Not Found | The provided ecosystem identifier does not exist. | Indicates that the ecosystem identifier specified was not found.               | 201 |
| Malformed Request	Request   | parameters are missing or incorrectly formatted.  | Indicates that the request lacks required parameters or contains invalid data. | 400 |

#### Authorization Query

In an authorization statement, an authority grants an authorization to an entity
under its authority. In the ToIP governance model, this entity is called a
governed party. This query serves the authority statements of the ecosystem. 

##### Request Parameters Table

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

##### Response

* Authorization Status: MUST be one of the following:

### Authorization Status Table

| Authorization | Status          | Description                                                    |
|---------------|-----------------|----------------------------------------------------------------|
|               | authorized      | The entity holds the requested authorization.                  | 
|               | not-authorized  | The entity does not hold the requested authorization.          | 
|               | revoked         | The authorization was previously granted but has been revoked. | 
|               | unknown-subject | The entity is not recognized or does not exist.                | 
|               | error           | An error occurred while evaluating the authorization query.    | 


* Optional Fields: Additional details, such as validity information or supporting proof references, MAY be included.

#### Behavior:
The system MUST clearly indicate whether the subject holds the specified authorization at the evaluated time.

#### Authorization Query Errors

| Error Name                   | When                                                        | Description                                                                        | Status Code |
|------------------------------|-------------------------------------------------------------|------------------------------------------------------------------------------------|-------------|
| Ecosystem ID Not Found       | The specified ecosystem identifier is not recognized.       | Indicates that the ecosystem identifier does not exist.                            | 201         |
| Invalid Authorization Type   | The provided authorization type does not match known types. | Indicates that the authorization type specified is invalid.                        | 200         |
| Authorization Type Not Found | The provided authorization type is not available.           | Indicates that the authorization type specified is not found.                      | 200         |
| Unknown Entity ID            | The provided entity identifier does not exist in records.   | Indicates that the entity ID is unknown.                                           | 200         |
| Invalid Time Requested       | The time parameter is invalid or incorrectly formatted.     | Indicates that the requested time does not conform to the expected RFC3339 format. | 400         |


### Ecosystem Recognition Query

In a recognition statement, one ecosystem governing authority recognizes another
ecosystem governing authority as a peer.

#### Ecosystem Recognition Models

#### Request Parameters Table

| Parameter           | Type   | Required? | Description                                                                                             | Example                |
|---------------------|--------|-----------|---------------------------------------------------------------------------------------------------------|------------------------|
| ecosystem_id        | string | Yes       | The identifier for the requesting ecosystem as defined in the TRQP Binding.                             | “ecosystem A”          |
| target_ecosystem_id | string | Optional  | Another ecosystem identifier against which recognition is being evaluated.                              | “ecosystem B”          |
| scope               | string | Optional  | A filter or context to narrow the recognition query; specific structure defined by individual profiles. | “financial-services”   |
| time                | string | Optional  | A timestamp in RFC3339 UTC format indicating when to evaluate the recognition query.                    | “2025-04-01T00:00:00Z” |

#### Example Request:

```json
{
  "ecosystem_id": "ecosystem A",
  "target_ecosystem_id": "ecosystem B",
  "scope": "financial-services",
  "time": "2025-04-01T00:00:00Z"
}
```

#### Response
	•	Recognition Status: MUST be one of the following:

#### Recognition Status Table

| Recognition Status | Description                                |
| accepted           | The recognition relationship is confirmed. |
| rejected           | The recognition relationship is denied.    |

	•	Optional Fields: Additional supporting details such as proof references or log entries MAY be included.

Behavior:

The system MUST return a clear yes/no answer regarding ecosystem recognition and MAY provide further explanatory details as specified in the TRQP Binding.

### Ecosystem Recognition Query Errors

| Error Name                    | When                                                        | Description                                                                        | Status Code |
|-------------------------------|-------------------------------------------------------------|------------------------------------------------------------------------------------|-------------|
| Ecosystem ID Not Found        | The requesting ecosystem identifier is not recognized.      | Indicates that the source ecosystem is not registered.                             | 201         |
| Target Ecosystem ID Not Found | The target ecosystem identifier is unknown or unrecognized. | Indicates that the target ecosystem does not exist.                                | 200         |
| Scope Not Found               | The specified scope does not match any known context.       | Indicates that the target ecosystem or scope is not found.                         | 200         |
| Malformed Recognition Request | Request parameters are incomplete or incorrectly formatted. | Indicates that essential elements of the recognition query are missing or invalid. | 400         |

#### Delegation Query

Note: The specifics of the Delegation Query model are pending further details. The following serves as a placeholder specification and should be expanded as additional requirements become available.

Delegation Query Models

Request Parameters Table

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

##### Response
	•	Delegation Status: MUST be one of the following:
	•	delegated
	•	not-delegated
	•	revoked
	•	error
	•	Optional Fields: Additional details or supporting information regarding the delegation MAY be included.

Delegation Query Errors

| Error Name                   | When                                                     | Description                                                                                    | Status Code |
|------------------------------|----------------------------------------------------------|------------------------------------------------------------------------------------------------|-------------|
| Delegator ID Not Found       | The specified delegator identifier is not recognized.    | Indicates that the delegator ID does not exist.                                                | TBD         |
| Delegatee ID Not Found       | The specified delegatee identifier is not recognized.    | Indicates that the delegatee ID does not exist.                                                | TBD         |
| Malformed Delegation Request | Request parameters are missing or incorrectly formatted. | Indicates that essential elements of the delegation query are missing or contain invalid data. | TBD         |
