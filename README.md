# PBKDF2-DoS-POC
Demonstration of PBKDF2's DoS vulnerability when implemented according to spec.

## Explanation
PBKDF2 performs an HMAC of the user's password on each iteration. This becomes an issue when users submit very large passwords, as it causes a CPU/memory bandwidth denial of service that could affect the functioning of a web service.

## Solution
As HMAC performs length normalization of its input in its first step, it is possible to perform this length normalization once and then reusing that normalized value for each iteration. This solves the denial of service issue but it is not mentioned in RFC 8081. In fact, it's non-compliant as it relies on utilising HMAC as the base authentication function. Similar optimizations are likely possible for other authentication hashes though. Most implementations of PBKDF2 already perform this optimization, although it's not very well documented.
