# go-keygrip

go-keygrip is a port of the node.js module for signing and verifying data through a rotating credential system, in which new server keys can be added and old ones removed regularly, without invalidating client credentials.

## API

```
keys = Keygrip(keylist)
```

This creates a new Keygrip based on the provided keylist.
```
import "github.com/17twenty/gokeygrip"
...
keys := gokeygrip.New(crypto.SHA1, BASE64,"06ae66fdc6c2faf5a401b70e0bf885cb")
withDefaults := gokeygrip.NewDefault("06ae66fdc6c2faf5a401b70e0bf885cb")
```
The keylist is an array of all valid keys for signing, in descending order of freshness; new keys should be unshifted into the array and old keys should be popped.

The tradeoff here is that adding more keys to the keylist allows for more granular freshness for key validation, at the cost of a more expensive worst-case scenario for old or invalid hashes.

Keygrip keeps a reference to this array to automatically reflect any changes. This reference is stored using a closure to prevent external access.
```
keys.Sign(data)
```
This creates a SHA1 HMAC based on the first key in the keylist, and outputs it as a 27-byte url-safe base64 digest (base64 without padding, replacing + with - and / with _).
```
keys.Index(data, digest)
```
This loops through all of the keys currently in the keylist until the digest of the current key matches the given digest, at which point the current index is returned. If no key is matched, -1 is returned.

The idea is that if the index returned is greater than 0, the data should be re-signed to prevent premature credential invalidation, and enable better performance for subsequent challenges.
```
keys.Verify(data, digest)
```
This uses index to return true if the digest matches any existing keys, and false otherwise.