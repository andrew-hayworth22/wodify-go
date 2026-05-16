# CLAUDE.md

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

### Rules
- Always use exact struct field names and function names from the source — do not guess or infer
- Always use exact make target names from the Makefile
- Do not remove existing README content unless the corresponding code has been removed