**Trust Registry Query Protocol (TRQP) Specification**
==================

- **Specification Version:** 2.0
- **Document Status:** Public Review 02

**Participate:**

~ [GitHub repo](https://github.com/trustoverip/tswg-trust-registry-protocol/tree/main)
~ [File an issue](https://github.com/trustoverip/tswg-trust-registry-protocol/issues)
~ [Start a discussion](https://github.com/trustoverip/tswg-trust-registry-protocol/discussions)

- **Editors:** Darrell O'Donnell, Andor Kesselman, Drummond Reed
- **Contributors:** Alex Tweeddale, Antti Kettunen, Christine Martin, Dave Poltorak, Eric Drury, Fabrice Rochette, Jacques Latour, Jesse Carter, Jeff Braswell, Jon Bauer, Makki Elfatih, Marcus Ubani, Markus Sabadello, Scott Perry, Sankarshan Mukhopadhyay, Subhasis Ojha, Tim Bouma

::: note
This specification has just completed Public Review 02 and is being dispositioned in the Trust Over IP (ToIP) [Trust Registry Task Force](https://wiki.trustoverip.org/display/HOME/Trust+Registry+Task+Force). 

For a complete overview of the motivations and core concepts behind TRQP, please see the [TRQP Overview page](https://lf-toip.atlassian.net/wiki/spaces/HOME/pages/22996548/ToIP+Trust+Registry+Query+Protocol+TRQP+Specification+Overview).
:::

### Intellectual Property Rights

This specification is provided under the [Joint Development Foundation](https://www.jointdevelopment.org/) (JDF) charter for the Trust Over IP Foundation and is subject to the intellectual property rights policy of the [Technology Stack Working Group](https://wiki.trustoverip.org/display/HOME/Technology+Stack+Working+Group):

- **Copyright:** [Creative Commons Attribution 4.0 International (CC BY 4.0)](https://creativecommons.org/licenses/by/4.0/). Contributors have signed the [OWF Contributor License Agreement 1.0 (Copyright)](COPYRIGHT_POLICY.md).
- **Patent:** [Open Web Foundation Final Specification Agreement 1.0 (OWFa 1.0 — Patent Only)](PATENT_LICENSING.md), consistent with the W3C Patent Mode.
- **Source Code:** [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0).

THESE MATERIALS ARE PROVIDED “AS IS.” The parties expressly disclaim any warranties (express, implied, or otherwise), including implied warranties of merchantability, non-infringement, fitness for a particular purpose, or title, related to the materials. The entire risk as to implementing or otherwise using the materials is assumed by the implementer and user. IN NO EVENT WILL THE PARTIES BE LIABLE TO ANY OTHER PARTY FOR LOST PROFITS OR ANY FORM OF INDIRECT, SPECIAL, INCIDENTAL, OR CONSEQUENTIAL DAMAGES OF ANY CHARACTER FROM ANY CAUSES OF ACTION OF ANY KIND WITH RESPECT TO THIS DELIVERABLE OR ITS GOVERNING AGREEMENT, WHETHER BASED ON BREACH OF CONTRACT, TORT (INCLUDING NEGLIGENCE), OR OTHERWISE, AND WHETHER OR NOT THE OTHER MEMBER HAS BEEN ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

### Change History

**Public Review 02** — December 2025

- Restructured specification into modular files (overview, definitions, high-level architecture, identifiers, API schemas, HTTPS binding, conformance, considerations, references).
- Adopted the PARC (Principal, Action, Resource, Context) information model for identifiers and query structure.
- Replaced `ecosystem_id` and `assertion_id` with `authority_id`, `entity_id`, `action`, and `resource` aligned to the PARC model.
- Added formal JSON schemas for authorization and recognition request/response messages.
- Added HTTPS binding section with detailed request/response examples and error handling using RFC 7807 Problem Details.
- Added `locator` as optional parameter to the `context` object for authorization queries.
- Added conformance criteria section defining TRQP Endpoint, TRQP Consumer, and HTTPS Binding conformance targets.
- Added formal Normative and Informative References sections.
- Added Versioning, Extensibility, and Backwards Compatibility section.
- Deferred delegation and description (metadata) query types to a future version.
- Extensive editorial revision throughout.

**Public Review 01** — April 2025

- Initial public review of TRQP v2.0 specification.
- Defined core protocol concepts: authorities, authority statements, trust registries, governance frameworks.
- Defined authorization and recognition query types.
- Established identifier requirements based on RFC 3986.
- Included Security, Privacy, and Implementation Considerations sections.

## Introduction

*This section is informative.*

The ToIP Trust Registry Query Protocol (TRQP) is a lightweight, read-only protocol for making fast, efficient queries for authoritative data from *trust registries*, also known as *trust lists*. To use an analogy, TRQP is to trust registries what DNS is to name servers.

The same way DNS name servers serve name domains, TRQP trust registries serve *trust domains*, also known as *digital trust ecosystems*. Four primary actors participate in the flow of verifiable data (including *verifiable credentials*) within the ecosystem: 1) data producers (issuers), 2) data subjects (holders), 3) data consumers (verifiers or relying parties), and 4) governing bodies (authorities). 

Authorities determine the policies governing which actors can perform what actions on what data within the ecosystem. These policies are typically published in a human-readable form called a *governance framework*, also known as a *trust framework*. To make these policies accessible to software agents, they are published in a machine-readable form known as *authority statements*. Authority statements can be published in a file, issued to individual actors as verifiable credentials, or published in a trust registry.

Digitally-verifiable authority statements can be expressed using various standards, including X.509 certificate hierarchies, OpenID Federations, EBSI Trust Chains, or TRAIN trust lists. Although these standards can work well for *[[ref:intra-ecosystem]]* authority verification, they are not optimized for *[[ref:inter-ecosystem]]* authority verification.

The purpose of TRQP is to bridge this gap by providing a standard protocol for querying authority statements from any TRQP-compliant trust registry. It specifies a standard data model, query vocabulary, and transport protocol binding that can be implemented by any ecosystem regardless of its internal trust architecture.

TRQP focuses on two query types:

1. **Authorization Queries:** “Has Authority A authorized Entity B to take Action X on Resource Y?”
2. **Recognition Queries**: "Does Authority X recognize Entity B as an authority to authorize taking Action X on Resource Y?”
