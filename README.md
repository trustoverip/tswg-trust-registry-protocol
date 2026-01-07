# Overview

The ToIP Trust Registry Query Protocol (TRQP) is a lightweight, read-only protocol for making fast, efficient queries for authoritative data from trust registries, also known as trust lists. To use an analogy, TRQP is to trust registries what DNS is to name servers.

The same way DNS name servers serve name domains, TRQP trust registries serve trust domains, also known as digital trust ecosystems. Four primary actors participate in the flow of verifiable data (including verifiable credentials) within the ecosystem: 1) data producers (issuers), 2) data subjects (holders), 3) data consumers (verifiers or relying parties), and 4) governing bodies (authorities).

Authorities determine the policies governing which actors can perform what actions on what data within the ecosystem. These policies are typically published in a human-readable form called a governance framework, also known as a trust framework. To make these policies accessible to software agents, they are published in a machine-readable form known as authority statements.

Authority statements can be published in a file, issued to individual actors as verifiable credentials, or published in a trust registry.

Digitally-verifiable authority statements can be expressed using various standards, including X.509 certificate hierarchies, OpenID Federations, EBSI Trust Chains, or TRAIN trust lists. Although these standards can work well for intra-ecosystem authority verification, they are not optimized for inter-ecosystem authority verification.

The purpose of TRQP is to bridge this gap by provide a standard protocol for querying authority statements from any TRQP-compliant trust registry. It specifies a standard data model, query vocabulary, and transport protocol binding that can be implemented by any ecosystem regardless of its internal trust architecture.

TRQP focuses on two query types:

1. Authorization Queries: “Has Entity X been granted Authorization Y under Ecosystem Governance Framework Z?”
2. Recognition Queries: "Is Ecosystem A recognized as having Authorization Y by Ecosystem C?” 

# specification-template

This specification is based on the [Trust Over IP Specification Template](https://github.com/trustoverip/specification-template).

The spec is written using [SpecUp](https://github.com/decentralized-identity/spec-up) which is maintained by the Decentralized Identity Foundation. 

To browse the spec, see the [rendering on GitHub pages](https://trustoverip.github.io/tswg-trust-registry-protocol/). To contribute to the spec, submit PRs that modify the .md files (in the `./spec` folder) that are used to generate the .html files in this folder.

Before submitting a PR, please see the [Editing The Spec](./EditingTheSpec.md) document for guidance on generating the specification locally for review.

## Rendering Spec-Up

To run Spec-up in live edit mode (will re-render upon save), in project folder run:

```
npm run edit
```

## Future Version Considerations:

The TRQP v2.0 specification is focused solely on Recognition and Authorization queries. Two other key areas were under discussion but didn't reach a point of closure, so have not been included in the v2.0 specification and may be addressed in further releases. These are:

- **Delegation Queries**: "Has Ecosystem A been delegated authority for Governance Framework D by Ecosystem C?"
- **Description (Metadata) Queries**: “What DID methods does Ecosystem A support?”

A further area about establishing a "query language" emerged from the efforts in the Trust Registry Task Force, which has taken on the [Trust Registry Query Language](https://lf-toip.atlassian.net/wiki/spaces/HOME/pages/149749777/TRQL+Trust+Registry+Query+Language) as a separate specification to consider.