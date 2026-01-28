## Identifiers

*This section is normative.*

Interoperability of TRQP across decentralized trust domains, communities, and ecosystems depends on globally unique identifiers in the same way interoperability of the Internet depends on IP addresses and DNS names.

The following requirements apply to all identifiers defined in this section:

1. The identifier MUST be represented as a single string conforming to [IETF RFC 3986](https://datatracker.ietf.org/doc/html/rfc3986).  
2. For globally unique identifiers, it is RECOMMENDED to use a [verifiable identifier](https://glossary.trustoverip.org/#term:verifiable-identifier) such as a W3C Decentralized Identifier (DID), a KERI autonomic identifier (AID), or an HTTPS URL so their authenticity can be verified by any relying party.

For additional assurance, it is RECOMMENDED to use multi-anchoring of identifiers as defined by the IETF [High Assurance DIDs using DNS specification](https://www.ietf.org/archive/id/draft-carter-high-assurance-dids-with-dns-03.html) or the work of the ToIP [High Assurance Verifiable Identifiers Task Force](https://lf-toip.atlassian.net/wiki/spaces/HOME/pages/32473104/High+Assurance+VID+Task+Force+HAVID?atlOrigin=eyJpIjoiMWJkOTU4MjI5NTdhNGU0ZTlhMmI3MGRlNWYwNmVmMGQiLCJwIjoiYyJ9).

::: note 

The TRQP information model is patterned after the [PARC model](https://docs.cedarpolicy.com/auth/authorization.html). The identifiers in this section map well to PARC: Principal ~= `entity_id`; Action ~= `action` Resource ~= `resource`; Context ~= `authority_id` as the *mandatory* context. An optional `context` object is available for additional query conditions, such as time.

:::

### `authority_id`

1. The `authority_id` MUST be the globally unique identifier of the authority in the context of an authority statement.  
2. It is RECOMMENDED that the `authority_id` be published in the governance framework for that ecosystem.  
3. It is RECOMMENDED that the TRQP service endpoint(s) for any authoritative trust registry be machine-discoverable via the `authority_id`. An example would be to publish either of the following in the DID document for the `authority_id`:   
   1. The authoritative TRQP service endpoint URL(s).  
   2. The DID(s) identifying authoritative trust registries. (In that case, the authoritative TRQP service endpoint URL(s) would be specified in the associated DID documents.)
4. The `ecosystem governance framework` MUST be discoverable via the `authority_id`. This can be established in the DIDDocument (in the case the `authority_id` is a DID, or another mechanism.)

### `entity_id`

1. The `entity_id` MUST be the identifier of the principal in an authority statement.  
2. For a recognition or delegation statement, the `entity_id` MUST also be an `authority_id`.
3. For an authorization statement about a [governed party](https://glossary.trustoverip.org/#term:governed-party), the `entity_id` MUST be unique in the scope of the authority. It is NOT REQUIRED for the `entity_id` to be globally unique.

### `action`

1. The identifier for an `action` in an authority statement MUST be a non-empty string conformant to the requirements in this section.  
2. It is NOT REQUIRED for the `action` identifier string to be globally unique.  
3. The `action` identifier string SHOULD:  
   1. Be specified in the authority’s governance framework.  
   2. Be machine-discoverable via a query to the authoritative trust registr(ies).

::: note

To encourage cross-ecosystem interoperability, in a future version of this specification (or a companion specification), the ToIP Trust Registry Task Force intends to publish a vocabulary of common `action` strings.

:::

### `resource`

1. The identifier for a `resource` in an authority statement MUST be a non-empty string conformant to the requirements in this section.  
2. It is NOT REQUIRED for the `resource` identifier string to be globally unique.  
3. In addition, the requirements for `resource` identifiers SHOULD:  
   1. Be defined in the authority’s governance framework.  
   2. Be machine-discoverable via a query to the authoritative trust registr(ies).

### `context` 

Because `authority_id` is the required context for all TRQP [[ref: authority statements]], a `context` object is OPTIONAL in a [[ref: TRQP query]]. If a `context` object is included, it MUST conform to the requirements in this section.

1. A `context` object MUST be a JSON object whose members convey other query conditions.   
2. If a `context` object needs to express a time-based condition:  
   1. It MUST include a `time` parameter.
   2. The value of this parameter MUST a time value expressed in [RFC 3339](https://trustoverip.github.io/tswg-trust-registry-protocol/%7B%7BRFC3339%7D%7D) format.  
   3. The value of the `time` parameter MUST be interpreted as the datetime as of which the target [[ref: authority statement]] is valid.  
3. Additional JSON object members specifying other conditions MAY be defined by TRQP profiles or bindings.