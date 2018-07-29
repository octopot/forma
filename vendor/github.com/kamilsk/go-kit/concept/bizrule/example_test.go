package bizrule_test

import "github.com/kamilsk/go-kit/concept/bizrule"

func Example() {
	ctx := bizrule.NewContext()
	ctx.Describe("handle http request", func(ctx *bizrule.Context) {
		bizrule.It("should be authorized", func() error {
			return nil
		})
	})

	rule1, out1 := bizrule.New(nil)
	rule2, out2 := bizrule.New(out1)
	rule3, _ := bizrule.New(bizrule.And(out1, out2))

	for _, rule := range [...]bizrule.Rule{rule1, rule2, rule3} {
		rule.Apply(ctx)
	}
}
