## High-Level Architecture

*This section is informative.*

This section presents the major conceptual components in TRQP architecture.

### Authorities

Authorities are the parties responsible for establishing the policies governing their ecosystem. An authority may take any legal form or may not be a formal legal entity at all. The only requirement is that the authority be recognized by the stakeholders in their ecosystem as being authoritative for the authority statements in the trust registry or registries serving that ecosystem.

Nation states, companies, NGOs, universities, churches, associations, social networks, online communities, and open source projects are all examples of entities who could serve as authorities.

### Authority Statements

An authority statement is a machine-readable representation of a policy governing an entity within the authority’s scope of authority. Trust registries serve as a mechanism for making authority statements accessible to parties who need to make trust decisions regarding those entities.

See the example policies in sections TODO for examples of authority statements.

### Governance Frameworks

Just as an authority may take any form, so may the policies governing its ecosystem. For the purposes of this specification, the collection of these policies (whether human-readable and/or machine-readable) is called the [governance framework](https://glossary.trustoverip.org/#term:governance-framework).

To facilitate [trust decisions](https://glossary.trustoverip.org/#term:trust-decision) by its stakeholders—or by any other relying party—the authority is responsible for publishing its governance framework. Although they are not normative requirements of this specification, the following recommendations apply:

1. The governance framework should be published using a [verifiable identifier](https://glossary.trustoverip.org/#term:verifiable-identifier) so its authenticity can be verified.  
2. The governance framework ID should be discoverable via the authority ID (section TODO).  
3. The governance framework should follow the recommendations of the [ToIP Governance Architecture Specification](https://trustoverip.org/wp-content/uploads/ToIP-Governance-Architecture-Specification-V1.0-2021-12-21.pdf) and [ToIP Governance Metamodel Specification](https://trustoverip.org/wp-content/uploads/ToIP-Governance-Metamodel-Specification-V1.0-2021-12-21.pdf).

### Trust Registries

In the context of this specification, a trust registry is a system accessible via a TRQP endpoint that can be queried for the authority statements published by one or more authorities. A trust registry is operated by a trust registry operator. The role of a trust registry operator may be performed directly by an authority or may be delegated to an independent trust registry operator who specializes in this function. In the latter case, from a ToIP governance architecture perspective, the trust registry operator is serving as an [administering authority](https://glossary.trustoverip.org/#term:administering-body).

The TRQP service endpoint for a trust registry may be published in the governance framework or discoverable from the authority\_id as described in section TODO.

### PARC Model

Our information model was heavily inspired by the [PARC model](https://docs.cedarpolicy.com/auth/authorization.html). The identifiers that follow map well to PARC:
* Principal - `entity_id`
* Action - `action`
* Resource - `resource`
* Context - `authority_id` is the mandatory context, and an optional `context` object is available for further refinement.



