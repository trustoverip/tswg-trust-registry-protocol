## Identifiers

*This section is normative.*

Interoperability of TRQP across decentralized trust domains, communities, and ecosystems depends on globally unique identifiers in the same way interoperability of the Internet depends on IP addresses and DNS names.

The following requirements apply to all identifiers defineddefining in this section:

1. The identifier MUST be represented as a single string conforming to [IETF RFC 3986](https://datatracker.ietf.org/doc/html/rfc3986) \[normative reference\].  
2. For globally unique identifiers, itIt is RECOMMENDED to use a [verifiable identifier](https://glossary.trustoverip.org/#term:verifiable-identifier) such as a W3C Decentralized Identifier (DID), a KERI autonomic identifier (AID), or an HTTPS URL so their authenticity can be verified by any relying party.

For additional assurance, it is RECOMMENDED to use multi-anchoring of identifiers as defined by the IETF [High Assurance DIDs using DNS specification](https://www.ietf.org/archive/id/draft-carter-high-assurance-dids-with-dns-03.html) \[normative reference\] or the work of the ToIP [High Assurance Verifiable Identifiers Task Force](https://lf-toip.atlassian.net/wiki/spaces/HOME/pages/32473104/High+Assurance+VID+Task+Force+HAVID?atlOrigin=eyJpIjoiMWJkOTU4MjI5NTdhNGU0ZTlhMmI3MGRlNWYwNmVmMGQiLCJwIjoiYyJ9) \[informative reference\].

### `authority_id`

1. The `authority_id` MUST be the globally unique identifier of the authority in the context of an authority statement.  
2. It is RECOMMENDED that the `authority_id` be published in the governance framework for that ecosystem.  
3. It is RECOMMENDED that the TRQP service endpoint(s) for any authoritative trust registry be machine-discoverable via the `authority_id`. An example would be to publish either of the following in the DID document for the `authority_id`:   
   1. The authoritative TRQP service endpoint URL(s).  
   2. The DID(s) identifying authoritative trust registries. (In that case, the authoritative TRQP service endpoint URL(s) would be specified in the associated DID documents.)

### `entity_id`

1. The `entity_id` MUST be the identifier of the principal in an authority statement.  
2. For a recognition or delegation statement, the `entity_id` MUST be an `authority_id`.  
3. For an authorization statement about a [governed party](https://glossary.trustoverip.org/#term:governed-party), the `entity_id` MUST be unique in the scope of the authority. It is NOT REQUIRED for the `entity_id` to be globally unique.

::: warning
point 2 above is WRONG?
> 2. For a recognition or delegation statement, the `entity_id` MUST be an `authority_id` ???in another ecoystem???.
:::

### `action`

1. The identifier for an `action` in an authority statement MUST be a non-empty string conformant to the requirements in this section.  
2. It is NOT REQUIRED for the `action` identifier string to be globally unique.  
3. The `action` identifier string SHOULD:  
   1. Be specified in the authority’s governance framework.  
   2. Be machine-discoverable via a query to the authoritative trust registr(ies).
4. ~~If the `action` is one of the actions defined in the Trust Registry Query Language (TRQL) specification to be defined by the ToIP Trust Registry Task Force, the action SHOULD use the action identifier string defined in that specification.~~ 


### `resource`

1. The identifier for a `resource` in an authority statement MUST be a non-empty string conformant to the requirements in this section.  
2. It is NOT REQUIRED for the `resource` identifier string to be globally unique.  
3. In addition, the requirements for `resource` identifiers SHOULD:  
   1. Be defined in the authority’s governance framework.  
   2. Be machine-discoverable via a query to the authoritative trust registr(ies).
4. ~~If the `resource` is one of the resource types defined in the Trust Registry Query Language (TRQL) specification to be defined by the ToIP Trust Registry Task Force, the resource identifier SHOULD follow the recommendations in that specification.~~  


::: warning
We have referenced TRQL in normative section. TODO decision to REMOVE.
:::

### Optional `context` Refinement

*This section is normative.*

In the PARC model, a context object is used to assert conditions that must be satisfied in order for an action to be permitted. As defined in section TODO, a TRQP query the `authority_id` is the primary context.  

Further refinement of the context is possible using an optional context parameter.

Other than this required parameter, in a TRQP query:

1. A `context` object is OPTIONAL.  
2. If included, a `context` object MUST be a JSON object whose members convey other conditions.   
3. If a `context` object includes a time value: a time-based condition:  
   1. The `time` parameter MUST be expressed using a time member whose value is in [RFC 3339](https://trustoverip.github.io/tswg-trust-registry-protocol/%7B%7BRFC3339%7D%7D) format.  
   2. The value of the `time` parameter MUST be interpreted as the datetime as of which the target authority statement is valid.  
4. Additional JSON object members specifying other conditions MAY be defined by TRQP profiles or bindings.