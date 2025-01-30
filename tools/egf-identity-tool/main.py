import streamlit as st
import json
import base64
from cryptography.hazmat.primitives.asymmetric import ed25519, x25519
from cryptography.hazmat.primitives import serialization
from base58 import b58encode

# Streamlit UI Setup
st.set_page_config(page_title="DID Peer 2 Generator")
st.title("DID Peer 2 Generator & Resolver")

# Abbreviations Mapping
full_to_abbreviation = {
    "profile": "p",
    "type": "t",
    "serviceEndpoint": "s",
    "routingKeys": "r",
    "accept": "a",
    "DIDCommMessaging": "dm",
    "integrity": "i",
    "uri": "u"
}

abbreviation_to_full = {v: k for k, v in full_to_abbreviation.items()}


def abbreviate_service(obj):
    """Convert full service fields to abbreviations."""
    if isinstance(obj, dict):
        return {full_to_abbreviation.get(k, k): abbreviate_service(v) for k, v in obj.items()}
    elif isinstance(obj, list):
        return [abbreviate_service(elem) for elem in obj]
    return obj


def expand_service(obj):
    """Convert abbreviated service fields to full names."""
    if isinstance(obj, dict):
        return {abbreviation_to_full.get(k, k): expand_service(v) for k, v in obj.items()}
    elif isinstance(obj, list):
        return [expand_service(elem) for elem in obj]
    return obj


def generate_did_peer2(config_data, method_prefix):
    """Generate a DID:peer:2 identifier with service endpoints."""
    # Generate Ed25519 Key
    ed_private_key = ed25519.Ed25519PrivateKey.generate()
    ed_public_key = ed_private_key.public_key()
    ed_pub_bytes = ed_public_key.public_bytes(
        encoding=serialization.Encoding.Raw, format=serialization.PublicFormat.Raw
    )

    # Generate X25519 Key
    x25519_private_key = x25519.X25519PrivateKey.generate()
    x25519_public_key = x25519_private_key.public_key()
    x25519_pub_bytes = x25519_public_key.public_bytes(
        encoding=serialization.Encoding.Raw, format=serialization.PublicFormat.Raw
    )

    ed_pub_mb = "z" + b58encode(ed_pub_bytes).decode()
    x25519_pub_mb = "z" + b58encode(x25519_pub_bytes).decode()

    did_parts = [method_prefix, "V" + ed_pub_mb, "E" + x25519_pub_mb]

    # Abbreviate and encode services
    for svc in config_data.get("services", []):
        abbrev = abbreviate_service(svc)
        svc_bytes = json.dumps(abbrev).encode()
        b64 = base64.urlsafe_b64encode(svc_bytes).decode().rstrip("=")
        did_parts.append("S" + b64)

    did = ".".join(did_parts)

    # Private keys in hex format
    ed_priv_hex = ed_private_key.private_bytes(
        encoding=serialization.Encoding.Raw,
        format=serialization.PrivateFormat.Raw,
        encryption_algorithm=serialization.NoEncryption()
    ).hex()

    x25519_priv_hex = x25519_private_key.private_bytes(
        encoding=serialization.Encoding.Raw,
        format=serialization.PrivateFormat.Raw,
        encryption_algorithm=serialization.NoEncryption()
    ).hex()

    return did, ed_priv_hex, x25519_priv_hex


