load("@io_bazel_rules_go//go:def.bzl", real_nogo = "nogo")
load("@com_github_sluongng_nogo_analyzer//:def.bzl", "nogo_config")

def nogo(name, deps, config, visibility):
    """
    nogo is a wrapper around the @go_rules/nogo rule.

    The benefits of this nogo are that we can ignore all external packages
    for linting without having to specify rules for every single one by
    defining the default config in the "overrides" arg as "_base".

    Example usage:
    nogo(
        name = "lint",
        config = {
            "_base": {
                "exclude_files": {
                    "external/": "third_party",
                    "proto/": "generated protobuf",
                },
            },
        },
        visibility = ["//visibility:public"],
        deps = TOOLS_NOGO,
    )

    Args:
        name: (Label) Name of nogo rule.
        deps: (LabelSet) Set of nogo analyzers to run.
        config: (struct) Struct in the form of the nogo json config.
        visibility: (Visibility) Visibility propogated to all the child targets.
    """

    # Define the names of the targets used below.
    config_label = "%s_nogo_config" % name

    # Generate the JSON for the override config.
    nogo_config(
        name = config_label,
        out = "%s_nogo_config.json" % name,
        analyzers = config.keys(),
        override = config,
    )

    real_nogo(
        name = name,
        deps = deps,
        config = config_label,
        visibility = visibility,
    )
