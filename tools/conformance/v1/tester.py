#!/usr/bin/env python3
import argparse
import sys
import requests

def test_get_entity_by_id(base_url, headers, entity_id):
    """
    Tests the GET /entitities/{entityid} endpoint.
    Returns True if the test passes, False otherwise.
    """
    url = f"{base_url}/entitities/{entity_id}"
    params = {
        "authorizationVID": "did:example:auth"  # example
    }
    print(f"--> Testing GET {url} with params={params}")

    try:
        resp = requests.get(url, headers=headers, params=params)
        print(f"    Status: {resp.status_code}")
        # Basic check
        if resp.status_code not in [200, 400, 401, 404]:
            print("    Unexpected status code.")
            return False

        if resp.status_code == 200:
            data = resp.json()
            # Quick check if 'entityVID' is present
            if "entityVID" not in data:
                print("    Missing 'entityVID' in JSON response.")
                return False
    except Exception as ex:
        print(f"    Exception occurred: {ex}")
        return False

    return True


def test_get_entity_authorization(base_url, headers, entity_vid):
    """
    Tests the GET /entities/{entityVID}/authorization endpoint.
    """
    url = f"{base_url}/entities/{entity_vid}/authorization"
    params = {
        "authorizationVID": "did:example:auth"
    }
    print(f"--> Testing GET {url} with params={params}")
    try:
        resp = requests.get(url, headers=headers, params=params)
        print(f"    Status: {resp.status_code}")
        if resp.status_code not in [200, 400, 401, 404]:
            print("    Unexpected status code.")
            return False
        if resp.status_code == 200:
            data = resp.json()
            if not isinstance(data, list):
                print("    Expected a list in the JSON response.")
                return False
    except Exception as ex:
        print(f"    Exception occurred: {ex}")
        return False

    return True


def test_get_entity_authorizations(base_url, headers, entity_vid):
    """
    Tests the GET /entities/{entityVID}/authorizations endpoint.
    """
    url = f"{base_url}/entities/{entity_vid}/authorizations"
    print(f"--> Testing GET {url}")
    try:
        resp = requests.get(url, headers=headers)
        print(f"    Status: {resp.status_code}")
        if resp.status_code not in [200, 400, 401, 404]:
            print("    Unexpected status code.")
            return False
        if resp.status_code == 200:
            data = resp.json()
            if not isinstance(data, list):
                print("    Expected a list in the JSON response.")
                return False
    except Exception as ex:
        print(f"    Exception occurred: {ex}")
        return False

    return True


def test_get_recognized_registries(base_url, headers):
    """
    Tests the GET /registries/recognized-registries endpoint.
    """
    url = f"{base_url}/registries/recognized-registries"
    params = {
        "namespace-VID": "did:example:namespace",
        "EGF-VID": "did:example:egf"
    }
    print(f"--> Testing GET {url} with params={params}")
    try:
        resp = requests.get(url, headers=headers, params=params)
        print(f"    Status: {resp.status_code}")
        if resp.status_code not in [200, 400, 401, 404]:
            return False
    except Exception as ex:
        print(f"    Exception occurred: {ex}")
        return False

    return True


def test_get_specific_recognized_registry(base_url, headers, registry_vid):
    """
    Tests the GET /registries/{registryVID}/recognized-registries endpoint.
    """
    url = f"{base_url}/registries/{registry_vid}/recognized-registries/"
    print(f"--> Testing GET {url}")
    try:
        resp = requests.get(url, headers=headers)
        print(f"    Status: {resp.status_code}")
        if resp.status_code not in [200, 400, 401, 404]:
            return False
    except Exception as ex:
        print(f"    Exception occurred: {ex}")
        return False

    return True


def test_get_registry(base_url, headers, registry_vid):
    """
    Tests the GET /registries/{registryVID}/ endpoint.
    """
    url = f"{base_url}/registries/{registry_vid}/"
    print(f"--> Testing GET {url}")
    try:
        resp = requests.get(url, headers=headers)
        print(f"    Status: {resp.status_code}")
        if resp.status_code not in [200, 400, 401, 404]:
            return False
    except Exception as ex:
        print(f"    Exception occurred: {ex}")
        return False

    return True


def test_lookup_authorizations(base_url, headers):
    """
    Tests the GET /lookup/authorizations endpoint.
    """
    url = f"{base_url}/lookup/authorizations"
    params = {
        "egfURI": "did:example:egf"
    }
    print(f"--> Testing GET {url} with params={params}")
    try:
        resp = requests.get(url, headers=headers, params=params)
        print(f"    Status: {resp.status_code}")
        if resp.status_code not in [200, 400, 401, 404]:
            return False
    except Exception as ex:
        print(f"    Exception occurred: {ex}")
        return False

    return True


def test_lookup_namespaces(base_url, headers):
    """
    Tests the GET /lookup/namespaces endpoint.
    """
    url = f"{base_url}/lookup/namespaces"
    print(f"--> Testing GET {url}")
    try:
        resp = requests.get(url, headers=headers)
        print(f"    Status: {resp.status_code}")
        if resp.status_code not in [200, 400, 401, 404]:
            return False
    except Exception as ex:
        print(f"    Exception occurred: {ex}")
        return False

    return True


