# stabl

**stabl** - Stabilise Traffic and Balance Load. `stabl` is a simple load balancer built in Go that utilizes a round-robin algorithm to distribute incoming traffic across multiple servers. This tool is designed to help manage and balance load, ensuring better performance and stability for your applications.

![stabl banner](https://dev.nnisarg.in/stabl/banner.jpg)

## Features

- **Round-Robin Load Balancing**: Distributes traffic evenly across multiple servers using a round-robin method.
- **Cross-Platform Support**: Compatible with Linux, macOS, and Windows.
- **Easy to Use**: Simple setup and configuration.
- **Customizable**: Configurable to suit your needs.
- **Simple Logging**: Write load balancer activity to a file.

## Installation

1. **Download the latest release:**

   Visit the [Releases page](https://github.com/nnisarggada/stabl/releases/latest) to download the latest version of `stabl` for your operating system.

2. **Extract the downloaded file (if necessary)** and move the executable to your desired directory.

3. **Place the `config.json` file** in the same directory as the downloaded executable. The `config.json` file should include the port for the load balancer and the list of servers.

   ### Example `config.json`

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

## Usage

After downloading and configuring `stabl`, you can pass the path to the config file, the port number, and the path to the log file as flags. The default config path is `./config.json`. The default port is `8080`.

**‼️ NOTE: Any values passed as flags will override the values in the config file.**

- **Linux**:

  ```bash
  ./stabl --config=path/to/config.json --port=8080 --logFile=path/to/log/file
  ```

- **macOS**:

  ```bash
  ./stabl_darwin --config=path/to/config.json --port=8080 --logFile=path/to/log/file
  ```

- **Windows**:

  ```bash
  stabl.exe --config=path/to/config.json --port=8080 --logFile=path/to/log/file
  ```

## Contributing

If you have suggestions, bug reports, or pull requests, please follow these steps:

1. Fork the repository.
2. Create a new branch for your changes.
3. Commit your changes and push the branch.
4. Open a pull request describing your changes.

## License

`stabl` is released under the [GNU General Public License v3.0](LICENSE).

## Contact

For questions or issues, please open an issue on the [GitHub repository](https://github.com/nnisarggada/stabl).
