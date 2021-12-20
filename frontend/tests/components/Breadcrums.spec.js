import Vuetify from 'vuetify'
import { createLocalVue, mount } from '@vue/test-utils'
import Breadcrumbs from '@/components/Breadcrumbs.vue'

describe('Breadcrumbs.vue', () => {

  const localVue = createLocalVue();
  let vuetify;

  beforeEach(() => {
    vuetify = new Vuetify();
  })

  it('is a Vue instance', () => {
    const wrapper = mount(Breadcrumbs, {
      localVue, vuetify,
      propsData: {
        path: "/hoge"
      }
    })
    expect(wrapper.vm).toBeTruthy()
  })
  it("items accept string", () => {
    const wrapper = mount(Breadcrumbs, {
      localVue, vuetify,
      propsData: {
        path: "/hoge",
      }
    })
    expect((wrapper.vm).items).toEqual([{
      text: "ホーム",
      disabled: false,
      href: "/"
    },
    {
      text: "hoge",
      disabled: true,
      href: "/hoge"
    }])
  })
  it("path has slash at eol", () => {
    const wrapper = mount(Breadcrumbs, {
      localVue, vuetify,
      propsData: {
        path: "/hoge/",
      }
    })
    expect((wrapper.vm).items).toEqual([{
      text: "ホーム",
      disabled: false,
      href: "/"
    },
    {
      text: "hoge",
      disabled: true,
      href: "/hoge"
    }])
  })
  it("multiple hierarchy", () => {
    const wrapper = mount(Breadcrumbs, {
      localVue, vuetify,
      propsData: {
        path: "/hoge/account/users",
      }
    })
    expect((wrapper.vm).items).toEqual([
      {
        text: "ホーム",
        disabled: false,
        href: "/"
      },
      {
        text: "hoge",
        disabled: false,
        href: "/hoge"
      },
      {
        text: "アカウント設定",
        disabled: false,
        href: "/hoge/account"
      },
      {
        text: "ユーザー",
        disabled: true,
        href: "/hoge/account/users"
      },
    ])
  })
})
