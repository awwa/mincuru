import Vuetify from 'vuetify'
import { createLocalVue, mount, shallowMount, config } from '@vue/test-utils'
import UserList from '@/components/UserList.vue'

const testUsers =[
  {
    id: 1,
    name: "hoge1",
    email: "hoge1@example.com",
    role: "admin",
  },
  {
    id: 2,
    name: "hoge2",
    email: "hoge2@example.com",
    role: "admin",
  }
]

describe('UserList.vue', () => {
  // methodsをmock化すると発生するエラーを一時的に回避
  // https://qiita.com/mejileben/items/32b1cbbb2f522601e63d
  config.showDeprecationWarnings = false

  const localVue = createLocalVue();
  let vuetify;

  beforeEach(() => {
    vuetify = new Vuetify();
  })

  it('is a Vue instance', () => {
    const wrapper = shallowMount(UserList, {
      localVue, vuetify,
      propsData: {
        users: []
      }
    })
    expect(wrapper.vm).toBeTruthy()
  })

  // const mockRouter = {
  //   push: (_id) => {
  //   }
  // }
  it("users accept array", () => {
    const wrapper = mount(UserList, {
      localVue, vuetify,
      propsData: {
        users: testUsers,
      },
      // mocks: {
      //   $router: mockRouter,
      // }
      // methods: {
      //   editItem() {}
      // }
    })

    // const spy = jest.spyOn(wrapper.vm, "editItem")
    // wrapper.vm.editItem(1)
    //console.log(wrapper.html())
    // console.log(wrapper.find("tbody tr").html())
    const editItem = jest.fn()
    wrapper.setMethods({ editItem })
    //wrapper.vm.editItem = jest.mock()
    wrapper.find("tbody tr").trigger("click")
    
    // wrapper.vm.$nextTick(() => {
    //   expect(spy).toHaveBeenCalled()
    //   done()
    // })
    // expect(spy).toHaveBeenCalled()

    expect(editItem).toBeCalled()
    // expect(wrapper.vm.editItem).toBeCalled()
    // expect((wrapper.vm).items).toEqual([{
    //   text: "ホーム",
    //   disabled: false,
    //   href: "/"
    // },
  //   {
  //     text: "hoge",
  //     disabled: true,
  //     href: "/hoge"
  //   }])
  // })
  })
})
