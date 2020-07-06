package e2e

// This test suite determines the behavior of the dependency resolver in OLM - how it resolves operator dependencies
// across different packages and catalogs. In the case where an operator has no external dependencies, resolution is
// straightforward. However, there are cases where operator depend on certain GVKs (CRDs) to be present in the cluster.
// Also, with semantic versioning, an operator can explicitly depend on a certain version of another operator. This suite
// attempts to ensure the new SAT-based resolver can resolve GVK based dependencies in the same way to the existing resolver.
// See https://github.com/operator-framework/enhancements/blob/master/enhancements/operator-dependency-resolution.md for more info.

import (
	. "github.com/onsi/ginkgo"
	"github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/clientset/versioned"
	"github.com/operator-framework/operator-lifecycle-manager/pkg/lib/operatorclient"
	"github.com/operator-framework/operator-lifecycle-manager/test/e2e/ctx"
)

var _ = Describe("GVK dependency resolution", func() {
	var (
		kubeClient     operatorclient.ClientInterface
		operatorClient versioned.Interface
	)

	BeforeEach(func() {
		kubeClient = ctx.Ctx().KubeClient()
		operatorClient = ctx.Ctx().OperatorClient()
	})

	AfterEach(func() {
		cleaner.NotifyTestComplete(true)
	})


	Context("an operator is being installed via a subscription (via dependency resolution)", func() {

		When("an operator has a single dependency on a GVK of another operator in the same catalog (baseline case)", func() {

			It("resolves the dependency using the operator in the same catalog", func() {
				// TODO
			})
		})

		When( "an operator has a single dependency on a GVK of another operator provided in multiple catalogs", func () {

			It("resolves the dependency using the operator in the same catalog", func () {
				// TODO
			})
		})

		When("an operator has multiple dependencies on GVKs of another operator provided in multiple catalogs", func() {

			It("resolves the dependencies using the operators in the same catalog", func() {
				// TODO
			})
		})

		
	})
})