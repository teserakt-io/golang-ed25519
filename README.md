# Current Readme

NOTE: this code is no longer strictly necessary and is not actively being 
maintained. It existed briefly, because agl/ed25519 was deprecated and 
contained Edwards-to-Montgomery conversion code that was useful. 

Now:

 * If you just want ed25519 signatures in Golang, that is [available in the 
   standard library now](https://golang.org/pkg/crypto/ed25519).
 * If you want to manipulate points on the edwards curve and in particular 
   convert to Montgomery coordinates (the reason for this short lived fork to 
   exist), there is a maintained library by Filippo and [a function for that](https://pkg.go.dev/filippo.io/edwards25519#Point.BytesMontgomery). 

As always with crypto code you should be sure you understand what you are doing 
and that you have read the necessary documentation and background material to 
make sure you avoid some of the subtle bugs that may exist.

# Previous Readme

This repository contains two things:

 1. The Ed25519 code from the Golang standard library.
 2. The Extra25519 code that didn't survive the move from AGL's repository.

Unfortunately in their infinite wisdom, the Golang Standard Library maintainers 
have "internal"d the code doing Field and Group element manipulation, 
necessary for the extra25519 point conversion. This means that it is 
impossible to do:

```
import ( 
    edwards25519 "golang.org/crypto/ed25519/internal/edwards25519"
)
```

Thus all the code must be duplicated.

I have backported all fixes.

**It would be infinitely better if the Golang Standard Library team included 
the extra25519 code in their repository**.
