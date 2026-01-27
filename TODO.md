**Copy‑paste this entire markdown block into your documentation or issue tracker.**  
Every line that begins with `- [ ]` is a **checkbox**; tick them off as you complete each step. The checklist has been updated so that _Task 1_ (Phase 0) now reflects the current design – there is no CSV master‑matrix file; all cyborg definitions live in `.proto` files and their generated `*.txtpb` artefacts, which must be listed in **TODO.md**.

---

# ✅ Full‑Scale Blueprint – “Cyborg Conductor Core” (CCC) Build Plan

## Phase 0 – Boilerplate & Tooling

| #    | Task    | Status |
|------|---------|--------|
| 0️⃣1 | **Update `TODO.md` to be inline with the current state of the codebase / design**. The "Full Master Matrix" CSV has been removed. Cyborg definitions are now stored as protocol‑buffer (`.proto`) files under `proto/` and their generated Go/textpb artefacts (`*.txtpb`). Edit **TODO.md** to list these protobuf schemas, the generated package paths (`pkg/proto/pb/…`), and remove any references to a CSV file. | ✅ Complete |
| 0️⃣2 | Clone / initialise repository – `git clone <repo‑url> && cd cyborg-conductor-core` _(skip if already present)_ | ✅ Complete |
| 0️⃣3 | Install Go **≥ 1.22** (verify with `go version`) | ✅ Complete |
| 0️⃣4 | Install protobuf tooling | ✅ Complete |
| 0️⃣5 | Create workspace folders – `mkdir -p pkg/proto/generated test/unit test/integration` | ✅ Complete |
| 0️⃣6 | Install extra code‑generators (e.g., `protoc-gen-grpc-web@latest`) | ✅ Complete |

---

## Phase 1 – Define & Store the Core Schemas

| #   | Task    | Status |
| --- | ------ | ------ |
| 1‑1 | **Create** `proto/cyborg.proto` containing the rebranded _CyborgDescriptor_ schema. | ✅ Complete |
| 1‑2 | Add all system‑envelope schemas to files under `proto/system_*.proto`. | ✅ Complete |
| 1‑3 | Generate any supporting documentation from these protobufs (e.g., Markdown via `cogenerate`). | ✅ Complete |

---

## Phase 2 – Compile Protobuf Definitions

