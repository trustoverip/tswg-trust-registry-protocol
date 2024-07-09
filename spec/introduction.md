
[//]: # (Pandoc Formatting Macros)

[//]: # (\mainmatter)

[//]: # (\doctitle)

## Introduction
*This section is non-normative*

A [[ref: trust registry]] is a resource that helps to bind governance (business, legal, and social mandates) for an ecosystem. A trust registry helps get the main answers that parties inside and outside of the ecosystem need to tie the governance into their own systems - both technically and on a governance (the information provided is created via a governed process).

It is crucially important to understand that a trust registry does not create trust, nor the conditions for trust, by itself. Trust and belief in the data provided by a trust registry is an outcome of governance. 

We need answers to a simple question:

> Does `Entity X` have `Authorization Y`, in the context of `Ecosystem Governance Framework Z`?

The Trust Registry Quert Protocol (TRQP) serves to provide a simple suite of adapted interfaces to enable querying of systems of record that provide the information that drives a trust registry. There are a plethora of systems that contain answers that are required to make trust decisions. The protocol is intended to make the communication with any particular system-of-record consistent and simple.

::: todo
**we should explain the "systems of record" content at high level**

** - trust regisry metadata information, type (primary, secondary, RoR), DID, EGF name & VID, assurance levels and metadate about the governance framework
** - and for each trust registrations, each registered entity, registration status, DID, type, etc... and a sample record? 
:::

It is intentionally simple to allow rapid integration into external systems.

The TRQP can:
  * via DNS queries, for DNS enabled/based DIDs/VID/VC, answer basic questions on current status of registered entities and associated credentials information to perform basic technical trust verifications
  * via RDAP queries, answer more advanced questions on configuration of the system relating to registered entities and the registry itself
  * via TRQP API queries (GET protocol), answer all in scope TRQP questions 
  
The TRQP does not:  
  * create a trust registry - it allows (read-only) access to a system-of-record that has the data needed to generate answers that a trust registry provides.
  * create new information - the Create, Update, and Delete of CRUD are not supported. Systems-of-record perform the full CRUD operations. The protocol provides a simple and consistent way of retrieving information from a system.
  * create nor implement governance - the system-of-record that supports the TRQP may have technical ways of doing this, supported by manual operations. Regardless, the TRQP has no opinion on how governance is implemented - just that the information retrieved complies with the stated EGF.
  * make decisions - the TRWP serves up data that are inputs to trust decisions.
  * assign Roles or Rights, though a consuming system may take information that is received via the TRQP and assign these.

It is most crucial to understand that a Trust Registry does NOT create authority. The authority of a trust registry is an outcome of governance.

The purpose of this [[xref: TOIP, ToIP specification]] is to define a standard interoperable protocol for querying a global web of [[xref: TOIP, peer]] [[xref: TOIP, trust registries]], each of which can answer queries about whether a particular [[xref: TOIP, entity]] holds an [[ref:authorization]], in a particular [[xref: TOIP, digital trust ecosystem]] (defined under an [[xref: TOIP, EGF]]), as well as which peer trust registries acknowledge each other.

### Trust Registry Query Protocol features
A core role within the ToIP stack is a [[xref: TOIP, trust registry]]. This is a network service that enables the [[xref:TOIP, governing authority]] for an [[xref: TOIP, EGF]] to share information about their ecosystem. In particular, which [[xref: TOIP, governed parties]] hold which [[ref: authorizations]] under the EGF.

A trust registry query protocol thus should provide the following features:

1. interface to query if a particular [[xref: TOIP, entity]] holds specific [[ref:authorization]] under a defined [[xref: TOIP, EGF]]? 
  - e.g.  "Does entity X hold the authorization of `canada.driver.license.issue` under Canadian Driver's license scheme?" 
2. interface to query what other trust registries are recognized by this trust registry?

### Read-only DNS query protocol

[https://datatracker.ietf.org/doc/draft-carter-high-assurance-dids-with-dns/] provides an outline on how the [[ref:DNS]] improves the authenticity,  discoverability, and portability of Decentralized Identifiers (DIDs) by utilizing the current DNS infrastructure and its technologies. This Internet Draft offers a straightforward procedure for a verifier to cryptographically cross-validate a DID using data stored in the DNS, separate from the DID document.

::: todo
(describe sample of DNS queries/answer and benefits).
DNS records signed with DNSSEC can be verified for authenticity
DNS can support high speed, can be cached and globally accessible to most browsers and applications
DNS read only view into a TR registry
:::

*  Query the entity via DNS to find the primary trust registry and credential information
*  Query the primary trusrt registry to find the registration status of the entity and credential information
*  If both the entity and trust registry concur on the information provided, then the technical trust of the entity can be ascertained.

todo: update https://datatracker.ietf.org/doc/draft-latour-dns-and-digital-trust/ with new language and remove high assurance did mentions.

Per [RFC8552], IANA is requested to add the following entries to the "Underscored and Globally Scoped DNS Node Names" registry:

```
 +---------+--------------------+-----------------------------------------+
 | RR Type | _NODE NAME         | Reference                               |
 +---------+--------------------+-----------------------------------------+
 | TLSA    | _trustregistration | [draft-latour-dns-and-digital-trust-02] |
 | URI     | _trustregistration | [draft-latour-dns-and-digital-trust-02] |
 | TLSA    | _trustregistry     | [draft-latour-dns-and-digital-trust-02] |
 | URI     | _trustregistry     | [draft-latour-dns-and-digital-trust-02] |
 +---------+--------------------+-----------------------------------------+
 ````

### Read-only RDAP query protocol

The Registration Data Access Protocol ([[ref:RDAP]]) is the new whois, it's used for IP addresses, Domain Name, Autonomous Networks, etc. Adding RDAP to trust registries makes the discovery process standard across all layers of the internet infrastructure.

To map the Trust Registry Query Protocol (TRQP) to the RDAP protocol using the /trust handle.

Define the **/trust** handle: We propose to define the use of a new "/trust" handle for the RDAP protocol that can be used to query trust registry information. Work is required to determine the object query structure

This handle would be appended to the RDAP base URL to form a new endpoint for trust registry queries.

Map TRQP operations to RDAP: Each operation in the TRQP (such as querying for a trust registry or verifying membership) would need to be mapped to an equivalent operation in RDAP. This could involve defining new RDAP query types or extending existing ones.

Define the response format: The format of the response returned by the /trust handle would need to be defined. This could be based on the existing RDAP response format, with additional fields added to include the trust registry information.

Please note that this is a high-level overview and the actual implementation and will require additional steps depending on the adoption, and specific requirements of the TRQP and the capabilities of the RDAP server. It’s also important to note that any modifications to the RDAP protocol would need to be in line with the IETF standards.




::: todo 
write IETF I-D on RDAP & Trust Registry Query Protocol support query mapping
::: 


### Read-only query TRQP API Protocol
The primary question (Does `Entity X` have `Authorization Y`, in the context of `Ecosystem Governance Framework Z`) we need an answer to when working in an ecosystem is in itself a simple query. Furthermore, it is read-only query and it doesn't modify any information in a system of record. It just makes data available.

In the web service world the TRQP is purely a GET protocol. 

Just as important it is to understand what the TRQP does NOT do. The TRQP does NOT:
* affect the operations and governance of the systems that support querying using the TRQP.
* create, update, or delete data in a system. In web services this means the TRQP does to PUT, POST, DELETE, and other non-GET operations.

As with all layers of the [[xref: TOIP, ToIP stack]], the purpose of a [[xref: TOIP, ToIP specification]] is to enable the technical interoperability necessary to support transitive trust within and between different [[xref: TOIP, trust communities]] implementing the [[xref: TOIP, ToIP stack]]. In this case, the desired interoperability outcome is a common query protocol that works between any number of decentralized peer trust registries operated by independent governing authorities** representing multiple legal and business jurisdictions.

### Registry of Registries
A Registry of Registries (RoR), is a form of [[xref: TOIP, trust registry]] that primarily serves information about other [[xref: TOIP, trust registries]]. 

1. What other [[xref: TOIP, governing authorities]] are known to the RoR. 
2. Which [[xref: TOIP, trust registry]] are known to be authoritative for particular actions. Examples:
	- Which trust registry is known to issue university diplomas for a particular jurisdiction?
  - Which trust registry is known to manage a list of professionals (e.g. CPAs, lawyers, engineers) that have particular signing rights (authorizations)?
3. Which [[xref: TOIP, trust registry]] are known to operate under a given [[xref: TOIP, EGF]].

The results on a [[xref: TOIP, trust decision]] based on input from a trust registry may range from:
* immediate decision that the entity meets or cannot meet the full requirement of the [[ref:trust relationship]]; or
* further input is required before trust decision can be made. 

These decisions relate to a determination that a relationship is (or is not) sufficiently [[ref: trustworthy]] to establish a [[ref: trust relationship]]. To reach that determination, each party may have its own way of determining the [[ref: trustworthiness]] of their counterparty for the [[ref: trust relationship]] that they require.
