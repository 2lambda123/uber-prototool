load("@bazel_gazelle//:deps.bzl", "go_repository")

def prototool_deps(**kwargs):
    go_repository(
        name = "co_honnef_go_tools",
        commit = "71123fcbb8fe",
        importpath = "honnef.co/go/tools",
    )
    go_repository(
        name = "com_github_burntsushi_toml",
        importpath = "github.com/BurntSushi/toml",
        tag = "v0.3.1",
    )
    go_repository(
        name = "com_github_client9_misspell",
        importpath = "github.com/client9/misspell",
        tag = "v0.3.4",
    )
    go_repository(
        name = "com_github_cpuguy83_go_md2man",
        importpath = "github.com/cpuguy83/go-md2man",
        tag = "v1.0.8",
    )
    go_repository(
        name = "com_github_davecgh_go_spew",
        importpath = "github.com/davecgh/go-spew",
        tag = "v1.1.1",
    )
    go_repository(
        name = "com_github_emicklei_proto",
        importpath = "github.com/emicklei/proto",
        tag = "v1.6.8",
    )
    go_repository(
        name = "com_github_fullstorydev_grpcurl",
        importpath = "github.com/fullstorydev/grpcurl",
        tag = "v1.1.0",
    )
    go_repository(
        name = "com_github_gobuffalo_flect",
        importpath = "github.com/gobuffalo/flect",
        tag = "v0.1.0",
    )
    go_repository(
        name = "com_github_gofrs_flock",
        importpath = "github.com/gofrs/flock",
        tag = "v0.7.0",
    )
    go_repository(
        name = "com_github_golang_glog",
        commit = "23def4e6c14b",
        importpath = "github.com/golang/glog",
    )
    go_repository(
        name = "com_github_golang_lint",
        commit = "8f45f776aaf1",
        importpath = "github.com/golang/lint",
    )
    go_repository(
        name = "com_github_golang_mock",
        importpath = "github.com/golang/mock",
        tag = "v1.1.1",
    )
    go_repository(
        name = "com_github_golang_protobuf",
        importpath = "github.com/golang/protobuf",
        tag = "v1.2.0",
    )
    go_repository(
        name = "com_github_inconshreveable_mousetrap",
        importpath = "github.com/inconshreveable/mousetrap",
        tag = "v1.0.0",
    )
    go_repository(
        name = "com_github_jhump_protoreflect",
        importpath = "github.com/jhump/protoreflect",
        tag = "v1.1.0",
    )
    go_repository(
        name = "com_github_kisielk_errcheck",
        importpath = "github.com/kisielk/errcheck",
        tag = "v1.2.0",
    )
    go_repository(
        name = "com_github_kisielk_gotool",
        importpath = "github.com/kisielk/gotool",
        tag = "v1.0.0",
    )
    go_repository(
        name = "com_github_kr_pretty",
        importpath = "github.com/kr/pretty",
        tag = "v0.1.0",
    )
    go_repository(
        name = "com_github_kr_pty",
        importpath = "github.com/kr/pty",
        tag = "v1.1.1",
    )
    go_repository(
        name = "com_github_kr_text",
        importpath = "github.com/kr/text",
        tag = "v0.1.0",
    )
    go_repository(
        name = "com_github_mitchellh_go_wordwrap",
        importpath = "github.com/mitchellh/go-wordwrap",
        tag = "v1.0.0",
    )
    go_repository(
        name = "com_github_pkg_errors",
        importpath = "github.com/pkg/errors",
        tag = "v0.8.1",
    )
    go_repository(
        name = "com_github_pmezard_go_difflib",
        importpath = "github.com/pmezard/go-difflib",
        tag = "v1.0.0",
    )
    go_repository(
        name = "com_github_russross_blackfriday",
        importpath = "github.com/russross/blackfriday",
        tag = "v1.5.2",
    )
    go_repository(
        name = "com_github_spf13_cobra",
        importpath = "github.com/spf13/cobra",
        tag = "v0.0.3",
    )
    go_repository(
        name = "com_github_spf13_pflag",
        importpath = "github.com/spf13/pflag",
        tag = "v1.0.3",
    )
    go_repository(
        name = "com_github_stretchr_objx",
        importpath = "github.com/stretchr/objx",
        tag = "v0.1.0",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        importpath = "github.com/stretchr/testify",
        tag = "v1.3.0",
    )
    go_repository(
        name = "com_google_cloud_go",
        importpath = "cloud.google.com/go",
        tag = "v0.26.0",
    )
    go_repository(
        name = "in_gopkg_check_v1",
        commit = "788fd7840127",
        importpath = "gopkg.in/check.v1",
    )
    go_repository(
        name = "in_gopkg_yaml_v2",
        importpath = "gopkg.in/yaml.v2",
        tag = "v2.2.2",
    )
    go_repository(
        name = "org_golang_google_appengine",
        importpath = "google.golang.org/appengine",
        tag = "v1.1.0",
    )
    go_repository(
        name = "org_golang_google_genproto",
        commit = "4b09977fb922",
        importpath = "google.golang.org/genproto",
    )
    go_repository(
        name = "org_golang_google_grpc",
        importpath = "google.golang.org/grpc",
        tag = "v1.18.0",
    )
    go_repository(
        name = "org_golang_x_exp",
        commit = "509febef88a4",
        importpath = "golang.org/x/exp",
    )
    go_repository(
        name = "org_golang_x_lint",
        commit = "8f45f776aaf1",
        importpath = "golang.org/x/lint",
    )
    go_repository(
        name = "org_golang_x_net",
        commit = "3a22650c66bd",
        importpath = "golang.org/x/net",
    )
    go_repository(
        name = "org_golang_x_oauth2",
        commit = "d2e6202438be",
        importpath = "golang.org/x/oauth2",
    )
    go_repository(
        name = "org_golang_x_sync",
        commit = "37e7f081c4d4",
        importpath = "golang.org/x/sync",
    )
    go_repository(
        name = "org_golang_x_sys",
        commit = "983097b1a8a3",
        importpath = "golang.org/x/sys",
    )
    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        tag = "v0.3.1-0.20180807135948-17ff2d5776d2",
    )
    go_repository(
        name = "org_golang_x_tools",
        commit = "dbb4d4be5353",
        importpath = "golang.org/x/tools",
    )
    go_repository(
        name = "org_uber_go_atomic",
        importpath = "go.uber.org/atomic",
        tag = "v1.3.2",
    )
    go_repository(
        name = "org_uber_go_multierr",
        importpath = "go.uber.org/multierr",
        tag = "v1.1.0",
    )
    go_repository(
        name = "org_uber_go_tools",
        commit = "ce2550dad714",
        importpath = "go.uber.org/tools",
    )
    go_repository(
        name = "org_uber_go_zap",
        importpath = "go.uber.org/zap",
        tag = "v1.9.1",
    )
