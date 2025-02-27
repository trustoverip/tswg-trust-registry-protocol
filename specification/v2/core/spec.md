**Trust Registry Query Protocol Core Specification**

**Specification Status:** Draft
**Version: 1.2:** Draft v1.2 

**Companion Docs**
~ [Overview](/v2/)
~ [Bindings](/v2/bindings)

**Participate:**
~ [GitHub repo](https://github.com/trustoverip/tswg-trust-registry-protocol/tree/main)
~ [File a bug](https://github.com/trustoverip/tswg-trust-registry-protocol/issues)

## **Introduction**

Modern digital ecosystems [[ref:Digital Trust Ecosystem]] rely on diverse a
**[[ref:Intra-Ecosystem Trust Framework]]s** (e.g., OpenID
Federation, X.509 Chains, EBSI Trust Chains, TRAIN). These frameworks are often
effective at verifying whether a subject in one ecosystem has a particular
authorization—but only within that ecosystem. When verifiers need to validate
authorizations **across** different ecosystems, they face interoperability
challenges due to incompatible data models, inconsistent APIs, and varying
governance rules.

The **Trust Registry Query Protocol (TRQP)**
addresses this gap by providing a standardized way to query and verify
authorizations and recognitions across ecosystems. It does not replace existing
intra-ecosystem solutions [[ref:Intra-Ecosystem Trust Framework]]; rather, it acts as a
**bridge** between them—a so-called “inter-trust framework [[ref:Inter-Ecosystem Trust Framework]].”
In practical terms, TRQP allows a verifier to answer questions such as:

- “Does **Entity X** have **Authorization Y** under **Ecosystem Z’s** governance
  framework [[ref:Ecosystem Governance Framework]]?” and
- “Is **Entity X** **Recognized** under **Ecosystem Y’s** governance framework
   [[ref:Ecosystem Governance Framework]] for **Z**?”

This specification describes the abstract rules, data models, and query flows
necessary to implement TRQP, leaving concrete details—such as transport
protocols, message formats, and discovery mechanisms to an ecosystem-specific or
domain-specific [[ref:TRQP Binding]]. By adhering to TRQP, implementers
ensure a consistent, secure, and interoperable means of authorization
verification **across** multiple trust frameworks.

## **Terms and Definitions**

 [[def:Authority Query, Authority Queries]]
~ A request (typically formal or protocol-based) that seeks to verify whether a
specific entity (subject) holds a particular authorization, credential, or right
within an ecosystem.

 [[def:Authority State, Authority States]]
~ A recorded or asserted status indicating whether an entity currently possesses
a valid authorization within an ecosystem. It reflects the definitive condition
of the entity’s rights or credentials at a given point in time.

 [[def:Authority Statement, Authority Statements]]
~ An authority statement is an assertion by an **authority** about either:
a) the **authorization** or **delegation** it grants to another party over which
it has authority, or b) the **recognition** it gives to a peer about the scope of
that peer’s authority.

 [[def:Digital Trust Ecosystem, Digital Trust Ecosystems]]
