#!/usr/bin/env bash
source ~/.profile
exec bazel run -- @io_bazel_rules_go//go/tools/gopackagesdriver "${@}"
