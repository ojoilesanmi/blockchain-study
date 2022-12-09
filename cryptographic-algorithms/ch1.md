# Chapter 1: Deep Dive into Cryptography

## Ciphers
A cipher is a system that is used to transform plain test into cipher text or cryptogram.

Message -> Cipher -> cryptogram

Encryption is the operation of transforming a plaintext `m` into cryptogram `c` using a function `f`.
Decryption is the opposite of the Encryption operation.

### Private and Public Keys
The private key is used to encrypt and/or decrypt a plaintext.
The public key is used to only encrypt a message and verify digital signatures of humans
and computers. These two keys determines the difference between symmetric and asymmetric
encryption.

### Methods of Encryption
Symmetric Encryption uses one shared key to both encrypt and decrypt plaintext.
Asymmetric Encryption uses more parameters to generate a public key (and encrypt the
plaintext) and one private key to decrypt the plaintext.

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
