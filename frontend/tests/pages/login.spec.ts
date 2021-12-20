// https://zenn.dev/fuqda/articles/a4d0bd213bf868
import { mount } from '@vue/test-utils'
import login from '@/pages/login.vue'

describe('login', () => {
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
})
