
[//]: # (::: forewordtitle)

[//]: # (Foreword)

[//]: # (:::)

[//]: # (\newpage)


## Foreword

ToIP (Trust Over IP Foundation) create a _____ 

* TODO: Preamble along the lines of an ISO Foreword.

List significant changes (non-normative):

* Shift away from a pure Issuer/Holder/Verifier approach to support non-credential use cases.
* Addition of namespacing concep to begin normalization of trust registries naming conventions.
* Enrichment of registry-of-registry concept to allow for registries that focus primarily on providing a list of registries.

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