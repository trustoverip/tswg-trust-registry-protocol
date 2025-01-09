
[//]: # (Pandoc Formatting Macros)

[//]: # (Portable Document Format)

[//]: # (blank)

[//]: # (: file format defined by ISO 32000-2)

## Terms & Definitions
The following terms are used to describe concepts in this specification.

[[def: authorization, authorizations]]:
~ Access privileges granted to an entity; conveys an “official” sanction to perform a cryptographic function or other sensitive activity. (Source: [NIST](https://csrc.nist.gov/glossary/term/permission) NIST SP 800-57 Part 2 Rev.1 under Authorization)

::: issue 
https://github.com/trustoverip/tswg-trust-registry-protocol/issues/6
- May need a `governed authorization` term to help link tech+governance.
:::

[[def: authorized trust registry, authorized trust registries]]
~ The primary trust registry plus all secondary trust registries are collectively referred to as the authorized trust registries.

[[def: authorization namespace]]:
~ A well-known string that is used in an EGF to indicate a discrete authorization. Examples (non-exhaustive): "canada:driver-license", "eu:trusted-list.authorized-timestamp", "global:tsm"

[[def: consuming party, consuming parties]]:
~ A party that consumes the services and information provided by a [[xref: TOIP, trust registry]] in order to make a trust decision.

[[def: registered entity, registered entities]]:
~ An [[xref: TOIP, entity]] that is listed in the system (i.e. the [[xref: TOIP, trust registry]]) that is being queried. 

[[def: permission]]:
~  Authorization to perform some action on a system. (Source: [NIST](https://csrc.nist.gov/glossary/term/permission))

[[def: primary trust registry]]:
~ The single [[xref: TOIP, trust registry]] that is considered the primary source for information of a particular type in an ecosystem.

[[def:secondary trust registry, secondary trust registries]]:
~ A trust registry that has copies of information based on the ecosystem's [[def:primary trust registry]]. 

[[def: service endpoint]]:
~ A network address, such as an HTTP URL, at which services operate on behalf of a DID subject. (Source: [[spec-norm:DID-CORE]])

[[def: service property]]:
~ in context of: [TRQP-1] ...MUST publish, in the [[xref: TOIP, DID document]] associated with the **DID** identifying its **EGF**, a [[ref: service property]] specifying the [[ref: service endpoint]]

[[def: trust registry, trust registries]]: 
~ A registry that serves as an authoritative source for trust graphs or other governed information describing one or more trust communities. A trust registry is typically authorized by a governance framework.  (See also: [[xref: TOIP, trust list]])

[[def: trust]]
~ A belief that an entity will behave in a predictable manner in specified circumstances. The entity may be a person, process, object or any combination of such components. The entity can be of any size from a single hardware component or software module, to a piece of equipment identified by make and model, to a site or location, to an organization, to a nation-state. Trust, while inherently a subjective determination, can be based on objective evidence and subjective elements. The objective grounds for trust can include for example, the results of information technology product testing and evaluation. Subjective belief, level of comfort, and experience may supplement (or even replace) objective evidence, or substitute for such evidence when it is unavailable. Trust is usually relative to a specific circumstance or situation (e.g., the amount of money involved in a transaction, the sensitivity or criticality of information, or whether safety is an issue with human lives at stake). Trust is generally not transitive (e.g., you trust a friend but not necessarily a friend of a friend). Finally, trust is generally earned, based on experience or measurement. (source: [NIST Special Publication 800-39](https://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-39.pdf) p.24)

[[def: trust relationship]]
~ An agreed upon relationship between two or more system elements that is governed by criteria for secure interaction, behavior, and outcomes relative to the protection of assets. (source: [NIST SP 800-160v1r1](https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-160v1r1.pdf))

[[def: trustworthy]]
~ Worthy of the confidence to others of the qualifications, capabilities, and reliability of that entity to perform specific tasks and fulfill assigned responsibilities. (note: based on the definition of [[ref: trustworthiness]]. note: from source "This refers to trust relationships between system elements implemented by hardware, firmware, and software" but the definition largely works.

[[def: trustworthiness]]
~ An attribute of a person or organization that provides confidence to others of the qualifications, capabilities, and reliability of that entity to perform specific tasks and fulfill assigned responsibilities. Trustworthiness is also a characteristic of information technology products and systems (see Section 2.6.2 on trustworthiness of information systems). The attribute of trustworthiness, whether applied to people, processes, or technologies, can be measured, at least in relative terms if not quantitatively.48 The determination of trustworthiness plays a key role in establishing trust relationships among persons and organizations. The trust relationships are key factors in risk decisions made by senior leaders/executives. NOTE: Current state-of-the-practice for measuring trustworthiness can reliably differentiate between widely different levels of trustworthiness and is capable of producing a trustworthiness scale that is hierarchical between similar instances of measuring activities (e.g., the results from ISO/IEC 15408 [Common Criteria] evaluations). (source: [NIST Special Publication 800-39](https://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-39.pdf) p.24)


[[def: trusted party]]:
~ A party that is trusted by an entity to faithfully perform certain services for that entity. An entity may choose to act as a trusted party for itself.(source: [NIST SP 800-56B Rev. 2](https://doi.org/10.6028/NIST.SP.800-56Br2) under Trusted party)

[[def: VID Type, VID Types]]:
~ A specific kind of [[xref: TOIP, VID]].


