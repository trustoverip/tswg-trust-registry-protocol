# Trust Over IP Trust Registry Query Protocol Specification

This specification is basked on the [Trust Over IP Specification
Template](https://github.com/trustoverip/specification-template).

The spec is written using [SpecUp](https://github.com/decentralized-identity/spec-up) which is maintained by the
Decentralized Identity Foundation. 


To browse the spec, see the [rendering on GitHub pages](https://trustoverip.github.io/tswg-trust-registry-protocol/). To
contribute to the spec, submit PRs that modify the .md files (in the `./spec` folder) that are used to generate the
.html files in this folder.

Before submitting a PR, please see the [Editing The Spec](./docs/EditingTheSpec.md) document for guidance on generating the
specification locally for review.

## Repository Structure

# Folder Structure

## `specification`
This folder contains the core specification documentation. The rendered version is available at  
[https://trustoverip.github.io/tswg-trust-registry-protocol/](https://trustoverip.github.io/tswg-trust-registry-protocol/).

## `docs`
This folder includes user guides and other supporting documentation to enhance usability and understanding.

## `profiles`
This folder provides a registry of TRQP profiles. These profiles are intended to demonstrate how a profile might be structured for learning purposes. They are not intended to represent definitive or production-ready profiles.

## `reference_implementation`
This folder contains a simple reference implementation using a JSON file as the registry. It serves as a teaching tool and is not designed for production environments.

## Rendering Spec-Up

To run Spec-up in live edit mode (will re-render upon save), in project folder run:

```
npm run edit
```
