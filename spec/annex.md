
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
| GA-1 | [[xref: TOIP, EGF]] MUST have exactly one [[ref: primary trust registry]]. | [#governing-authorities-ga-] |
| GA-2 | [[xref: TOIP, EGF]] MAY have one or more [[ref: secondary trust registries]].| [[#governing-authorities-ga-]|
|A.3|MUST publish an [[xref: TOIP, EGF]] that meets the **requirements** in: 
|A.3.1|    This specification. | [LINK]
|A.3.2| The [ToIP Governance Architecture Specification](https://wiki.trustoverip.org/pages/viewpage.action?pageId=71241). Note that this includes the requirement that the **EGF** and all **governed parties** (which includes **authorized issuers** and **authorized verifiers**) |[LINK]|


## Annex B: OpenAPI Specification

The OpenAPI Specification (v3.1.0) is the first "concrete" API specification. 

It is provided as an Open API Specification v3 YAML file. 

[OAS (.yaml) for TRP v2](../api/toip-tswg-trustregistryprotocol-v2.yaml). 

There are several renderings of the OAS specification:

* Inline - this rendering is managed in this repository [Redoc Rendering (static HTML) of specification](./api/redoc-static.html)
* SwaggerHub - this rendering is manually updated from time to time and may be out of date: [SwaggerHub](https://app.swaggerhub.com/apis/CULedger/CULedger.Identity/0.3.1-oas3.1) 


## Annex C - Uses and Data Model Reference

### Use of the Trust Registry Protocol.

The TRP is intended to be used in at least two key ways:

* Native Support - systems may directly implement access using the TRP.
* Bridged - systems may create access "bridges" that provide TRP access to their systems.

![C4 Systems Model - showing native TRP support on one system, bridged support to two other systems (e.g. TRAIN and EU Trusted List ARF)](./images/puml/protocol-bridging.png).


### Object Model

We provide a high-level object model (NOTE: source of truth is the Swagger as this diagram may be out of date during development)

![High Level Object Model](./images/puml/highlevel.png)

## Annex D - Guides (for future breakout)

We will need to provide guides and other thought pieces that explain many aspects of trust registries. A notional (short bullet) list of items could include:
* "why do I need a trust registry?" - blog article or position paper to explain why trust registries help.
* "I have the data, but how do I use the TRP?" - paper about how adding TRP to a bridge or native integration.
* "where do I learn about the governance changes that I have?"