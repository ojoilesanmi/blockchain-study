# Chapter 1: Deep Dive into Cryptography

## Ciphers
A cipher is a system that is used to transform plain test into cipher text or cryptogram.

Message -> Cipher -> cryptogram

Encryption is the operation of transforming a plaintext `m` into cryptogram `c` using a function `f`.
Decryption is the opposite of the Encryption operation.

### Private and Public Keys
The private key `Kpr` is used to encrypt and/or decrypt a plaintext.
The public key `Kpu` is used to only encrypt a message and verify digital signatures of humans
and computers. These two keys determines the difference between symmetric and asymmetric
encryption.

### Methods of Encryption
Symmetric Encryption uses one shared key to both encrypt and decrypt plaintext.
Asymmetric Encryption uses more parameters to generate a public key (and encrypt the
plaintext) and one private key to decrypt the plaintext.

Asymmetric encryption is used to exchange the key.

### Caesar Cipher
Shifting three places to encrypt, returning three places to decrypt. An implementation
is done in Go and Rust.

## Fermat's Last Theorem
In number theory, no three positive integers a, b, and c satisfy the equation `a^n + b^n = c^n`
for any integer value of n greater than 2. The cases n = 1 and n = 2 have infinitely many solutions.

If `p` is a prime number, then for any integer `(a)` elevated to the prime
number `(p)`, we find `(a)` as the result of the following equation:
```math
a^p = a (mod p)
```
For example:
If `a = 2` and `p = 3`, then `2^3 = 2 (mod 3)`.

Clearly, from n > 3, no integer a, b, or c, verifies the sum:
```math
a^n + b^n = c^n
```

### Beale Cipher
Involved using different letters in different texts to encrypt a message. Only
the second of the 3 pages have been cracked.

### Vernam Cipher
Vernam (a.k.a One Time Pad (OTP)) cipher has the highest degree of security a cipher can have. It is the 
perfect cipher. With Vernham, you use the key only once per session, making the 
algorithm invulnerable to attacks.

Steps involved:
1. Transform the plaintext into a string of bits using ASCII code
2. Generate a random key of the same length as the plaintext
3. Encrypt by adding modulo 2 (XOR) of the plaintext bitwise to the key and obtain the ciphertext
4. Decrypt by making the inverse operation of adding the ciphertext to the key and obtain the plaintext again.

