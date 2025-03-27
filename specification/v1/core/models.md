## Metadata Models

### Trust Registry
_This section is normative._

* **Properties**
  * **id: MUST** be a globally unique identifier for the registry (e.g., URI, DID, UUID)
  * **ecosystem: SHOULD** indicate which ecosystem(s) the registry serves or recognizes
  * **controller: SHOULD** reference the entity that manages or operates the registry

### Ecosystem
_This section is normative_

* **Properties**
  * **id: MUST** be a globally unique identifier for the registry (e.g., URI, DID, UUID)
  * **egf_id: MUST** specify a *resolvable* EGF identifier referencing the official EGF document
  * **trustregistries: MUST** provide a list of authorized Trust Registries that serve the ecosystem authority state [[ref:Authority State]]
    * Each registry **MUST** have the following properties:
      * **endpoint**: The address (URL, DID, etc.) for TRQP queries [[ref:Authority Query]] / [[ref:Recognition Query]]
    * Each registry **MAY** also be scoped to a particular set of authorization states as defined in the Binding [[ref:TRQP Binding]]
  * **controller: SHOULD** include a method of validating ecosystem controllers

## Baseline Requirements For Conformance

### Trust Registry
_This section is normative_

* All TRQP registries **MUST** provide an addressable endpoint resolvable as defined by the Implementation Profile
* All Trust Registries **MUST** supply the required interfaces described in the Required Interfaces section over the *same* addressable endpoint to be TRQP conformant

### TRQP Binding
_This section is normative_

* All compliant [[ref:TRQP Binding]]s **MUST** support the required interfaces described in the Required Interfaces section
* A compliant [[ref:TRQP Binding]] **MUST** comply with [[ref:TRQP Core]] requirements
* A compliant [[ref:TRQP Binding]] **MUST** support versioning using [Semantic Versioning 2.0](https://semver.org/)

### TRQP Profiles
_This section is normative_

* All TRQP profiles **MUST** specify a compliant [[ref:TRQP Binding]]

## Required Interfaces
_This section is normative_

Below are abstract API methods that **MUST** be exposed. Each **[[ref:TRQP Binding]]** **MUST** define a binding (e.g., REST, gRPC, DIDComm) that maps these methods to actual endpoints.

```mermaid
sequenceDiagram
    participant C as Client
    participant R as Trust Registry

    C->>R: Metadata Query<br/>(registryIdentifier)
    R-->>C: Metadata Response<br/>(auth types, credential formats, version info)<br/>MUST provide details for further queries

    C->>R: Authorization Query<br/>(subject, authorizationType, context?)
    R-->>C: Authorization Response<br/>(authorized / not-authorized / revoked / error)<br/>MUST indicate authorization status clearly

    C->>R: Ecosystem Recognition Query<br/>(ecosystemIdentifier, governanceFrameworkRef)
    R-->>C: Recognition Response<br/>(yes/no + reasons for rejection)<br/>MUST indicate acceptance or rejection
```

**Figure 5:** Sequence diagram showing interactions between a client and Trust Registry for the required interfaces.

### Metadata Query
_This section is normative_

* **Request**:  
  There are no mandatory request parameters.  
  * Optionally, an `ecosystem_id` can specify that the metadata request should be interpreted within a specific ecosystem's governance framework context [[ref:Ecosystem Governance Framework]]

* **Response**:  
  * `id`: string. Uniquely identifies the registry. If an `ecosystem_id` is provided, the response must clearly reflect that the returned data is scoped to the specified ecosystem (e.g., "ecosystem A").

### Authorization Query
_This section is normative_

* **Request**:
  * **ecosystem_id**: string. An ecosystem identifier as defined in the TRQP Binding.  
    - *Example*: `"ecosystem_id": "ecosystem A"`
  * **authorization_id**: string. Must match one of the defined authorization types in the TRQP Binding.  
    - *Example*: `"authorization_id": "credential-A-issuer"`
  * **entity_id**: string. Identifies the entity for which the authorization is being queried.  
    - *Example*: `"entity_id": "random-id-1234"`
  * **time**: string (optional). Specifies when the Trust Registry should evaluate the authority query.
    * If supplied, must adhere to the required time format (e.g., RFC3339 UTC)
    * If omitted, the system must use the current time and include that timestamp in its response

* **Response**:  
  A status indicating the entity's authorization:
  - **authorized**
  - **not-authorized**
  - **revoked**
  - **unknown-subject**
  - **error**  
  
  Optionally, the response may include additional details on validity or supporting proof references.

* **Behavior**:  
  The system **MUST** clearly indicate whether the subject holds the specified authorization.

### Ecosystem Recognition Query
_This section is normative_

* **Request**:
  * **ecosystem_id**: string. The identifier for the ecosystem, defined in the TRQP Binding.  
    - *Example*: `"ecosystem_id": "ecosystem A"`
  * **target_ecosystem_id**: string (optional). Another ecosystem identifier against which recognition is being evaluated.
  * **scope**: string (optional). Can filter or narrow the request. The specification doesn't enforce a specific structure for scopes, but individual profiles may define their own conventions.
  * **time**: string (optional). May be provided as described in the TRQP Binding guidelines.

* **Response**:  
  The recognition status of the ecosystem:
  - **accepted** (if recognized)
  - **rejected** (if not)  
  
  Additional supporting details, such as proof references or log entries, may be included.

* **Behavior**:  
  The system **MUST** return a clear yes/no answer regarding ecosystem recognition, and **MAY** provide further explanation as specified in the TRQP Binding.

