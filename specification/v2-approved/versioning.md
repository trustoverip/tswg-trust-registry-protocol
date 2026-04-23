## Versioning, Extensibility, and Backwards Compatibility

*This section is normative.*

### Protocol Versioning

1. The TRQP protocol version follows a **major.minor** numbering scheme (e.g., 2.0, 2.1, 3.0).
2. A **major version** change (e.g., 2.x to 3.0) indicates changes that are not backwards-compatible with the prior major version.
3. A **minor version** change (e.g., 2.0 to 2.1) indicates additions or clarifications that remain backwards-compatible with the current major version.
4. Implementations SHOULD include the TRQP specification version they conform to in their documentation or service metadata.

### Extensibility

1. The `context` object in TRQP queries is the primary extensibility mechanism. Because the `context` object permits additional JSON members beyond `time`, [[ref:TRQP bindings]] and profiles MAY define additional `context` members for domain-specific query conditions.
2. A [[ref:TRQP endpoint]] that receives a `context` member it does not recognize MUST ignore that member and process the query using only the members it supports.
3. Future [[ref:TRQP bindings]] MAY define additional query types beyond authorization and recognition, provided they conform to the identifier requirements in the [Identifiers](#identifiers) section.
4. Extensions MUST NOT redefine or alter the semantics of the fields defined in this specification.

### Backwards Compatibility

The following rules define what constitutes a backwards-compatible change to this specification:

1. **Backwards-compatible changes** (permitted in minor versions):
   1. Adding new OPTIONAL fields to query or response schemas.
   2. Adding new OPTIONAL `context` members.
   3. Adding new query types.
   4. Adding new informative sections or clarifications to normative text that do not change its meaning.
2. **Breaking changes** (requiring a new major version):
   1. Removing or renaming existing required fields in query or response schemas.
   2. Changing the semantics of existing fields.
   3. Adding new required fields to query schemas.
   4. Changing the identifier requirements in a way that invalidates previously conformant identifiers.
3. Implementations conforming to a given major version MUST accept and process queries from any minor version within that major version, ignoring any unrecognized OPTIONAL fields.
