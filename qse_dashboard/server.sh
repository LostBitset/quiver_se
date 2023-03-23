#!/bin/bash

python "$(git rev-parse --show-toplevel)/qse_dashboard/manage.py" runserver

