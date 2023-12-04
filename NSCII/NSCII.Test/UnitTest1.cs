using NUnit.Framework;
using NSCII.Service;

namespace NSCII.Test
{
    public class NSCIITests
    {
        [Test]
        public void TestEncryption()
        {
            string input = "EncryptedStringHere";

            string encrypted = App.EncryptNSCII(input);
            
            string decrypted = App.DecryptNSCII(encrypted);

            Assert.That(decrypted, Is.EqualTo(input));
        }

        [Test]
        public void TestDecryption()
        {
            string encryptedInput = "EncryptedStringHere";

            string decrypted = App.DecryptNSCII(encryptedInput);

            string reEncrypted = App.EncryptNSCII(decrypted);

            Assert.That(reEncrypted, Is.EqualTo(encryptedInput));
        }
    }
}
