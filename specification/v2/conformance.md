## Conformance

*This section is normative.*

As well as sections marked *normative*, all authoring guidelines, diagrams, examples, and notes in this specification are non-normative. Everything else in this specification is normative.

The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED", "MAY", and "OPTIONAL" in this specification are to be interpreted as described in [RFC 2119](https://datatracker.ietf.org/doc/html/rfc2119).

### Conformance Targets

This specification defines two conformance targets:

1. **TRQP Endpoint** — a network service that receives TRQP queries and returns TRQP responses on behalf of a [[ref:trust registry]].
2. **[[ref: TRQP Consumer]]** — a network client or server that sends TRQP queries to a [[ref:TRQP endpoint]].

A conformant implementation MUST satisfy the requirements for at least one of these targets using at least one [[ref:TRQP binding]].

### TRQP Endpoint Conformance

A conformant TRQP endpoint MUST:

1. Support at least one of the following query types:
   1. **Authorization queries** — accepting queries and returning responses that conform to the authorization query and response schemas defined in this specification.
   2. **Recognition queries** — accepting queries and returning responses that conform to the recognition query and response schemas defined in this specification.
2. Accept and validate identifiers conforming to the requirements in the [Identifiers](#identifiers) section, including:
   1. All identifiers MUST be represented as single strings conforming to [RFC 3986](https://datatracker.ietf.org/doc/html/rfc3986).
   2. The `authority_id` MUST be a globally unique identifier.
   3. The `action` and `resource` MUST each be non-empty strings.
   4. If a `context` object is provided, it MUST be processed as a JSON object. If the `context` includes a `time` parameter, its value MUST be interpreted as an [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339) datetime.
3. Return responses that conform to the response schemas defined in this specification for each supported query type.
4. Make the `ecosystem governance framework` discoverable via the `authority_id` as required in the [Identifiers](#identifiers) section.

### TRQP Consumer Conformance

A conformant [[ref: TRQP consumer]] MUST:

1. Send queries that conform to the query schemas defined in this specification for at least one of the supported query types (authorization, recognition, or both).
2. Construct queries using identifiers that conform to the requirements in the [Identifiers](#identifiers) section, including:
   1. All identifiers MUST be represented as single strings conforming to [RFC 3986](https://datatracker.ietf.org/doc/html/rfc3986).
   2. The `action` and `resource` MUST each be non-empty strings.
   3. If a `context` object is included, it MUST be a JSON object conforming to the requirements in this specification.
3. Be capable of processing responses that conform to the response schemas defined in this specification for each query type it supports.

### TRQP HTTPS Binding Conformance

A conformant implementation of the TRQP HTTPS Binding MUST additionally satisfy the following requirements:

1. All HTTPS requests MUST include a `Content-Type: application/json` header.
2. Authorization queries MUST be sent as `POST /authorization` requests with a JSON body conforming to the authorization query schema.
3. Recognition queries MUST be sent as `POST /recognition` requests with a JSON body conforming to the recognition query schema.
4. TRQP endpoints MUST return error responses using the [Problem Details for HTTP APIs](https://datatracker.ietf.org/doc/html/rfc7807) format with appropriate HTTP status codes.
5. TRQP endpoints MUST return HTTP 200 for successful queries with a JSON body conforming to the appropriate response schema.
