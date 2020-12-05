<template>
  <img alt="Vue logo" src="./assets/logo.png">
  <HelloWorld msg="hello world"/>
  <h2>hello 欢迎光临</h2>
  <div>请选择菜品</div>
  <div>
    <button 
      @click="selectMenuFun(index)"
      v-for="(item, index) in menu"
      :key="index">
      {{index}} : {{item}}
    </button>
  </div>
  <div>你选择了{{selectMenu}}</div>
  <button @click="overActive">点餐完毕</button>
  <div>{{overText}}</div>
</template>

<script lang="ts">
import {
  ref,
  reactive,
  toRefs,
  onBeforeMount,
  onMounted,
  onBeforeUpdate,
  onUpdated,
  // onRenderTracked,
  onRenderTriggered,
  watch
  } from 'vue'
import HelloWorld from './components/HelloWorld.vue';

//reactive 把普通的数据变成模板中有响应能力的数据

// @Options({
//   setup () {
//     const meaus = ref(['花菜','小炒肉'])
//     debugger
//     return {
//       meaus
//     }
//   },
// })
interface DataProps {
  menu: string[];
  selectMenu: string;
  selectMenuFun: (index: number) => void;
}
export default {
  name:'app',
  components: {
    HelloWorld
  },
  setup () {
    console.log('1----开始创建组件执行-------setup()')
    const data: DataProps = reactive ({
      menu:['小炒肉','花菜'],
      selectMenu: '',
      selectMenuFun: (index: number) => {
        data.selectMenu = data.menu[index];
      }
    })
    onBeforeMount (() => {
      console.log('2-----组件挂载到页面之前执行-------onBeforeMount()')
    });
    onMounted (() => {
        console.log('3-----组件挂载到页面之后执行-------onMounted()')
    });
    onBeforeUpdate (() => {
      console.log('组件更新之前执行-------onBeforeUpdate()')
    });
    onUpdated (() => {
        console.log('组件更新之后执行-------onUpdated()')
    });
    // onRenderTracked ((event) => { // 每一个都跟踪
    //   console.log("状态跟踪钩子函数--------")
    //   console.log("event",event)
    // });
    onRenderTriggered ((event) => { // 新值旧值都会出现
      console.log("状态跟踪钩子函数--------")
      console.log("event",event)
    });
    // setup() 开始创建之前 在beforeCreate之前执行
    // onbeforeUnmount() 组件卸载完成之后执行的函数
    // onUnmounted () 组件卸载完成之后执行的函数
    // onActivated () 被包含在keep-alive中的组件，会多出两个生命周期钩子函数，被激活时执行
    // onDeactivated() 比如从A组件,切换到B组件，A组件消失时执行
    // onErrorCaptured () 当捕获一个来自子孙组件的异常时激活钩子函数


    // 调试钩子 onRenderTracked 状态跟踪 onRenderTriggered




    // const menu = ref(['小炒肉','花菜']);
    // const selectMenu = ref("");
    // const selectMenuFun = (index: number) => {
    //   selectMenu.value = menu.value[index];
    // }
    console.log(data)
    const refData = toRefs(data)
    console.log(refData)

    const overText = ref("湘味缘")
    const overActive = () => {
      overText.value = "点餐完成" + overText.value
      // document.title = overText.value
    };
    watch(() => {
      return [overText,data.selectMenu]
    }, ([newValue,newValue2], [oldValue,oldValue2]) => {
      console.log(`new-----${newValue}${newValue2}`);
      console.log(`old-----${oldValue}${oldValue2}`);
      // document.title= newValue
    })
    return {
      ...refData,
      overText,
      overActive
    }
  }
}


//生命周期
//钩子函数
// beforeCreate created    setup 钩子是在berforeCreate前执行

// 调试钩子 onRenderTracked onRenderTriggered

</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
