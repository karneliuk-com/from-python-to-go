"""From Python to Go: Python: 019 - REST API and GNMI."""

# Modules
import datetime
import os
import sys
from dataclasses import dataclass
from typing import List, Tuple
import difflib
import pygnmi.client
import httpx


# Classes
@dataclass
class Credentials:
    """Class to store credentials."""
    username: str
    password: str


@dataclass
class InventoryCredentials:
    """Class to store credentials."""
    url: str
    token: str


@dataclass
class Instruction:
    """Class to store instructions."""
    command: List[str]
    config: List[tuple]


@dataclass
class Result:
    """Class to store command execution data."""
    instruction: Instruction
    diff: str
    timestamp: datetime.datetime


class Device:
    """Class to interact with netowrk device."""
    def __init__(self, hostname: str, ip_address: str, port: int, platform: str, credentials: Credentials):
        self.hostname = hostname
        self.ip_address = ip_address
        self.port = port
        self.platform = platform
        self.credentials = credentials

        self.results: List[Result] = []

    def execute_change(self, instruction: Instruction) -> None:
        """Method to execute change."""

        # Connect to device
        with pygnmi.client.gNMIclient(
            target=(self.ip_address, self.port),
            username=self.credentials.username,
            password=self.credentials.password,
            skip_verify=True,
            timeout=5,
        ) as gconn:
            # Get state before change
            before = gconn.get(path=instruction.command, datatype="config")
            before_stringified = self.dict_to_xpath(before)

            # Apply change
            config_result = gconn.set(update=instruction.config, encoding="json_ietf")
            print(f"{config_result=}")

            # Get state after change
            after = gconn.get(path=instruction.command, datatype="config")
            after_stringified = self.dict_to_xpath(after)

            # Diff
            diff = "\n".join(
                difflib.context_diff(
                    before_stringified,
                    after_stringified,
                    lineterm="",
                )
            )

            self.results.append(
                Result(
                    instruction=instruction,
                    diff=diff,
                    timestamp=datetime.datetime.now(),
                )
            )

    def dict_to_xpath(self, data: dict) -> list:
        """Method to convert dict to xpath."""
        result = []

        if isinstance(data, str):
            return data

        for key, value in data.items():
            if isinstance(value, list):
                for ind, item in enumerate(value):
                    tr = self.dict_to_xpath(item)
                    result.extend([f"{key}/{ind}/{_}" for _ in tr])

            elif isinstance(value, dict):
                tr = self.dict_to_xpath(value)
                result.extend([f"{key}/{_}" for _ in tr])

            else:
                result.append(f"{key} = {value}")

        return result


# Functions
def load_inventory(inventory: InventoryCredentials, credentials: Credentials) -> List[Device]:
    """Function to load inventory data."""
    # Create HTTP client and set headers
    hclient = httpx.Client(
        base_url=inventory.url.rstrip("/"),
        headers={"Authorization": f"Token {inventory.token}"},
    )
    # Retrieve data from REST API
    try:
        response = hclient.get(
            "/api/dcim/devices/",
            params={
                "site": "kblog",
            }
        )
        response.raise_for_status()
        data = response.json()

    except Exception as e:
        print(e)
        sys.exit(1)

    # Populate list of devices
    result = []
    for device in data["results"]:
        result.append(
            Device(
                hostname=device["name"],
                ip_address=device["primary_ip"]["address"].split("/")[0],
                port=device["custom_fields"].get("gnmi_port", 50051),
                platform=device["platform"]["slug"],
                credentials=credentials,
            )
        )

    return result


def get_credentials() -> Tuple[Credentials, InventoryCredentials]:
    """Function to get credentials."""
    return (
        Credentials(
            os.getenv("AUTOMATION_USER"),
            os.getenv("AUTOMATION_PASS"),
        ),
        InventoryCredentials(
            os.getenv("AUTOMATION_INVENTORY_URL"),
            os.getenv("AUTOMATION_INVENTORY_TOKEN"),
        ),
    )


# Main code
if __name__ == "__main__":
    # Get credentials
    credentials, inventory_credentials = get_credentials()

    # Load inventory
    devices = load_inventory(inventory_credentials, credentials=credentials)

    # Config
    instruction = Instruction(
        command=["/openconfig-interfaces:interfaces"],
        config=[
            (
                "/openconfig-interfaces:interfaces",
                {
                    "interface": [
                        {
                            "name": "Loopback 23",
                            "config": {
                                "name": "Loopback 23",
                                "description": "Test-gnmi-python-2",
                            }
                        },
                    ],
                },
            ),
        ],
    )

    # Execute command
    for device in devices:
        device.execute_change(instruction)

    # Print results
    for device in devices:
        print(f"Device: {device.hostname}")
        for result in device.results:
            print(f"Config: {result.instruction.config}", f"Impact: {result.diff}", f"Timestamp: {result.timestamp}", sep="\n")
