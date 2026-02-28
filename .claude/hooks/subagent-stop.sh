#!/bin/bash
# Hook: SubagentStop
# Fires when a subagent (spawned via the Agent tool) finishes its turn.
#
# Stub — no subagent coordination needed for deck development.

set -euo pipefail

echo "[hook:subagent-stop] Subagent finished"

exit 0