~ A [digital ecosystem](https://glossary.trustoverip.org/#term:digital-ecosystem)
in which the participants are one or more interoperating
[trust communities](https://glossary.trustoverip.org/#term:trust-communities).
Governance of the various
[roles](https://glossary.trustoverip.org/#term:roles) of
[governed parties](https://glossary.trustoverip.org/#term:governed-parties)
within a digital trust ecosystem (e.g.,
[issuers](https://glossary.trustoverip.org/#term:issuers),
[holders](https://glossary.trustoverip.org/#term:holders),
[verifiers](https://glossary.trustoverip.org/#term:verifiers),
[certification bodies](https://glossary.trustoverip.org/#term:certification-bodies),
[auditors](https://glossary.trustoverip.org/#term:auditors)) is typically managed
by a [governing body](https://glossary.trustoverip.org/#term:governing-body)
using a [governance framework](https://glossary.trustoverip.org/#term:governance-framework)
as recommended in the [ToIP Governance Stack](https://glossary.trustoverip.org/#term:toip-governance-stack).
Many digital trust ecosystems will also maintain one or more
[trust lists](https://glossary.trustoverip.org/#term:trust-lists) and/or
[trust registries](https://glossary.trustoverip.org/#term:trust-registries).

 [[def:Ecosystem Governance Framework, Ecosystem Governance Frameworks]]
~ A [governance framework](https://glossary.trustoverip.org/#term:governance-framework)
for a [digital trust ecosystem](https://glossary.trustoverip.org/#term:digital-trust-ecosystem).
An ecosystem governance framework may incorporate, aggregate, or reference other
types of governance frameworks such as a
[credential governance framework](https://glossary.trustoverip.org/#term:credential-governance-framework)
or a
[utility governance framework](https://glossary.trustoverip.org/#term:utility-governance-framework).

 [[def:Inter-Ecosystem Trust, Inter-Ecosystem Trusts]]
~ The confidence and assurance established between two or more distinct
ecosystems or governance frameworks. This type of trust enables cross-recognition
of rules, trust registries, and authorization states among separate ecosystems.

 [[def:Intra-Ecosystem Trust Framework, Intra-Ecosystem Trust Frameworks]]
~ The confidence and assurance maintained within a single ecosystem or governance
framework. It applies to entities operating under the same set of rules, trust
registries, and oversight mechanisms.

 [[def:Hierarchical Authority Relationship, Hierarchical Authority Relationships]]
~ A relationship between an authority and another party subject to that authority.
This relationship is unilateral and exclusive, meaning that the authority is the
only one in control, i.e., the only one who can grant authorization to or revoke
authorization from the authorized party.

 [[def:Recognition Relationship, Recognition Relationships]]
~ A heterarchical authority relationship between two authorities who are peers,
each authoritative for their own ecosystems. The relationship between these two
authorities can be either unilateral or bilateral and it is non-exclusive. One
authority is attesting to the other’s authority in one direction or both
directions.

 [[def:Recognition Query, Recognition Queries]]
~ A request to a network that enables a verifier to check the Recognition
Relationship [[ref:Recognition Relationship]] of an ecosystem in relation to
another ecosystem.

 [[def:TRQP Core]]
~ A foundational specification that defines the core data models, queries, and
security requirements necessary for consistent, interoperable trust interactions
across different systems and ecosystems. It sets the universal “language” for
TRQP implementations, ensuring that all compliant solutions share a common
framework for exchanging and validating trust information.

 [[def:TRQP Binding, TRQP Bindings]]
~ A technical specification document outlining the precise requirements for
implementing interoperability via the base TRQP interfaces and data models. It
dictates how systems should interact and exchange trust information to remain
compliant with TRQP standards.

 [[def:TRQP Bridge, TRQP Bridges]]
~ A software or infrastructure component that connects a System of Record
 [[ref:System of Record]] to a specified TRQP Binding [[ref:TRQP Binding]],
enabling seamless data exchange and interoperability. It serves as the interface
through which trust and authority data move between the system and the
TRQP-compliant environment.

 [[def:System of Record, Systems of Record]]
~ An authoritative source that manages and maintains authority and recognition
statuses for participants within an ecosystem. It is responsible for preserving
the integrity and continuity of records, including authorizations and trust
credentials.

## **Scope**

This specification is primarily focused on defining the
**[[ref:TRQP Core]]** framework. While we recognize the importance of
**[[ref:TRQP Binding]]** and **[[ref:TRQP Bridge]]** in enabling
interoperability, their specific implementations are left to the discretion of
ecosystems and implementers. **[[ref:TRQP Binding]]** extend the core
abstractions to concrete implementations, while **[[ref:TRQP Bridge]]**
connect ecosystems to their **[[ref:System of Record]]**.
However, this specification does not prescribe how they should be designed or
deployed, allowing flexibility for diverse use cases and ecosystem needs.

![images/trqp_layers.png](images/trqp_layers.png)

Fig 1: This specification is focused specifically on addressing the core
requirements for a binding specification to be TRQP compliant. It is up to
ecosystems to build their own bindings.

## **Problem Statement**

Modern digital ecosystems often rely on **[[ref:Intra-Ecosystem Trust Framework]]** (e.g., OpenID Federation, X.509 Chains, EBSI
Trust Chains, TRAIN) to manage authorization within their own boundaries. While
these frameworks are effective at validating whether an entity is authorized
**within** a single ecosystem, they do not easily extend to other ecosystems. As
a result, organizations face considerable challenges when attempting to verify
authorizations **across** different frameworks. The core issues include:

* **Siloed Trust Frameworks**: Each ecosystem typically operates in isolation,
  lacking a common or standardized method to verify whether a subject in one
  ecosystem has valid authorization recognized by another.
* **Inconsistent Interfaces**: Every ecosystem defines its own APIs, credential
  formats, and governance rules, which forces implementers to work with multiple
  disparate interfaces and data models.

When attempting to establish **[[ref:Inter-Ecosystem Trust Framework]]**,
verifiers face two fundamental questions:

1. **Ecosystem Recognition**:

   *“Do I recognize the governance framework [[ref:Ecosystem Governance Framework]]
   of the other ecosystem?”*

   This question, also called “ecosystem recognition,” is inherently complex and
   ultimately depends on human policy decisions. It is not easily automated.

2. **Entity Authorization**:

   *“Is the issuer authorized to issue this type of data under the ecosystem’s
   governance framework [[ref:Ecosystem Governance Framework]]?”*

   This question takes the form of an **[[ref:Authority Query]]**,
   which is a formal or protocol-based request to confirm whether a given entity
   (subject) holds a specific authorization, credential, or right within the
   ecosystem.

![images/authority_questions.png](images/authority_questions.png)

Figure 2: The two fundamental queries required cross-ecosystem authority
verification.

TRQP tackles the [[ref:Inter-Ecosystem Trust Framework]] problem
by allowing verifiers outside an ecosystem to request an
**[[ref:Authority Query]]** and **[[ref:Recognition Query]]** to any TRQP compliant network. The specification
must work independently of any particular Systems of Record [[ref:System of Record]]
and intra-trust frameworks, ensuring trust can be established across different
ecosystems without uplifting a current authority system.

## **High-Level Architecture**

The **TRQP** architecture is designed to enable standardized cross-ecosystem
queries regarding trust registry information, authorization, and recognition. At
its heart, TRQP comprises:

1. An **abstract specification** (the *Core*) defining data models, query flows,
   and security considerations. (i.e. **[[ref:TRQP Core]]**)
2. One or more **concrete bindings [[ref:TRQP Binding]]** that map the abstract
   specification to specific transport protocols (e.g., HTTPS, DIDComm, TSP).
3. **[[ref:TRQP Bridge]]** that connect TRQP queries to particular trust
   frameworks (OIDF Federation, x.509, etc.).
4. **[[ref:System of Record]]**—the actual trust frameworks or
   registries (e.g., x.509 Ecosystem, OIDF Federation) responsible for issuing
   or validating trust information.

This layered approach allows implementers to select or build only what they
need. If a trust framework has not implemented TRQP, integrators can connect a
new **[[ref:TRQP Bridge]]** to a **[[ref:TRQP Binding]]** for it,
as long as they follow the core specification and a compatible binding. The spec
focuses on the abstract specification layer *(i.e Core)* and will not go into
detail on anything lower in the stack.

![images/system_boundaries.png](images/system_boundaries.png)

Fig 3: TRQP Architecture has three layers: Core, Bindings, and Bridges. Profiles
can be built on top to enable networks.

Details of *Bridges [[ref:TRQP Bridge]]*, *Systems of Record [[ref:System of Record]]*,
and *Bindings [[ref:TRQP Binding]]* are out of scope for this specification, but
defined in concept for other specifications to describe in detail.

We will briefly go through each of the layers at a high level, and expand on the
5.1 **[[ref:TRQP Core]]** layer in more detail in the subsequent
section.

### **TRQP Core**

* **What it is**: The TRQP Core [[ref:TRQP Core]] is an **abstract** specification
  that defines:
  * **Data Models**: Metadata, authorization, ecosystem recognition, etc.
  * **Required Queries**: MetadataQuery, AuthorizationQuery, and
    EcosystemRecognitionQuery.
* **Role**: It ensures every TRQP-based implementation speaks the same
  “language” (even if actual messages go over different transports).

### **TRQP Bindings**

* **What they are**: Concrete mappings of the Core specification onto specific
  transports and protocols. For example:
  * **HTTPS Binding**: Illustrates how to send TRQP queries over HTTPS.
* **Role**: A TRQP Binding [[ref:TRQP Binding]] ensures that an abstract query
  from the Core spec is transformed into real network requests and responses in
  a standardized way.

### **TRQP Bridges**

* **What they are**: Adapters or connectors that apply a chosen
  Binding [[ref:TRQP Binding]] to a specific trust framework (x.509, OIDF, DIF
  CTE, etc.).
* **Examples**:
  * **x.509 Bridge**: Translates TRQP queries into x.509 certificate validations
    and chain checks.
  * **OIDF Bridge**: Leverages OpenID Federation endpoints to answer TRQP queries
    about OIDC-based trust relationships.
  * **CTE Bridge**: Adapts TRQP queries to DIF’s Credential Trust Establishment
    protocols.
* **Role**: A TRQP Bridge [[ref:TRQP Bridge]] “bridges” existing frameworks into
  TRQP by implementing the relevant Binding [[ref:TRQP Binding]] and mapping
  framework-specific data.

### **Systems of Record**

* **Definition**: Real-world trust frameworks or registries storing authoritative
  data. Examples include:
  * **OIDF Federation** (Profiles 1, N)
  * **x.509 Ecosystem** (with a CA and certificate hierarchy)
  * **TRAIN** (some trust registry or network)
  * **EU Trusted List** (an EU-level trust list or EBSI-based registry)
* **Role**: The ultimate source of truth for whether an entity is recognized,
  authorized, or otherwise valid within a particular ecosystem. (i.e.
  **[[ref:System of Record]]**)

### **TRQP Profiles**

* **Definition:** TRQP Profiles specify the implementation details necessary for
  aligning a trust network with TRQP standards.
* **Examples:**
  * **Identifier Design:** How entities are uniquely identified within the
    system.
  * **Resolution Paths:** The process for resolving trust queries within a given
    framework.
* **Role:** Profiles guide the adaptation of TRQP to different ecosystems,
  ensuring that queries, identifiers, and resolution mechanisms conform to
  standardized practices.

## **Abstract Metadata Models**

![images/ecosystem_model.png](images/ecosystem_model.png)

Fig 4: A trust registry may serve an ecosystem Authority Statement. A trust registry may serve multiple ecosystems. An ecosystem may have multiple trust registries.

### **Interpretation of the Diagram: Ecosystem and Trust Registry Relationship**

Each Ecosystem is represented by a yellow box in the diagram and consists of:

* An identifier (green box) – A globally unique reference that distinguishes the
  ecosystem.
* An EGF Document (green box) – The Ecosystem Governance Framework
   [[ref:Ecosystem Governance Framework]] (EGF), which defines governance terms,
  policies, and operational rules for the ecosystem.

The dashed arrows from these Ecosystems point to one or more Trust Registries
below, indicating that:

* Each Ecosystem explicitly references the Trust Registry(ies) it recognizes for
  managing and verifying authority-related queries [[ref:Authority Query]].
* Trust Registries are designated by the Ecosystem metadata and are responsible
  for enforcing the rules and policies outlined in the EGF document.

Each Trust Registry (represented as a green box in the lower row) consists of:

* An identifier (yellow box) – A unique reference for the registry.
* References to one or more Ecosystems – If supported by the metadata, a Trust
  Registry may explicitly list the Ecosystem(s) it serves.

### **Role of the Trust Registry in Ecosystem Governance**

A Trust Registry manages and serves authority statements [[ref:Authority Statement]]
across one or more Ecosystems by:

* Maintaining a structured record of trust relationships – storing authoritative
  data on recognized entities and their authorization statuses.
* Handling authority queries [[ref:Authority Query]] (as described in Section 8)
  – providing verified responses regarding entity recognition and authorization.
* Operating under the governance of the ecosystem – with the governing body
  defining the policies and processes for registering entries into the ecosystem.

### **Scalability and Multi-Ecosystem Trust Registries**

* A single Trust Registry may serve multiple Ecosystems, acting as a shared
  infrastructure for trust and authorization across different governance models.
* An Ecosystem may rely on multiple Trust Registries to provide redundancy,
  distribute authority management, or allow for diverse verification approaches.
* The Trust Registry is not independent but operates within the authority scope
  defined by the EGF Document of each ecosystem it serves.

### **Trust Registry**

* **Properties**
  * **id: MUST** be a globally unique identifier for the registry (e.g., URI,
    DID, UUID).
  * **ecosystem: SHOULD** indicate which ecosystem(s) the registry serves or
    recognizes.
  * **controller: SHOULD** reference the entity (individual, organization, or
    automated system) that manages or operates the registry.

### **Ecosystem**

* **Properties**
  * **id: MUST** be a globally unique identifier for the registry (e.g., URI,
    DID, UUID).
  * **egf_id: MUST** specify a *resolvable* EGF identifier referencing the
    official EGF document or descriptor.
  * **trustregistries: MUST** provide a list of authorized Trust Registries that
    serve the ecosystem authority state [[ref:Authority State]].
    * Each registry **MUST** have the following properties:
      * **endpoint**: The address (URL, DID, etc.) for TRQP
        queries [[ref:Authority Query]] /  [[ref:Recognition Query]].
    * Each registry **MAY** also be scoped to a particular set of authorization
      states and is defined in the Binding [[ref:TRQP Binding]].
  * **controller: SHOULD** include a method of validating controllers of an
    ecosystem.

## **Baseline Requirements For Conformance**

### **Trust Registry**

* All TRQP registriesqueries **MUST** provide an addressable endpoint that can be
  resolvable as defined by the Implementation Profile.
* All Trust Registries **MUST** supply the required interfaces described in
  Section 8 over the *same* addressable endpoint to be TRQP conformant.

### **TRQP Binding**

* All compliant binding [[ref:TRQP Binding]] **MUST** support the required
  interfaces described in Section 8.
* A compliant binding [[ref:TRQP Binding]] **MUST** be compliant with
  TRQP Core [[ref:TRQP Core]] requirements.
* A compliant binding [[ref:TRQP Binding]] **MUST** support versioning using
  [Semantic Versioning 2.0](https://semver.org/)

### **TRQP Profiles**

* All TRQP profiles **MUST** specify a compliant binding [[ref:TRQP Binding]].

## **Required Interfaces**

Below are abstract API methods that **MUST** be exposed;
**[[ref:TRQP Binding]]** **MUST** define a binding (e.g., REST, gRPC,
DIDComm) that maps these methods to actual endpoints.

```mermaid
sequenceDiagram
    participant C as Client
    participant R as Trust Registry

    C->>R: 9.1 Metadata Query<br/>(registryIdentifier)
    R-->>C: Metadata Response<br/>(auth types, credential formats, version info)<br/>MUST provide details for further queries

    C->>R: 9.2 Authorization Query<br/>(subject, authorizationType, context?)
    R-->>C: Authorization Response<br/>(authorized / not-authorized / revoked / error)<br/>MUST indicate authorization status clearly

    C->>R: 9.3 Ecosystem Recognition Query<br/>(ecosystemIdentifier, governanceFrameworkRef)
    R-->>C: Recognition Response<br/>(yes/no + reasons for rejection)<br/>MUST indicate acceptance or rejection
```
Below is a rewritten version with additional context and sample values:

---

### **Metadata Query**

* **Request**:  
  There are no mandatory request parameters.  
  * Optionally, an `ecosystem_id` can be supplied to indicate that the metadata request should be interpreted within the context of a specific ecosystem’s governance framework (see [[ref:Ecosystem Governance Framework]]).

* **Response**:  
  * `id`: string. This value uniquely identifies the registry. If an `ecosystem_id` is provided, the response should clearly reflect that the returned data is scoped to the specified ecosystem (e.g., "ecosystem A").

---
### **Authorization Query**

* **Request**:
  * **ecosystem_id**: string. This must be an ecosystem identifier as defined in the TRQP Binding.  
    - *Example*: `"ecosystem_id": "ecosystem A"`
  * **authorization_id**: string. This must match one of the defined authorization types in the TRQP Binding.  
    - *Example*: `"authorization_id": "credential-A-issuer"`
  * **entity_id**: string. This identifier specifies the entity for which the authorization is being queried.  
    - *Example*: `"entity_id": "random-id-1234"`
  * **time**: string (optional). May be provided according to the TRQP Binding guidelines and describes the time at which the Trust Registry should evaluate the authority query.
    * If supplied, the `time` value must adhere to the required time format (e.g., RFC3319 UTC).
    * If omitted, the system must use the current time and include that timestamp in its response.

* **Response**:  
  A status indicating the entity's authorization, such as:
  - **authorized**
  - **not-authorized**
  - **revoked**
  - **unknown-subject**
  - **error**  
  Optionally, the response may include additional details on validity or supporting proof references.

* **Behavior**:  
  The system **MUST** clearly indicate whether the subject holds the specified authorization.

---

### **Ecosystem Recognition Query**

* **Request**:
  * **ecosystem_id**: string. This is the identifier for the ecosystem, defined in the TRQP Binding.  
    - *Example*: `"ecosystem_id": "ecosystem A"`
  * **target_ecosystem_id**: string (optional). This may be another ecosystem identifier against which recognition is being evaluated.
  * **scope**: string (optional). This parameter may be used to filter or narrow the request. The specification does not enforce a specific structure for scopes, but individual profiles may define their own conventions.
  * **time**: string (optional). May be provided as described in the TRQP Binding guidelines.

* **Response**:  
  The response indicates the recognition status of the ecosystem, for example:
  - **accepted** (if the ecosystem is recognized)
  - **rejected** (if it is not)  
  Additional supporting details, such as proof references or log entries, may also be included.

* **Behavior**:  
  The system **MUST** return a clear yes/no answer regarding ecosystem recognition, and it **MAY** provide further explanation or details as specified in the TRQP Binding.

---
## **Security Considerations**

All implementers (“bindings [[ref:TRQP Binding]]” and “bridges
 [[ref:TRQP Bridge]]”) of TRQP **SHOULD** take the following threats into
account and implement appropriate controls:

* **Trust Anchor Hijacking**: Use strong cryptography and rotate keys regularly.
* **Trust Registry Bugs**: Conduct code reviews, vulnerability scans, and robust
  QA.
* **Trust Anchor Spoofing**: Verify responses using known cryptographic anchors
  or certificate chains.
* **Domain Hijacking**: Protect DNS entries; if DNS-based discovery is used,
  consider DNSSEC or other verification.
* **Replay Attacks**: Use timestamps, nonces, and short-lived tokens.
* **Data Integrity**: Sign or hash data at rest, use TLS or equivalent in
  transit.
* **Denial of Service (DoS)**: Rate-limit queries, monitor usage, scale
  infrastructure appropriately.
* **Insufficient Data Validation**: Enforce strict schema checks and reject
  malformed data with clear error messages.
* **Trust Anchor Compromise**: Implement multi-tier trust anchors, have a plan
  to revoke or replace compromised keys quickly.
* **Logging and Auditing**: Log all access, changes, and suspicious activities;
  adopt real-time monitoring.
* **Protocol Downgrade Attacks**: Default to the latest secure version, disallow
  fallback to insecure versions.
* **Privacy Concerns**: Encrypt sensitive or personally identifiable
  information, comply with relevant data protection laws.
* **Timing Attacks**: Where feasible, adopt constant-time operations for
  cryptographic and authorization checks.

## **Conclusion**

The TRQP enables a consistent, abstract protocol for verifying cross-ecosystem
authorizations and ecosystem recognition. By adhering to these requirements—
data models, flows, security measures—implementers provide a standardized bridge
across diverse trust frameworks.

* **MUST** statements are mandatory for compliance.
* **SHOULD** statements are recommended best practices for enhanced
  interoperability and security.
* **MAY** statements indicate optional features or extensions.

When combined with ecosystem-specific bindings [[ref:TRQP Binding]] and
governance framework [[ref:Ecosystem Governance Framework]] definitions, the TRQP
can dramatically streamline multi-ecosystem [[ref:Inter-Ecosystem Trust Framework]]
integrations, ensuring reliable and secure authorization checks across
organizational and technological boundaries.
