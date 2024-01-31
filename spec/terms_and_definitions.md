
[//]: # (Pandoc Formatting Macros)

[//]: # (Portable Document Format)

[//]: # (blank)

[//]: # (: file format defined by ISO 32000-2)



## Terms & Definitions


The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED", "NOT RECOMMENDED", "MAY", and "OPTIONAL" in this document are to be interpreted as described in [[spec-inform:RFC2119]] when, and only when, they appear in all capitals, as shown here.

[[def: assurance levels]]
~ TODO: 

[[def: authentication]] (copied from ToIP Glossary)
~ Verifying the identity of a user, process, or device, often  as a prerequisite to allowing access to resources in an information system. source: [NIST Special Publication 800-39](https://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-39.pdf)

[[def: authenticity]] (copied from ToIP Glossary)
~ The property of being genuine and being able to be verified and trusted; confidence in the validity of a transmission, a message, or message originator. source: [NIST Special Publication 800-39](https://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-39.pdf)



[[def: authorization]]
~ Access privileges granted to an entity; conveys an “official” sanction to perform a cryptographic function or other sensitive activity.
* source: [NIST](https://csrc.nist.gov/glossary/term/permission) NIST SP 800-57 Part 2 Rev.1 under Authorization
::: issue 
https://github.com/trustoverip/tswg-trust-registry-protocol/issues/6
- May need a `governed authorization` term to help link tech+governance.
:::

[[def:authorized trust registries]]
~ The primary trust registry plus all secondary trust registries are collectively referred to as the authorized trust registries.

[[def: action]]
~ a discrete property (string) that an entity can be authorized for, in the form of a [permission](https://trustoverip.github.io/ctwg-main-glossary/#term:permission) response.

[[def: action namespace]]
~ A well-known string that is used in an EGF to indicate a discrete authorization. Examples (non-exhaustive): "canada:driver-license", "eu:trusted-list.authorized-timestamp", "global:tsm"

[[def: ecosystem governance framework, ecosystem governance frameworks, EGF]]
~ TODO: replace this ChatGPT definiton: refers to a structured set of principles, rules, and mechanisms that guide and regulate the management and decision-making processes within an ecosystem. Ecosystem governance is typically associated with natural or environmental systems, where various stakeholders, such as governments, communities, businesses, and non-governmental organizations, work together to sustainably manage and protect ecosystems.

[[def: registered entity]]
~ An entity that is listed in the system (i.e. the [[ref: trust registry]]) that is being queried. 

[[def: permission]]
~  Authorization to perform some action on a system.

* source: [NIST](https://csrc.nist.gov/glossary/term/permission)

[[def: primary trust registry]]
~ TODO:

[[def:secondary trust registry, secondary trust registries]]
~ TODO: 

[[def: trust decision]]
~ A decision that a party needs to make about whether to engage in a specific interaction or transaction with another entity that involves real or perceived risks. source: [ToIP Glossary](https://docs.google.com/document/d/1fZByfuSOwszDRkE7ARQLeElSYmVznoOyJK4sxRvJpyM/edit#heading=h.m8c86ccqis9r)

[[def: trust list]]
~ A one-dimensional trust graph in which an authoritative source publishes a list of entities that are trusted in a specific trust context. A trust list can be considered a simplified form of a trust registry.

[[def: trust registry, trust registries]] 
~ A registry that serves as an **authoritative source** for **trust graphs** or other **governed information** describing one or more **trust communities**. A trust registry is typically **authorized** by a **governance framework**.  See also: trust list

[[def: trusted party]]
~ A party that is trusted by an entity to faithfully perform certain services for that entity. An entity may choose to act as a trusted party for itself.
- source: [NIST SP 800-56B Rev. 2](https://doi.org/10.6028/NIST.SP.800-56Br2) under Trusted party

[[def: VID Type]]
~ TODO: 

