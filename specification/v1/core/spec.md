## Terms and Definitions
_This section is informative._

 [[def:authorization relationship, authorization relationships]]
~ An authority statement asserting the authorization an authority grants to another party over which it has authority.

 [[def:authorization query, authorization queries]]
~ A request for an authority statement asserting an authorization relationship or a delegation relationship.

 [[def:authority state, authority states]]
~ The set of authority statements that describe a particular entity at a particular point in time.

 [[def:authority statement, authority statements]]
~ An assertion by an authority about either: a) the authorization or delegation it grants to another party over which it has authority, or b) the recognition it gives to a peer about the scope of that peer's authority.

 [[def:delegation relationship, delegation relationships]]
~ An authority statement asserting the rights an authority delegates to another party over which it has authority.

 [[def:digital trust ecosystem, digital trust ecosystems]]
~ A [digital ecosystem](https://glossary.trustoverip.org/#term:digital-ecosystem) in which participants are one or more interoperating [trust communities](https://glossary.trustoverip.org/#term:trust-communities). Governance of various [roles](https://glossary.trustoverip.org/#term:roles) within this ecosystem is typically managed by a [governing body](https://glossary.trustoverip.org/#term:governing-body) using a [governance framework](https://glossary.trustoverip.org/#term:governance-framework). Many digital trust ecosystems maintain one or more [trust registries](https://glossary.trustoverip.org/#term:trust-registries).

 [[def:ecosystem]]
~ See [[ref:digital trust ecosystem]].

 [[def:ecosystem governance framework, ecosystem governance frameworks]]
~ A [governance framework](https://glossary.trustoverip.org/#term:governance-framework) for a [digital trust ecosystem](https://glossary.trustoverip.org/#term:digital-trust-ecosystem). This may incorporate other types of frameworks such as [credential governance frameworks](https://glossary.trustoverip.org/#term:credential-governance-framework).

 [[def:inter-ecosystem]]
~ An adjective describing relationships and data exchanges between participants in two or more separate ecosystems operating under separate governance frameworks.

 [[def:intra-ecosystem]]
~ An adjective describing relationships and data exchanges between participants operating within the same ecosystem and the same governance frameworks.

 [[def:hierarchical authority relationship, hierarchical authority relationships]]
~ A unilateral and exclusive relationship between an authority and another party subject to that authority. The authority is the only one who can grant or revoke authorization from the authorized party.

 [[def:metadata query, metadata queries]]
~ A request for an authority statement describing an entity.

 [[def:recognition relationship, recognition relationships]]
~ A heterarchical authority relationship between two peer authorities, each authoritative for their own ecosystem. This relationship can be unilateral or bilateral and is non-exclusive. One authority attests to the other's authority in one or both directions.

 [[def:recognition query, recognition queries]]
~ A request for an authority statement asserting an recognition relationship.

 [[def:TRQP binding, TRQP bindings]]
~ A technical specification defining how to implement the TRQP Core protocol over a specific transport protocol.

 [[def:TRQP bridge, TRQP bridges]]
~ A system that connects a [[ref:TRQP endpoint]] to a [[ref:system of record]]. The bridge transforms a TRQP query into the query format supported by the system of record. It also performs the reverse mapping for the response.

 [[def:TRQP Core]]
~ The foundational specification that defines core data models, queries, and security requirements for the Trust Registry Query Protocol.

 [[def:TRQP consumer]]
~ A network device (client or server) that send TRQP queries to a TRQP endpoint.

 [[def:TRQP endpoint]]
~ The network service endpoint for trust registry that speaks TRQP.

 [[def:system of record, systems of record]]
~ An authoritative source for the authority statements governing the participants in a digital trust ecosystem.

## Scope
_This section is informative._

Figure 1 illustrates the four primary components involved with TRQP architecture.

![TRQP primary components](images/trqp_components.png)

**Figure 1:** The primary components involved in TRQP architecture.

The scope of this specification is limited to the TRQP protocol operating between TRQP consumers and TRQP endpoints representing addressable TRQP trust registries. The following are out-of-scope:

* **Systems of record**. This specification casts no requirements on how the system of record is designed or deployed. Also, because TRQP is read-only, this specification does not address create, update, or delete operations for the system of record.
* **TRQP bridges**. If the system of record is not a native TRQP trust registry, a TRQP bridge is needed to transform a TRQP query into the query format supported by the system of record. Seperate specifications may be published for popular TRQP bridges, however they are out-of-scope for this specification.

## High-Level Architecture 
*This section is informative.*

Figure 2 illustrates the relationships between the core concepts in TRQP architecture.

![images/authority_model.png](images/authority_model.png)

*Figure 2: Overview of the core concepts in TRQP architecture*

### Ecosystem Governing Authorities and Trust Registry Operators

At the top of Figure&nbsp;2 are the two primary actors involved in TRQP infrastructure—the **ecosystem governing authority** and the **trust registry operator**. From a legal standpoint, they are the real-world entities with ultimate responsibility for the infrastructure that will serve the authority statements. Key considerations about these two roles:

- **Both roles can be played by the same entity.** Although the roles are shown separately in Figure&nbsp;2, the ecosystem [governing authority](https://glossary.trustoverip.org/#term:governing-body) may also serve as the trust registry operator. If the ecosystem chooses to use a separate trust registry operator, then from a ToIP governance architecture standpoint, the operator serves as an [administering authority](https://glossary.trustoverip.org/#term:administering-body).
- **The legal responsibilities of these actors—including liability and indemnity—depend on the ecosystem governance framework** and any specific contractual terms it requires. Those considerations are out-of-scope for this specification.
- **An ecosystem may be served by multiple trust registries and a trust registry may serve multiple ecosystems.** This multiplicity can be especially helpful when designing a group of related ecosystems.
- **Both roles publish authority statements—however it is important to distinguish between them.** The ecosystem governing authority is authoritative for statements describing or implementing the policies in the ecosystem governance framework, while the trust registry operator is authoritative for metadata statements describing the capabilities and operations of the trust registry itself (those that are under the operator’s sole control).

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

### Error Response Considerations

_this section is normative_

#### Query Error Handling Guidelines
_this section is informative_

This document outlines general guidelines for handling errors in responses to queries within the Trust Registry Query Protocol. The approach described here is abstracted from any specific transport or protocol (such as HTTP) to offer guidance applicable across various implementations.

While this does not require HTTP, the error codes are loosely aligned with [HTTP Error Codes](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status) and [DNS Codes](https://help.dnsfilter.com/hc/en-us/articles/4408415850003-DNS-return-codes).

Currently, this section is informative.

#### General Data Model for Errors

Error responses should provide the following structured information which MUST be described in the binding.

- **code** *(number)*: A numeric code identifying the type of error.  
- **message** *(string)*: A clear and descriptive explanation for developers and implementers.  
- **details** *(optional, object)*: Additional context that aids in diagnosing or rectifying the issue.

The following section describes the suggested code number and the situations when you should use the response.

#### Recommendations for Implementers
_this section is non-normative_

- Error responses should be consistent and predictable.  
- Clearly differentiate between recoverable errors (such as malformed requests) and terminal conditions (such as missing resources).  
- Include contextual information whenever possible to expedite issue resolution.

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

#### Metadata Query Models

* **Request**:  
  There are no mandatory request parameters.  
  * Optionally, an `ecosystem_id` can specify that the metadata request should be interpreted within a specific ecosystem's governance framework context [[ref:Ecosystem Governance Framework]]

* **Response**:  
  * `id`: string. Uniquely identifies the registry. If an `ecosystem_id` is provided, the response must clearly reflect that the returned data is scoped to the specified ecosystem (e.g., "ecosystem A").

#### Metadata Query Errors

- **Ecosystem Identifier Not Found** 
  - **When:** The provided registry identifier does not exist. 
  - **Description:** Indicates the registry identifier specified in the query was not found.
  - **Code Number:** 404
- **Malformed Request**
  - **When:** Request parameters are missing or incorrectly formatted.
  - **Description:** Indicates the request lacks required parameters or contains invalid data.
  - **Code Number:** 400


### Authorization Query
_This section is normative_

#### Authorization Model

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

#### Authorization Query Errors

- **Ecosystem ID Not Found** 
  - **When:** The specified ecosystem ID is not recognized by the registry.
  - **Description:** Indicates the ecosystem identifier does not exist in the registry.
  - **Code Number:** 404
- **Invalid Authorization Type** 
  - **When:** Authorization type provided does not match known types.  
  - **Description:** Indicates the authorization type specified is invalid or unrecognized.  
  - **Code Number:** 400 
- **Authorization Type Not Found** 
  - **When:** Authorization type provided does not match known types.  
  - **Description:** Indicates the authorization type specified is not available.  
  - **Code Number:** 404 
- **Unknown Entity ID** 
  - **When:** The provided entity ID does not exist in registry records.  
  - **Description:** Indicates the entity ID provided in the query is unknown.  
  - **Code Number:** 404 
- **Invalid Time Requested** 
  - **When:** The time parameter provided is invalid or incorrectly formatted.  
  - **Description:** Indicates the requested time parameter does not conform to expected formats.  
  - **Code Number:** 400

### Ecosystem Recognition Query
_This section is normative_

#### Ecosystem Recognition Models

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

#### Ecosystem Recognition Query Errors

- **Ecosystem ID Not Found** 
  - **When:** The ecosystem ID of the requesting ecosystem is not recognized. 
  - **Description:** Indicates that the source ecosystem specified is not registered or recognized. 
  - **Code Number:** 404 
- **Target Ecosystem ID Not Found** 
  - **When:** The ecosystem ID of the target ecosystem is unknown or unrecognized. 
  - **Description:** Indicates the target ecosystem specified in the query does not exist. 
  - **Code Number:** 404 
- **Scope Not Found** 
  - **When:** The ecosystem ID of the target ecosystem is not found. 
  - **Description:** Indicates the target ecosystem specified in the query does not exist. 
  - **Code Number:** 404 
- **Malformed Recognition Request** 
  - **When:** Request parameters are incomplete or incorrectly formatted.  
  - **Description:** Indicates essential elements of the recognition request are missing or invalid.  
  - **Code Number:** 400
