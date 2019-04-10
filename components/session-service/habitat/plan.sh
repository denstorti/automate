pkg_name=session-service
pkg_description="A2 session service"
pkg_origin=chef
pkg_version="0.1.0"
pkg_maintainer="Chef Software Inc. <support@chef.io>"
pkg_license=('Chef-MLSA')
pkg_upstream_url="http://github.com/chef/automate/components/session-service"
pkg_deps=(
  core/bash
  core/curl # health_check hook
  core/glibc
  chef/mlsa
  ${local_platform_tools_origin:-chef}/automate-platform-tools
)
pkg_build_deps=(
  core/gcc
  core/git # for ref in version
)
pkg_exports=(
  [port]=service.port # default service is http
)
pkg_exposes=(
  port
)
pkg_binds=(
  [automate-dex]="port"
  [automate-pg-gateway]="port"
  [pg-sidecar-service]="port"
)
pkg_bin_dirs=(bin)
pkg_scaffolding="${local_scaffolding_origin:-chef}/automate-scaffolding-go"
scaffolding_go_base_path=github.com/chef
scaffolding_go_repo_name=automate
scaffolding_go_import_path="${scaffolding_go_base_path}/${scaffolding_go_repo_name}/components/${pkg_name}"
scaffolding_go_binary_list=(
  "${scaffolding_go_import_path}/cmd/${pkg_name}"
)

do_strip() {
    return 0;
}

do_prepare() {
  GIT_SHA=$(git rev-parse HEAD)
  GO_LDFLAGS=" -X ${scaffolding_go_base_path}/automate/lib/version.Version=${pkg_release}"
  GO_LDFLAGS="${GO_LDFLAGS} -X ${scaffolding_go_base_path}/automate/lib/version.GitSHA=${GIT_SHA}"
  GO_LDFLAGS="${GO_LDFLAGS} -X ${scaffolding_go_base_path}/automate/lib/version.BuildTime=${pkg_release}"
  export GO_LDFLAGS
  build_line "Setting GO_LDFLAGS=${GO_LDFLAGS}"
}

do_install() {
  # Go scaffolding install callback
  scaffolding_go_install

  build_line "Copying migration files"
  cp -r migration/sql "${pkg_prefix}/migrations"
}
