# Hash Functions
Any size input, fixed size output, and the output is "random" looking.
You can do everything with just hash functions. Signatures, encryptiong, etc.
You have a data, run it to a hash function, and get an output.
```math
hash(data) -> output
```
SHA256 is used for hash function, in the following spaces:
- cryptography in blockchains
- database file encryption

# Hashing  Properties
The Avalanche Effect:
Changing 1 bit of the input results in a change in about half of the output bits.

Preimage Resistance:
Given `y`, you can't find any `x` such that:
```math
hash(x) == y
```
You can find it eventually, but that will take `2^256` operations (`10^78`).

Collision Resistance:
There is no `x`, `z` pair such that `x != z` but
```math
hash(x) == hash(z)
```
You can find them eventually, but not before `2^256`.

Collision Resistance have been broken for popular algorithms like `SHA-1`,  and `MD5`, but the Preimage Resistance property still remains intact.

## Uses of Hash functions
1. Naming files (where immutability is important)
2. As references or pointers in algorithms
3. As commitments like zero-knowledge proofs
4. Merkle Tree: a binary tree of hashes, where you take data, and hash it again and again.  It's a useful data structure.
5. A blockchain: is a chain of hashes.

## Best practices
1. Salt your hash functions.

Questions:
1. Where can you use Hash Functions since it can be used everywhere?

# Signatures
These are useful messages signed by someone. Signature schemes are created using 3 functions:
- `GenerateKeys()`: receives randomness, and returns a private and public key pair.
- `Sign(secretKey, message)`: receives a secret/private key and a message, and returns a signature.
- `Verify(publicKey, message, signature)`: verifies a signature on a message from a public key.
It takes a public key, message, and signature, and it returns a boolean value whether it worked or not.

**Note:** The `publicKey` is your identity, and the `secretKey` is a private key used to prove your identity.

Examples of Signature schemes include:
- DSA
- ElGamal
- RSA
- Elliptic curve signatures.

You can generate signature schemes from hashes. This is called "Lamport Signatures"

# TO DO EXERCISES
* [ ] Implement Signature in Go and Rust
* [ ] Implement Hashing in Go and Rust
* [ ] Pset: Implement a signature system using only hashes.
