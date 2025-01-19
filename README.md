
# Go Blockchain 🚀

**Go Blockchain** is a simple implementation of a blockchain using Golang. The goal of this project is to demonstrate how a blockchain can be built and validated in a simple and practical way.

## 🔧 Technologies used

- Golang
- Blockchain

## ⬇️ How to install and run

1. Clone this repository to your local machine:
   ```bash
   git clone https://github.com/Fepipo/go-blockchain.git
   ```
2. Navigate to the folder where the `main.go` file is located:
   ```bash
   cd go-blockchain
   ```
3. Run the project with the command below, passing as arguments the number of blocks to create and the difficulty (number of 0s at the beginning of the hash):
   ```bash
   go run main.go <number_of_blocks> <difficulty>
   ```
   Example:
   ```bash
   go run main.go 10 4
   ```
   This will create 10 blocks with a difficulty of 4 (4 zeros at the beginning of the hash).

## ⚙️ Key Features

- **Create blocks:** The project allows the creation of blocks that form the chain.
- **Validate blocks:** Each new block is validated to ensure the integrity of the blockchain.

## 🤝 How to contribute

Currently, the project does not have a specific contribution guide. However, you can contribute by opening issues or creating pull requests. Any contributions are welcome!

## 📝 License

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for more information.

## 📞 Contact

If you have any questions or would like to get in touch, feel free to email me at: [felipemmarcioli@gmail.com](mailto:felipemmarcioli@gmail.com)
