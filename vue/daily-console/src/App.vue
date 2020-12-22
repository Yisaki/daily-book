<template>
    <div id="app">
        <router-view />

        <!--<TabBar>
            <TabBarItem>
                <img slot="item-icon" src="@/assets/img/p.jpg" alt="" />
                <div slot="item-text">首页</div>
            </TabBarItem>
            <TabBarItem>
                <img slot="item-icon" src="@/assets/img/p.jpg" alt="" />
                <div slot="item-text">分类</div>
            </TabBarItem>
        </TabBar>-->
        <el-menu :default-active="'/period'" mode="horizontal" @select="handleSelect" id="global-nav" router >
          <el-menu-item index="/period">小裴</el-menu-item>
          <el-menu-item index="/pet">小祖</el-menu-item>
        </el-menu>
    </div>
</template>
<script>
import TabBar from '@/components/tabbar/TabBar'
import TabBarItem from '@/components/tabbar/TabBarItem'
export default {
    components: {
        // TabBar,
        // TabBarItem
    },
    data() {
      return {

      };
    },
    methods: {
      handleSelect(key, keyPath) {
        console.log(key, keyPath);
      }
    }
}
</script>
<style scoped>
@import './assets/css/base.css';
#global-nav{
  display: flex;
  text-align: center;
  position: fixed;
  left: 0;
  bottom: 0;
  right: 0;
}
#global-nav .el-menu-item{
  flex: 1;
}
</style>
