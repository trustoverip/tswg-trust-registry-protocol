## High-Level Architecture

*This section is informative.*

This section presents the major conceptual components in TRQP architecture.

### Authorities

Authorities are the parties responsible for establishing the policies governing their ecosystem. An authority may take any legal form or may not be a formal legal entity at all. The only requirement is that the authority be recognized by the stakeholders in their ecosystem as being authoritative for the authority statements in the trust registry or registries serving that ecosystem.

Nation states, companies, NGOs, universities, churches, associations, social networks, online communities, and open source projects are all examples of entities who could serve as authorities.

### Authority Statements

An authority statement is a machine-readable representation of a policy governing an entity within the authority’s scope of authority. Trust registries serve as a mechanism for making authority statements accessible to parties who need to make trust decisions regarding those entities.

TRQP supports two basic types of authority statements;

#### Authorization Authority Statements

An [[ref: authorization authority statement]] expresses a hierarchical relationship between the `authority_id` and the `entity_id`. It represents a declaration by an authority that an entity under its jurisdiction or sphere-of-control is authorized to take a specific action on a specific resource.

**Example Authorization Statement (Pseudocode)**
```sh
authority_id: American Association of Motor Vehicle Administrators (AAMVA)
entity_id : Department Of Motor Vehicles (DMV)
action: issue
resource: DriversLicense
```

In English, this corresponds to the statement: "AAMVA has authorized the DMV to issue Drivers Licenses."

#### Recognition Authority Statements

A [[ref: recognition authority statement]] expresses that one authority recognizes the authority of another as a **peer**. Such a recognition relationship may be unilateral or bilateral and is non-exclusive.

::: note

Unlike an [[ref: authorization authority statement]], the authority making a [[ref: recognition authority statement]] is **not** asserting authority over the target authority. Rather it is a referral from one peer to another.

:::

**Example Recognition Statement (Pseudocode)**
```sh
authority_id: France
entity_id : Germany
action: issue
resource: Passport
```

In English, this corresponds to the statement: "France recognizes Germany to issue Passports."

### Governance Frameworks

Just as an authority may take any form, so may the policies governing its ecosystem. For the purposes of this specification, the collection of these policies (whether human-readable and/or machine-readable) is called the [governance framework](https://glossary.trustoverip.org/#term:governance-framework).

To facilitate [trust decisions](https://glossary.trustoverip.org/#term:trust-decision) by its stakeholders—or by any other relying party—the authority is responsible for publishing its governance framework. Although they are not normative requirements of this specification, the following recommendations apply:

1. The governance framework should be published using a [verifiable identifier](https://glossary.trustoverip.org/#term:verifiable-identifier) so its authenticity can be verified.  
2. The governance framework ID should be discoverable via the authority ID.
3. The governance framework should follow the recommendations of the [ToIP Governance Architecture Specification](https://trustoverip.org/wp-content/uploads/ToIP-Governance-Architecture-Specification-V1.0-2021-12-21.pdf) and [ToIP Governance Metamodel Specification](https://trustoverip.org/wp-content/uploads/ToIP-Governance-Metamodel-Specification-V1.0-2021-12-21.pdf).

### Trust Registries

In the context of this specification, a trust registry is a system accessible via a TRQP endpoint that can be queried for the authority statements published by one or more authorities. A trust registry is operated by a trust registry operator. The role of a trust registry operator may be performed directly by an authority or may be delegated to an independent trust registry operator who specializes in this function. In the latter case, from a ToIP governance architecture perspective, the trust registry operator is serving as an [administering authority](https://glossary.trustoverip.org/#term:administering-body).

The TRQP service endpoint for a trust registry may be published in the governance framework or discoverable from the authority\_id as described in the [identifiers section](#identifiers).
