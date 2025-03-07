"""From Python to Go: Python: 016 - Advanced SSH."""

# Modules
import argparse
import datetime
import os
import sys
import difflib
from dataclasses import dataclass
from typing import List
import yaml
from scrapli import Scrapli


# Classes
@dataclass
class Credentials:
    """Class to store credentials."""
    username: str
    password: str


@dataclass
class Instruction:
    """Class to store instructions."""
    command: str
    config: List[str]


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
        with Scrapli(
            host=self.ip_address,
            auth_username=self.credentials.username,
            auth_password=self.credentials.password,
            platform=self.platform,
            auth_strict_key=False,
        ) as conn:
            # Get state before change
            before = conn.send_command(instruction.command)

            # Apply change
            conn.send_configs(instruction.config)

            # Get state after change
            after = conn.send_command(instruction.command)

            # Diff
            diff = "\n".join(
                difflib.context_diff(
                    before.result.splitlines(),
                    after.result.splitlines(),
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
        command="show interfaces description",
        config=["interface Loopback 23", "description Test"],
    )

    # Execute command
    for device in devices:
        device.execute_change(instruction)

    # Print results
    for device in devices:
        print(f"Device: {device.hostname}")
        for result in device.results:
            print(f"Config: {result.instruction.config}", f"Impact: {result.diff}", f"Timestamp: {result.timestamp}", sep="\n")
