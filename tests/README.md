# Integration tests

## Writing e2e tests

We aim to run e2e tests in parallel by default. As such the following rules should be followed:
 * Use cirros and alpine VMs for testing wherever possible. If you have to use
   use another OS, discuss on the PR why it is needed.
 * Stay within the boundary of the Testnamespaces which we prove. If you have
   to create resources outside of the test namespaces, discuss potential
   solutions on such a PR.
 * If you really have to run tests serial (destructive tests, infra-tests,
  ...), mark the test with a `[Serial]` tag.
 * If tests are not using the default cleanup code, additional custom
   preparations may be necessary.

The following types of tests need to be marked as `[Serial]` right now:

 * Tests which use PVCs or DataVolumes (parallelizing these is on the way).
 * Tests which use `BeforeAll`.

Additional suggestions:

 * The `[Disruptive]` tag is recognized by the test suite but is not yet
   mandatory. Feel free to set it on destructive tests.
 * Conformance tests need to be marked with a `[Conformance]` tag.

## Test Namespaces

If tests are executed in parallel, every test gets its unique set of namespaces
to test in. If you write a test and reference the namespaces
`test.NamespaceTestDefault`, `test.NamespaceTestAlternative` or
`tests.NamespaceTestOperator`, you get different values, based on the ginkgo
execution node.

For as long as test resources are created by referencing these namespaces,
there is no test conflict to expect.

## Running integration tests

Integration tests require a running Kubevirt cluster.  Once you have a running
Kubevirt cluster, you can use the `-master` and the `-kubeconfig` flags to
point the tests to the cluster.

## Running networking tests for outside connectivity

When running the tests with no internet connection,
some networking tests ,that test outside connectivity, might fail,
and you might want to skip them.
For that some additional flags are needed to be passed.
In addition, if you'd like to test outside connectivity
using different addresses than the default
(`8.8.8.8`, `2001:db8:1::1` and `google.com`), you can achive that with the 
designated flags as well.

For each method detailed below, there is note about the needed flags
to the outside connectivity tests and how to pass them.

## Run them on an arbitrary KubeVirt installation

```
cd tests # from the git repo root folder
go test -kubeconfig=path/to/my/config -config=default-config.json
```

>**outside connectivity tests:** The tests will run by default.
>To skip the outside connectivity tests add
>`-ginkgo.skip='\[outside_connectivity\]'` To your go command.
>To change the IPV4, IPV6 or DNS used for outside connectivity tests,
>add `conn-check-ipv4-address`,
>`conn-check-ipv6-address` or `conn-check-dns` to your go command,
>with the desired value.
>For example:
>```
>go test -kubeconfig=$KUBECONFIG -config=default-config.json \
>-conn-check-ipv4-address=8.8.4.4 -conn-check-ipv6-address=2620:119:35::35 \
>-conn-check-dns=amazon.com \
>```


## Run them on one of the core KubeVirt providers

There is a make target to run this with the config
taken from hack/config.sh:

```
# from the git repo root folder
make functest
```

>**outside connectivity tests:** The tests will run by default. To skip
>the tests export `KUBEVIRT_E2E_SKIP='\[outside_connectivity\]'` 
>environment variable before running the tests.
>To change the IPV4, IPV6 or DNS used for outside connectivity tests,
>you can export `CONN_CHECK_IPV4_ADDRESS`, `CONN_CHECK_IPV6_ADDRESS` and  
> `CONN_CHECK_DNS` with the desired values. For example:
>```
>export CONN_CHECK_IPV4_ADDRESS=8.8.4.4
>export CONN_CHECK_IPV6_ADDRESS=2620:119:35::35
>export CONN_CHECK_DNS=amazon.com
>```

## Run them anywhere inside of container

```
# Create directory for data / results / kubectl binary
mkdir -p /tmp/kubevirt-tests-data
# Make sure that eveybody can write there
setfacl -m d:o:rwx /tmp/kubevirt-tests-data
setfacl -m o:rwx /tmp/kubevirt-tests-data

docker run \
    -v /tmp/kubevirt-tests-data:/home/kubevirt-tests/data:rw,z --rm \
    kubevirt/tests:latest \
        --kubeconfig=data/openshift-master.kubeconfig \
        --container-tag=latest \
        --container-prefix=quay.io/kubevirt \
        --test.timeout 180m \
        --junit-output=data/results/junit.xml \
        --deploy-testing-infra \
        --path-to-testing-infra-manifests=data/manifests
```

## Test lane labels

**UPGRADE**: Use ginkgo label decorators to mark if certain tests belong to a specific test lane or not. The tests under consideration from `automation/test.sh` are:

1. `GPU` tests in test-lane `$TARGET =~ gpu.*`
2. `MediatedDevices` tests in test lane `$TARGET =~ vgpu.*`
3. `Migration tests` (which need at least 2 nodes and a CPU manager) in test lane `$TARGET =~ sig-compute-migrations`
4. Non-root tests if `$KUBEVIRT_NONROOT =~ true`

This adds new labels to the group of tests written in the test files mentioned below:

1. Label- `needs-gpu` in file `kubevirt/tests/vmi_gpu_test.go`
2. Label- `needs-mdev-gpu` in file `kubevirt/tests/mdev_configuration_allocation_test.go`. 
3. Label- `needs-two-nodes-with-cpumanager` in file `kubevirt/tests/migration_test.go`.
4. Label- `needs-non-root` in file `kubevirt/tests/nonroot_test.go`.

This eliminates the dependency on the use of random keywords matchers.
--label-filter is then used with gingko params in automation/test.sh, to indicate all test lanes that require these tests.

**Note**: For these tests to pass, the TARGET must have at least one node on the cluster that contains an Nvidia GPU. For mediated devices tests, requirement is of an Nvidia GPU + mediated device.

**RUNNING TESTS** : There are two ways to run the tests. Using:
1. `automation/test.sh` : This changes apply to this mode of running test.
2. `make functest` : This is used for local testing. The file `hack/functests.sh` has not been updated for anything.