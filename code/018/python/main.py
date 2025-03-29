"""From Python to Go: Python: 017 - NETCONF."""

# Modules
import argparse
import datetime
import os
import sys
from dataclasses import dataclass
from typing import List
import difflib
import pygnmi.client
import yaml
import pygnmi
# from scrapli_netconf.driver import NetconfDriver
# import xmltodict


# Classes
@dataclass
class Credentials:
    """Class to store credentials."""
    username: str
    password: str


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
def read_args() -> argparse.Namespace:
    """Helper function to read CLI arguments."""
    parser = argparse.ArgumentParser(description="User input.")
    parser.add_argument("-i", "--inventory", type=str, help="Path to inventory file.")
    return parser.parse_args()


def load_inventory(filename: str, credentials: Credentials) -> List[Device]:
    """Function to load inventory data."""
    # Open file
    try:
        with open(filename, "r", encoding="utf-8") as f:
            data = yaml.safe_load(f)

    except FileNotFoundError as e:
        print(e)
        sys.exit(1)

    # Populate list of devices
    result = []
    for device in data:
        result.append(Device(credentials=credentials, **device))

    return result


def get_credentials() -> Credentials:
    """Function to get credentials."""
    username = os.getenv("AUTOMATION_USER")
    password = os.getenv("AUTOMATION_PASS")
    return Credentials(username, password)


# Main code
if __name__ == "__main__":
    # Read CLI arguments
    args = read_args()

    # Get credentials
    credentials = get_credentials()

    # Load inventory
    devices = load_inventory(args.inventory, credentials=credentials)

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
