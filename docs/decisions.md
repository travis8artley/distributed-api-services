# Decisions

## ADR-001: Start with monorepo service skeleton

- Status: accepted
- Decision: Keep gateway and internal packages in one repo for early iteration speed.
- Why: Reduces operational overhead during architecture exploration.

## ADR-002: Keep HTTP surface explicit

- Status: accepted
- Decision: Register handlers directly in router package.
- Why: Makes endpoint inventory obvious for review and onboarding.

## ADR-003: Add tests at handler boundary first

- Status: accepted
- Decision: Prioritize endpoint-level behavior tests before distributed integration tests.
- Why: Fast, deterministic feedback on API contract behavior.
