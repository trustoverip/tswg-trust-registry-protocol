
[//]: # (Pandoc Formatting Macros)

[//]: # (# Requirements)

[//]: # (:::)

## Requirements

### Registry Queries [RQ-*]

The following queries relate to receiving answers related to entities and other trust registries.

* [RQ-1] MUST support query operations for the current status of a **registered entity**.
* [RQ-2] MUST support querying about ---TODO: 

### Configuration Queries [CQ-*]

The following queries relate to configuration of systems that will interact with the trust registry.


* [CQ-1] MUST provide a list of [[ref: action namespace]] that are supported by the responding system.
* [CQ-2] MUST provide 
* [CQ-2] MUST provide a list of [[ref: VID Type]] (i.e. VID Types) that are supported by the responding system.
* [CQ-3] MUST provide a list of  [[ref: assurance level]] that are supported by the responding system.
* 


### Metadata Queries [MQ-*]

* [MQ-1] MUST provide a list of [[ref: ecosystem governance frameworks]] (EGFs) that the system is operating under. This data will be comprised of the following elements:
  * [MQ-1-1] MUST provide the VID of the EGF.
  * [MQ-1-2] MAY provide the name of the EGF.
1. [MQ-2] SHOULD provide the legal name and jurisdiction of the **governing authority** for the **trust registry** service.
2. [MQ-3] SHOULD provide the legal name and jurisdiction of the **administering authority** for the **trust registry** operator (if different from **governing authority**).
3. [MQ-4] SHOULD provide a textual description of the trust registry mandate.



### Governing Authorities [GA-*]

**Governing authorities** compliant with this specification:

* [GA-1] MUST have exactly one [[ref: primary trust registry]].
* [GA-2] MAY have one or more [[ref: secondary trust registries]].

> The [[ref:primary trust registry]] plus all [[ref: secondary trust registries]] are collectively referred to as the [[ref:authorized trust registries]].

* [GA-3] MUST publish an [[ref: EGF]] that meets the **requirements** of:
  * [GA-3-1] This specification.
  * [GA-3-2] The [ToIP Governance Architecture Specification](https://wiki.trustoverip.org/pages/viewpage.action?pageId=71241). Note that this includes the requirement that the **EGF** and all **governed parties** (which includes **authorized issuers** and **authorized verifiers**) must be identified with a **DID**.

TODO: Add normative ref to [ToIP Governance Architecture Specification](https://wiki.trustoverip.org/pages/viewpage.action?pageId=71241)

* [GA-4] MUST publish, in the **DID document** associated with the **DID** identifying its **EGF**, a **service property **specifying the **service endpoint** for its **primary trust registry** that meets the **requirements** in the _[Trust Registry Service Property](#trust-registry-service-property)_ section.
[GA-5] MUST publish in its **EGF** a list of any other EGFs governing **secondary trust registries.**
[GA-6] MUST specify in the EGF any additional **requirements** for an **authorized trust registry**. This data will be comprised of the following elements::

    * [GA-6-1] SHOULD provide **Information trust requirements**.
    * [GA-6-2] SHOULD provide Technical **requirements**.
    * [GA-6-3] SHOULD provide Operational **requirements**.
    * [GA-6-4] MAY provide Legal contracts.
* [GA-7] MUST specify in its **EGF** (or in any referenced documents) **requirements** for:
    - [GA-7-1] MUST provide all [[ref: authorization]] values that are used by the trust registry.
    - [GA-7-2] MUST provide all [[ref: assurance levels]], specified with unique names, that are service by the trust registry, and what [[ref: authorization]] values they apply to.
    - [GA-7-3] MUST provide a list of all [[ref: VID Types]] that are supported by the ecosystem, and serviced by the trust registry.
    - [GA-7-4] SHOULD provide `resources (TODO: TERM IS VAGUE)` that are required by systems integrating into the ecosystem that the system serves. 
    - [GA-7-5] `???any metadata required by implementors (e.g. claim name that is mandatory if pointing a credential back to an EGF.) [this is a weak example]???`
    - [GA-7-6] `???a statement about the basis the trust registry claims to be authoritative???`
    - [GA-7-7] `???means by which others are able to verify the asserted authority???`
* [GA-8] SHOULD specify in the **EGF** the following **requirements** for an **authorized trust registry** and any **registered party** (i.e., issuer, verifier, or peer trust registry):
    - [GA-8-1] The **requirements** to become authorized.
    - [GA-8-2] How to request registration.
    - [GA-8-3] The **requirements** for assignment of each **authorization** for a **registry entry**.
    - [GA-8-4] Any access limitations (e.g. unrestricted public access, authentication-limited access).
    - [GA-8-5] How to request access where unrestricted public access is not available.


### Trust Registry Service Property [TRSP-*] 

The **DID document** for the **DID** that identifies an **EGF** compliant with this specification MUST include a service property that meets the **requirements** [in section 5.4 of the W3C Decentralized Identifiers (DIDs) 1.0 specification](https://www.w3.org/TR/did-core/#services) plus the following additional **requirements**:

* The value of the `type` property MUST be `TrustRegistry`.
* The value of the `serviceEndpoint` property MUST be exactly one HTTPS URI.


[`TODO:` reconcile above with Profiles concept. ]

[`TODO:` The issuer/verifier needs to state their primary trust registry affiliation (a trust relationship) - is this a new section?]

### Trust Registry Protocol [TRP-*]

The authoritative technical specifications for the API calls in the ToIP Trust Registry Protocol V1 are specified in Appendix A (OpenAPI YAML file). This section contains a textual description of the **requirements**.

**Trust registries** implementing this protocol:

* [TRP-1] MUST maintain the service implementing this protocol at the HTTPS URI specified in the _[Trust Registry Service Property](#trust-registry-service-property)_ section.
* [TRP-2] MUST return responses to queries for the **status value** of a **registry entry** that satisfies one or more of the following sets of query parameters:

* Entity
* Entity Authorization
* Registry


    - i. **Entity Authorization**: entityDID, authorization
    - ii. **Recognized Registry:** entityDID
3. MUST return responses using the data model specified in the _[Data Model](#data-model)_ section.
4. MUST return exactly one of the following **status values** for a **registry entry** satisfying the query parameters:
    - i. `Not found` (http 404)
    - ii. `Current`
    - iii. `Expired` (not renewed after the previous valid registration period)
    - iv. `Terminated` (voluntary termination by the **registered party**)
    - v. `Revoked` (involuntary termination by the **governing authority**)
5. For queries returning a **status value** other than `Not Found`, the response MUST return the following values:
    - i. The parameter values exactly as supplied in the query (so responses can be stateless).
   - ii. The **status value**.
   - iii. Exactly two **datetime values** conforming to the following requirements:
        - a. The value labels MUST be:
            - i. `AuthorizationStartDate`
            - ii. `AuthorizationEndDate`
        - b. The values MUST be formatted to comply with [RFC 3339](https://tools.ietf.org/html/rfc3339) in the UTC/Z time zone with no offset.
        - c. The `AuthorizationStartDate` MUST be the date that the **registered party’s** authorization began.
        - d. The `AuthorizationEndDate` MUST be either:
            - i. `Null` for an entry whose **status value** is `Current` at the time of the query.
            - ii. A specific date value if the **registered party’s** **status value** is `Expired`, `Terminated` or `Revoked.`
        - e. If a **registered party** has multiple entries (representing an authorization history), the most recent value MUST be returned. 

### Anti-Requirements

[AR-1] SHALL NOT support query operations for the history of a [[ref: registered entity]].   
        
[AR-2] SHALL NOT include support for a DIDComm interface, only a RESTful (i.e. OpenAPI Specification) interface. When a repeatable **trust task** specification approach is created, a DIDComm/**trust task** approach should be considered as a work effort.

[AR-3]]SHALL NOT support automated **rules** processing.

[AR-4] Anyting other than read-only INSERT, UPDATE and DELETE operations. The TRP is a read-only (RETRIEVE in the CRUD sense) protocol.


