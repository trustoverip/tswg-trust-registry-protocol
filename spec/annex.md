
[//]: # (Pandoc Formatting Macros)

[//]: # (# This is an annex {#sec:annexA .normative})

[//]: # (With some text)

[//]: # (# This is another annex {#sec:annexB .informative})

[//]: # (With some more text)

##  Annex A: Consolidated Requirements

For ease of reference, the following table consolidates all normative requirements in this specification. Each requirement is linked to the section in which it appears.

`THE FOLLOWING REQUIREMENTS IN THE TABLE ARE JUST EXAMPLES FOR NOW.`

TODO: Finalize table once requirements (earlier).

| Req # | Description | Section |
|---------|--------------|-----------|
| | **Governing Authority Requirements**| |
| GA-1 | [[ref: EGF]] MUST have exactly one [[ref: primary trust registry]]. | [#governing-authorities-ga-] |
| GA-2 | [[ref: EGF]] MAY have one or more [[ref: secondary trust registries]].| [[#governing-authorities-ga-]|
|A.3|MUST publish an [[ref: EGF]] that meets the **requirements** in: 
|A.3.1|    This specification. | [LINK]
|A.3.2| The [ToIP Governance Architecture Specification](https://wiki.trustoverip.org/pages/viewpage.action?pageId=71241). Note that this includes the requirement that the **EGF** and all **governed parties** (which includes **authorized issuers** and **authorized verifiers**) |[LINK]|


## Annex B: OpenAPI Specification

The OpenAPI Specification (v3.0.1) is the first "concrete" API specification. 

It is provided as an Open API Specification v3 YAML file. 

[OAS (.yaml) for TRP v2](../api/toip-tswg-trustregistryprotocol-v2.yaml). 

[Redoc Rendering (static HTML) of specification](../api/redoc-static.html)


## Annex C - Uses and Data Model Reference

### Use of the Trust Registry Protocol.

The TRP is intended to be used in at least two key ways:

* Native Support - systems may directly implement access using the TRP.
* Bridged - systems may create access "bridges" that provide TRP access to their systems.

![C4 Systems Model - showing native TRP support on one system, bridged support to two other systems (e.g. TRAIN and EU Trusted List ARF)](../assets/out/diagrams/protocol-bridging/protocol-bridging.png).


### Object Model

We provide a high-level object model (NOTE: source of truth is the Swagger as this diagram may be out of date during development)

![High Level Object Model](../assets/out/diagrams/highlevel/highlevel.png)