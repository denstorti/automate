package backup

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chef/automate/components/automate-deployment/pkg/services"
	"github.com/chef/automate/lib/platform/command"
)

// TestDefaultSpecs tests that all service specs have their desired backup
// operations.
func TestDefaultSpecs(t *testing.T) {
	chefServerEnabled := true
	workflowEnabled := true
	t.Run("authn-service", func(t *testing.T) {
		spec := testSpecFor(t, "authn-service", chefServerEnabled, workflowEnabled)
		testRequireMetadata(t, spec)
		testRequireDb(t, spec, "chef_authn_service", "authn")
	})
	t.Run("authz-service", func(t *testing.T) {
		spec := testSpecFor(t, "authz-service", chefServerEnabled, workflowEnabled)
		testRequireMetadata(t, spec)
		testRequireDb(t, spec, "chef_authz_service", "authz")
	})
	t.Run("automate-dex", func(t *testing.T) {
		spec := testSpecFor(t, "automate-dex", chefServerEnabled, workflowEnabled)
		testRequireMetadata(t, spec)
		testRequireDb(t, spec, "dex", "dex")
	})
	t.Run("automate-elasticsearch", func(t *testing.T) {
		testSpecFor(t, "automate-elasticsearch", chefServerEnabled, workflowEnabled)
	})
	t.Run("automate-gateway", func(t *testing.T) {
		testSpecFor(t, "automate-gateway", chefServerEnabled, workflowEnabled)
	})
	t.Run("automate-load-balancer", func(t *testing.T) {
		testSpecFor(t, "automate-load-balancer", chefServerEnabled, workflowEnabled)
	})
	t.Run("automate-postgresql", func(t *testing.T) {
		testSpecFor(t, "automate-postgresql", chefServerEnabled, workflowEnabled)
	})
	t.Run("automate-ui", func(t *testing.T) {
		testSpecFor(t, "automate-ui", chefServerEnabled, workflowEnabled)
	})
	t.Run("compliance-service", func(t *testing.T) {
		spec := testSpecFor(t, "compliance-service", chefServerEnabled, workflowEnabled)
		testRequireMetadata(t, spec)
		testRequireDb(t, spec, "chef_compliance_service", "compliance")
		testRequireEs(t, spec, "compliance-service", "comp-*")
	})
	t.Run("config-mgmt-service", func(t *testing.T) {
		testSpecFor(t, "config-mgmt-service", chefServerEnabled, workflowEnabled)
	})
	t.Run("data-lifecycle-service", func(t *testing.T) {
		testSpecFor(t, "data-lifecycle-service", chefServerEnabled, workflowEnabled)
	})
	t.Run("deployment-service", func(t *testing.T) {
		spec := testSpecFor(t, "deployment-service", chefServerEnabled, workflowEnabled)
		testRequireMetadata(t, spec)
		testRequirePath(t, spec, "/hab/user/deployment-service/config")
		testRequirePath(t, spec, "/hab/svc/deployment-service/data", Exclude("bolt.db"))
		testRequireCmd(t, spec, "deployment-service dumpdb", "deployment-service restoredb")
	})
	t.Run("es-sidecar-service", func(t *testing.T) {
		testSpecFor(t, "es-sidecar-service", chefServerEnabled, workflowEnabled)
	})
	t.Run("ingest-service", func(t *testing.T) {
		spec := testSpecFor(t, "ingest-service", chefServerEnabled, workflowEnabled)
		testRequireMetadata(t, spec)
		testRequirePath(t, spec, "/hab/svc/ingest-service/data")
		testRequireEs(t, spec, "ingest-service", "node-state-5,node-attribute,converge-history-*,actions-*")
	})
	t.Run("license-control-service", func(t *testing.T) {
		spec := testSpecFor(t, "license-control-service", chefServerEnabled, workflowEnabled)
		testRequireMetadata(t, spec)
		testRequirePath(t, spec, "/hab/svc/license-control-service/data")
		testRequirePath(t, spec, "/hab/svc/license-control-service/config", Include("opt_out"), Exclude("*"))
	})
	t.Run("local-user-service", func(t *testing.T) {
		testSpecFor(t, "local-user-service", chefServerEnabled, workflowEnabled)
	})
	t.Run("notifications-service", func(t *testing.T) {
		spec := testSpecFor(t, "notifications-service", chefServerEnabled, workflowEnabled)
		testRequireMetadata(t, spec)
		testRequireDb(t, spec, "notifications_service", "notifications")
	})
	t.Run("teams-service", func(t *testing.T) {
		spec := testSpecFor(t, "teams-service", chefServerEnabled, workflowEnabled)
		testRequireMetadata(t, spec)
		testRequireDb(t, spec, "chef_teams_service", "teams")
	})
	t.Run("session-service", func(t *testing.T) {
		spec := testSpecFor(t, "session-service", chefServerEnabled, workflowEnabled)
		testRequireMetadata(t, spec)
		testRequireDb(t, spec, "chef_session_service", "session")
	})
	t.Run("automate-cs-bookshelf", func(t *testing.T) {
		spec := testSpecFor(t, "automate-cs-bookshelf", chefServerEnabled, workflowEnabled)
		testRequireMetadata(t, spec)
		testRequireDb(t, spec, "automate-cs-bookshelf", "automate-cs-bookshelf")
	})
	t.Run("automate-cs-oc-bifrost", func(t *testing.T) {
		spec := testSpecFor(t, "automate-cs-oc-bifrost", chefServerEnabled, workflowEnabled)
		testRequireMetadata(t, spec)
		testRequireDb(t, spec, "automate-cs-oc-bifrost", "automate-cs-oc-bifrost")
	})
	t.Run("automate-cs-oc-erchef", func(t *testing.T) {
		spec := testSpecFor(t, "automate-cs-oc-erchef", chefServerEnabled, workflowEnabled)
		testRequireMetadata(t, spec)
		testRequireDb(t, spec, "automate-cs-oc-erchef", "automate-cs-oc-erchef")
		testRequireEs(t, spec, "automate-cs-oc-erchef", "chef")
	})
	t.Run("nodemanager-service", func(t *testing.T) {
		spec := testSpecFor(t, "nodemanager-service", chefServerEnabled, workflowEnabled)
		testRequireMetadata(t, spec)
		testRequireDb(t, spec, "nodemanager_service", "nodemanager")
	})
}

