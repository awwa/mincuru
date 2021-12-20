import Vuetify from 'vuetify'
import { createLocalVue, mount } from '@vue/test-utils'
import NuxtLogo from '@/components/NuxtLogo.vue'

describe('NuxtLogo.vue', () => {

  const localVue = createLocalVue();
  let vuetify;

  beforeEach(() => {
    vuetify = new Vuetify();
  })

  it('is a Vue instance', () => {
    const wrapper = mount(NuxtLogo, {
      localVue,
      vuetify,
      propsData: { title: "hoge"}
    })
    expect(wrapper.vm).toBeTruthy()
    const title = wrapper.find(".title")
    expect(title.text()).toBe("hoge")
  })
})
