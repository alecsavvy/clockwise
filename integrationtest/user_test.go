package integrationtest

import (
	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo/v2"
	"github.com/stretchr/testify/require"
)

var _ = Describe("Users", func() {
    It("should be able to create users", func() {
        sdk := newSdk()

				account, _ := generateWallet()

				handle := faker.Username()
				address := account.Address.Hex()
				bio := faker.Sentence()
			
				user, err := sdk.CreateUser(
					handle,
					address,
					bio,
				)

				require.Nil(GinkgoT(), err)
				require.EqualValues(GinkgoT(), user.CreateUser.Address, address)
    })
})
