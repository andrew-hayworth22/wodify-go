# Contributing to wodify-go

wodify-go is a community-maintained Go SDK for the Wodify API. It is not officially supported by Wodify. Contributions, bug reports, and feedback of any kind are welcome and appreciated.

---

## Getting started

**Prerequisites**

- Go 1.26.2+
- A Wodify API key

**Clone and set up**

```bash
git clone https://github.com/andrew-hayworth22/wodify-go.git
cd wodify-go
cp .env.example .env   # fill in your WODIFY_API_KEY
go mod download
```

**Run the tests**

```bash
make test
```

**Run an example**

```bash
make leads-crud
make clients-search
# see Makefile for the full list
```

---

## Project structure

```
wodify-go/
├── wodify.go          # Top-level client; entry point for SDK users
├── options.go         # WithAPIKey, WithBaseURL, etc.
├── errors.go          # Sentinel error values
├── requests.go        # NewPaginationRequest, SortAscending, SortDescending helpers
├── version.go         # SDK version constant
├── models/            # Shared model types (Client, Lead, Date, etc.)
├── leads/             # Leads domain (methods, request/response types)
├── clients/           # Clients domain
├── utils/             # Utility lookups (countries, genders, timezones, etc.)
├── internal/
│   ├── httpclient/    # HTTP transport, retries, error parsing
│   ├── query/         # Query builder for search requests
│   ├── request/       # Generic list/search request types
│   ├── sort/          # Sort helpers
│   ├── testutil/      # Shared test utilities (not part of public API)
│   └── version/       # SDK version string
└── examples/          # Runnable usage examples per domain
```

Each domain package (e.g. `leads`, `clients`) owns its own methods, request structs, response structs, and field sentinel values. The `internal/` packages are shared infrastructure not exposed to SDK users.

---

## Opening a pull request

**Before you start**

For anything beyond a small bug fix, open an issue first to discuss the approach. This avoids wasted effort if the change isn't a good fit.

**Branch naming**

```
feat/add-appointments-domain
fix/pagination-off-by-one
chore/update-go-version
```

**PR checklist**

- [ ] Tests added or updated for the changed code
- [ ] `make test` passes with no failures
- [ ] New public types and functions have doc comments
- [ ] Examples added under `examples/` for any new domain methods
- [ ] No unrelated changes bundled in

**PR description**

Include:
- What the change does and why
- Any Wodify API endpoints added/changed (link to docs if available)
- How to test it manually (e.g. `make leads-crud`)

---

## Adding a new domain

The pattern for a new domain (e.g. `appointments`) is:

1. Create `appointments/` with at minimum:
   - `http_client.go` — the `Client` struct and constructor
   - `appointments.go` — methods, request types, response types, and field sentinels
   - `appointments_test.go` — tests
2. Wire it into the top-level `Client` in `wodify.go`
3. Add `make test-appointments` to the Makefile
4. Add at least one runnable example under `examples/appointments/`

Follow the existing `leads` package as the reference implementation.

---

## Reporting issues

Open a GitHub issue with:
- A minimal reproduction (code snippet or `make` target)
- The Go version (`go version`) and wodify-go version
- The API endpoint involved, if known

---

## License

By contributing you agree that your work will be licensed under the same [MIT License](LICENSE) as this project.