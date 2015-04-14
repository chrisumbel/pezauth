package pezauth_test

import (
	"github.com/go-martini/martini"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/pezauth/service"
)

var _ = Describe("AuthKeyController", func() {
	var (
		controller Controller
		fakeGUID   = "123asdkghasdiawlkehgaweh"
		domain     = "pivotal.io"
	)
	setGetUserInfo(domain)

	Context("PUT", func() {
		Context("called successfully", func() {
			BeforeEach(func() {
				kg := getKeygen(false, fakeGUID)
				controller = NewAuthKeyV1(kg)
			})

			It("Should NOT result in panic", func() {
				Ω(func() {
					controller.Put()
				}).ShouldNot(Panic())
			})

			Context("Handler function", func() {
				It("Should yeild a valid status code and response object", func() {
					var ph AuthPutHandler
					ph = controller.Put().(AuthPutHandler)
					render := &mockRenderer{}
					ph(martini.Params{UserParam: "test"}, nil, render, new(mockTokens))
					Ω(render.StatusCode).Should(Equal(200))
					Ω(render.ResponseObject.(Response).ApiKey).Should(Equal(fakeGUID))
				})
			})
		})

		Context("called with failure", func() {
			BeforeEach(func() {
				kg := getKeygen(true, fakeGUID)
				controller = NewAuthKeyV1(kg)
			})

			Context("Handler function", func() {
				It("Should yeild a error status code and error response object", func() {
					var ph AuthPutHandler
					ph = controller.Put().(AuthPutHandler)
					render := &mockRenderer{}
					ph(martini.Params{UserParam: "test"}, nil, render, new(mockTokens))
					Ω(render.StatusCode).Should(Equal(403))
					Ω(render.ResponseObject.(Response).ApiKey).Should(Equal(""))
					Ω(render.ResponseObject.(Response).ErrorMsg).ShouldNot(Equal(""))
				})
			})
		})
	})

	Context("POST", func() {
		Context("called successfully", func() {
			BeforeEach(func() {
				kg := getKeygen(false, fakeGUID)
				controller = NewAuthKeyV1(kg)
			})

			It("Should NOT result in panic", func() {
				Ω(func() {
					controller.Post()
				}).ShouldNot(Panic())
			})

			Context("Handler function", func() {
				It("Should yeild a valid status code and response object", func() {
					var ph AuthPostHandler
					ph = controller.Post().(AuthPostHandler)
					render := &mockRenderer{}
					ph(martini.Params{UserParam: "test"}, nil, render, new(mockTokens))
					Ω(render.StatusCode).Should(Equal(200))
					Ω(render.ResponseObject.(Response).ApiKey).Should(Equal(fakeGUID))
				})
			})
		})

		Context("called with failure", func() {
			BeforeEach(func() {
				kg := getKeygen(true, fakeGUID)
				controller = NewAuthKeyV1(kg)
			})

			Context("Handler function", func() {
				It("Should yeild a error status code and error response object", func() {
					var ph AuthPostHandler
					ph = controller.Post().(AuthPostHandler)
					render := &mockRenderer{}
					ph(martini.Params{UserParam: "test"}, nil, render, new(mockTokens))
					Ω(render.StatusCode).Should(Equal(403))
					Ω(render.ResponseObject.(Response).ApiKey).Should(Equal(""))
					Ω(render.ResponseObject.(Response).ErrorMsg).ShouldNot(Equal(""))
				})
			})
		})
	})

	Context("GET", func() {
		Context("called successfully", func() {
			BeforeEach(func() {
				kg := getKeygen(false, fakeGUID)
				controller = NewAuthKeyV1(kg)
			})

			It("Should NOT result in panic", func() {
				Ω(func() {
					controller.Get()
				}).ShouldNot(Panic())
			})

			Context("Handler function", func() {
				It("Should yeild a valid status code and response object", func() {
					var ph AuthGetHandler
					ph = controller.Get().(AuthGetHandler)
					render := &mockRenderer{}
					ph(martini.Params{UserParam: "test"}, nil, render, new(mockTokens))
					Ω(render.StatusCode).Should(Equal(200))
					Ω(render.ResponseObject.(Response).ApiKey).Should(Equal(fakeGUID))
				})
			})
		})

		Context("called with failure", func() {
			BeforeEach(func() {
				kg := getKeygen(true, fakeGUID)
				controller = NewAuthKeyV1(kg)
			})

			Context("Handler function", func() {
				It("Should yeild a error status code and error response object", func() {
					var ph AuthGetHandler
					ph = controller.Get().(AuthGetHandler)
					render := &mockRenderer{}
					ph(martini.Params{UserParam: "test"}, nil, render, new(mockTokens))
					Ω(render.StatusCode).Should(Equal(403))
					Ω(render.ResponseObject.(Response).ApiKey).Should(Equal(""))
					Ω(render.ResponseObject.(Response).ErrorMsg).ShouldNot(Equal(""))
				})
			})
		})
	})

	Context("DELETE", func() {
		Context("called successfully", func() {
			BeforeEach(func() {
				kg := getKeygen(false, fakeGUID)
				controller = NewAuthKeyV1(kg)
			})

			It("Should NOT result in panic", func() {
				Ω(func() {
					controller.Delete()
				}).ShouldNot(Panic())
			})

			Context("Handler function", func() {
				It("Should yeild a valid status code and response object", func() {
					var ph AuthDeleteHandler
					ph = controller.Delete().(AuthDeleteHandler)
					render := &mockRenderer{}
					ph(martini.Params{UserParam: "test"}, nil, render, new(mockTokens))
					Ω(render.StatusCode).Should(Equal(200))
					Ω(render.ResponseObject.(Response).ApiKey).ShouldNot(Equal(fakeGUID))
					Ω(render.ResponseObject.(Response).ApiKey).Should(Equal(""))
				})
			})
		})

		Context("called with failure", func() {
			BeforeEach(func() {
				kg := getKeygen(true, fakeGUID)
				controller = NewAuthKeyV1(kg)
			})

			Context("Handler function", func() {
				It("Should yeild a error status code and error response object", func() {
					var ph AuthDeleteHandler
					ph = controller.Delete().(AuthDeleteHandler)
					render := &mockRenderer{}
					ph(martini.Params{UserParam: "test"}, nil, render, new(mockTokens))
					Ω(render.StatusCode).Should(Equal(403))
					Ω(render.ResponseObject.(Response).ApiKey).Should(Equal(""))
					Ω(render.ResponseObject.(Response).ErrorMsg).ShouldNot(Equal(""))
				})
			})
		})
	})
})