## Prologue:

| Padding bits at end | Prefix bits you write | Meaning                   |
| ------------------- | --------------------- | ------------------------- |
| 0                   | `11111110`            | no padding (8 valid bits) |
| 1                   | `1111110`             | 7 valid bits              |
| 2                   | `111110`              | 6 valid bits              |
| 3                   | `11110`               | 5 valid bits              |
| 4                   | `1110`                | 4 valid bits              |
| 5                   | `110`                 | 3 valid bits              |
| 6                   | `10`                  | 2 valid bits              |
| 7                   | `0`                   | 1 valid bit               |


# Huffman coding

## Important resources
- for everything https://en.wikipedia.org/wiki/Huffman_coding
- for weird https://en.wikipedia.org/wiki/Canonical_Huffman_code
- for zip https://pkware.cachefly.net/webdocs/casestudies/APPNOTE.TXT
- for storing tree (VERY IMPORTANT) https://stackoverflow.com/questions/759707/efficient-way-of-storing-huffman-tree

## How it works (stolen from wiki)

Given string:
```py
"A_DEAD_DAD_CEDED_A_BAD_BABE_A_BEADED_ABACA_BED""
```

1. Characters sorted by freq:
```
C: 2
B: 6
E: 7
_: 10
D: 10
A: 11
```

2. Construct tree from first 2 (`C`, `B`):
```
  CB 8
 /0   \1
C 2  B: 6
```

And move it back to it's place:
```
E: 7
CB: 8
_: 10
D: 10
A: 11
```

3. Construct tree from first 2 (`E`, `CB`)
```
  ECB 15
 /0    \1
E 7   CB 8
      /0 \1
    C 2   B 6
```

4. Continue till it's:
```
      _DAECB 46
     /0        \1
   _D 20        AECB 26
  /0   \1      /0     \1
_ 10   D 10   A 11    ECB 15
                      /0    \1
                     E 7    CB 8
                           /0   \1
                         C 2     B 6
```
And you get:
```
_: 00
D: 01
A: 10
E: 110
C: 1110
B: 1111
```
brilliant!
