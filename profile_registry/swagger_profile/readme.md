

## Rendering OpenAPI Specification

In the `./api` folder, run:

```
npx @redocly/cli build-docs toip-tswg-trustregistryprotocol-v2.yaml --output ../docs/api/redoc-static.html
```

NOTE: the `--output` switch is required as SpecUp only takes content from the `./docs` folder (and below).

That produces the [Redoc Rendered OpenAPI Spec](redoc-static.html) (NOTE: this is a local link)  

A reference to Redoc can be found [here](https://github.com/Redocly/redoc).  

