# go-keygrip

go-keygrip is a port of the node.js module for signing and verifying data through a rotating credential system, in which new server keys can be added and old ones removed regularly, without invalidating client credentials.

## API

```
keys = Keygrip(keylist)
```

This creates a new Keygrip based on the provided keylist.
```
var Keygrip = require('keygrip')
keys = Keygrip(["SEKRIT2", "SEKRIT1"])
```
The keylist is an array of all valid keys for signing, in descending order of freshness; new keys should be unshifted into the array and old keys should be popped.

The tradeoff here is that adding more keys to the keylist allows for more granular freshness for key validation, at the cost of a more expensive worst-case scenario for old or invalid hashes.

Keygrip keeps a reference to this array to automatically reflect any changes. This reference is stored using a closure to prevent external access.

When using Keygrip to encrypt and decrypt data, each key's length is important, as it should be at least the minimum key length for the cipher you are using, otherwise it'll be padded with NULs. The default cipher, AES 256, should be a 32 character string key, for example.
```
var buf = keys.sign(data)
```
This creates a HMAC based on the first key in the keylist, and outputs it as a buffer.

Uses .hash= as the underlying algorithm.
```
var index = keys.indexOf(data)
```
This loops through all of the keys currently in the keylist until the digest of the current key matches the given digest, at which point the current index is returned. If no key is matched, -1 is returned.

The idea is that if the index returned is greater than 0, the data should be re-signed to prevent premature credential invalidation, and enable better performance for subsequent challenges.
```
var bool = keys.verify(data)
```
This uses index to return true if the digest matches any existing keys, and false otherwise.
```
var buf = keys.encrypt(message, [iv])
```
Creates an encrypted message as a buffer based on the first key in the keylist and optionally based on an initialization vector.

Uses .cipher= as the underlying algorithm. Note that iv length is important.
```
var [buf, i] = keys.decrypt(message, [iv])
```
Decrypts a message, optionally with an initialization vector. Returns a buffer as buf. Also returns i, the index of the key used. If i !== 0, you may want to re-encrypt the message to use the latest key.
```
keys.hash=
```
Set the hashing algorithm for signing, defaulting to sha256.
```
.cipher=
```
Set the algorithm used for message encryption, defaulting to aes-256-cbc.