| #    | Task    | Status |
|------|---------|--------|
| 2️⃣1 | Export output paths | ✅ Complete |
| 2️⃣2 | **Generate Go bindings** from every `.proto` file | ✅ Complete |
| 2️⃣3 | **Generate textpb bindings** (if you use Go's `txtpb` library) | ✅ Complete |
| 2️⃣4 | Run a quick compile sanity check | ✅ Complete |

---

## Phase 3 – Core Data Structures & Generated Types

| #   | Task    | Status |
| --- | ------ | ------ |
| 3‑1 | Add generated protobuf packages to **go.mod** | ✅ Complete |
| 3‑2 | Create `pkg/core/pb/definition.go` with strong‑typed structs | ✅ Complete |
| 3‑3 | Implement **registry** (`pkg/core/pb/registry.go`) containing: `Register(*pb.CyborgDescriptor) error` `Get(id string) (*pb.CyborgDescriptor, bool)` `List() []*pb.CyborgDescriptor` (mutex‑protected). | ✅ Complete |
| 3‑4 | Add **MemoryManager** in `pkg/memory` | ✅ Complete |
| 3‑5 | Stub **ContextOverlayEngine** (`internal/context/overlay.go`) with a method: `GetSnapshot(ctx, cyborgID) []byte` | ✅ Complete |

---

## Phase 4 – Implement Core Runtime Modules

| #   | Task    | Status |
| --- | ------ | ------ |
| 4‑1 | Write **Scheduler** (`pkg/core/orchestrator/scheduler.go`) that holds a back‑pressure aware worker pool and selects a cyborg based on capability match, latency budget, and current load. | ⏳ In Progress |
| 4‑2 | Implement **SubprocessRunner** in `internal/runner/exec_manager.go` (timeout handling via `context.WithTimeout`, capture stdout/stderr). | ✅ Complete |
| 4‑3 | Add **Python adapter glue** (`adapters/python/exec_wrapper.py`) plus Bash entrypoint (`runner.sh`). | ✅ Complete |
| 4‑4 | Add **Node.js adapter glue** (`adapters/node/src/index.ts`) exposing an async `runScript(name:string, args:[]string): Promise<{out:string, err:string}>`. | ✅ Complete |
| 4‑5 | Wire the **ContextOverlayEngine** to load immutable evidence snapshots from `$EVIDENCE_ROOT` and pass buffers to outgoing protobuf messages. | ⏳ In Progress |
| 4‑6 | Build **MemoryCacheManager** (`context/manager.go`) that enforces a maximum context size, automatically truncates excess data, and compresses with LZ4 when needed. | ⏳ In Progress |

---

## Phase 5 – Register & Discover Cyborgs

| #   | Task    | Status |
| --- | ------ | ------ |
| 5‑1 | Create the **registration service** in `cmd/server/main.go` that reads _all_ protobuf definitions from the generated Go packages and builds a full **CyborgDescriptor** for each. | ✅ Complete |
| 5‑2 | Add a gRPC method `RegisterCyborg(req *pb.RegisterRequest) (*pb.RegisterResponse, error)` to validate uniqueness, non‑empty tags, parsable deployment spec, etc. | ⏳ In Progress |
| 5‑3 | Write & run a **test registration script** (e.g., Python client iterating over generated descriptors). | ⏳ In Progress |
| 5‑4 | Add a lightweight **watchdog** that logs `$EVIDENCE_ROOT` size vs. `MAX_CONTEXT_BYTES`. | ⏳ In Progress |

---

## Phase 6 – Adaptive Scheduler & Back‑Pressure

| #   | Task    | Status |
| --- | ------ | ------ |
| 6‑1 | Implement scoring logic in `scheduler.go` | ⏳ In Progress |
| 6‑2 | Read **resource quotas** (`max_concurrent`, etc.) from each cyborg's `config_blob` | ⏳ In Progress |
| 6‑3 | Integrate back‑pressure | ⏳ In Progress |
| 6‑4 | Write unit tests | ⏳ In Progress |

---

## Phase 7 – LLM Streaming Sessions

| #   | Task    | Status |
| --- | ------ | ------ |
| 7‑1 | Create a **LLMClient** struct in `pkg/llm-client/` | ⏳ In Progress |
| 7‑2 | Implement an HTTP streaming wrapper | ⏳ In Progress |
| 7‑3 | Build a **session manager** (`internal/llm/session.go`) | ⏳ In Progress |
| 7‑4 | Unit‑test the flow | ⏳ In Progress |

---

## Phase 8 – Observability & Admin Interfaces

| #   | Task    | Status |
| --- | ------ | ------ |
| 8‑1 | Expose **Prometheus `/metrics`** endpoint | ✅ Complete |
| 8‑2 | Add an admin HTTP endpoint (`/api/v1/status`) | ✅ Complete |
| 8‑3 | Provide a health probe `/healthz` | ✅ Complete |
| 8‑4 | Create a minimal Dockerfile for the server binary | ✅ Complete |
| 8‑5 | Add a sample `docker-compose.yml` in `test/integration/` | ⏳ In Progress |
| 8‑6 | Confirm that all metrics are non‑zero and health checks pass | ⏳ In Progress |
| 8-7 | Implement proper structured logging with zap library | ⏳ In Progress |

---

## Phase 9 – Full Test Suite

| #   | Task    | Status |
| --- | ------ | ------ |
| 9️⃣1 | Run **all unit tests** with coverage | ⏳ In Progress |
| 9️⃣2 | Ensure overall coverage ≥ 95 % | ⏳ In Progress |
| 9️⃣3 | Execute the **integration suite** | ⏳ In Progress |
| 9️⃣4 | Add a GitHub Actions workflow to publish a coverage badge | ✅ Complete |

---

## Phase 10 – Production Deployment Blueprint

| #   | Task    | Status |
| --- | ------ | ------ |
| 10‑1 | **Write Helm chart** (`charts/cyborg-conductor-core`) | ⏳ In Progress |
| 10‑2 | Add a **ConfigMap** containing the compiled `.proto` files | ⏳ In Progress |
| 10️⃣3 | Deploy to Kubernetes | ⏳ In Progress |
| 10‑4 | Implement a **blue‑green upgrade** script | ⏳ In Progress |
| 10‑5 | Enable automatic restarts and watchdog monitoring | ⏳ In Progress |

---

## ✅ Current Status

The core system has the basic framework implemented but is **NOT yet production-ready**. Several critical features are still incomplete:

- ✅ Fixed config loading bug
- ✅ Implemented Cyborg Registry with txtpb loading
- ✅ Added CI/CD pipeline with GitHub Actions
- ✅ Added linter configuration
- ✅ Updated documentation
- ✅ Created Ops README
- ✅ Complete protobuf integration
- ✅ Proper error handling and logging

**Next Steps (to achieve production readiness):**

1. Complete Scheduler implementation
2. Add gRPC registration service
3. Implement LLM streaming features
4. Complete Prometheus metrics
5. Add comprehensive unit tests (coverage ≥95%)
6. Deploy Helm chart
7. Add integration tests
8. Add HTTPS/TLS support
9. Implement proper structured logging (currently basic logging)
10. Enhance health checks with DB connectivity verification

**System is not yet production-ready** - several key components are missing or incomplete. 🛠️