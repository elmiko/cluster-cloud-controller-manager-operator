# Cluster Cloud Controller Manager Operator

The Cluster Cloud Controller Manager operator (CCCMO) manages and updates the various [Cloud Controller Managers](https://kubernetes.io/docs/concepts/architecture/cloud-controller/) deployed on top of [OpenShift](https://openshift.io). The operator is based on the [Kubebuilder](https://kubebuilder.io/) framework and [controller-runtime](https://github.com/kubernetes-sigs/controller-runtime) libraries. It is installed via [Cluster Version Operator](https://github.com/openshift/cluster-version-operator) (CVO).

## Operator Status and Supported Platforms

Kubernetes is in the process of migrating its cloud controller functionality out of the Kubernetes core and into separate Cloud Controller Managers (see [KEP 2395](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cloud-provider/2395-removing-in-tree-cloud-providers) for more information). As this process is an on-going effort, we will document the status and progress of this operator, as well as the supported platforms, until the operator has gone into general availability (GA) in OpenShift.

### Operator Status

**Stable**

This operator is considered stable with unit and end-to-end testing against the latest OpenShift. It is currently available as a released component in OpenShift versions >= 4.7.

### Supported Platforms

| Platform                                                             | Included in Operator | Tested in CI  |
|----------------------------------------------------------------------|----------------------| ------------  |
| [AWS](https://github.com/openshift/cloud-provider-aws)               | Yes                  | Yes           |
| [Azure](https://github.com/openshift/cloud-provider-azure)           | Yes                  | Yes           |
| [GCP](https://github.com/openshift/cloud-provider-gcp)               | Yes                  | Yes           |
| [OpenStack](https://github.com/openshift/cloud-provider-openstack)   | Yes                  | Yes           |
| [vSphere](https://github.com/openshift/cloud-provider-vsphere)       | Yes                  | Yes           |
| [IBM](https://github.com/openshift/cloud-provider-ibm)               | Yes                  |               |
| [Alibaba](https://github.com/openshift/cloud-provider-alibaba-cloud) | Yes                  | Yes           |
| [Power VS](https://github.com/openshift/cloud-provider-powervs)      | Yes                  |               |
| [Nutanix](https://github.com/openshift/cloud-provider-nutanix)       | Yes                  | Yes           |

## Deploying and Running CCCMO

The CCCMO deploys controllers which provide a central core component of Kubernetes. As such, its deployment and operation is highly sensitive to cluster bootstrapping and initial payload deployment. In general, it is best to allow the OpenShift installer to manage its operation.

To better understand how this operator is deployed, please see the [`manifests`](/manifests) directory. It contains a series of Kubernetes YAML manifests which are deployed by the CVO during installation. Additionally, [this OpenShift enhancement(currently in review)](https://github.com/openshift/enhancements/pull/463/) provides detailed information about how the operator will be deployed, operated, and upgraded.

More detailed guide is in [#hacking-guide](./docs/dev/hacking-guide.md)

## Development

**Prerequisites**

* Go language 1.20+
* GNU Make

All development related tasks can be run through the `Makefile`. Supplemental scripts can be found in the `hack` directory.

If you do not have the necessary tools for building, but do have access to Podman or Docker, you may use the `hack/container-run.sh` script to run Makefile targets in a container. See the `container-run.sh` file for usage instructions.

### Building the Operator

To build the operator binary, and run related linting tests, type `make` or `make build` from the root of the project.

### Adding support for new cloud provider

This step is extensively covered in [Integrating a new cloud provider in Cluster-cloud-controller-manager-operator (CCCMO)](/docs/dev/cloud-provider-integration.md)

### Vendoring Dependencies

After adding or updating dependencies in the `go.mod` file, run `make vendor` to ensure that all new dependencies are added to the `vendor` directory. It is also useful to run `make vet` to ensure that no build-time errors have been introduced during the vendor process.

### Running Tests

The CCCMO has multiple levels of testing: unit tests, and end to end (e2e) functionality tests. As a developer you should run the unit tests locally to ensure that your changes do not break the tests. Although running the e2e tests manually can be a rewarding experience, it is also complicated to configure and maintain. For these reasons it is often better to let the continuous integration systems run the e2e tests automatically for you on pull requests.

#### Local Unit Tests

To invoke the unit tests, run `make unit`. If you wish to also run the code generation and verification steps, run `make test`.

#### End to End Tests

The CCCMO e2e tests are configured and deployed from the [OpenShift Release repository](https://github.com/openshift/release). You will find the CCCMO specific configurations in the [`release/ci-operator/config/openshift/cluster-cloud-controller-manager-operator`](https://github.com/openshift/release/tree/master/ci-operator/config/openshift/cluster-cloud-controller-manager-operator) directory. For more information about these tests and how they are run and configured, please see the [OpenShift CI Docs](https://docs.ci.openshift.org/docs/).

