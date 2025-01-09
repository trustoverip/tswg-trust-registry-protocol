# Trust Registry Reference Implementation

This PR provides a super simple ( as simple as possible ) trust registry managed by a git JSON document. It supports
TRQP out of the box, and could easily become a reference implementation for the TRQP as well. 

**Features**

* TRQP API w/ Redoc Frontend
* Entire Registry Managed Over a JSON file
  * Multiple EGF's allowed
  * Organization and Ecosystem Registration Allowed w/ 2 participants; GAN as an organization and ecosystem, and
    Velocity Network, as an organization and ecosystem. 
  *  Sample Namespacing provided.

The registry is described in the `data/registry.json` file. The `registry.json` file is read and then the output is
provided over the TRQP handlers. 

**Design Considerations:** 

The current demo app adds a _lot_ of complexity, and the point of this was to _simplify_ as much as possible. This was particularly important when thinking about the data models. This is intentionally as simple as possible and intended to help us explore the minimum viable implementation for Phase 1.  

**To Use**

1. `go mod tidy`
2. `go run main.go`
