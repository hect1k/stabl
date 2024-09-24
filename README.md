# stabl - Stabilize Traffic and Balance Load

**stabl** is a lightweight load balancer built in Go, using a round-robin algorithm to efficiently distribute incoming traffic across multiple servers. This tool helps to ensure better performance and reliability by stabilizing traffic and balancing load across your application servers.

![stabl banner](https://dev.nnisarg.in/stabl/banner.jpg)

# Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Features

- **Round-Robin Load Balancing**: Evenly distributes incoming requests to servers using a round-robin method.
- **Cross-Platform Compatibility**: Works on Linux, macOS, and Windows.
- **Easy Setup**: Minimal configuration required to get started.
- **Customizable Settings**: Modify load balancing behavior through a simple configuration file.
- **Logging**: Logs load balancer activity to a file for easy monitoring and debugging.

## Installation

### Step 1: Download the latest release

Go to the [Releases page](https://github.com/nnisarggada/stabl/releases/latest) and download the appropriate version of `stabl` for your operating system.

### Step 2: Extract and move the executable

If required, extract the downloaded file and place the `stabl` executable in your preferred directory.

### Step 3: Prepare the configuration file

Create or move the `config.json` file to the same directory as the executable. The configuration file should define the port for the load balancer and the list of servers you want to balance.

## Configuration

An example `config.json` file:

```json
{
  "port": 8080,
  "checkAfter": 5,
  "servers": [
    "http://server1.example.com",
    "http://server2.example.com",
    "http://server3.example.com"
  ],
  "logFile": "stabl.log"
}
```

- **`port`**: The port on which `stabl` will listen for incoming traffic.
- **`checkAfter`**: Interval in seconds after which the health of the servers is checked.
- **`servers`**: List of backend servers to distribute traffic.
- **`logFile`**: File path for storing log information about load balancer activity.

## Usage

After downloading and configuring `stabl`, you can run it by providing the path to the configuration file, the port, and the log file path using command-line flags. The default config path is `./config.json`, and the default port is `8080`.

**Note:** Command-line flags will override the settings in the `config.json` file.

### Linux

```bash
./stabl --config=path/to/config.json --port=8080 --logFile=path/to/log/file
```

### macOS

```bash
./stabl_darwin --config=path/to/config.json --port=8080 --logFile=path/to/log/file
```

### Windows

```bash
stabl.exe --config=path/to/config.json --port=8080 --logFile=path/to/log/file
```

## Contributing

We welcome contributions to improve `stabl`. Here's how you can contribute:

1. Fork the repository on GitHub.
2. Create a new branch for your changes.
3. Make your modifications and commit them.
4. Push your branch to GitHub.
5. Open a pull request with a description of the changes you've made.

## License

This project is licensed under the [GNU General Public License v3.0](LICENSE).

## Contact

For questions, suggestions, or reporting issues, please open an issue on our [GitHub repository](https://github.com/nnisarggada/stabl).
