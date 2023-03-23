# from django.shortcuts import render
from django.http import HttpResponse, JsonResponse

import pathlib

# Create your views here.

script_dir = pathlib.Path(__file__).parent.resolve()

index_html = None
with open(script_dir.joinpath("index.html"), "r") as f:
    index_html = f.read()

def index(req):
    return HttpResponse(index_html)

def live_json(req):
    return JsonResponse({"todo": True})

