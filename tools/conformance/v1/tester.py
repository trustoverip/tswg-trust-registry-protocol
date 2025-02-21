#!/usr/bin/env python3

import argparse
import sys
import requests
import json


def test_get_metadata(base_url, headers, egf_id=None):
    """
    Tests the GET /metadata endpoint as specified in the TRQP spec.
    Optionally passes egf_id as a query parameter.
    """
    url = f"{base_url}/metadata"
    params = {}
    if egf_id:
        params["egf_id"] = egf_id

    print(f"--> Testing GET {url} with params={params}")
    try:
        resp = requests.get(url, headers=headers, params=params)
        print(f"    Status: {resp.status_code}")

        # The TRQP spec expects possible responses: 200, 401, 404
        if resp.status_code not in [200, 401, 404]:
            print("    Unexpected status code.")
            return False

        if resp.status_code == 200:
            data = resp.json()
            # Minimal checks based on TRQP "TrustRegistryMetadata" schema:
            for required_field in ["id", "description", "name", "controllers"]:
                if required_field not in data:
                    print(f"    Missing '{required_field}' in JSON response.")
                    return False
    except Exception as ex:
        print(f"    Exception occurred: {ex}")
        return False

    return True


def test_check_ecosystem_recognition(
    base_url, headers, ecosystem_id, egf_id, time_val=None
):
    """
    Tests the GET /registries/{ecosystem_id}/recognition endpoint.
    Required query param: egf_id
    Optional query param: time
    """
    url = f"{base_url}/registries/{ecosystem_id}/recognition"
    params = {"egf_id": egf_id}
    if time_val:
        params["time"] = time_val

    print(f"--> Testing GET {url} with params={params}")
    try:
        resp = requests.get(url, headers=headers, params=params)
        print(f"    Status: {resp.status_code}")

        # The TRQP spec expects possible responses: 200, 401, 404
        if resp.status_code not in [200, 401, 404]:
            print("    Unexpected status code.")
            return False

        if resp.status_code == 200:
            data = resp.json()
            # Minimal checks based on TRQP "RecognitionResponse" schema:
            for required_field in [
                "recognized",
                "message",
                "evaluated_at",
                "response_time",
            ]:
                if required_field not in data:
                    print(f"    Missing '{required_field}' in JSON response.")
                    return False
    except Exception as ex:
        print(f"    Exception occurred: {ex}")
        return False

    return True


def test_check_authorization_status(
    base_url, headers, entity_id, authorization_id, egf_id, all_bool=True, time_val=None
):
    """
    Tests the GET /entities/{entity_id}/authorization endpoint.
    Required query params: authorization_id, egf_id, all
    Optional query param: time
    """
    url = f"{base_url}/entities/{entity_id}/authorization"
    params = {
        "authorization_id": authorization_id,
        "egf_id": egf_id,
        "all": str(
            all_bool
        ).lower(),  # or simply all_bool if your server can parse boolean
    }
    if time_val:
        params["time"] = time_val

    print(f"--> Testing GET {url} with params={params}")
    try:
        resp = requests.get(url, headers=headers, params=params)
        print(f"    Status: {resp.status_code}")

        # The TRQP spec expects possible responses: 200, 401, 404
        if resp.status_code not in [200, 401, 404]:
            print("    Unexpected status code.")
            return False

        if resp.status_code == 200:
            # Could be oneOf: a single AuthorizationResponse or an array of them
            # We'll handle both possibilities:
            try:
                data = resp.json()
            except json.JSONDecodeError:
                print("    Invalid JSON in response.")
                return False

            if isinstance(data, list):
                # If it returns a list, check each item
                for item in data:
                    for required_field in [
                        "recognized",
                        "authorized",
                        "message",
                        "evaluated_at",
                        "response_time",
                    ]:
                        if required_field not in item:
                            print(
                                f"    Missing '{required_field}' in one of the array items."
                            )
                            return False
            elif isinstance(data, dict):
                # If it returns a single JSON object
                for required_field in [
                    "recognized",
                    "authorized",
                    "message",
                    "evaluated_at",
                    "response_time",
                ]:
                    if required_field not in data:
                        print(f"    Missing '{required_field}' in JSON response.")
                        return False
            else:
                print("    Response not JSON object or array of objects.")
                return False

    except Exception as ex:
        print(f"    Exception occurred: {ex}")
        return False

    return True


def run_all_tests(base_url, bearer_token):
    """
    Runs all test functions, collects results, and prints a summary.
    """
    headers = {"Accept": "application/json"}
    # Attach bearer token if provided
    if bearer_token:
        headers["Authorization"] = f"Bearer {bearer_token}"

    tests = [
        # 1) /metadata (GET)
        ("test_get_metadata", test_get_metadata(base_url, headers, egf_id=None)),
        # 2) /registries/{ecosystem_id}/recognition (GET)
        (
            "test_check_ecosystem_recognition",
            test_check_ecosystem_recognition(
                base_url,
                headers,
                ecosystem_id="did:example:some-ecosystem",
                egf_id="did:example:some-egf",
            ),
        ),
        # 3) /entities/{entity_id}/authorization (GET)
        (
            "test_check_authorization_status",
            test_check_authorization_status(
                base_url,
                headers,
                entity_id="did:example:some-entity",
                authorization_id="did:example:auth",
                egf_id="did:example:some-egf",
                all_bool=True,
            ),
        ),
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
    parser = argparse.ArgumentParser(description="TRQP Conformance Test Script")
    parser.add_argument(
        "--base-url", required=True, help="The base URL of the TRQP API."
    )
    parser.add_argument(
        "--bearer-token",
        required=False,
        default="",
        help="Bearer token for authorization (optional).",
    )
    args = parser.parse_args()

    run_all_tests(args.base_url, args.bearer_token)


if __name__ == "__main__":
    main()
