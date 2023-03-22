import subprocess

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

    target = None
    for line in output.split("\n"):
        print(line)
        if not line.startswith("[REPORT]"):
            continue
        

