# Review Committee Checklist {#review-committee-checklist}

This MUST be submitted with the [Review Committee Report](#review-committee-report) and SHOULD use [this Google Form](https://docs.google.com/forms/d/e/1FAIpQLSf2lfZ9aLCOtA5U4LhiMiTX0IFJk2sdX-1fHwpCE_w7jxs52Q/viewform). One Review Committee Checklist MUST be completed for each deliverable.

## Document Readiness {#document-readiness}

### Content Completeness {#content-completeness}

1. All sections fully completed (no TODOs, placeholders, or “TBD”).  
2. All diagrams, examples, tables, and references included and final.  
3. All normative statements use IETF [RFC 2119](https://datatracker.ietf.org/doc/html/rfc2119) keywords (MUST/SHOULD/MAY).  
4. If the deliverable is a specification, all sections (except the introduction and boilerplate sections) are marked as Normative or Informative.  
5. All references are included and properly classified as Normative and Informative.  
6. Security and Privacy Consideration sections are finalized and complete.  
7. Glossary and definitions are complete and consistent and reference the ToIP Glossary family whenever possible.

### Internal Consistency {#internal-consistency}

1. Terminology is consistent across the spec.  
2. There are no contradictions between sections.  
3. All references to numbered items (e.g., sections, algorithms, figures) are correct.  
4. All examples match the normative definitions and schemas.

## Technical Quality {#technical-quality}

### Normative Content {#normative-content}

1. Requirements are testable and unambiguous.  
2. All algorithms or flows have clearly defined inputs, outputs, and error behavior.  
3. Schemas validate correctly (if machine-readable schemas are included).

### Interoperability & Implementability {#interoperability-&-implementability}

1. Mandatory-to-implement features are clear.  
2. Optional features are clearly labeled, with rules for negotiation or detection.  
3. Versioning, extensibility, and backwards-compatibility rules are explicit.  
4. Conformance criteria align with test suite requirements.

## Editorial Quality {#editorial-quality}

### Style & Formatting {#style-&-formatting}

1. The deliverable follows the ToIP style guide.  
2. All abbreviations are defined in the Terminology section and at first use.  
3. Typography, heading structure, and captioning are consistent.  
4. The document uses valid HTML or Markdown (depending on publishing platform).

### Clarity & Readability {#clarity-&-readability}

1. Long paragraphs broken into logical chunks.  
2. Complex concepts explained with examples.  
3. Diagrams improve comprehension and match the text.

## Metadata & Ancillary Materials {#metadata-&-ancillary-materials}

### Document Metadata {#document-metadata}

1. Title, version number, date, are correct.  
2. The document has been assigned a ToIP permalink conformant with the template specified on [this wiki page](https://lf-toip.atlassian.net/wiki/spaces/HOME/pages/22986818/File+Names+and+Permalinks).  
3. The status of the document is a Working Group Approved Deliverable.  
4. The editors and contributors are listed with proper attribution.  
5. Change history / release notes are current.

### Links and References {#links-and-references}

1. All external references include stable links (ideally permalinks).  
2. All normative references are final versions (not drafts).  
3. Informative references are clearly marked.

## Publishing Readiness {#publishing-readiness}

### IPR and Licensing {#ipr-and-licensing}

1. Specification licensing text is included and correct.  
2. All contributors acknowledged as per Working Group policy.  
3. Any patent disclosures are filed (if required under the JDF process and WG patent licensing policy).  
4. There are no licensing conflicts with referenced content.

### Artifacts {#artifacts}

1. Supporting files (schemas, JSON-LD contexts, vocabularies, examples, test artifacts) are published and linked.  
2. Any accompanying test suite is current and aligned with the normative statements in the specification.  
3. Any implementation or interoperability reports (if required) are complete.

### Final Checks {#final-checks}

1. Spellcheck and grammar are clean.  
2. The document passes broken link checks.  
3. The final PDF/HTML rendered versions have been reviewed manually.  
4. The publishing pipeline or repo build passes cleanly.
