## Trust Registry Query Protocol RESTful Binding
_this section is normative_

RESTful TRQP Bindings specification that implements the core specification. 

The following [OpenAPI](./swagger.yaml) Document describes the RESTful endpoints that are required for the TRQP RESTful binding. 

* The `/metadata` endpoint is aligned to the MetadataQuery.
* The `/registries/{ecosystem_id}/recognition` maps to the RecognitionQuery. 
* The `/entities/{entity_id}/authorization` maps to the AuthorizationQuery. 

### HTTP Error Code Mapping
_This section is normative_

The following mapping of error codes to HTTP Status is provided for http-based implementations:

| Return Code            | Return Message          | HTTP Status | HTTP Reason             | 
| -----------            | --------------          | ----------- | -----------             |
| `statuscode`           | `message`               |             |                         |
| TRQP-0                 | success                 | 200         | OK                      |
| TRQP-100               | error                   | 500         | Internal Server Error   |
| TRQP-200               | notfound                | 404         | Not Found               |
| TRQP-300               | unauthorized            | 401         | Not Authorized          |
| TRQP-400               | invalidrequest          | 400         | Invalid request         |

**Additional Information:**

* Error Codes and further detail are represented using [Problem Details described in rfc7807](https://datatracker.ietf.org/doc/html/rfc7807).
* Authorization and Recognition Queries both take timestamps as a required parameter to resolve.
* Time parameters are in the form [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) and MUST be sent in UTC time. 
* If `jws` field in response is provided, verifiers are recommended to use that to verify the response payload controller.

Security considerations are left to the implementation profile to describe. 
Identifier requirements are left to the implementation profile to describe. 
Resolution paths are left to the implementation profile to describe. 

### Swagger

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
        - name: ecosystem_id
          in: query
          required: true
          description: Unique identifier of the ecosystem being queried.
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
        - name: ecosystem_id
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

