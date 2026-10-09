# runtime: preserve libc signal handlers and enable alternate signal stacks

## Summary

Update the Linux Go runtime to ensure that existing libc handlers for `SIGSETXID` and `SIGCANCEL` execute on an alternate signal stack when one is available.

These signals support glibc's thread-wide credential synchronization and pthread cancellation. Because libc may deliver them asynchronously to threads executing Go code, handlers that lack `SA_ONSTACK` can consume space on a Go-managed stack, risking stack overflow and memory corruption.

## Implementation requirements

* Preserve the original libc signal-handler function pointers.
* Enable `SA_ONSTACK` without replacing the handlers with Go runtime handlers.
* Preserve the existing signal mask and unrelated signal-action flags.
* Follow the runtime's established signal-stack initialization and platform-specific conventions.
* Avoid hard-coded signal numbers when portable runtime constants or existing abstractions are available.
* Ensure repeated initialization does not unintentionally remove required flags or overwrite handler state.

## Validation

* Build the runtime for supported Linux architectures, including 32-bit targets.
* Run relevant runtime and cgo tests.
* Verify that handler pointers and unrelated signal-action flags remain unchanged.
* Test mixed Go/C execution where practical.
* Confirm that the change does not regress ordinary signal handling or thread cancellation.

## Expected result

Improved stack safety for glibc signal handlers in mixed Go/C processes, without taking ownership of libc's internal signal handling.
