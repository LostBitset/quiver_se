import os
import subprocess
import typing as t
import time
from tabulate import tabulate

class CsvWriter:
    class TwoHeadingsException(BaseException): pass
    class NoHeadingException(BaseException): pass
    
    def __init__(self, filename):
        self.filename = filename
        self.has_heading = False
    
    def write_heading(self, heading: t.List[str]):
        if self.has_heading:
            raise CsvWriter.TwoHeadingsException
        self.write_unchecked(heading)
        self.has_heading = True
    
    def write(self, data: t.List[str]):
        if not self.has_heading:
            raise CsvWriter.NoHeadingException
        self.write_unchecked(data)

    def write_unchecked(self, data: t.List[str]):
        line = ",".join(data) + "\n"
        with open(self.filename, "a") as f:
            f.write(line)

class EvaluationData:
    class UnknownAlgException(BaseException): pass

    # algnames: i.e. [("dse", "DSE"), ("simreq:simple", "SiMReQ (Simple)"), ...]
    def __init__(self, algnames: t.List[t.Tuple[str, str]]):
        self.counts = [] # i.e. [{"dse": 1, "simreq": 2}, ...]
        self.algnames = algnames
        self.algname_map = { k: v for (k, v) in algnames }
    
    def parse_eval_output(self, output: str) -> t.List[int]:
        count = {
            internal_name: 0
            for (internal_name, _) in self.algnames
        }
        watch_prefix = "[REPORT] __FOUND_A_BUG__"
        for line in output.split("\n"):
            if not line.startswith(watch_prefix):
                continue
            algname = line[len(watch_prefix):]
            if algname not in count:
                raise EvaluationData.UnknownAlgException
            count[algname] += 1
        self.counts.append(count)
        return [
            count[internal_name]
            for (internal_name, _) in self.algnames
        ]
    
    def display(self):
        counts = self.counts
        counts_last_run = [
            counts[len(counts)-1][internal_name]
            for (internal_name, _) in self.algnames
        ]
        counts_total = [
            sum([ count[internal_name] for count in counts ])
            for (internal_name, _) in self.algnames
        ]
        os.system("clear")
        results = [
            ["Last Run", *counts_last_run],
            ["In Total", *counts_total],
        ]
        print(tabulate(
            results,
            headers=["Bug counts...", *[ full_name for (_, full_name) in self.algnames ]],
            tablefmt="fancy_grid",
        ))

class EvaluationProxy:
    DEFAULT_ALGNAMES = [
        ("dse", "DSE"),
        ("simreq:simple", "SiMReQ (Simple)"),
        ("simreq:jitdse", "SiMReQ (JIT DSE)"),
    ]

    def __init__(self, csv_writer: CsvWriter, algnames = DEFAULT_ALGNAMES):
        self.csv_writer = csv_writer
        self.data = EvaluationData(algnames)
        csv_writer.write_heading([ full_name for (_, full_name) in algnames ])
    
    def run_once(self, display: bool = False):
        cmd = "./bin_synthetic"
        output = None
        try:
            output_bytes = subprocess.check_output(cmd, shell=True)
            output = output_bytes.decode("utf-8")
        except subprocess.CalledProcessError as e:
            output = e.output.decode("utf-8")
        if "exit status 2" in output:
            return None
        if "__FOUND_A_BUG__" in output:
            list_form = self.data.parse_eval_output(output)
            self.csv_writer.write([ str(i) for i in list_form ])
            if display:
                self.data.display()
    
    def run_forever(self, **kwargs):
        while True:
            self.run_once(**kwargs)
            os.system("rm /tmp/temp_qse-*")
            os.system("rm /tmp/go-build* -r")

if __name__ == "__main__":
    ep = EvaluationProxy(
        CsvWriter(f"eval_log-{int(time.time()*50)}.csv")
    )
    ep.run_forever(display=True)

