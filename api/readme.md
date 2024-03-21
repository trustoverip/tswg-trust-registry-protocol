

## Rendering OpenAPI Specification

In the `./api` folder, run:

```
npx @redocly/cli build-docs toip-tswg-trustregistryprotocol-v2.yaml --output ../spec/api/redoc-static.html
```

NOTE: the `--output` switch is required as SpecUp only takes content from the `./spec` folder (and below).

That produces the [Redoc Rendered OpenAPI Spec](redoc-static.html)