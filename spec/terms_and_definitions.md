
[//]: # (Pandoc Formatting Macros)

[//]: # (Portable Document Format)

[//]: # (blank)

[//]: # (: file format defined by ISO 32000-2)



## Terms & Definitions


The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED", "NOT RECOMMENDED", "MAY", and "OPTIONAL" in this document are to be interpreted as described in [[spec-inform:RFC2119]] when, and only when, they appear in all capitals, as shown here.

[[def: authorization, authorizations]]
~ Access privileges granted to an entity; conveys an “official” sanction to perform a cryptographic function or other sensitive activity.
* source: [NIST](https://csrc.nist.gov/glossary/term/permission) NIST SP 800-57 Part 2 Rev.1 under Authorization
::: issue 
https://github.com/trustoverip/tswg-trust-registry-protocol/issues/6
- May need a `governed authorization` term to help link tech+governance.
:::

[[def: authorized trust registry, authorized trust registries]]
~ The primary trust registry plus all secondary trust registries are collectively referred to as the authorized trust registries.


[[def: authorization namespace]]
~ A well-known string that is used in an EGF to indicate a discrete authorization. Examples (non-exhaustive): "canada:driver-license", "eu:trusted-list.authorized-timestamp", "global:tsm"


[[def: registered entity, registered entities]]:
~ An [[xref: TOIP, entity]] that is listed in the system (i.e. the [[xref: TOIP, trust registry]]) that is being queried. 

[[def: permission]]:
~  Authorization to perform some action on a system.

* source: [NIST](https://csrc.nist.gov/glossary/term/permission)

[[def: primary trust registry]]
~ TODO:

[[def:secondary trust registry, secondary trust registries]]:
~ TODO: 

[[def: service endpoint]]:
~ TODO: 

[[def: service property]]:
~ TODO: 



[[def: trust registry, trust registries]]: 
~ A registry that serves as an authoritative source for trust graphs or other governed information describing one or more trust communities. A trust registry is typically authorized by a governance framework.  
* See also: [[xref: TOIP, trust list]]

[[def: trusted party]]
~ A party that is trusted by an entity to faithfully perform certain services for that entity. An entity may choose to act as a trusted party for itself.
- source: [NIST SP 800-56B Rev. 2](https://doi.org/10.6028/NIST.SP.800-56Br2) under Trusted party

[[def: VID Type, VID Types]]
~ A specific kind of [[xref: TOIP, VID]].

