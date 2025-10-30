**Trust Registry Query Protocol (TRQP) Specification**
==================

- **Specification Version:** 2.0
- **Document Status:** Public Review 01

**Participate:**
~ [GitHub repo](https://github.com/trustoverip/tswg-trust-registry-protocol/tree/main)
~ [File an issue](https://github.com/trustoverip/tswg-trust-registry-protocol/issues)
~ [Start a discussion](https://github.com/trustoverip/tswg-trust-registry-protocol/discussions)

- **Editors:** Darrell O’Donnell, Andor Kesselman, Drummond Reed, Antti Kettunen,  
- **Contributors:** Alex Tweeddale, Christine Martin, Dave Poltorak, Eric Drury, Fabrice Rochette, Jacques Latour, Jesse Carter, Jeff Braswell, Jon Bauer, Makki Elfatih, Marcus Ubani, Markus Sabadello, Scott Perry, Sankarshan Mukhopadhyay, Subhasis, Tim Bouma

::: note
This specification is currently a Working Draft of the Trust Over IP (ToIP) [Trust Registry Task Force](https://wiki.trustoverip.org/display/HOME/Trust+Registry+Task+Force). Feedback is welcome - see instructions [here](https://github.com/trustoverip/tswg-trust-registry-protocol)

For a complete overview of the motivations and core concepts behind TRQP, please see the [TRQP Overview page](https://lf-toip.atlassian.net/wiki/spaces/HOME/pages/22996548/ToIP+Trust+Registry+Query+Protocol+TRQP+Specification+Overview).
:::

## Introduction
_This section is informative._

The ToIP Trust Registry Query Protocol (TRQP) is a lightweight, read-only protocol for making fast, efficient queries for authoritative data from *trust registries*, also known as *trust lists*. To use an analogy, TRQP is to trust registries what DNS is to name servers.

The same way DNS name servers serve name domains, TRQP trust registries serve *trust domains*, also known as *digital trust ecosystems*. Four primary actors participate in the flow of verifiable data (including *verifiable credentials*) within the ecosystem: 1) data producers (issuers), 2) data subjects (holders), 3) data consumers (verifiers or relying parties), and 4) governing bodies (authorities). 

Authorities determine the policies governing which actors can perform what actions on what data within the ecosystem. These policies are typically published in a human-readable form called a *governance framework*, also known as a *trust framework*. To make these policies accessible to software agents, they are published in a machine-readable form known as *authority statements*. Authority statements can be published in a file, issued to individual actors as verifiable credentials, or published in a trust registry.

Digitally-verifiable authority statements can be expressed using various standards, including X.509 certificate hierarchies, OpenID Federations, EBSI Trust Chains, or TRAIN trust lists. Although these standards can work well for *intra-ecosystem* authority verification, they are not optimized for *inter-ecosystem* authority verification.

The purpose of TRQP is to bridge this gap by provide a standard protocol for querying authority statements from any TRQP-compliant trust registry. It specifies a standard data model, query vocabulary, and transport protocol binding that can be implemented by any ecosystem regardless of its internal trust architecture.

TRQP focuses on two query types:

1. **Authorization Queries:** “Has Entity X been granted Authorization Y under Ecosystem Governance Framework Z?”
2. **Recognition Queries**: "Is Ecosystem A recognized as an authority for Governance Framework B by Ecosystem C?"