def test_lookup_vidmethods(base_url, headers):
    """
    Tests the GET /lookup/vidmethods endpoint.
    """
    url = f"{base_url}/lookup/vidmethods"
    params = {
        "egfURI": "did:example:egf"
    }
    print(f"--> Testing GET {url} with params={params}")
    try:
        resp = requests.get(url, headers=headers, params=params)
        print(f"    Status: {resp.status_code}")
        if resp.status_code not in [200, 400, 401, 404]:
            return False
    except Exception as ex:
        print(f"    Exception occurred: {ex}")
        return False

    return True


def test_lookup_assurancelevels(base_url, headers):
    """
    Tests the GET /lookup/assurancelevels endpoint.
    """
    url = f"{base_url}/lookup/assurancelevels"
    params = {
        "egfURI": "did:example:egf"
    }
    print(f"--> Testing GET {url} with params={params}")
    try:
        resp = requests.get(url, headers=headers, params=params)
        print(f"    Status: {resp.status_code}")
        if resp.status_code not in [200, 400, 401, 404]:
            return False
    except Exception as ex:
        print(f"    Exception occurred: {ex}")
        return False

    return True


def test_get_metadata(base_url, headers):
    """
    Tests the GET /metadata endpoint.
    """
    url = f"{base_url}/metadata"
    print(f"--> Testing GET {url}")
    try:
        resp = requests.get(url, headers=headers)
        print(f"    Status: {resp.status_code}")
        if resp.status_code not in [200, 400, 401, 404]:
            return False
    except Exception as ex:
        print(f"    Exception occurred: {ex}")
        return False

    return True


def test_get_offline_exportfile(base_url, headers):
    """
    Tests the GET /offline/exportfile endpoint.
    """
    url = f"{base_url}/offline/exportfile"
    print(f"--> Testing GET {url}")
    try:
        resp = requests.get(url, headers=headers)
        print(f"    Status: {resp.status_code}")
        if resp.status_code not in [200, 400, 401, 404]:
            return False
        if resp.status_code == 200:
            data = resp.json()
            # Quick sanity check
            if "extractdatetime" not in data:
                print("    Missing 'extractdatetime' in response.")
                return False
    except Exception as ex:
        print(f"    Exception occurred: {ex}")
        return False

    return True


def test_get_offline_ted(base_url, headers):
    """
    Tests the GET /offline/trustestablishmentdocument endpoint.
    """
    url = f"{base_url}/offline/trustestablishmentdocument"
    print(f"--> Testing GET {url}")
    try:
        resp = requests.get(url, headers=headers)
        print(f"    Status: {resp.status_code}")
        if resp.status_code not in [200, 400, 401, 404]:
            return False
        if resp.status_code == 200:
            data = resp.json()
            if "TBD" not in data:
                print("    Missing 'TBD' field in trust establishment document.")
                return False
    except Exception as ex:
        print(f"    Exception occurred: {ex}")
        return False

    return True


def run_all_tests(base_url, bearer_token):
    """
    Runs all test functions, collects results, and prints a summary.
    """
    headers = {
        "Accept": "application/json",
        "Authorization": f"Bearer {bearer_token}" if bearer_token else ""
    }

    tests = [
        ("test_get_entity_by_id", test_get_entity_by_id(base_url, headers, "did:example:123")),
        ("test_get_entity_authorization", test_get_entity_authorization(base_url, headers, "did:example:123")),
        ("test_get_entity_authorizations", test_get_entity_authorizations(base_url, headers, "did:example:123")),
        ("test_get_recognized_registries", test_get_recognized_registries(base_url, headers)),
        ("test_get_specific_recognized_registry", test_get_specific_recognized_registry(base_url, headers, "did:example:registry")),
        ("test_get_registry", test_get_registry(base_url, headers, "did:example:registry")),
        ("test_lookup_authorizations", test_lookup_authorizations(base_url, headers)),
        ("test_lookup_namespaces", test_lookup_namespaces(base_url, headers)),
        ("test_lookup_vidmethods", test_lookup_vidmethods(base_url, headers)),
        ("test_lookup_assurancelevels", test_lookup_assurancelevels(base_url, headers)),
        ("test_get_metadata", test_get_metadata(base_url, headers)),
        ("test_get_offline_exportfile", test_get_offline_exportfile(base_url, headers)),
        ("test_get_offline_ted", test_get_offline_ted(base_url, headers)),
    ]

    overall_success = True
    print("\n==================== Test Results ====================")
    for test_name, result in tests:
        status = "PASS" if result else "FAIL"
        print(f"{test_name}: {status}")
        if not result:
            overall_success = False

    print("=====================================================")
    if overall_success:
        print("ALL TESTS PASSED.")
        sys.exit(0)
    else:
        print("ONE OR MORE TESTS FAILED.")
        sys.exit(1)


def main():
    parser = argparse.ArgumentParser(description="Trust Registry Testing Tool")
    parser.add_argument("--base-url", required=True,
                        help="The base URL of the Trust Registry API.")
    parser.add_argument("--bearer-token", required=False, default="",
                        help="Bearer token for authorization (optional).")

    args = parser.parse_args()

    run_all_tests(args.base_url, args.bearer_token)


if __name__ == "__main__":
    main()
