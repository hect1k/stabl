# stabl

**stabl** - Stabilise Traffic and Balance Load. `stabl` is a simple load balancer built in Go that utilizes a round-robin algorithm to distribute incoming traffic across multiple servers. This tool is designed to help manage and balance load, ensuring better performance and stability for your applications.

![stabl banner](https://dev.nnisarg.in/stabl/banner.jpg)

## Features

- **Round-Robin Load Balancing**: Distributes traffic evenly across multiple servers using a round-robin method.
- **Cross-Platform Support**: Compatible with Linux, macOS, and Windows.
- **Easy to Use**: Simple setup and configuration.

## Installation

1. **Download the latest release:**

   Visit the [Releases page](https://github.com/nnisarggada/stabl/releases/latest) to download the latest version of `stabl` for your operating system.

2. **Extract the downloaded file (if necessary)** and move the executable to your desired directory.

3. **Place the `config.json` file** in the same directory as the downloaded executable. The `config.json` file should include the port for the load balancer and the list of servers.

   ### Example `config.json`

   ```json
   {
     "port": 8080,
     "servers": [
       "http://server1.example.com",
       "http://server2.example.com",
       "http://server3.example.com"
     ]
   }
   ```

## Usage

After downloading and configuring `stabl`, you can pass the path to the config file as a flag. The default config path is `./config.json`.

- **Linux**:

  ```bash
  ./stabl_linux --config=path/to/config.json
  ```

- **macOS**:

  ```bash
  ./stabl_mac --config=path/to/config.json
  ```

- **Windows**:

  ```bash
  stabl_windows.exe --config=path/to/config.json
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
