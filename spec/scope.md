
[//]: # (Pandoc Formatting Macros)

[//]: # (\mainmatter)

[//]: # (\doctitle)

## Scope

The Trust Registry Protocol serves to provide a simple interface to enable access to systems of record that provide the information that drives a trust registry. There are a plethora of systems that contain answers that are required to make trust decisions. The protocol is intended to make the communication with any particular system-of-record consistent and simple.

The TRP does not:  
  * create a trust registry - it allows (read-only) access to a system-of-record that has the data needed to generate answers that a trust registry provides.
  * create new information - the Create, Update, and Delete of CRUD are not supported. Systems-of-record perform the full CRUD operations. The protocol provides a simple and consistent way of retrieving information from a system.
  * create nor implement governance - the system-of-record that supports the TRP may have technical ways of doing this, supported by manual operations. Regardless, the TRP has no opinion on how governance is implemented - just that the information retrieved complies with the stated EGF.
  * make decisions - the TRP serves up data that are inputs to trust decisions.
  * assign Roles or Rights, though a consuming system may take information that is received via the TRP and assign these.

It is most crucial to understand that a Trust Registry does NOT create authority. As Jacques Latour says "the authority of a trust registry is an outcome of governance". 

### Purpose

The purpose of this [[xref: TOIP, ToIP specification]] is to define a standard interoperable protocol for interacting with a global web of [[xref: TOIP, peer]] [[xref: TOIP, trust registries]], each of which can answer queries about whether a particular [[xref: TOIP, entity]] holds an [[ref:authorization]], in a particular [[xref: TOIP, digital trust ecosystem]] (defined under an [[xref: TOIP, EGF]]), as well as which peer trust registries acknowledge each other.

### Motivations

A core role within the ToIP stack is a [[xref: TOIP, trust registry]]. This is a network service that enables the [[xref:TOIP, governing authority]] for an [[xref: TOIP, EGF]] to specify which [[xref: TOIP, governed parties]] hold which [[ref: authorizations]] under the EGF. For example:

1. Which [[xref: TOIP, entities]] hold an [[ref:authorization]] under an EGF? 
  - e.g.  "Does entity X hold the authorization of `canada.driver.license.issue`" - equating to the authority to "issue" a "driver license in Canada"; 
2. What other trust registries are recognized by this particular trust registry?

As with all layers of the [[xref: TOIP, ToIP stack]], the purpose of a [[xref: TOIP, ToIP specification]] is to enable the technical interoperability necessary to support transitive trust within and between different [[xref: TOIP, trust communities]] implementing the [[xref: TOIP, ToIP stack]]. In this case, the desired interoperability outcome is a common protocol that works between any number of decentralized peer trust registries operated by independent governing authorities** representing multiple legal and business jurisdictions. One specific example of this need is the digital trust ecosystem defined by the [Interoperability Working Group for Good Health Pass](https://wiki.trustoverip.org/pages/viewpage.action?pageId=73790) (GHP). 

A Registry of Registries (RoR), is a form of [[xref: TOIP, trust registry]] that primarily serves information about other [[xref: TOIP, trust registries]]. 

1. What other [[xref: TOIP, governing authorities]] are known to the RoR. 
2. Which [[xref: TOIP, trust registry]] are known to be authoritative for particular actions. Examples:
	- Which trust registry is known to issue university diplomas for a particular jurisdiction?
  - Which trust registry is known to manage a list of professionals (e.g. CPAs, lawyers, engineers) that have particular signing rights (authorizations)?
3. Which [[xref: TOIP, trust registry]] are known to operate under a given [[xref: TOIP, EGF]].


