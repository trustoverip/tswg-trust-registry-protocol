
[//]: # (Pandoc Formatting Macros)

[//]: # (\mainmatter)

[//]: # (\doctitle)

## Scope

The Trust Registry Protocol (TRP) serves to provide a simple interface to enable querying of systems of record that provide the information that drives a trust registry. There are a plethora of systems that contain answers that are required to make trust decisions. The protocol is intended to make the communication with any particular system-of-record consistent and simple.

It is intentionally simple to allow rapid integration into external systems.

The TRP does not:  
  * create a trust registry - it allows (read-only) access to a system-of-record that has the data needed to generate answers that a trust registry provides.
  * create new information - the Create, Update, and Delete of CRUD are not supported. Systems-of-record perform the full CRUD operations. The protocol provides a simple and consistent way of retrieving information from a system.
  * create nor implement governance - the system-of-record that supports the TRP may have technical ways of doing this, supported by manual operations. Regardless, the TRP has no opinion on how governance is implemented - just that the information retrieved complies with the stated EGF.
  * make decisions - the TRP serves up data that are inputs to trust decisions.
  * assign Roles or Rights, though a consuming system may take information that is received via the TRP and assign these.

It is most crucial to understand that a Trust Registry does NOT create authority. As Jacques Latour says "the authority of a trust registry is an outcome of governance". 

### Purpose

The purpose of this [[xref: TOIP, ToIP specification]] is to define a standard interoperable protocol for querying a global web of [[xref: TOIP, peer]] [[xref: TOIP, trust registries]], each of which can answer queries about whether a particular [[xref: TOIP, entity]] holds an [[ref:authorization]], in a particular [[xref: TOIP, digital trust ecosystem]] (defined under an [[xref: TOIP, EGF]]), as well as which peer trust registries acknowledge each other.

The Trust Registry Protocol enables an interested party to ask the following question:

> Does `Entity X` have `Authorization Y`, in the context of `Ecosystem Governance Framework Z`?




### Motivations

A core role within the ToIP stack is a [[xref: TOIP, trust registry]]. This is a network service that enables the [[xref:TOIP, governing authority]] for an [[xref: TOIP, EGF]] to share information about their ecosystem. In particular, which [[xref: TOIP, governed parties]] hold which [[ref: authorizations]] under the EGF. For example:

1.  Does a particular [[xref: TOIP, entity]] hold a particular [[ref:authorization]] under an EGF? 
  - e.g.  "Does entity X hold the authorization of `canada.driver.license.issue` under an EGF? - notionaly equating to the authority to "issue" a "driver license in Canada"; 
2. What other trust registries are recognized by this trust registry?

The Trust Registry Protocol is focused on answering those questions.


### Querying Consistently Requires a (Read-Only) Protocol

One primary question we need an answer to when working in an ecosystem is (as stated throughout this specification):

> Does `Entity X` have `Authorization Y`, in the context of `Ecosystem Governance Framework Z`?


This is a query - a simple one. Further, it is read-only. It doesn't modify any information in a system of record. It just makes data available.

In the web service world the TRP is purely a GET protocol. 

Just as important, is to understand what the Trust Registry Protocol does NOT do. The TRP does NOT:

* affect the operations and governance of the systems that support querying using the TRP.
* create, update, or delete data in a system. In web services this means the TRP does to PUT, POST, DELETE, and other non-GET operations.

### Querying a Trust Registry

As with all layers of the [[xref: TOIP, ToIP stack]], the purpose of a [[xref: TOIP, ToIP specification]] is to enable the technical interoperability necessary to support transitive trust within and between different [[xref: TOIP, trust communities]] implementing the [[xref: TOIP, ToIP stack]]. In this case, the desired interoperability outcome is a common query protocol that works between any number of decentralized peer trust registries operated by independent governing authorities** representing multiple legal and business jurisdictions.

### Registry of Registry (MetaRegistry)

A Registry of Registries (RoR), is a form of [[xref: TOIP, trust registry]] that primarily serves information about other [[xref: TOIP, trust registries]]. 

1. What other [[xref: TOIP, governing authorities]] are known to the RoR. 
2. Which [[xref: TOIP, trust registry]] are known to be authoritative for particular actions. Examples:
	- Which trust registry is known to issue university diplomas for a particular jurisdiction?
  - Which trust registry is known to manage a list of professionals (e.g. CPAs, lawyers, engineers) that have particular signing rights (authorizations)?
3. Which [[xref: TOIP, trust registry]] are known to operate under a given [[xref: TOIP, EGF]].


### On Trust, Trustworthy, and Trustworthiness

The term [[ref:trust]] is loaded with varied meanings that often conflict. In the context of [[ref:trust registries]] we need to establish the scope of what we are talking about when we apply the term "trust" to trust registires. There are baseline definitions that follow this limiting scope. 

A trust registry does not create trust. The decision for one entity to "trust" another is their decision. A trust registry may provide information that helps the *consuming party*  in deciding that an entity is [[ref: trustworthy]]. 

::: todo 
  define term "*consuming party*" - OR find better term and capture definition.
:::

The results on a [[xref: TOIP, trust decision]] based on input from a trust registry may range from:
* immediate decision that the entity meets or cannot meet the full requirement of the [[ref:trust relationship]]; or
* further input is required before trust decision can be made. 

These decisions relate to a determination that a relationship is (or is not) sufficiently [[ref: trustworthy]] to establish a [[ref: trust relationship]]. To reach that determination, each party may have its own way of determining the [[ref: trustworthiness]] of their counterparty for the [[ref: trust relationship]] that they require.

The following terms are presented to help create a general understanding and may be only indirectly related to trust registry efforts:

[[def: trust]]
~ A belief that an entity will behave in a predictable manner in specified circumstances. The entity may be a person, process, object or any combination of such components. The entity can be of any size from a single hardware component or software module, to a piece of equipment identified by make and model, to a site or location, to an organization, to a nation-state. Trust, while inherently a subjective determination, can be based on objective evidence and subjective elements. The objective grounds for trust can include for example, the results of information technology product testing and evaluation. Subjective belief, level of comfort, and experience may supplement (or even replace) objective evidence, or substitute for such evidence when it is unavailable. Trust is usually relative to a specific circumstance or situation (e.g., the amount of money involved in a transaction, the sensitivity or criticality of information, or whether safety is an issue with human lives at stake). Trust is generally not transitive (e.g., you trust a friend but not necessarily a friend of a friend). Finally, trust is generally earned, based on experience or measurement.
- source: [NIST Special Publication 800-39](https://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-39.pdf) p.24

[[def: trust relationship]]
~ An agreed upon relationship between two or more system elements that is governed by criteria for secure interaction, behavior, and outcomes relative to the protection of assets.
- source: [NIST SP 800-160v1r1](https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-160v1r1.pdf)

[[def: trustworthy]]
~ Worthy of the confidence to others of the qualifications, capabilities, and reliability of that entity to perform specific tasks and fulfill assigned responsibilities. (note: based on the definition of [[ref: trustworthiness]]. note: from source "This refers to trust relationships between system elements implemented by hardware, firmware, and software" but the definition largely works.

[[def: trustworthiness]]
~ An attribute of a person or organization that provides confidence to others of the qualifications, capabilities, and reliability of that entity to perform specific tasks and fulfill assigned responsibilities. Trustworthiness is also a characteristic of information technology products and systems (see Section 2.6.2 on trustworthiness of information systems). The attribute of trustworthiness, whether applied to people, processes, or technologies, can be measured, at least in relative terms if not quantitatively.48 The determination of trustworthiness plays a key role in establishing trust relationships among persons and organizations. The trust relationships are key factors in risk decisions made by senior leaders/executives. NOTE: Current state-of-the-practice for measuring trustworthiness can reliably differentiate between widely different levels of trustworthiness and is capable of producing a trustworthiness scale that is hierarchical between similar instances of measuring activities (e.g., the results from ISO/IEC 15408 [Common Criteria] evaluations). 
- source: [NIST Special Publication 800-39](https://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-39.pdf) p.24
