---
description: Guided Go debugging session
agent: ask
---

Help me debug this Go problem: ${input:problem}

Do not immediately provide the fix.

Guide me through:

1. expected vs actual behavior;
2. assumptions;
3. hypotheses;
4. the smallest experiment that distinguishes the hypotheses;
5. diagnosis;
6. fix and why it works;
7. regression test or experiment.

For concurrency problems, explicitly consider races, deadlocks, goroutine leaks, cancellation, and ownership.