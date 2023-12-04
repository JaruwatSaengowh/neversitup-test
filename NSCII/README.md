# NSCII Project

## Project Structure

`NSCII/
|-- NSCII/
|   |-- Program.cs
|-- NSCII.Test/
|   |-- UnitTest1.cs
|-- NSCII.sln
|-- README.md` 

## Description

The NSCII project is designed to perform encryption and decryption using the Natnuphe Standard Code for Information Interchange (NSCII). This README provides an overview of the project structure, usage instructions, and details about the encryption algorithm.

## Usage

To use the NSCII application:

1.  Clone the repository.
2.  Open the solution file `NSCII.sln` in your preferred IDE.
3.  Run the `NSCII` project.

Follow the on-screen instructions to enter information for NSCII encryption or decryption. Choose the appropriate option and view the results.

## Project Structure

### `NSCII.Service.App`

#### `Main(string[] args)`

-   Entry point of the application.
-   Accepts user input for NSCII encryption or decryption.
-   Handles user choices and calls appropriate methods.

#### `EncryptNSCII(string inputNSCII)`

-   Encrypts the given NSCII string using the Natnuphe encryption algorithm.

#### `DecryptNSCII(string encryptedNSCII)`

-   Decrypts the given encrypted NSCII string using the Natnuphe encryption algorithm.

### `NSCII.Test.NSCIITests`

#### `TestEncryption()`

-   Tests the `EncryptNSCII` method to ensure correct encryption.

#### `TestDecryption()`

-   Tests the `DecryptNSCII` method to ensure correct decryption.

## Natnuphe Encryption Algorithm

The Natnuphe encryption algorithm operates as follows:

-   Each character is mapped to a corresponding index in the encoder string.
-   The encryption involves adding 1 to 26 to each character sequentially.
-   If the sum exceeds 26, it wraps around to 1.
-   Decryption involves subtracting 1 to 26 from each character sequentially.

## Example

Consider the following standards for NSCII:

`9..0A..Zz..a` 


-   `9 + 1` results in `8`
-   `8 + 2` results in `6`
-   `9 - 1` results in `a`
-   `A + 1` results in `B`
-   `Z + 1` results in `z`
-   `z + 1` results in `y`

The Natnuphe encryption involves continuously adding 1 to 26 to each character, wrapping around when necessary.

This example illustrates how the Natnuphe encryption works for NSCII.

Feel free to customize this template further to include specific details about your project or additional information that you find relevant.# NSCII Project