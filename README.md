## OpenMyShell

OpenMyShell is a Go-based tool designed to quickly create custom OpenVPN configuration files with embedded shell payloads. It's perfect for security researchers, network administrators, and penetration testers who need an efficient way to inject shell commands into `.ovpn` files.

## How It Works

1. **Install**: 
   - Ensure you have Go installed on your system.
   - Clone the repository and build the program.

   ```sh
   git clone https://github.com/yourusername/OpenMyShell.git
   cd OpenMyShell
   ```

2. **Usage**:
   - Run the program with command-line arguments:

   ```sh
   go run main.go <ip> <port> <filename>
   ```

3. **(TO ADD) XOR Encryption to the Payload**:
   - Encrypt the payload using XOR before writing it to the file to evade security measures, then decrypt it at runtime.


This will create `myconfig.ovpn` with an embedded reverse shell payload targeting `192.168.1.10` on port `4444`.

## Contributing

Feel free to contribute by submitting issues or pull requests.
