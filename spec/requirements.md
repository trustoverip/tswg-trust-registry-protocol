
[//]: # (Pandoc Formatting Macros)

[//]: # (# Requirements)

[//]: # (:::)

## Requirements

### Registry Queries [RQ-*]

The following queries relate to receiving answers related to entities and other trust registries.

* [RQ-1] The system MUST support query operations for the current status of a [[ref: registered entity]].
* [RQ-2] The system SHOULD support query operations for a list of related [[xref: TOIP, trust registries]].

### Configuration Queries [CQ-*]

The following queries relate to configuration of systems that will interact with the trust registry.

* [CQ-1] MUST provide a list of [[ref: authorization]] namespaces that are supported by the responding system.
* [CQ-2] MUST provide list of additional [[xref: TOIP, EGFs]] that the trust registry operates under.
* [CQ-2] MUST provide a list of [[ref: VID Type]] (i.e. VID Types) that are supported by the responding system.
* [CQ-3] MUST provide a list of  [[xref: TOIP, assurance levels]] that are supported by the responding system.

### Metadata Queries [MQ-*]

* [MQ-1] MUST provide a list of [[xref: TOIP, ecosystem governance frameworks]] (EGFs) that the system is operating under. This data will be comprised of the following elements:
  * [MQ-1-1] MUST provide the VID of the EGF.
  * [MQ-1-2] MAY provide the name of the EGF.
1. [MQ-2] SHOULD provide the legal name and jurisdiction of the [[xref: TOIP, governing authority]] for the [[ref: trust registry]] service.
2. [MQ-3] SHOULD provide the legal name and jurisdiction of the [[xref: TOIP, administering authority]] for the [[xref:TOIP, trust registry]] operator (if different from [[xref: TOIP, governing body]]).
3. [MQ-4] SHOULD provide a textual description of the trust registry mandate.

### Governing Authorities [GA-*]

**Governing authorities** compliant with this specification:

* [GA-1] MUST have exactly one [[ref: primary trust registry]].
* [GA-2] MAY have one or more [[ref: secondary trust registries]].

> The [[ref:primary trust registry]] plus all [[ref: secondary trust registries]] are collectively referred to as the [[ref:authorized trust registries]].

* [GA-3] MUST publish an [[xref: TOIP, EGF]] that meets the requirements of:
  * [GA-3-1] This specification.
  * [GA-3-2] The [ToIP Governance Architecture Specification](https://wiki.trustoverip.org/pages/viewpage.action?pageId=71241). Note that this includes the requirement that the [[xref: TOIP, EGF]] and all [[xref: TOIP, governed parties]] must be identified with a [[xref: TOIP, DID]].

::: todo 
Add normative ref to [ToIP Governance Architecture Specification](https://wiki.trustoverip.org/pages/viewpage.action?pageId=71241)
:::

* [GA-4] MUST publish, in the [[xref: TOIP, DID document]] associated with the **DID** identifying its **EGF**, a [[ref: service property]] specifying the [[ref: service endpoint]] for its [[ref: primary trust registry]] that meets the requirements in the _[Trust Registry Service Property](#trust-registry-service-property)_ section.

* [GA-5] MUST publish in its EGF a list of any other EGFs governing [[ref: secondary trust registries]].

* [GA-6] MUST specify in the EGF any additional requirements for an [[ref: authorized trust registry]]. This data will be comprised of the following elements::

    * [GA-6-1] SHOULD provide Information Trust requirements.
    * [GA-6-2] SHOULD provide Technical requirements.
    * [GA-6-3] SHOULD provide Operational requirements.
    * [GA-6-4] MAY provide Legal contracts.
* [GA-7] MUST specify in its **EGF** (or in any referenced documents) requirements for:
    - [GA-7-1] MUST provide all [[ref: authorization]] values that are used by the trust registry.
    - [GA-7-2] MUST provide all [[xref: TOIP, assurance levels]], specified with unique names, that are service by the trust registry, and what [[ref: authorization]] values they apply to.
    - [GA-7-3] MUST provide a list of all [[ref: VID Types]] that are supported by the ecosystem, and serviced by the trust registry.
    - [GA-7-4] SHOULD provide `resources` (e.g. logo files, documents, interoperability profile information) that are required by systems integrating into the ecosystem that the system serves. 
    - [GA-7-5] `???any metadata required by implementors (e.g. claim name that is mandatory if pointing a credential back to an EGF.) [this is a weak example]???`
    - [GA-7-6] `???a statement about the basis the trust registry claims to be authoritative???`
    - [GA-7-7] `???means by which others are able to verify the asserted authority???`
* [GA-8] SHOULD specify in the **EGF** the following requirements for an **authorized trust registry** and any **registered party** (i.e., issuer, verifier, or peer trust registry):
    - [GA-8-1] The **requirements** to become authorized.
    - [GA-8-2] How to request registration.
    - [GA-8-3] The **requirements** for assignment of each **authorization** for a **registry entry**.
    - [GA-8-4] Any access limitations (e.g. unrestricted public access, authentication-limited access).
    - [GA-8-5] How to request access where unrestricted public access is not available.


### Trust Registry Service Property [TRSP-*] 

The [[xref: TOIP, DID document]] for the **DID** that identifies an **EGF** compliant with this specification MUST include a service property that meets the **requirements** in section 5.4 of [[spec-norm:DID-CORE]] plus the following additional **requirements**:

* [TRSP-1] The value of the `type` property MUST be `TrustRegistry`.
* [TRSP-2] The value of the `serviceEndpoint` property MUST be exactly one HTTPS URI.


::: todo
FIX
:::

[[ref: Registered entities]] MUST indicate which registries they are part of. 
* [TRSP-3] Registered entities MUST indicate the [[ref: primary trust registry]]] for a particular [[ref: authorization]].



#### Service Profile Recommendation

_the following recommendation is non-normative_


It is recommended that the service leverage the [Service Profile
Specification](https://github.com/trustoverip/tswg-trust-registry-service-profile/blob/main/spec.md).
Trust Over IP hosts a [Service Profile]() with the following pointer: 

```json
{
  "integrity": "<>",
  "profile": "<>"
  "uri": "<your service endpoint uri here>"
}
```

By implementing service profiles, it enables easier interoperability and
discovery of service capabilities for the trust registry being implemented.

### Trust Registry Query Protocol [TRQP-*]

The authoritative technical specifications for the API calls in the ToIP Trust Registry Query Protocol are specified in Appendix A (OpenAPI YAML file). This section contains a textual description of the **requirements**.

**Trust registries** implementing this protocol:

* [TRQP-1] MUST maintain the service implementing this protocol at the HTTPS URI specified in the _[Trust Registry Service Property](#trust-registry-service-property)_ section.

* [TRQP-2] The system SHOULD support queries that are at a point in time in the past. 
  - [TRQP-2-1] The parameter for the point in time must be named `queryTime`.
  - [TRQP-2-2] The datetime value provided MUST be formatted per [[spec-norm:RFC3339]] using the UTC (i.e. Z for Zulu) zero offset (e.g. "2018-03-20T09:12:28Z". 
  - [TRQP-2-3] If the system does not support non-current data, and the the `queryTime` parameter is present, the system MUST NOT return entity data and must se http error code 405 (Method not allowed).
  
* [TRQP-3] MUST return responses to queries for the **status value** of a **registry entry** that satisfies one or more of the following sets of query parameters:
    - [TRQP-3-1] **Entity Authorization**: Given the `entityDID`, and `authorization` return the status of that registered entity, MUST return exactly one of the following **status values** for a **registry entry** satisfying the query parameters:
      - `Not Found` + http code 404 - entry not found.
      - `Current` + http code 200 - authorization for the registered entity is current as of the time of query, or as of the time requested.
      - `Expired` + http code 200 - authorization has expired (e.g. not renewed after the previous valid registration period)
      - `Terminated` + http code 200 - authorization was terminated (e.g. voluntary termination by the **registered entity**)
      - `Revoked` + http code 200 - authorization was revoked (e.g. involuntary termination by the **governing authority**) 
    - [TRQP-3-2] **Entity Authorizations**: Given only the `entityDID` the system SHOULD return the array of Authorization objects for the entity identified by `entityDID`. 
    - [TRQP-3-3] **Recognized Registry:** Given the entityDID the system SHOULD return the list of [[def:trust registries]] that the entity has indicated it is registered in. 
      - [TRQP-3-3-1] The system MUST NOT return more than one trust registry in the array designated as a [[def: primary registry]].

::: TODO: 
  Align VID and/or DID terminology.
:::

* [TRQP-4] MUST return responses using the data model specified in the OpenAPI Specification . 

* [TRQP-5] For queries returning a **status value** other than `Not Found`, the response MUST return the following values:
  - [TRQP-5-1] The system must return the parameter values exactly as supplied in the query (so responses can be stateless).
  - [TRQP-5-2] The system must return the **status value** for the entity (per TRP-3-1).
  - [TRQP-5-3] The system must return exactly two **datetime values** conforming to the following requirements:
    - [TRQP-5-3-1]The value labels MUST be:
      - i. `AuthorizationStartDate`
      - ii. `AuthorizationEndDate`
    - [TRQP-5-3-2] The datetime values MUST be formatted to comply with [[spec-norm:RFC3339]] in the UTC/Z time zone with no offset.
    - [TRQP-5-3-3] The `AuthorizationStartDate` MUST be the date that the **registered entity** authorization began.
    - [TRQP-5-3-4] The `AuthorizationEndDate` MUST be either:
      - [TRQP-5-3-4-1] `Null` for an entry whose **status value** is `Current` at the time of the query.
      - [TRQP-5-3-4-2] A specific datetime value if the **registered entity** **status value** is `Expired`, `Terminated` or `Revoked`.
    - [TRQP-5-3-5] If a **registered entity** has multiple entries in the system (representing an authorization history), the value that is active at the time indicated must be returned:
      - [TRQP-5-3-5-1] when no `queryTime` value is provided the value that is active at time of the query MUST be returned.
      - [TRQP-5-3-5-2] when a `queryTime` parameter is provided the entry that is active at that time (i.e. indicted by `queryTime`) MUST be returned. 

### Anti-Requirements

The following are considered anti-requirements in that they have been considered in the current design of the TRQP:

* [AR-1] SHALL NOT support query operations for the history of a [[ref: registered entity]].   
        
* [AR-2] SHALL NOT include support for a DIDComm interface, only a RESTful (i.e. OpenAPI Specification) interface. When a repeatable **trust task** specification approach is created, a DIDComm/**trust task** approach should be considered as a work effort.

* [AR-3]]SHALL NOT support automated rules processing in the protocol. A rules engine can certainly use the protocol.

* [AR-4] Anything other than read-only operations. The TRQP is a read-only (RETRIEVE in the CRUD sense) protocol.


