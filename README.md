[![GoDoc](https://godoc.org/github.com/opencoff/go-sign?status.svg)](https://godoc.org/github.com/opencoff/go-sign)

# README for sigtool


## What is this?
`sigtool` is an opinionated tool to generate keys, sign, verify, encrypt &
decrypt files using Ed25519 signature scheme.  In many ways, it is like
like OpenBSD's [signify][1] -- except written in Golang and definitely
easier to use.

It can sign and verify very large files - it prehashes the files
with SHA-512 and then signs the SHA-512 checksum. The keys and signatures
are YAML files and so, human readable.

It can encrypt files for multiple recipients - each of whom is identified
by their Ed25519 public key. The encryption by default generates ephmeral
Curve25519 keys and creates pair-wise shared secret for each recipient of
the encrypted file. The caller can optionally use a specific secret key
during the encryption process - this has the benefit of also authenticating
the sender (and the receiver can verify the sender if they possess the
corresponding public key).


## How do I build it?
With Go 1.5 and later:

    git clone https://github.com/opencoff/sigtool
    cd sigtool
    make

The binary will be in `./bin/$HOSTOS-$ARCH/sigtool`.
where `$HOSTOS` is the host OS where you are building (e.g., openbsd)
and `$ARCH` is the CPU architecture (e.g., amd64).

## How do I use it?
Broadly, the tool can:

- generate new key pairs (public key and private key)
- sign a file
- verify a file against its signature
- encrypt a file
- decrypt a file

### Generate Key pair
To start with, you generate a new key pair (a public key used for
verification and a private key used for signing). e.g.,

    sigtool gen /tmp/testkey

The tool then generates */tmp/testkey.pub* and */tmp/testkey.key*.  The secret
key (".key") can optionally be encrypted with a user supplied pass
phrase - which the user has to enter via interactive prompt:

    sigtool gen -p /tmp/testkey

### Sign a file
Signing a file requires the user to provide a previously generated
Ed25519 private key.  The signature (YAML) is written to STDOUT.
e.g.,  to sign `archive.tar.gz` with private key `/tmp/testkey.key`:

    sigtool sign /tmp/testkey.key archive.tar.gz

If *testkey.key* was encrypted with a user pass phrase:

    sigtool sign -p /tmp/testkey.key archive.tar.gz


The signature can also be written directly to a user supplied output
file.

    sigtool sign -p -o archive.sig /tmp/testkey.key archive.tar.gz


### Verify a signature against a file
Verifying a signature of a file requires the user to supply three
pieces of information:

- the Ed25519 public key to be used for verification
- the Ed25519 signature
- the file whose signature must be verified

e.g., to verify the signature of *archive.tar.gz* against
*testkey.pub* using the signature *archive.sig*

    sigtool verify /tmp/testkey.pub archive.sig archive.tar.gz

### Encrypt a file by authenticating the sender
If the sender wishes to prove to the recipient that they  encrypted
a file:

   sigtool encrypt -s sender.key to.pub -o archive.tar.gz.enc archive.tar.gz


This will create an encrypted file *archive.tar.gz.enc* such that the
recipient in possession of *to.key* can decrypt it. Furthermore, if
the recipient has *sender.pub*, they can verify that the sender is indeed
who they expect.

### Decrypt a file and verify the sender
If the receiver has the public key of the sender, they can verify that
they indeed sent the file by cryptographically checking the output:

   sigtool decrypt -o archive.tar.gz -v sender.pub to.key archive.tar.gz.enc

Note that the verification is optional and if the `-v` option is not
used, then decryption will proceed without verifying the sender.

### Encrypt a file *without* authenticating the sender
`sigtool` can generate ephemeral keys for encrypting a file such that
the receiver doesn't need to authenticate the sender:

   sigtool encrypt to.pub -o archive.tar.gz.enc archive.tar.gz

This will create an encrypted file *archive.tar.gz.enc* such that the
recipient in possession of *to.key* can decrypt it.

## Technical Details

### How is the private key protected?
The Ed25519 private key is encrypted using a key derived from the
user supplied pass phrase. This pass phrase is used to derive an
encryption key using the Scrypt key derivation algorithm. The
resulting derived key is XOR'd with the Ed25519 private key before
being committed to disk. To protect the integrity of the process,
the essential parameters used for deriving the key, and the derived
key are hashed via SHA256 and stored along with the encrypted key.

As an additional security measure, the user supplied pass phrase is
hashed with SHA512.

### How is the Encryption done?
The file encryption uses AES-GCM-256 in AEAD mode. The encryption uses
a random 32-byte AES-256 key. The input is broken into chunks and
each chunk is individually AEAD encrypted. The default chunk size
is 4MB (4 * 1048576 bytes). Each chunk generates its own nonce
from a global salt. The nonce is calculated as a SHA256 hash of
the salt, the chunk length and the block number.

### What is the public-key cryptography used?
`sigtool` uses Curve25519 ECC to generate shared secrets between
pairs of sender & recipients. This pairwise shared secret is expanded
using HKDF to generate a key-encryption-key. The file-encryption key
is AEAD encrypted with this key-encryption-key. Thus, each recipient
has their own individual encrypted key blob.

The Ed25519 keys generated by `sigtool` are transformed to their 
corresponding Curve25519 points in order to generate the shared secret.
This elliptic co-ordinate transform follows [FiloSottile's writeup][2].

### Format of the Encrypted File
Every encrypted file starts with a header:

    7 byte magic ("SigTool")
    1 byte version number
    4 byte header length (big endian encoding)
    32 byte SHA256 of the encryption-header

The encryption-header is described as a protobuf file (sign/hdr.proto):

```protobuf
    message header {
        uint32 chunk_size = 1;
        bytes  salt = 2;
        repeated wrapped_key keys = 3;
    }

    message wrapped_key {
        bytes pk_hash = 1; // hash of Ed25519 PK
        bytes pk = 2;       // curve25519 PK
        bytes nonce = 3;    // AEAD nonce
        bytes key = 4;      // AEAD encrypted key
    }
```

The encrypted data immediately follows the headers above. Each encrypted
chunk is encoded the same way:

```C
    4 byte chunk length (big endian encoding)
    chunk data
    AEAD tag
```

The chunk data and AEAD tag are treated as an atomic unit for AEAD
decryption.

## Understanding the Code
`src/sign` is a library to generate, verify and store Ed25519 keys
and signatures.  It uses the extended library (golang.org/x/crypto)
for the underlying operations.

`src/crypt.go` contains the encryption & decryption code.

The generated keys and signatures are proper YAML files and human
readable.

The signature file contains a hash of the public key - so that at
verification time, the right private key may be used (in situations
where there are lots of keys).

Signatures on large files are calculated efficiently by reading them
in memory mapped mode (```mmap(2)```) and hashing the file contents
using SHA-512. The Ed25519 signature is calculated on the file-hash.

## Example of Keys, Signature

### Ed25519 Public Key
A serialized Ed25519 public key looks like so:

    pk: uxpDh+gqXojAmxA/6vxZHzA+Uk+8wogUwvEhPBlWgvo=

### Ed25519 Private Key
And, a serialized Ed25519 private key looks like so:

```yaml

    esk: t3vfqHbgUiA733KKPymFjWT8DdnBEkiMfsDHolPUdQWpvVn/F1Z4J6KYV3M5rGO9xgKxh5RAmqt+6LKgOiJAMQ==
    salt: pPHKG55UJYtJ5wU0G9hBvNQJ0DvT0a7T4Fmj4aPB84s=
    algo: scrypt-sha256
    verify: JvjRjJMKhJhBmZngC3Pvq7x3KCLKt7gar1AAz7HB4qM=
    Z: 131072
    r: 16
    p: 1
```

The Ed25519 private key is encrypted using Scrypt password hashing
mechanism. A user supplied passphrase to protect the private key
is first pre-hashed using SHA-512 before being used in
```scrypt()```. In pseudo code, this operation looks like below:

    passphrase = get_user_passphrase()
    hpass      = SHA512(passphrase)
    salt       = randombytes(32)
    xorkey     = Scrypt(hpass, salt, N, r, p)
    verify     = SHA256(salt, xorkey)
    esk        = ed25519_private_key ^ xorkey

Where, ```N```, ```r```, ```p``` are Scrypt parameters. In our
implementation:

    N = 131072
    r = 16
    p = 1

```verify```  is used during the decryption of the Ed25519 private
key - *before* actually doing the "xor" operation. This check
ensures that the supplied passphrase yields the same value as
```verify```.

### Ed25519 Signature
A generated signature looks like below after serialization:

```yaml

    comment: inpfile=/tmp/file.txt
    pkhash: 36z9tCwTIVNwwDlExrB0SQ==
    signature: ow2oBP+buDbEvlNakOrsxgB5Yc/7PYyPVZCkfyu7oahw8BakF4Qf32uswPaKGZ8RVz4uXboYHdZtfrEjCgP/Cg==
```

Here, ```pkhash`` is a SHA256 of the public key needed to verify
this signature.

## Licensing Terms
The tool and code is licensed under the terms of the
GNU Public License v2.0 (strictly v2.0). If you need a commercial
license or a different license, please get in touch with me.

See the file ``LICENSE.md`` for the full terms of the license.

## Author
Sudhi Herle <sw@herle.net>

[1]: https://www.openbsd.org/papers/bsdcan-signify.html
[2]: https://blog.filippo.io/using-ed25519-keys-for-encryption/
