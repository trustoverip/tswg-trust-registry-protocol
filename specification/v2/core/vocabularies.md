This is supplementary material that is being drafted on the Trust Over IP Wiki [here](https://lf-toip.atlassian.net/wiki/spaces/HOME/pages/149749777/TRQP+Query+Vocabulary)

### Authority IDs and Entity IDs

Every authority statement is made by an authority **about** an entity. Therefore, the characteristics of these two identifiers are especially important:

##### Authority ID
- **MUST** be a globally unique identifier for the authority making the statement.  
- **MUST** be represented as a single string conforming to [IETF RFC&nbsp;3986](https://datatracker.ietf.org/doc/html/rfc3986).  
- Should be a cryptographically verifiable identifier (e.g., DID or AID) or an HTTPS URL, so that verifiers can resolve and verify public keys, TRQP service endpoints, and other metadata describing the authority.  
- It is recommended to use multi-anchoring of the verifiable identifier for additional assurance (for example, using [High Assurance DIDs using DNS](https://www.ietf.org/archive/id/draft-carter-high-assurance-dids-with-dns-03.html) or equivalent).

##### Entity ID
- **MUST** be unique **within** the ecosystem.  
  - If correlation across ecosystems is desirable (for example, to establish the reputation of a credential issuer), then it is recommended to be globally unique.  
  - If correlation across ecosystems is **not** desirable (for example, for individual privacy), then it is recommended to be a locally unique identifier.  
- **MUST** be represented as a single string conforming to [IETF RFC&nbsp;3986](https://datatracker.ietf.org/doc/html/rfc3986).  
- Should be a cryptographically verifiable identifier (e.g., DID or AID) or an HTTPS URL, so that verifiers can resolve public keys, service endpoints, and other metadata describing the entity.

### Authorization Statements

In an authorization statement, an ecosystem grants an authorization to an entity under its authority. In the ToIP governance model, this entity is called a [governed party](https://glossary.trustoverip.org/#term:governed-party).

#### ABNF (Authorization Assertions)
```txt
auth-assertion = authorization “/” scope
authorization  = segment         ; as defined in RFC 3986
scope          = URI-reference   ; as defined in RFC 3986
```

#### Authorization Assertions for Verifiable Digital Credentials

The following enumerated strings SHOULD be used for authorization assertions governing verifiable digital credentials:

`auth-string	= "issue" / "verify"`

- The `issue` string SHOULD be used for an entity authorized to act in the role of issuing a digital credential as defined in the specification for the relevant digital credential format.
- The `verify` string SHOULD be used for an entity authorized to act in the role of verifying a digital credential as defined in the specification for the relevant digital credential format.
- The `scope` of an digital credential authorization assertion SHOULD define the type of digital credential the entity is authorized to issue or verify. It is RECOMMENDED to:

1. Use the same URI string that uniquely defines the credential type as defined in the appropriate credential specification or type definition.  
2. Publish that URI string in the ecosystem governance framework and any associated type catalogues.

#### Authorization Assertions for Other Verifiable Data

The following enumerated strings SHOULD be used for authorization assertions governing other forms of verifiable data:

`auth-string	= "publish" / "consume"`

- The `publish` string SHOULD be used for an entity authorized to act in the role of publishing and digitally signing verifiable data that is not in a digital credential format.
- The `consume` string SHOULD be used for an entity authorized to act in the role of requesting, verifying, and using digitally signing verifiable data that is not in a digital credential format.
- The `scope` of a verifiable data authorization assertion SHOULD define the type of verifiable data an entity is authorized to issue or verify. It is RECOMMENDED to:

1. Use the same URI string that uniquely defines the verifiable data type as defined in the appropriate specification or type definition.  
2. Publish that URI string in the ecosystem governance framework and any associated type catalogues.


### Recognition Statements

In a recognition statement, one ecosystem governing authority recognizes another ecosystem governing authority as a peer.

#### ABNF

The assertion in a TRQP recognition statement MUST be a string that conforms to the following ABNF:

`recog-string	= "recognizes"`

#### Recognition Assertions

Recognition assertions do not have a `scope` because by definition both authorities are sovereign—neither controls the scope of the other. If two authorities have a control relationship, it **MUST** be expressed using a delegation statement.

### Delegation Statements

A delegation statement expresses a control relationship between two ecosystem governing authorities where one delegates a specific scope of authority to the other.

#### ABNF

The assertion in an TRQP delegation statement **MUST** be a string that conforms to the following ABNF:

```txt
deleg-assertion	= delegation "/" scope
delegation		= segment			; as defined in RFC 3986  
scope			= URI-reference		; as defined in RFC 3986
```
#### Delegation Assertions
    
The following enumerated strings SHOULD be used for delegation assertions:

`deleg-string	= "delegates" / "delegated-by"`

The `delegates` string SHOULD when one ecosystem governing authority is delegating authority for a specific `scope` to another ecosystem governing authority.

The `delegated-by` string SHOULD be used to express the precise inverse delegation relationship as that expressed by the `delegates` string.

It is RECOMMENDED that a TRQP consumer verifies a `delegated-by` statement by querying the TRQP endpoint of the delegating authority for the inverse `delegates` statement.

The `scope` of a delegation assertion SHOULD identify the governance framework (or subset of a governance framework) for which authority is being delegated. It is RECOMMENDED to define and publish this URI string in the governance framework as a self-reference.


### Description (Metadata) Statements

A description statement identifies metadata stored in the trust registry that describes the identified entity. If the authority ID for the description statement is the same as the entity ID, then the description is self-asserted. If the authority ID and entity ID are different, then the authority is making an assertion that the metadata describes the entity.

Description statements can be used to lookup any relevant metadata about an entity, including metadata about the ecosystem governing authority, the ecosystem governance framework, the trust registry operator, the trust registry, or any governed party in the ecosystem.


#### ABNF (Description Assertions)

The assertion in a TRQP description statement **MUST** be a string that conforms to the following ABNF: 

```txt
desc-assertion = description
description    = URI-reference   ; as defined in RFC 3986
```

### Generic Metadata Assertions

The following enumerated string SHOULD be used to query for generic metadata:

`desc-string	= "metadata"`

A TRQP description query with the `metadata` assertion **MUST** return the metadata
for the entity ID as asserted by the authority ID.

#### Verification Metadata Assertions

Certain types of verifiable identifiers, called [self-certifying identifiers](https://glossary.trustoverip.org/#term:self-certifying-identifier) (SCIDs), have defined formats for their associated verification metadata. Two examples are KERI AIDs \[normative-reference\] and [did:scid DIDs](https://lf-toip.atlassian.net/wiki/spaces/HOME/pages/88572360/DID+SCID+Method+Specification) \[normative-reference\]. The following enumerated string SHOULD be used to query for the verification metadata for a SCID:

> **Important:** For a SCID, “self-asserted” vs. “witnessed” verification metadata must be queried by specifying the correct authority ID (the SCID itself vs. the witness).

#### Verifiable Credential Assertions

Trust registries can also store verifiable credentials describing any entity. To query for a specific credential type:

1. The authority ID should be the **issuer** of the credential.  
2. The assertion should be the **URI reference** identifying the credential type.  
3. The entity ID must be the **subject** of the credential.



