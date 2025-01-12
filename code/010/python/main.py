"""From Python to Go: Python: 010 - Text files"""

import re


# Body
if __name__ == "__main__":
    # Get paths
    file_to_open = "../data/file.txt"
    file_to_save = "../data/output.txt"

    # Read file
    with open(file_to_open, "rt", encoding="utf-8") as f:
        data = f.read()

    # Print the raw text file
    print(data)

    # File is a multiline string, so split it to lines
    new_data = []
    for ind, line in enumerate(data.splitlines()):
        print(f"line {ind:>03}: {line}")

        # Make FQDN
        if re.match("^hostname:\s+.*$", line) and "network.karneliuk.com" not in line:
            line += ".network.karneliuk.com"

        # Copy line to new output
        new_data.append(line)

    # Save result file
    with open(file_to_save, "wt", encoding="utf-8") as f:
        f.write("\n".join(new_data))
