# CAPVCD fork notes

We maintain a fork of the **archived** upstream `vmware/cluster-api-provider-cloud-director`
(last upstream release `v1.3.2`). It exists only so the `VCDCluster` API types keep compiling
against newer `controller-runtime` / `cluster-api`; we consume it purely as an API-types dependency.

## Branches

- **`main`** — tracks the old upstream line. It has **diverged** from the v1.3.x releases and is
  *not* what we release or consume. Do not base fixes here.
- **`1.3.z`** — the v1.3.x maintenance line and the source of every `v1.3.x` tag. **This is where
  our patches go.**

## Tags / releases

Releases are cut from `1.3.z`. Each tag (`v1.3.2`, `v1.3.3`, …) is a plain patch bump carrying
compatibility fixes. `v1.3.3` adds: webhook migration to controller-runtime v0.24, and repointed
`api/` imports for the cluster-api v1.13 package layout (`api/v1beta1` → `api/core/v1beta1`).

To ship a new fix: commit onto `1.3.z`, then `git tag -a v1.3.N` on the new tip and push the tag.

## How it's consumed

`cluster-api-cleaner-cloud-director/go.mod` uses a `replace` directive to redirect the archived
`vmware` module path to this fork at a tagged version:

```
replace github.com/vmware/cluster-api-provider-cloud-director => github.com/giantswarm/cluster-api-provider-cloud-director v1.3.3
```

The `replace` is required because the code still imports the `vmware/...` path (the fork keeps that
module path in its own `go.mod`). Bump the version here when a new fork tag is released.

