// https://zenn.dev/fuqda/articles/a4d0bd213bf868
import { createLocalVue, mount, shallowMount, config } from '@vue/test-utils'
import login from '@/pages/login.vue'
import Vuetify from 'vuetify'

describe('login.vue', () => {
  const wrapper = mount(login)
  test('is a Vue instance', () => {
    expect(wrapper.vm).toBeTruthy()
  })
  test('is empty email', () => {
    expect(wrapper.find('#email').text()).toBe("")
  })
  test('is empty password', () => {
    expect(wrapper.find('#password').text()).toBe("")
  })
  test('is enable submit', () => {
    expect(wrapper.find('#submit').attributes()).not.toHaveProperty("disabled")
  })
  test("is empty error", () => {
    expect(wrapper.find("#error").text()).toBe("")
  })

  const localVue = createLocalVue()
  let vuetify: Vuetify

  beforeEach(() => {
    vuetify = new Vuetify()
  })

  it("validation error when email and password are empty", () => {
    const wrapper = shallowMount(login, {
      localVue, vuetify
    })
    expect(wrapper.vm).toBeTruthy()
    wrapper.find("#submit").trigger("submit.prevent")

  })
})
