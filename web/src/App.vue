<template>
  <div id="app" v-if="!loading">
    <div>矿机列表</div>
    <div>共计：{{count}}个</div>
    <ul>
      <li class="item" v-for="(item, index) in list" :key="index">
        <div>{{item._id}}</div>
      </li>
    </ul>
  </div>
</template>

<script>

export default {
  name: 'app',
  data () {
    return {
      count: 0,
      list: [],
      loading: true
    }
  },
  mounted() {
    fetch('/MiningMachine/list').then(res => {
      if (res.status !== 200) return alert(`请求出错：错误码${res.status}`)
      return res.json()
    }).then(data => {
      console.log(data); // {name: 'test', age: 1}
      this.list = data.data
      this.count = data.count
      this.loading = false
    })
  },
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
