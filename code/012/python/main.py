"""From Python to Go: Python: 012 - User input."""

# Modules
import argparse
from getpass import getpass
from dataclasses import dataclass
import sys


# Classes
@dataclass
class Credentials:
    """Class to store credentials."""
    username: str
    password: str


# Functions
def read_args() -> argparse.Namespace:
    """Helper function to read CLI arguments."""
    parser = argparse.ArgumentParser(description="User input.")
    parser.add_argument("-p", "--path", type=str, help="Path to the input file.")
    return parser.parse_args()


def load_file(path: str) -> str:
    """Function to load a file."""
    with open(path, "r", encoding="utf-8") as file:
        return file.read()


# Main
if __name__ == "__main__":
    # Get arguments
    args = read_args()

    # Load file
    if args.path:
        print(load_file(args.path))

    # Exit if no path provided
    else:
        sys.exit("No path provided.")

    # Get user input
    creds = Credentials(username=input("Username: "), password=getpass("Password: "))
    print(f"{creds=}")
