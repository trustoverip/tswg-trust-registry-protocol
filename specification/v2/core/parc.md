
## The PARC Information Model

The [PARC authorization model](https://docs.cedarpolicy.com/auth/authorization.html) is a framework for making authority statements using four core components: Principal, Action, Resource, and Context. In English, this corresponds to the question:

“Can this principal take this action on this resource in this context?”. 

The simplicity and flexibility of the PARC model, which is the basis for some of the largest authorization systems in use today, allows for expressing many types of authorization policies including Attribute-Based Access Control (ABAC), Role-Based Access Control (RBAC), and Relationship-Based Access Control (ReBAC). 

### Principal

The principal is the entity that is the object of the authority statement, i.e., the entity whose actions the authority statement permits or denies. In the general PARC model, the principal could be a user, a group, a role, or another service or API.

In the context of TRQP, a principal is usually:

* An [issuer](https://glossary.trustoverip.org/#term:issuer) of credentials.  
* A [verifier](https://glossary.trustoverip.org/#term:verifier) of credentials.  
* A trust registry.  
* Another authority.

### Action

The action is the specific operation or verb the principal is permitted to perform. When the PARC model is applied to access control, examples include read, write, delete, or view.

In the context of TRQP authorization statements, example action values might be:

* issue  
* verify  
* revoke

In the context of TRQP recognition and delegation statements, where the principal is another authority, example action values might be:

* recognize  
* parent-of  
* child-of

### Resource

The resource is the object, data, or entity that the principal is permitted to act upon. With PARC-based access control, typical examples are a file, a database entry, or an API.

In the context of TRQP authorization statements, example resources might be:

* credential type IDs  
* trust registry IDs

::: note
In the context of TRQP recognition and delegation statements, the resource would be the governance framework ID or ecosystem ID being recognized or being delegated to/from. The Resource (and Action) values indicated must align with the delegation approach used in the governance framework. 
:::

### Context

The context specifies constraints or conditions that must be met for the policy to apply. With PARC-based access control, typical examples might be the time of day, the network IP address, the principal's group membership, or attributes of the resource itself.

When the PARC model is used within the scope of a single authorization system (such as for access control within a company), context is optional because the authority for the policies can be assumed. The opposite is true of TRQP: as a decentralized trust registry query protocol, context must be supplied in the form of the authority ID of the authority responsible for the authority statements being queried. 

A TRQP query may also include other context parameters, such as time, as specified in section TODO.

#### Context Objects

*This section is normative.*

In the PARC model, a context object is used to assert conditions that must be satisfied in order for an action to be permitted. As defined in section TODO, a TRQP query the `authority_id` is the primary context. Further refinement of the context is possible using an optional context parameter.

::: warning
the use of `authority_id` above is confusing. 

the requirements here are better handled in the Identifiers area?
:::

Other than this required parameter, in a TRQP query:

1. A context object is OPTIONAL.  
2. If included, a context object MUST be a JSON object whose members convey other conditions.   
3. If a context object includes a time value: a time-based condition:  
   1. The time parameter MUST be expressed using a time member whose value is in [RFC 3339](https://trustoverip.github.io/tswg-trust-registry-protocol/%7B%7BRFC3339%7D%7D) format.  
   2. The value of the time parameter MUST be interpreted as the datetime as of which the target authority statement is valid.  
4. TODO: DO WE NEED TO PUT REQUIREMENTS IN FOR locator?  
5. Additional JSON object members specifying other conditions MAY be defined by TRQP profiles or bindings.
