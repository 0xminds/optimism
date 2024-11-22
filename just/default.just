set shell := ["bash", "-c"]

PARALLEL_JOBS := num_cpus()

MAP_JUST := "/usr/bin/env -S parallel --shebang --jobs " + PARALLEL_JOBS + " --colsep ' ' -r " + just_executable()
