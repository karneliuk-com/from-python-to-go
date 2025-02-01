"""From Python to Go: Python: 013 - Exception handling."""

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
    try:
        with open(path, "r", encoding="utf-8") as file:
            return file.read()

    except FileNotFoundError:
        sys.exit(f"File not found: {path}. Check the path and try again.")

    except Exception as e:
        sys.exit(f"Error: {e}")


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
    try:
        creds = Credentials(username=input("Username: "), password=getpass("Password: "))
        if not creds.password:
            raise Exception("No password is provided!")

    except Exception as e:
        print(f"Recovering from: {e}")

    print(f"{creds=}")
