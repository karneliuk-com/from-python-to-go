"""From Python to Go: Python: 017 - NETCONF."""

# Modules
import argparse
import datetime
import os
import sys
from dataclasses import dataclass
from typing import List
import difflib
import yaml
from scrapli_netconf.driver import NetconfDriver
import xmltodict


# Classes
@dataclass
class Credentials:
    """Class to store credentials."""
    username: str
    password: str


@dataclass
class Instruction:
    """Class to store instructions."""
    command: dict
    config: dict


@dataclass
class Result:
    """Class to store command execution data."""
    instruction: Instruction
    diff: str
    timestamp: datetime.datetime


class Device:
    """Class to interact with netowrk device."""
    def __init__(self, hostname: str, ip_address: str, platform: str, credentials: Credentials):
        self.hostname = hostname
        self.ip_address = ip_address
        self.platform = platform
        self.credentials = credentials

        self.results: List[Result] = []

    def execute_change(self, instruction: Instruction) -> None:
        """Method to execute change."""

        # Connect to device
        with NetconfDriver(
            host=self.ip_address,
            port=830,
            auth_username=self.credentials.username,
            auth_password=self.credentials.password,
            auth_strict_key=False,
            transport="ssh2",
        ) as conn:
            filter_ = xmltodict.unparse(instruction.command).splitlines()[1]
            # Get state before change
            before = conn.get_config(source="running", filter_=filter_, filter_type="subtree")
            before_stringified = self.dict_to_xpath(xmltodict.parse(before.result))

            # Apply change
            config_ = "\n".join(xmltodict.unparse(instruction.config).splitlines()[1:])
            change_result = conn.edit_config(target="candidate", config=config_)

            if change_result.failed:
                print(f"Error: {change_result.result}")
            else:
                commit_result = conn.commit()
                if commit_result.failed:
                    print(f"Error: {commit_result.result}")

            # Get state after change
            after = conn.get_config(source="running", filter_=filter_, filter_type="subtree")
            after_stringified = self.dict_to_xpath(xmltodict.parse(after.result))

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
        command={"interfaces": {"@xmlns": "http://openconfig.net/yang/interfaces"}},
        config={
            "config": {
                "interfaces": {
                    "@xmlns": "http://openconfig.net/yang/interfaces",
                    "interface": [
                        {
                            "name": "Loopback 23",
                            "config": {
                                "name": "Loopback 23",
                                "description": "Test-netconf-python-2",
                            }
                        },
                    ],
                },
            },
        },
    )

    # Execute command
    for device in devices:
        device.execute_change(instruction)

    # Print results
    for device in devices:
        print(f"Device: {device.hostname}")
        for result in device.results:
            print(f"Config: {result.instruction.config}", f"Impact: {result.diff}", f"Timestamp: {result.timestamp}", sep="\n")
