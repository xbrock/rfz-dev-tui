# Component Gallery Rendering Issues (Lite)

Fix two critical visual rendering bugs in the Component Gallery demo: (1) TuiBox borders appear disconnected/floating when boxes are placed side-by-side, and (2) TuiButton rows show staggered/misaligned layout instead of horizontal alignment. Root cause is using string concatenation instead of `lipgloss.JoinHorizontal()` for multi-line component layout.

**3 stories, ~100 LOC changes, Severity: Critical, Priority: Urgent**

## Quick Summary

| Issue | Component | Fix |
|-------|-----------|-----|
| Broken box borders | gallery.go renderBoxSection() | Use lipgloss.JoinHorizontal(lipgloss.Top, ...) |
| Staggered buttons | gallery.go renderButtonSection() | Use lipgloss.JoinHorizontal(lipgloss.Top, ...) |

## Stories

1. **FIX-001**: Fix TuiBox Border Rendering (S, 2 SP)
2. **FIX-002**: Fix TuiButton Layout Alignment (S, 2 SP)
3. **FIX-003**: Add Regression Tests (XS, 1 SP)

## Verification

```bash
go build ./cmd/rfz/...
./rfz  # Visual check
go test ./internal/ui/components/... -v
golangci-lint run ./internal/ui/components/...
```
