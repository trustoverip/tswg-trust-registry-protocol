## Authority Statements
*This section is normative.*

Authority statements are the machine-readable data stored in trust registries to communicate authoritative information about authorizations, delegations, recognitions, and descriptions (metadata).

### Standard Structure

To enable interoperability of authority statements across ecosystems, TRQP authority statements are structured in three standard parts as shown in Figure&nbsp;3:

![images/authority_statements.png](images/authority_statements.png)

*Figure 3: The standard three-part structure of TRQP authority statements*

These three strings are simple yet flexible enough to express all types of [[ref:authority statements]]. This standard structure should also enable [[ref:TRQP bridges]] to perform deterministic transformations regardless of the structure of the data in an underlying [[ref:system of record]].

### Query Vocabulary

:::note
While this specification defines requirements for the [[ref:authority ID]] and [[ref:entity ID]] components of [[ref: authority statements]], it does not yet specify requirements for the assertion component. The The ToIP Trust Registry Task Force is currently developing standardized query vocabulary to support interoperability of assertions across ecosystems. This query vocabulary may be published in a subsequent draft of this specification or as a separate specification. We invite comments on [the TRQP Query Vocabulary draft currently posted on the ToIP wiki](https://lf-toip.atlassian.net/wiki/spaces/HOME/pages/149749777/TRQP+Query+Vocabulary).
:::