func TestSetCommandExecutor(t *testing.T) {
	testCmdExec := command.NewMockExecutor(t)
	t.Run("sets sync path defaults", func(t *testing.T) {
		s := Spec{Name: "test",
			SyncPaths: []PathCopyOperation{
				{Name: "foo", SrcPath: "/tmp/test"},
			},
		}
		SetCommandExecutor(s, testCmdExec)

		require.Equal(t, testCmdExec, s.SyncPaths[0].cmdExecutor, "command executor must match")
	})
	t.Run("sets sync cmd defaults", func(t *testing.T) {
		s := Spec{Name: "postgres", SyncCmds: []CommandExecuteOperation{
			{Name: "sync-pgdump", Cmd: Cmd{
				Name: "sync", Dump: []string{"pgdump"},
			}},
		}}
		SetCommandExecutor(s, testCmdExec)

		require.Equal(t, testCmdExec, s.SyncCmds[0].cmdExecutor, "command executor must match")
	})
}

func TestSetDefaults(t *testing.T) {
	t.Run("sets sync path defaults", func(t *testing.T) {
		s := Spec{Name: "test",
			SyncPaths: []PathCopyOperation{
				{Name: "foo", SrcPath: "/tmp/test"},
			},
		}
		setDefaults(s)

		require.Equal(t, []string{"test", "foo"}, s.SyncPaths[0].ObjectName)
	})
	t.Run("sets sync cmd defaults", func(t *testing.T) {
		s := Spec{Name: "postgres", SyncCmds: []CommandExecuteOperation{
			{Name: "sync-pgdump", Cmd: Cmd{
				Name: "sync", Dump: []string{"pgdump"},
			}},
		}}
		setDefaults(s)

		require.Equal(t, []string{"postgres", "sync-pgdump"}, s.SyncCmds[0].ObjectName)
		require.Equal(t, "chef", s.SyncCmds[0].PkgOrigin)
		require.Equal(t, "postgres", s.SyncCmds[0].PkgName)
	})
}

func TestChefServerDisabled(t *testing.T) {
	usualServices, err := services.ServicesInCollection("automate-full")
	require.NoError(t, err)

	chefServerServices, err := services.ServicesInCollection("chef-server")
	require.NoError(t, err)
	chefServerEnabled := false
	workflowEnabled := false

	// Expect all of the usual a2 services to have a backup spec.
	for _, svcpkg := range usualServices {
		svcname := svcpkg.Name()
		require.Truef(t, hasSpec(svcname, chefServerEnabled, workflowEnabled), "%s should have a spec", svcname)
	}

	// When the chef-server is not enabled, expect all of the chef-server services not
	// to have a backup spec
	for _, svcpkg := range chefServerServices {
		svcname := svcpkg.Name()
		assert.Falsef(t, hasSpec(svcname, chefServerEnabled, workflowEnabled), "%s should not have a spec", svcname)
	}
}

