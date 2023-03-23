import itertools
import subprocess
import os

def run_once():
    cmd = "go run ."
    output = None
    try:
        output_bytes = subprocess.check_output(cmd, shell=True)
        output = output_bytes.decode("utf-8")
    except subprocess.CalledProcessError as e:
        output = e.output.decode("utf-8")

    if "exit status 2" in output:
        return None

    n_bugs = {"dse": 0, "simreq": 0}
    n_sq = {"dse": 0, "simreq": 0}
    n_exec = {"dse": 0, "simreq": 0}
    target = None
    for line in output.split("\n"):
        if not line.startswith("[REPORT]"):
            continue
        evaluating = "[REPORT] [EVALUATING]"
        if line.startswith(evaluating):
            target = line[len(evaluating)+1:]
        elif "__FOUND_A_BUG__dse" in line:
            n_bugs["dse"] += 1
        elif "__FOUND_A_BUG__simreq" in line:
            n_bugs["simreq"] += 1
        elif "[EVAL-INFO] SOLVERQUERY" in line:
            n_sq[target] += 1
        elif "[EVAL-INFO] EXECUTION" in line:
            n_exec[target] += 1
    if n_bugs['dse'] == n_bugs['simreq']:
        if n_bugs['dse'] == 0:
            return None
    return {"bugs": n_bugs, "queries": n_sq, "executions": n_exec}

total_bugs = {"dse": 0, "simreq": 0}
total_queries = {"dse": 0, "simreq": 0}
total_execs = {"dse": 0, "simreq": 0}
n_samples = 0
for i in itertools.count():
    r = run_once()
    if r == None:
        continue
    n_samples += 1
    for alg in ["dse", "simreq"]:
        total_bugs[alg] += r["bugs"][alg]
        total_queries[alg] += r["queries"][alg]
        total_execs[alg] += r["executions"][alg]
    os.system("clear")
    print("(bug)\t\tDSE\tSiMReQ")
    print(f"BUGS FOUND\t{r['bugs']['dse']}\t{r['bugs']['simreq']}")
    print("(tot)\t\tDSE\tSiMReQ")
    print(f"BUGS FOUND\t{total_bugs['dse']}\t{total_bugs['simreq']}")
    print(f"SMT QUERIES\t{total_queries['dse']}\t{total_queries['simreq']}")
    #print(f"EXCS\t{total_execs['dse']}\t{total_execs['simreq']}")
    print("--- --- ---")

