"""From Python to Go: Python: 014 - Templating configuration."""

# Modules
import argparse
from typing import List, Union
import sys
import os
import yaml
from pydantic import BaseModel
import jinja2


# Classes
class IPAddress(BaseModel):
    """Class to store IP address data."""
    address: str
    prefix: int


class Interface(BaseModel):
    """Class to store interface data."""
    name: str
    description: Union[str, None] = None
    ip4: Union[IPAddress, None] = None
    enabled: bool = False


class Device(BaseModel):
    """Class to store credentials."""
    hostname: str
    interfaces: List[Interface]


# Functions
def read_args() -> argparse.Namespace:
    """Helper function to read CLI arguments."""
    parser = argparse.ArgumentParser(description="User input.")
    parser.add_argument("-d", "--data", type=str, help="Path to the input file.")
    parser.add_argument("-t", "--template", type=str, help="Path to the template.")
    return parser.parse_args()


def load_inventory(filename: str) -> List[Device]:
    """Function to load inventory data."""
    # Open file
    try:
        with open(filename, "r", encoding="utf-8") as f:
            data = yaml.safe_load(f)

    except FileNotFoundError as e:
        print(e)
        sys.exit(1)

    # Populate list of devices
    devices = []
    for device in data:
        devices.append(Device(**device))

    return devices


def load_template(filename: str) -> jinja2.Template:
    """Function to load Jinja2 template."""
    # Open file
    try:
        with open(filename, "r", encoding="utf-8") as file:
            return jinja2.Template(file.read())

    except FileNotFoundError as e:
        print(e)
        sys.exit(1)


def create_configuration(devices: List[Device], t: jinja2.Template) -> bool:
    """Function to create configuration files."""
    # Render template
    os.makedirs("output", exist_ok=True)
    try:
        for device in devices:
            with open(f"output/{device.hostname}.txt", "w", encoding="utf-8") as f:
                f.write(t.render(device=device))

        return True

    except Exception as e:
        print(e)
        return False


# Main
if __name__ == "__main__":
    # Get arguments
    args = read_args()

    # Load inventory
    try:
        inventory = load_inventory(args.data)
    except FileNotFoundError as e:
        print(e)
        sys.exit(1)

    # Load template
    template = load_template(args.template)

    # Create configuration
    if create_configuration(inventory, template):
        print("Configuration files created.")
    else:
        print("Something went wrong.")
        sys.exit(1)