func TestWorkflowDisabled(t *testing.T) {
	usualServices, err := services.ServicesInCollection("automate-full")
	require.NoError(t, err)

	workflowServices, err := services.ServicesInCollection("workflow")
	require.NoError(t, err)
	chefServerEnabled := false
	workflowEnabled := false

	// Expect all of the usual a2 services to have a backup spec.
	for _, svcpkg := range usualServices {
		svcname := svcpkg.Name()
		require.Truef(t, hasSpec(svcname, chefServerEnabled, workflowEnabled), "%s should have a spec", svcname)
	}

	// When Workflow is not enabled, expect all of the Workflow services not
	// to have a backup spec
	for _, svcpkg := range workflowServices {
		svcname := svcpkg.Name()
		assert.Falsef(t, hasSpec(svcname, chefServerEnabled, workflowEnabled), "%s should not have a spec", svcname)
	}
}

func testRequirePath(t *testing.T, spec Spec, path string, matchers ...RsyncMatcher) {
	var p PathCopyOperation
	for _, sp := range spec.SyncPaths {
		if sp.SrcPath == path {
			p = sp
		}
	}
	require.Equal(t, path, p.SrcPath, "could not find path matching %s", path)
	for _, dm := range matchers {
		match := false
		for _, pm := range p.RsyncMatchers {

			pmFlag, pmArgs := pm.RsyncArgs()
			dmFlag, dmArgs := dm.RsyncArgs()

			if pmFlag == dmFlag && pmArgs == dmArgs {
				match = true
			}
		}
		f, a := dm.RsyncArgs()
		require.Truef(t, match, "could not find required matcher (%s %s) for path %s", f, a, path)
	}
}

func findDbByName(name string, dbs []DatabaseDumpOperationV2) DatabaseDumpOperationV2 {
	for _, d := range dbs {
		if d.Name == name {
			return d
		}
	}

	return DatabaseDumpOperationV2{}
}

func testRequireDb(t *testing.T, spec Spec, name string, user string) {
	p := findDbByName(name, spec.SyncDbsV2)
	assert.Equal(t, name, p.Name, "could not find database matching %s", name)
	assert.Equal(t, user, p.User, "username for database %s doesn't match %s", name, user)
}

func testRequireCmd(t *testing.T, spec Spec, dumpCmd, restoreCmd string) {
	dc := strings.Split(dumpCmd, " ")
	rc := strings.Split(restoreCmd, " ")

	var c CommandExecuteOperation
	for _, sc := range spec.SyncCmds {
		if reflect.DeepEqual(sc.Cmd.Dump, dc) && reflect.DeepEqual(sc.Cmd.Restore, rc) {
			c = sc
		}
	}
	require.Truef(t, reflect.DeepEqual(dc, c.Cmd.Dump),
		"could not find command with Dump matching %s and Restore matching %s",
		dumpCmd, restoreCmd,
	)
}

func testRequireMetadata(t *testing.T, spec Spec) {
	require.True(t, spec.WriteMetadata, "spec should have WriteMetadata enabled")
}

func findEsByName(name string, esops []ElasticsearchOperation) ElasticsearchOperation {
	for _, e := range esops {
		if e.ServiceName == name {
			return e
		}
	}

	return ElasticsearchOperation{}
}

func testRequireEs(t *testing.T, spec Spec, serviceName string, multiIndexSpec string) {
	op := findEsByName(serviceName, spec.SyncEsIndices)
	assert.Equal(t, serviceName, op.ServiceName, "could not find es operation for %s", serviceName)
	assert.Equal(t, multiIndexSpec, op.MultiIndexSpec, "incorrect index spec for %s", serviceName)
}

func testSpecFor(t *testing.T, desired string, chefServerEnabled bool, workflowEnabled bool) Spec {
	for _, s := range DefaultSpecs(chefServerEnabled, workflowEnabled) {
		if s.Name == desired {
			return s
		}
	}

	require.Truef(t, false, "could not find spec for %s", desired)
	return Spec{}
}

func hasSpec(desired string, chefServerEnabled bool, workflowEnabled bool) bool {
	for _, s := range DefaultSpecs(chefServerEnabled, workflowEnabled) {
		if s.Name == desired {
			return true
		}
	}

	return false
}
