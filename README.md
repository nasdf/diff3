# Diff3

A diff3 text merge implementation in Go based on the awesome paper below.

["A Formal Investigation of Diff3" by Sanjeev Khanna, Keshav Kunal, and Benjamin C. Pierce](https://www.cis.upenn.edu/~bcpierce/papers/diff3-short.pdf)

## Usage

```go
import "github.com/nasdf/diff3"
textO := "original"
textA := "changesA"
textB := "changesB"
merge := diff3.Merge(textO, textA, textB)
```

### Customize seperators

```go
diff3.Sep1 = "$$$$$$$"
diff3.Sep2 = "@@@@@@@"
diff3.Sep3 = "*******"
```

### Customize DiffMatchPatch settings

```go
diff3.DiffMatchPatch.DiffTimeout = time.Second
diff3.DiffMatchPatch.DiffEditCost = 4
diff3.DiffMatchPatch.MatchThreshold = 0.5
diff3.DiffMatchPatch.MatchDistance = 1000
diff3.DiffMatchPatch.PatchDeleteThreshold = 0.5
diff3.DiffMatchPatch.PatchMargin = 4
diff3.DiffMatchPatch.MatchMaxBits = 32
```

## License

MIT