def resolve_did_peer2(did_str, method_prefix):
    """Resolve a DID:peer:2 string into a DID Document."""
    if not did_str.startswith(method_prefix):
        raise ValueError(f"DID '{did_str}' does not match method prefix '{method_prefix}'")

    doc = {
        "@context": [
            "https://www.w3.org/ns/did/v1",
            "https://w3id.org/security/multikey/v1"
        ],
        "id": did_str
    }

    vm_list, auth, assertion, key_agreement, cap_inv, cap_del, services = [], [], [], [], [], [], []
    service_count = 0

    body = did_str[len(method_prefix):].lstrip('.')
    parts = body.split('.')
    key_index = 1

    for seg in parts:
        if len(seg) < 2:
            raise ValueError(f"Malformed segment: '{seg}'")
        purpose, rest = seg[0], seg[1:]

        if purpose in ['V', 'A', 'E', 'I', 'D']:
            key_id = f"#key-{key_index}"
            key_index += 1
            vm_list.append({
                "id": key_id,
                "type": "Multikey",
                "controller": did_str,
                "publicKeyMultibase": rest
            })

            if purpose == 'V':
                auth.append(key_id)
            elif purpose == 'A':
                assertion.append(key_id)
            elif purpose == 'E':
                key_agreement.append(key_id)
            elif purpose == 'I':
                cap_inv.append(key_id)
            elif purpose == 'D':
                cap_del.append(key_id)

        elif purpose == 'S':
            padding_needed = (4 - len(rest) % 4) % 4
            svc_bytes = base64.urlsafe_b64decode(rest + ("=" * padding_needed))
            expanded = expand_service(json.loads(svc_bytes))

            if not isinstance(expanded, dict):
                raise ValueError("Decoded service is not an object")

            expanded["id"] = f"#service-{service_count}" if service_count else "#service"
            service_count += 1
            services.append(expanded)

        else:
            raise ValueError(f"Unknown purpose code '{purpose}' in segment '{seg}'")

    if vm_list:
        doc["verificationMethod"] = vm_list
    if auth:
        doc["authentication"] = auth
    if assertion:
        doc["assertionMethod"] = assertion
    if key_agreement:
        doc["keyAgreement"] = key_agreement
    if cap_inv:
        doc["capabilityInvocation"] = cap_inv
    if cap_del:
        doc["capabilityDelegation"] = cap_del
    if services:
        doc["service"] = services

    return doc


# Streamlit UI
mode = st.radio("Select mode:", ["Generate", "Resolve"])

if mode == "Generate":

    method_prefix = st.text_input("Method Prefix", value="did:peer:2")
    st.subheader("Generate a DID")

    config_file = st.file_uploader("Upload config.json", type="json")

    st.write("**Or manually enter service details:**")
    egf_uri = st.text_input("EGF URI", value="https://localhost:3000/terms")
    trqp_uris_str = st.text_area("TRQP URIs (comma or newline-separated)", value="http://localhost:8080")

    if st.button("Generate DID"):
        try:
            if config_file:
                config_data = json.loads(config_file.read().decode("utf-8"))
            else:
                trqp_uris = [u.strip() for u in trqp_uris_str.replace(",", "\n").split("\n") if u.strip()]
                config_data = {
                    "services": [
                        {
                            "id": "#egfURI",
                            "type": "egfURI",
                            "serviceEndpoint": {
                                "profile": "https://trustoverip.org/profiles/trp/egfURI/v1",
                                "uri": egf_uri,
                                "integrity": "122041dd7b6443542e75701aa98a0c235951a28a0d851b11564d20022ab11d2589a8"
                            }
                        },
                        {
                            "id": "#tr-1",
                            "type": "TRQP",
                            "serviceEndpoint": {
                                "profile": "https://trustoverip.org/profiles/trp/v2",
                                "uri": trqp_uris,
                                "integrity": "122041dd7b6443542e75701aa98a0c235952a28a0d851b11564d20022ab11d2589a8"
                            }
                        }
                    ]
                }

            did_str, ed_hex, x_hex = generate_did_peer2(config_data, method_prefix)
            st.success("DID Generation Successful!")
            st.code(did_str)
        except Exception as e:
            st.error(f"Error: {str(e)}")

if mode == "Resolve":
    st.subheader("Resolve a DID")

    did_str = st.text_input("DID", value="")
    if st.button("Resolve DID"):
        try:
            doc = resolve_did_peer2(did_str, "did:peer:2")
            st.success("DID Resolution Successful!")
            st.json(doc)
        except Exception as e:
            st.error(f"Error: {str(e)}")
