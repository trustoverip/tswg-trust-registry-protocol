## Related Specification (Trust Registry Query Language (TRQL))

The Trust Registry Query Language (TRQL, pronounced “turkle”) is a companion specification to TRQP, developed to provide a standardized language for expressing queries against trust registries. TRQL was established by the Trust Registry Task Force (TRTF) to enable rapid evolution of query capabilities while keeping the TRQP specification stable. TRQL defines the vocabulary, syntax, and semantics for querying authority, recognition, delegation, and metadata statements in TRQP-compliant registries. It supports queries for authorizations, recognitions, delegated authorities, descriptions, and verifiable credential assertions, using structured strings and URI-references to ensure interoperability across ecosystems.

For details on the Trust Registry Query Language (TRQL), which defines the query syntax and semantics used by TRQP, see the [TRQL specification](https://lf-toip.atlassian.net/wiki/spaces/HOME/pages/149749777/TRQL+Trust+Registry+Query+Language).

## Security Considerations

All implementers (“bindings [[ref:TRQP Binding]]” and “bridges [[ref:TRQP Bridge]]”) of TRQP **SHOULD** take the following threats into account and implement appropriate controls:

- **Trust Anchor Hijacking**: Use strong cryptography and rotate keys regularly.
- **Trust Registry Bugs**: Conduct code reviews, vulnerability scans, and robust QA.
- **Trust Anchor Spoofing**: Verify responses using known cryptographic anchors or certificate chains.
- **Domain Hijacking**: Protect DNS entries; if DNS-based discovery is used, consider DNSSEC or other verification.
- **Replay Attacks**: Use timestamps, nonces, and short-lived tokens.
- **Data Integrity**: Sign or hash data at rest, use TLS or equivalent in transit.
- **Denial of Service (DoS)**: Rate-limit queries, monitor usage, scale infrastructure appropriately.
- **Insufficient Data Validation**: Enforce strict schema checks and reject malformed data with clear error messages.
- **Trust Anchor Compromise**: Implement multi-tier trust anchors, and have a plan to revoke or replace compromised keys quickly.
- **Logging and Auditing**: Log all access, changes, and suspicious activities; adopt real-time monitoring.
- **Protocol Downgrade Attacks**: Default to the latest secure version, disallow fallback to insecure versions.
- **Privacy Concerns**: Encrypt sensitive or personally identifiable information, and comply with relevant data protection laws.
- **Timing Attacks**: Where feasible, adopt constant-time operations for cryptographic and authorization checks.

## Privacy Considerations

Implementers must design the system so that the handling of authorizations and identity information minimizes the risk of exposing sensitive details. In addition to data minimization and regulatory compliance, pay special attention to the following:

- **Careful Handling of Authorizations & Identities**:  
  - Ensure that authorization tokens and identity data (such as DIDs) do not leak more information than necessary.  
  - Avoid exposing internal structures or hierarchies that could be exploited by correlating queries or monitoring network traffic.
- **Correlation Risk Mitigation**:  
  - Recognize that even though only ecosystem DIDs are transmitted in recognition queries, combining this data with network-level information (such as IP addresses) can lead to correlation attacks.  
  - Consider implementing techniques such as query obfuscation, randomized response timings, or other privacy-preserving measures to reduce the risk of linking requests back to a particular requester.

## Implementation Considerations

Implementing the TRQP for the Ayra Trust Network requires a dual focus: establishing secure bridging mechanisms between heterogeneous trust frameworks and ensuring that trust registries accurately represent the ecosystem model. Key aspects include:

- **Bridging Across Ecosystems**:  
  - **Adapter Development**: As the TRQP is designed to bridge various intra-trust frameworks (such as Open ID Federation, X509 Chains, or TRAIN), implementers should develop adapters tailored to their ecosystem. These adapters translate local trust assertions into the common TRQP language.  
  - **Interoperability Focus**: The bridge should speak the “lowest common denominator” across frameworks. This means abstracting complex internal authorization models into a standardized format that can be reliably queried by external verifiers.
  - **Flexible Query Handling**: Ensure your bridging solution can handle both Recognition and Authorization Queries seamlessly. For example, when verifying whether an entity is authorized under a particular Ecosystem Governance Framework (EGF), the adapter must translate internal rules into a compliant TRQP response.
- **Trust Registries & Ecosystem Representation**:  
  - **Trust Registry as the Ecosystem State Holder**: A Trust Registry is not the ecosystem itself but a capability within the ecosystem that serves as the authoritative source for its authorization state. It must be capable of answering queriesthat reflect the current state of authority in the ecosystem.
  - **Ecosystem Model Mapping**:  
    - Trust registries should maintain a clear mapping of internal ecosystem authorizations to the standardized TRQP model. This means maintaining up-to-date relationships between entities, their roles, and the associated authorizations.  
    - Represent the ecosystem model in a way that abstracts the underlying complexity—ensuring that verifiers only need to interact with a uniform interface regardless of the diversity of the underlying trust frameworks.
  - **Bridging Governance & Registries**:  
    - The registry should integrate with the broader ecosystem governance framework, adhering to the TRQP’s requirements for identifier creation (using compliant DID methods) and service endpoint specifications.  
    - Document how the registry bubbles up the state of authorizations, including how updates and revocations are handled to maintain an accurate and timely reflection of the ecosystem’s trust landscape.

