import streamlit as st
import json
import base64
from cryptography.hazmat.primitives.asymmetric import ed25519, x25519
from cryptography.hazmat.primitives import serialization
from base58 import b58encode

############################################
# A Python Streamlit app replicating the Go code for
# DID:peer:2 generation & resolution, with manual or file-based config.
############################################

st.set_page_config(page_title="DID Peer 2 App")
st.title("TRQP EGF DID Creator")

################################################################################
# Utility: Abbreviate and expand fields for services
################################################################################

full_to_abbreviation = {
    "type": "t",
    "serviceEndpoint": "s",
    "routingKeys": "r",
    "accept": "a",
    "DIDCommMessaging": "dm"
}

abbreviation_to_full = {
    "t": "type",
    "s": "serviceEndpoint",
    "r": "routingKeys",
    "a": "accept",
    "dm": "DIDCommMessaging"
}


def abbreviate_service(obj):
    if isinstance(obj, dict):
        new_map = {}
        for k, v in obj.items():
            abbr_key = full_to_abbreviation.get(k, k)
            if k == "type" and isinstance(v, str) and v in full_to_abbreviation:
                v = full_to_abbreviation[v]
            new_map[abbr_key] = abbreviate_service(v)
        return new_map
    elif isinstance(obj, list):
        return [abbreviate_service(elem) for elem in obj]
    else:
        return obj


def expand_service(obj):
    if isinstance(obj, dict):
        new_map = {}
        for k, v in obj.items():
            expanded_key = abbreviation_to_full.get(k, k)
            if expanded_key == "type" and isinstance(v, str) and v in abbreviation_to_full:
                v = abbreviation_to_full[v]
            new_map[expanded_key] = expand_service(v)
        return new_map
    elif isinstance(obj, list):
        return [expand_service(elem) for elem in obj]
    else:
        return obj

################################################################################
# Generate DID: function
################################################################################

def generate_did_peer2(config_data, method_prefix):
    # 1. Generate Ed25519 keys
    ed_private_key = ed25519.Ed25519PrivateKey.generate()
    ed_public_key = ed_private_key.public_key()
    ed_pub_bytes = ed_public_key.public_bytes(
        encoding=serialization.Encoding.Raw,
        format=serialization.PublicFormat.Raw
    )

    # 2. Generate X25519 keys
    x25519_private_key = x25519.X25519PrivateKey.generate()
    x25519_public_key = x25519_private_key.public_key()
    x25519_pub_bytes = x25519_public_key.public_bytes(
        encoding=serialization.Encoding.Raw,
        format=serialization.PublicFormat.Raw
    )

    ed_pub_mb = "z" + b58encode(ed_pub_bytes).decode()
    x25519_pub_mb = "z" + b58encode(x25519_pub_bytes).decode()

    # Build DID e.g.: did:peer:2.Vz6Mkj3PU...Ez6LSg8zQ...
    did_parts = [method_prefix, "V" + ed_pub_mb, "E" + x25519_pub_mb]

    # Abbreviate services and append each as .S...
    for svc in config_data.get("services", []):
        abbrev = abbreviate_service(svc)
        svc_bytes = json.dumps(abbrev).encode()
        b64 = base64.urlsafe_b64encode(svc_bytes).decode().rstrip("=")
        did_parts.append("S" + b64)

    did = ".".join(did_parts)

    # Hex-encoded private keys
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

################################################################################
# Resolve DID: function
################################################################################

def resolve_did_peer2(did_str, method_prefix):
    if not did_str.startswith(method_prefix):
        raise ValueError(f"DID '{did_str}' does not match method prefix '{method_prefix}'")

    doc = {
        "@context": [
            "https://www.w3.org/ns/did/v1",
            "https://w3id.org/security/multikey/v1"
        ],
        "id": did_str
    }

    vm_list = []
    auth = []
    assertion = []
    key_agreement = []
    cap_inv = []
    cap_del = []

    services = []
    service_count = 0

    # Remove prefix => everything after method_prefix
    body = did_str[len(method_prefix):]
    if body.startswith('.'):
        body = body[1:]

    parts = body.split('.')

    key_index = 1

    for seg in parts:
        if len(seg) < 2:
            raise ValueError(f"Malformed segment: '{seg}'")
        purpose = seg[0]
        rest = seg[1:]

        if purpose in ['V', 'A', 'E', 'I', 'D']:
            # Build a verification method
            key_id = f"#key-{key_index}"
            key_index += 1
            vm = {
                "id": key_id,
                "type": "Multikey",
                "controller": did_str,
                "publicKeyMultibase": rest
            }
            vm_list.append(vm)

            # Now add references to verification relationships
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
            # Service segment
            padding_needed = (4 - len(rest) % 4) % 4
            svc_b64 = rest + ("=" * padding_needed)
            svc_bytes = base64.urlsafe_b64decode(svc_b64)

            raw = json.loads(svc_bytes)
            expanded = expand_service(raw)
            if not isinstance(expanded, dict):
                raise ValueError("Decoded service not an object")

            if "id" not in expanded:
                if service_count == 0:
                    expanded["id"] = "#service"
                else:
                    expanded["id"] = f"#service-{service_count}"
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

################################################################################
# Streamlit UI
################################################################################

mode = st.radio("Select mode:", ["Generate", "Resolve"])

method_prefix = st.text_input("Method Prefix", value="did:peer:2")

if mode == "Generate":
    st.subheader("Generate a DID")
    st.write(
        "You can either upload a config.json or manually provide 'egfURI' + 'uri' (array). "
        "If you do both, the uploaded file overrides the manual fields."
    )

    config_file = st.file_uploader("Upload config.json", type="json")

    st.write("**Or** manually set service info (used only if no config file is uploaded):")
    egf_uri = st.text_input("egfURI", value="https://my-tr-service/egfURI")
    uri_list_str = st.text_area(
        "uri (one per line or comma-separated)",
        value="https://my-tr-service/egfURI\nhttps://my-tr-service/another"
    )

    if st.button("Generate DID"):
        try:
            config_data = None
            if config_file is not None:
                # Use the uploaded config
                config_data = json.loads(config_file.read().decode("utf-8"))
            else:
                # Construct config_data from the fields
                if "," in uri_list_str:
                    raw_uris = [u.strip() for u in uri_list_str.split(",") if u.strip()]
                else:
                    raw_uris = [u.strip() for u in uri_list_str.split("\n") if u.strip()]

                config_data = {
                    "services": [
                        {
                            "id": "#egfURI",
                            "type": "egfURI",
                            "serviceEndpoint": {
                                "profile": "https://trustoverip.org/profiles/trp/egfURI/v1",
                                "uri": raw_uris,
                                "integrity": "122041dd7b6443542e75701aa98a0c235951a28a0d851b11564d20022ab11d2589a8"
                            }
                        }
                    ]
                }

            did_str, ed_hex, x_hex = generate_did_peer2(config_data, method_prefix)
            st.success("DID Generation Successful!")
            st.write("**DID**:")
            st.code(did_str)
            st.write("**Ed25519 Private Key (hex)**:")
            st.code(ed_hex)
            st.write("**X25519 Private Key (hex)**:")
            st.code(x_hex)
        except Exception as e:
            st.error(f"Error: {str(e)}")

elif mode == "Resolve":
    st.subheader("Resolve a DID")
    did_input = st.text_input("Enter DID string")

    if st.button("Resolve DID"):
        try:
            doc = resolve_did_peer2(did_input, method_prefix)
            st.success("DID Resolution Successful!")
            st.json(doc)
        except Exception as e:
            st.error(f"Error: {str(e)}")
