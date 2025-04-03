## Authority Statements And Query Vocabulary

_this section is normative._

Authority statements are the heart of TRQP architecture—the “lingua franca” of digital trust verification. From a logical standpoint, they represent the content available in any TRQP-compliant trust registry, regardless of the data model, structure, or technology used in the underlying system of record.

### Standard Structure

To enable semantics to be shared across ecosystems, TRQP authority statements are structured in three standard parts as shown in figure 4:

![images/authority_statements.png](images/authority_statements.png)
*Figure 4: The standard three-part structure of TRQP authority statements*

These three strings are simple yet flexible enough to express all types of authority statements, including authorization, recognition, delegation, and description (metadata) as defined in this section.

### Authorization Statements

In an authorization statement, an authority grants an authorization to an entity under its authority. In the ToIP governance model, this entity is called a [governed party](https://glossary.trustoverip.org/#term:governed-party).

#### ABNF

The assertion in a TRQP authorization statement MUST be a string that conforms to the following ABNF:

`auth-assertion	= auth-type "/" scope`  
`auth-type		= segment			; as defined in RFC 3986`  
`scope			= URI-reference		; as defined in RFC 3986`

#### Authorization Assertions for Verifiable Digital Credentials

The following enumerated strings SHOULD be used for authorization assertions governing verifiable digital credentials:

`auth-string	= "issue" / "verify"`

The `issue` string SHOULD be used for an entity authorized to act in the role of issuing a digital credential as defined in the specification for the relevant digital credential format.

The `verify` string SHOULD be used for an entity authorized to act in the role of verifying a digital credential as defined in the specification for the relevant digital credential format.

The `scope` of an digital credential authorization assertion SHOULD define the type of digital credential the entity is authorized to issue or verify. It is RECOMMENDED to:

1. Use the same `URI-reference` that uniquely defines the credential type as defined in the appropriate credential specification or type definition.  
2. Publish that `URI-reference` in the ecosystem governance framework and any associated type catalogues.

#### Authorization Assertions for Other Verifiable Data

The following enumerated strings SHOULD be used for authorization assertions governing other forms of verifiable data:

`auth-string	= "publish" / "consume"`

The `publish` string SHOULD be used for an entity authorized to act in the role of publishing and digitally signing verifiable data that is not in a digital credential format.

The `consume` string SHOULD be used for an entity authorized to act in the role of requesting, verifying, and using digitally signed verifiable data that is not in a digital credential format.

The `scope` of a verifiable data authorization assertion SHOULD define the type of verifiable data an entity is authorized to issue or verify. It is RECOMMENDED to:

1. Use the same `URI-reference` that uniquely defines the verifiable data type as defined in the appropriate specification or type definition.  
2. Publish that `URI-reference` in the ecosystem governance framework and any associated type catalogues.

### Recognition Statements

In a recognition statement, one ecosystem governing authority recognizes another ecosystem governing authority as a peer.

#### ABNF

The assertion in a TRQP recognition statement MUST be a string that conforms to the following ABNF:

`recog-string	= "recognizes”`

#### Recognition Assertions

Recognition assertions do not have a `scope` because by definition both authorities are sovereign—neither controls the scope of the other. If two authorities have a control relationship, it MUST be expressed using a delegation statement.

### Delegation Statements

A delegation statement expresses a control relationship between two ecosystem governing authorities where one delegates a specific scope of authority to the other.

#### ABNF

The assertion in an TRQP delegation statement MUST be a string that conforms to the following ABNF:

`deleg-assertion	= delegation "/" scope`  
`delegation		= segment			; as defined in RFC 3986`  
`scope			= URI-reference		; as defined in RFC 3986`

#### Delegation Assertions

The following enumerated strings SHOULD be used for delegation assertions:

`deleg-string	= "delegates" / "delegated-by"`

The `delegates` string SHOULD when one ecosystem governing authority is delegating authority for a specific `scope` to another ecosystem governing authority.

The `delegated-by` string SHOULD be used to express the precise inverse delegation relationship as expressed by the `delegates` string.

It is RECOMMENDED that a TRQP consumer verifies a `delegated-by` assertion by querying the TRQP endpoint of the delegating authority for the inverse `delegates` assertion.

The `scope` of a delegation assertion MUST identify the governance framework (or subset of a governance framework) for which authority is being delegated. It is RECOMMENDED to define and publish this `URI-reference` in the governance framework as a self-reference.

### Description (Metadata) Statements

An accompanying description statement MAY contain any type of metadata
describing the entity ID. If the authority ID for the description statement is
the same as the entity ID, then the description is considered self-asserted by
the authority. If the authority ID and entity ID are different, then the
authority is asserting the metadata as a set of claims about the entity ID. 

Description statements can be used to access any relevant metadata about an
entity, including metadata about the ecosystem governing authority, the
ecosystem governance framework, the trust registry operator, the trust registry,
or any governed party.

#### ABNF

The assertion in a TRQP description statement MUST be a string that conforms to the following ABNF:

`desc-assertion	= URI-reference		; as defined in RFC 3986`

#### Generic Metadata Assertions

The following enumerated string SHOULD be used to query for generic metadata:

`desc-string	= "metadata"`

A TRQP description query with the `metadata` assertion MUST return the metadata for the entity ID as asserted by the authority ID.

#### Verification Metadata Assertions

Certain types of verifiable identifiers, called [self-certifying identifiers](https://glossary.trustoverip.org/#term:self-certifying-identifier) (SCIDs), have defined formats for their associated verification metadata. Examples are KERI AIDs \[normative-reference\] and [did:scid DIDs](https://lf-toip.atlassian.net/wiki/spaces/HOME/pages/88572360/DID+SCID+Method+Specification) \[normative-reference\]. The following enumerated string SHOULD be used to query for the verification metadata for a SCID:

`desc-string	= "verification-metadata"`

A TRQP description query with the `verification-metadata` assertion MUST return the verification metadata for the entity ID as asserted by the authority ID.

:::note
IMPORTANT: There are two types of verification metadata for a SCID: Self-asserted—the verification metadata is published directly by the SCID controller. Witnessed—the verification metadata has been verified by an authority independent from the SCID controller. To query for self-asserted verification metadata, both the authority ID and the entity ID MUST be the SCID. To query for witnessed verification metadata, the authority ID MUST be the witness and the entity ID MUST be the SCID. 
:::

#### Verifiable Credential Assertions

Trust registries MAY also store verifiable credentials describing any entity in an ecosystem. To query for a specific type of verifiable credential describing an entity:

1. The authority ID SHOULD be the issuer of the credential.  
2. The assertion SHOULD be the `URI-reference` identifying the credential type.  
3. The entity ID MUST be the subject of the credential.
