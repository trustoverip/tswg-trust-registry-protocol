
[//]: # (::: forewordtitle)

[//]: # (Foreword)

[//]: # (:::)

[//]: # (\newpage)


## Foreword

ToIP (Trust Over IP Foundation) create a _____ 

::: todo 
Preamble along the lines of an ISO Foreword.
:::

List significant changes (non-normative):

* Shift away from a pure Issuer/Holder/Verifier approach to support non-credential use cases.
* Addition of namespacing concep to begin normalization of trust registries naming conventions.
* Enrichment of registry-of-registry concept to allow for registries that focus primarily on providing a list of registries.

### On Trust, Trustworthy, and Trustworthiness

The term [[ref:trust]] is loaded with varied meanings that often conflict. In the context of [[ref:trust registries]] we need to establish the scope of what we are talking about when we apply the term "trust" to trust registires. There are baseline definitions that follow this limiting scope. 

A trust registry does not create trust. The decision for one entity to "trust" another is their decision. A trust registry may provide information that helps the *consuming party*  in deciding that an entity is [[ref: trustworthy]]. 

::: todo 
  define term "*consuming party*" - OR find better term and capture definition.

### Copyright Notice

This specification is subject to the **OWF Contributor License Agreement 1.0 - Copyright** available at
[https://www.openwebfoundation.org/the-agreements/the-owf-1-0-agreements-granted-claims/owf-contributor-license-agreement-1-0-copyright](https://www.openwebfoundation.org/the-agreements/the-owf-1-0-agreements-granted-claims/owf-contributor-license-agreement-1-0-copyright).

If source code is included in the specification, that code is subject to the Apache 2.0 license unless otherwise marked. In the case of any conflict or confusion within this specification between the OWF Contributor License and the designated source code license, the terms of the OWF Contributor License shall apply.

These terms are inherited from the Technical Stack Working Group at the Trust over IP Foundation. [Working Group Charter](https://trustoverip.org/wp-content/uploads/TSWG-2-Charter-Revision.pdf)


### Terms of Use

These materials are made available under and are subject to the [OWF CLA 1.0 - Copyright & Patent license](https://www.openwebfoundation.org/the-agreements/the-owf-1-0-agreements-granted-claims/owf-contributor-license-agreement-1-0-copyright-and-patent). Any source code is made available under the [Apache 2.0 license](https://www.apache.org/licenses/LICENSE-2.0.txt).

THESE MATERIALS ARE PROVIDED “AS IS.” The Trust Over IP Foundation, established as the Joint Development Foundation Projects, LLC, Trust Over IP Foundation Series ("ToIP"), and its members and contributors (each of ToIP, its members and contributors, a "ToIP Party") expressly disclaim any warranties (express, implied, or otherwise), including implied warranties of merchantability, non-infringement, fitness for a particular purpose, or title, related to the materials. The entire risk as to implementing or otherwise using the materials is assumed by the implementer and user. 

IN NO EVENT WILL ANY ToIP PARTY BE LIABLE TO ANY OTHER PARTY FOR LOST PROFITS OR ANY FORM OF INDIRECT, SPECIAL, INCIDENTAL, OR CONSEQUENTIAL DAMAGES OF ANY CHARACTER FROM ANY CAUSES OF ACTION OF ANY KIND WITH RESPECT TO THESE MATERIALS, ANY DELIVERABLE OR THE ToIP GOVERNING AGREEMENT, WHETHER BASED ON BREACH OF CONTRACT, TORT (INCLUDING NEGLIGENCE), OR OTHERWISE, AND WHETHER OR NOT THE OTHER PARTY HAS BEEN ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

### Introduction

*This section is non-normative*

A [[ref: trust registry]] is a resource that helps to bind governance (business, legal, and social mandates) for an ecosystem. A trust registry helps get the main answers that parties inside and outside of the ecosystem need to tie the governance into their own systems - both technically (it is a protocol) and on a governance (the information provided is created via a governed process).

It is crucially important to understand that a trust registry does not create trust, nor the conditions for trust, by itself. Trust and belief in the data provided by a trust registry is an outcome of governance. 

We need answers to a simple question:

> Does `Entity X` have `Authorization Y`, in the context of `Ecosystem Governance Framework Z`?


[[def: trustworthiness]]
~ An attribute of a person or organization that provides confidence to others of the qualifications, capabilities, and reliability of that entity to perform specific tasks and fulfill assigned responsibilities. Trustworthiness is also a characteristic of information technology products and systems (see Section 2.6.2 on trustworthiness of information systems). The attribute of trustworthiness, whether applied to people, processes, or technologies, can be measured, at least in relative terms if not quantitatively.48 The determination of trustworthiness plays a key role in establishing trust relationships among persons and organizations. The trust relationships are key factors in risk decisions made by senior leaders/executives. NOTE: Current state-of-the-practice for measuring trustworthiness can reliably differentiate between widely different levels of trustworthiness and is capable of producing a trustworthiness scale that is hierarchical between similar instances of measuring activities (e.g., the results from ISO/IEC 15408 [Common Criteria] evaluations). 
- source: [NIST Special Publication 800-39](https://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-39.pdf) p.24
