
[//]: # (Pandoc Formatting Macros)

[//]: # (\mainmatter)

[//]: # (\doctitle)

## Introduction
*This section is non-normative*

A [[ref: trust registry]] is a resource that helps to bind governance (business, legal, and social mandates) for an ecosystem. A trust registry helps get the main answers that parties inside and outside of the ecosystem need to tie the governance into their own systems - both technically and on a governance (the information provided is created via a governed process).

It is crucially important to understand that a trust registry does not create trust, nor the conditions for trust, by itself. Trust and belief in the data provided by a trust registry is an outcome of governance. 

We need answers to a simple question:

> Does `Entity X` have `Authorization Y`, in the context of `Ecosystem Governance Framework Z`?

The Trust Registry Protocol (TRP) serves to provide a simple interface to enable querying of systems of record that provide the information that drives a trust registry. There are a plethora of systems that contain answers that are required to make trust decisions. The protocol is intended to make the communication with any particular system-of-record consistent and simple.

It is intentionally simple to allow rapid integration into external systems.

The TRP does not:  
  * create a trust registry - it allows (read-only) access to a system-of-record that has the data needed to generate answers that a trust registry provides.
  * create new information - the Create, Update, and Delete of CRUD are not supported. Systems-of-record perform the full CRUD operations. The protocol provides a simple and consistent way of retrieving information from a system.
  * create nor implement governance - the system-of-record that supports the TRP may have technical ways of doing this, supported by manual operations. Regardless, the TRP has no opinion on how governance is implemented - just that the information retrieved complies with the stated EGF.
  * make decisions - the TRP serves up data that are inputs to trust decisions.
  * assign Roles or Rights, though a consuming system may take information that is received via the TRP and assign these.

It is most crucial to understand that a Trust Registry does NOT create authority. The authority of a trust registry is an outcome of governance.

The purpose of this [[xref: TOIP, ToIP specification]] is to define a standard interoperable protocol for querying a global web of [[xref: TOIP, peer]] [[xref: TOIP, trust registries]], each of which can answer queries about whether a particular [[xref: TOIP, entity]] holds an [[ref:authorization]], in a particular [[xref: TOIP, digital trust ecosystem]] (defined under an [[xref: TOIP, EGF]]), as well as which peer trust registries acknowledge each other.

### Trust Registry Protocol features
A core role within the ToIP stack is a [[xref: TOIP, trust registry]]. This is a network service that enables the [[xref:TOIP, governing authority]] for an [[xref: TOIP, EGF]] to share information about their ecosystem. In particular, which [[xref: TOIP, governed parties]] hold which [[ref: authorizations]] under the EGF.

A trust registry protocol thus should provide the following features:

1. interface to query if a particular [[xref: TOIP, entity]] holds specific [[ref:authorization]] under a defined [[xref: TOIP, EGF]]? 
  - e.g.  "Does entity X hold the authorization of `canada.driver.license.issue` under Canadian Driver's license scheme?" 
2. interface to query what other trust registries are recognized by this trust registry?

### Read-only query Protocol
The primary question (Does `Entity X` have `Authorization Y`, in the context of `Ecosystem Governance Framework Z`) we need an answer to when working in an ecosystem is in itself a simple query. Furthermore, it is read-only query and it doesn't modify any information in a system of record. It just makes data available.

In the web service world the TRP is purely a GET protocol. 

Just as important it is to understand what the Trust Registry Protocol does NOT do. The TRP does NOT:
* affect the operations and governance of the systems that support querying using the TRP.
* create, update, or delete data in a system. In web services this means the TRP does no PUT, POST, DELETE, and other non-GET operations.

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
