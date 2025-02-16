"""From Python to Go: Python: 015 - Basic SSH."""

# Modules
import argparse
import datetime
import os
import sys
import time
import re
from dataclasses import dataclass
from typing import List
import yaml
import paramiko


# Classes
@dataclass
class Credentials:
    """Class to store credentials."""
    username: str
    password: str


@dataclass
class Result:
    """Class to store command execution data."""
    command: str
    output: str
    timestamp: datetime.datetime


class Device:
    """Class to interact with netowrk device."""
    def __init__(self, hostname: str, ip_address: str, credentials: Credentials):
        self.hostname = hostname
        self.ip_address = ip_address
        self.credentials = credentials

        self.results: List[Result] = []

    def execute_command(self, command: str) -> None:
        """Method to execute a command."""
        # Create a new SSH client
        client = paramiko.SSHClient()
        client.set_missing_host_key_policy(paramiko.AutoAddPolicy())

        # Connect to the device
        client.connect(
            self.ip_address,
            username=self.credentials.username,
            password=self.credentials.password,
            look_for_keys=False,
            allow_agent=False,
        )

        # Invoke the session
        session = client.invoke_shell()
        session.recv(65535)

        # Execute the command
        session.send(command + "\n")
        output = ""
        regex = re.compile(rf"{self.hostname}[#>]", flags=re.MULTILINE)
        print(f"{regex=}")
        while not regex.search(output):
            time.sleep(.1)
            output += session.recv(65535).decode("utf-8")
            print(output)

        # Store the result
        self.results.append(Result(command, output, datetime.datetime.now()))

        # Close the connection
        session.close()


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

    # Execute command
    for device in devices:
        device.execute_command("show version")

    # Print results
    for device in devices:
        print(f"Device: {device.hostname}")
        for result in device.results:
            print(f"Command: {result.command}", f"Output: {result.output}", f"Timestamp: {result.timestamp}", sep="\n")
