
[//]: # (::: forewordtitle)

[//]: # (Foreword)

[//]: # (:::)

[//]: # (\newpage)

## Scope

The usefulness of an ecosystem is largely dependant on its ability to assert trust for and between its members. This is true for traditional ecosystems, but even more so with digital ecosystems. With the growing trend on decentralisation of digital services, we are also seeing increased need to transitively assert trust across ecosystems. 

The term [[ref:trust]] is loaded with varied meanings that may conflict. In the context of trust registries we want to be clear what we mean, when we apply the term “trust”. A trust registry does not create trust by itself. The decision for one entity to “trust” another is each party's own decision. The purpose of the trust registry is to provide access to a system of record that contains answers to questions that help drive those trust decisions.

A trust registry may provide information that helps the [[ref:consuming party]] in deciding that an entity is [[ref:trustworthy]]. 
The ToIP Trust Registry Query Protocol helps ecosystems create the foundation of trust within its governed domain, by providing a common protocol for querying information that helps the consuming party in deciding that an entity is [[ref: trustworthy]].

In addition to providing information on its own ecosystem, the Trust Registry Query Protocol (TRQP) enables creation of a registry of registries. This is done by allowing an ecosystem to assert trust to other trust registries, and thus ecosystems. This can be achieved by allowing a governance entity to assert that consuming parties that rely on the trust registry, may also utilize information from another trust registry for additional assertions. This effectively creates transitive trust across ecosystems to achieve wider reach.

In essence, the TRQP establishes a specification for seamless communication across intra-ecosystem trust frameworks. It outlines an abstract data model, ecosystem requirements, and interaction patterns needed to answer the core question:

> Does `Entity X` have `Authorization Y` within the context of `Ecosystem Governance Framework Z`?

This specification does not cover the specific implementation of ecosystem "drivers" to support the TRQP. While these drivers are vital for enabling ecosystem-specific functionality, their implementation details are left to each ecosystem to define according to their unique requirements. 

The TRQP provides a general framework for facilitating inter-ecosystem trust, offering flexibility in how individual ecosystems bridge intra- and inter-trust frameworks. This approach ensures adaptability while promoting consistent communication standards across diverse ecosystems.

## Foreword
This specification is subject to the **OWF Contributor License Agreement 1.0 - Copyright** available at
[https://www.openwebfoundation.org/the-agreements/the-owf-1-0-agreements-granted-claims/owf-contributor-license-agreement-1-0-copyright](https://www.openwebfoundation.org/the-agreements/the-owf-1-0-agreements-granted-claims/owf-contributor-license-agreement-1-0-copyright).

If source code is included in the specification, that code is subject to the Apache 2.0 license unless otherwise marked. In the case of any conflict or confusion within this specification between the OWF Contributor License and the designated source code license, the terms of the OWF Contributor License shall apply.

These terms are inherited from the Technical Stack Working Group at the Trust over IP Foundation. [Working Group Charter](https://trustoverip.org/wp-content/uploads/TSWG-2-Charter-Revision.pdf)

THESE MATERIALS ARE PROVIDED “AS IS.” The Trust Over IP Foundation, established as the Joint Development Foundation Projects, LLC, Trust Over IP Foundation Series ("ToIP"), and its members and contributors (each of ToIP, its members and contributors, a "ToIP Party") expressly disclaim any warranties (express, implied, or otherwise), including implied warranties of merchantability, non-infringement, fitness for a particular purpose, or title, related to the materials. The entire risk as to implementing or otherwise using the materials is assumed by the implementer and user. 

IN NO EVENT WILL ANY ToIP PARTY BE LIABLE TO ANY OTHER PARTY FOR LOST PROFITS OR ANY FORM OF INDIRECT, SPECIAL, INCIDENTAL, OR CONSEQUENTIAL DAMAGES OF ANY CHARACTER FROM ANY CAUSES OF ACTION OF ANY KIND WITH RESPECT TO THESE MATERIALS, ANY DELIVERABLE OR THE ToIP GOVERNING AGREEMENT, WHETHER BASED ON BREACH OF CONTRACT, TORT (INCLUDING NEGLIGENCE), OR OTHERWISE, AND WHETHER OR NOT THE OTHER PARTY HAS BEEN ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 
### Conventions
The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED", "NOT RECOMMENDED", "MAY", and "OPTIONAL" in this document are to be interpreted as described in BCP 14 [RFC2119] [RFC8174] when, and only when, they appear in all capitals, as shown here.


