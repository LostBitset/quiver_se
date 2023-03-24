import subprocess
import typing as t

class CsvWriter:
    class TwoHeadingsException(BaseException): pass
    class NoHeadingException(BaseException): pass

    def __init__(self, filename):
        self.filename = filename
        self.has_heading = False
    
    def write_heading(self, heading: t.List[string]):
        if self.has_heading:
            raise TwoHeadingsException
        self.write_unchecked(heading)
    
    def write(self, data: t.List[string]):
        if not self.has_heading:
            raise NoHeadingException
        self.write_unchecked(data)

    def write_unchecked(self, data: t.List[string])
        line = ",".join(data) + "\n"
        with open(self.filename, "a") as f:
            f.write(line)

class EvaluationData:
    # algnames: i.e. [("dse", "DSE"), ("simreq:simple", "SiMReQ (Simple)"), ...]
    def __init__(self, algnames: t.List[(string, string)]):
        self.counts = [] # i.e. [{"dse": 1, "simreq": 2}, ...]
        self.algnames = algnames
        self.algname_map = { k: v for (k, v) in algnames }
    
    def parse_eval_output(self, output: string) -> t.List[int]:
        count = {}
        watch_prefix = "[REPORT] __FOUND_A_BUG__"
        for line in output.split("\n"):
            if not line.startswith(watch_prefix):
                continue
            algname = line[watch_prefix:]
            if algname in count:
                count[algname] = 0
            count[algname] += 1
        self.counts.append(count)
        return [ count[internal_name] for (internal_name, _) in self.algnames ]
    
    def display(self):
        ... # TODO

class EvaluationProxy:
    DEFAULT_ALGNAMES = [
        ("dse", "DSE"),
        ("simreq:simple", "SiMReQ (Simple)"),
        ("simreq:jitdse", "SiMReQ (JIT DSE)"),
    ]

    def __init__(self, csv_writer: CsvWriter, algnames = self.__class__.DEFAULT_ALGNAMES):
        self.csv_writer = csv_writer
        self.data = EvaulationData(algnames)
        csv_writer.write_heading([ full_name for (_, full_name) in algnames ])
    
    def run_once(self):
        cmd = "go run ."
        output = None
        try:
            output_bytes = subprocess.check_output(cmd, shell=True)
            output = output_bytes.decode("utf-8")
        except subprocess.CalledProcessError as e:
            output = e.output.decode("utf-8")
        if "exit status 2" in output:
            return None
        list_form = self.data.parse_eval_output(output)
        self.csv_writer.write([ str(i) for i in list_form ])
