package user_test

import (
	. "cf/commands/user"
	"cf/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	testapi "testhelpers/api"
	testassert "testhelpers/assert"
	testcmd "testhelpers/commands"
	testconfig "testhelpers/configuration"
	testreq "testhelpers/requirements"
	testterm "testhelpers/terminal"
)

func callSetOrgRole(args []string, reqFactory *testreq.FakeReqFactory, userRepo *testapi.FakeUserRepository) (ui *testterm.FakeUI) {
	ui = new(testterm.FakeUI)
	ctxt := testcmd.NewContext("set-org-role", args)

	config := testconfig.NewRepositoryWithDefaults()

	cmd := NewSetOrgRole(ui, config, userRepo)
	testcmd.RunCommand(cmd, ctxt, reqFactory)
	return
}

var _ = Describe("Testing with ginkgo", func() {
	It("TestSetOrgRoleFailsWithUsage", func() {
		reqFactory := &testreq.FakeReqFactory{}
		userRepo := &testapi.FakeUserRepository{}

		ui := callSetOrgRole([]string{"my-user", "my-org", "my-role"}, reqFactory, userRepo)
		Expect(ui.FailedWithUsage).To(BeFalse())

		ui = callSetOrgRole([]string{"my-user", "my-org"}, reqFactory, userRepo)
		Expect(ui.FailedWithUsage).To(BeTrue())

		ui = callSetOrgRole([]string{"my-user"}, reqFactory, userRepo)
		Expect(ui.FailedWithUsage).To(BeTrue())

		ui = callSetOrgRole([]string{}, reqFactory, userRepo)
		Expect(ui.FailedWithUsage).To(BeTrue())
	})
	It("TestSetOrgRoleRequirements", func() {

		reqFactory := &testreq.FakeReqFactory{}
		userRepo := &testapi.FakeUserRepository{}

		reqFactory.LoginSuccess = false
		callSetOrgRole([]string{"my-user", "my-org", "my-role"}, reqFactory, userRepo)
		Expect(testcmd.CommandDidPassRequirements).To(BeFalse())

		reqFactory.LoginSuccess = true
		callSetOrgRole([]string{"my-user", "my-org", "my-role"}, reqFactory, userRepo)
		Expect(testcmd.CommandDidPassRequirements).To(BeTrue())

		Expect(reqFactory.UserUsername).To(Equal("my-user"))
		Expect(reqFactory.OrganizationName).To(Equal("my-org"))
	})
	It("TestSetOrgRole", func() {

		org := models.Organization{}
		org.Guid = "my-org-guid"
		org.Name = "my-org"
		user := models.UserFields{}
		user.Guid = "my-user-guid"
		user.Username = "my-user"
		reqFactory := &testreq.FakeReqFactory{
			LoginSuccess: true,
			UserFields:   user,
			Organization: org,
		}
		userRepo := &testapi.FakeUserRepository{}

		ui := callSetOrgRole([]string{"some-user", "some-org", "OrgManager"}, reqFactory, userRepo)

		testassert.SliceContains(ui.Outputs, testassert.Lines{
			{"Assigning role", "OrgManager", "my-user", "my-org", "my-user"},
			{"OK"},
		})
		Expect(userRepo.SetOrgRoleUserGuid).To(Equal("my-user-guid"))
		Expect(userRepo.SetOrgRoleOrganizationGuid).To(Equal("my-org-guid"))
		Expect(userRepo.SetOrgRoleRole).To(Equal(models.ORG_MANAGER))
	})
})
