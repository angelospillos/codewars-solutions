// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo></http:>
// Gomega Matcher Library <http://onsi.github.io/gomega></http:>

package kata_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/kata"
)
var _ = Describe("Test Example", func() {
  It("should return the correct value", func() {
    Expect(DuplicateEncode("din")).To(Equal("((("))
    Expect(DuplicateEncode("recede")).To(Equal("()()()"))
    Expect(DuplicateEncode("(( @")).To(Equal("))(("))
  })

  It("should ignore case", func() {
    Expect(DuplicateEncode("Success")).To(Equal(")())())"))
  })
})
