**Trust Registry Query Protocol (TRQP) Overview**
==================

- **Specification Status:** Draft
- **Version:** 1.2 
- **Status:** Draft v1.2 

**Participate:**
~ [GitHub repo](https://github.com/trustoverip/tswg-trust-registry-protocol/tree/main)
~ [File a bug](https://github.com/trustoverip/tswg-trust-registry-protocol/issues)


::: note
This specification is currently a DRAFT.
:::

## Overview

Many digital ecosystems use intra-ecosystem trust frameworks—for example, OpenID
Federation, X.509 certificate hierarchies, EBSI Trust Chains, or TRAIN—to
confirm whether an entity holds a specific authorization under a given
governance framework. Although these frameworks excel within their own
ecosystems, they often lack native interoperability when verifiers seek to
verify trust and authorization across ecosystems.

The Trust Registry Query Protocol (TRQP) aims to solve this interoperability gap
by acting as an inter-trust framework. It specifies a standardized set of
queries and data models that different ecosystems can implement. This approach
permits verifiers to retrieve information from an external trust registry—even
one based on a completely different internal architecture—without forcing any
ecosystem to abandon or overhaul its existing trust model.

TRQP focuses on three main query types:

1. **Authorization Query:** “Does Entity X hold Authorization Y under Ecosystem
 Z’s governance framework?”
2. **Recognition Query**: “Does Ecosystem A recognize or accept the governance
 framework of Ecosystem B?”
3. **Metadata Query**: “What trust capabilities, data models, or policies does a
 given registry support?”

By standardizing these queries, TRQP ensures that trust verification can be
performed uniformly, even if each ecosystem operates with different governance
logic, cryptographic primitives, and internal policies.
