using System;

namespace NSCII.Service
{
    public class App
    {
      public static void Main(string[] args)
        {
            var Line = "---------------------------------------------------------------------------------";

            Console.WriteLine(Line);

            Console.WriteLine("Please enter information to encrypt/decrypt NSCII : ");

            string input = Console.ReadLine();

            Console.WriteLine(Line);

            Console.WriteLine("Enter '1' for encryption and '2' for decryption:");

            int choice = int.Parse(Console.ReadLine());

            Console.WriteLine(Line);

            if (choice == 1)
            {
                var encryptedNSCII = EncryptNSCII(input);
                Console.WriteLine("Encrypted NSCII: " + encryptedNSCII);
            }
            else if (choice == 2)
            {
                var decryptedNSCII = DecryptNSCII(input);
                Console.WriteLine("Decrypted NSCII: " + decryptedNSCII);
            }
            else
            {
                Console.WriteLine("Invalid choice. Please enter '1' for encryption or '2' for decryption.");
            }
            Console.WriteLine(Line);

        }

    public static string EncryptNSCII(string inputNSCII)
        {
            var encoder = "9876543210ABCDEFGHIJKLMNOPQRSTUVWXYZzyxwvutsrqponmlkjihgfedcba";
            var result = "";
            int num_plus = 1;

            for (int i = 0; i < inputNSCII.Length; i++)
            {
                char character = inputNSCII[i];
                var index = encoder.IndexOf(character);

                if (num_plus == 27)
                {
                    num_plus = 1;
                }

                var int_nscii = (index + num_plus) % encoder.Length;

                result = result + encoder[int_nscii];
                num_plus++;
            }

            return result;
        }

    public static string DecryptNSCII(string encryptedNSCII)
        {
            var encoder = "9876543210ABCDEFGHIJKLMNOPQRSTUVWXYZzyxwvutsrqponmlkjihgfedcba";
            var result = "";
            int num_plus = 1;

            for (int i = 0; i < encryptedNSCII.Length; i++)
            {
                char character = encryptedNSCII[i];
                var index = encoder.IndexOf(character);

                if (num_plus == 27)
                {
                    num_plus = 1;
                }

                var int_nscii = (index - num_plus + encoder.Length) % encoder.Length;

                result = result + encoder[int_nscii];
                num_plus++;
            }

            return result;
        }

     
    }
}
