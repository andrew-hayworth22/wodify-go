# CLAUDE.md

## Behavior

- Act as a conversational assistant, not an autonomous agent.
- Do **not** make code changes unless the user explicitly asks you to.
- When showing code suggestions, print them in the terminal for the user to read — do not write them to files unless asked.

## Task: Update the README

When asked to update the README, check each of the following sources and reflect any changes.

### 1. Makefile
Read `Makefile` and check for new or removed targets. Update these README sections accordingly:
- New targets under `# Examples` → add to the **Examples** section
- New targets under `# Tests` → add to the **Testing** section
- New targets under `# Profiling` → add to the **Profiling** section

### 2. Lead service methods
Read `leads/service.go` and check for new or removed methods on `*Client`. For each new method:
- Add a usage example to the **Leads** section of the README
- Read `leads/request.go` to get the exact request struct name and field names for the example
- Read `leads/response.go` to get the exact response struct name if needed

### 3. New domain packages
If a new top-level domain package has been added (e.g. `appointments`, `clients`):
- Add a new section to the README for that domain, following the same pattern as the **Leads** section
- Check the domain's `service.go`, `request.go`, and `response.go` for method signatures and struct names

### 4. Error sentinels
Read `errors.go` in the root package. If new sentinel errors have been added, update the sentinel table in the **Error Handling** section.

### 5. Configuration options
Read `options.go`. If new `With*` options have been added, update the configuration table in the **Configuration** section.

### 6. CONTRIBUTING.md
After updating the README, check whether any structural changes warrant updating `CONTRIBUTING.md` as well:
- New or removed Makefile targets used as examples → update the **Getting started** section
- New domain packages added → update the **Project structure** section
- Go version change in `go.mod` → update the **Prerequisites** line

### Rules
- Always use exact struct field names and function names from the source — do not guess or infer
- Always use exact make target names from the Makefile
- Do not remove existing README content unless the corresponding code has been removed
- **Do NOT update the Domain Coverage table** unless the user explicitly asks you to update it
- When updating the Domain Coverage table, always order rows by status: `✅ complete` first, then `🚧 in-development`, then `⏳ coming soon`

### Domain Coverage table structure

The table uses exactly three possible status values:

| Status | Meaning |
|---|---|
| `✅ complete` | Domain is fully implemented in the SDK |
| `🚧 in-development` | Domain is partially implemented / actively being worked on |
| `⏳ coming soon` | Domain is not yet implemented |