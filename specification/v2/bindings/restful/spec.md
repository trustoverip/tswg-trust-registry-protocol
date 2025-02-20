# **RESTful Binding**

RESTful TRQP Bindings specification that implements the core specification. 

The following [OpenAPI](./v2.yaml) Document describes the RESTful endpoints that are required for the TRQP RESTful binding. 

* The `/metadata` endpoint is aligned to the MetadataQuery.
* The `/registries/{ecosystem_id}/recognition` maps to the RecognitionQuery. 
    * The `/entities/{entity_id}/authorization` maps to the AuthorizationQuery. 

Additional Information:

* Error Codes are represented using [Problem Details described in rfc7807](https://datatracker.ietf.org/doc/html/rfc7807).
* Authorization and Recognition Queries both take timestamps as a required parameter to resolve.
* Time parameters are in the form [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) and MUST be sent in UTC time. 
* If `jws` field in response is provided, verifiers are recommended to use that to verify the response payload controller.

Security considerations are left to the implementation profile to describe. 
Identifier requirements are left to the implementation profile to describe. 
Resolution paths are left to the implementation profile to describe. 
