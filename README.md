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
