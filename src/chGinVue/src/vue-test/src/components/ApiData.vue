<template>
  <div class="hello">
    <h1>{{msg}}</h1>
    <h1>site:{{site}}</h1>
    <h1>url:{{url}}</h1>
    <h3>{{details()}}</h3>
    <h1 v-for="data in ydata" :key="data">{{data}}</h1>
    <h3 v-for="item in xdata" :key="item">{{item}}</h3>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'apidata',
  data () {
    return {
      msg: 'hello, MX',
      site: 'MX',
      url: 'https://maicoder.github.io',
      xdata: null,
      ydata: null
    }
  },
  methods: {
    details: function () {
      return this.site
    }
  },
  mounted () {
    var a = this
    axios.get('http://localhost:8000/v1/line')
      .then(function (response) {
        a.xdata = response.data.legend_data
        a.ydata = response.data.xAxis_data
      })
  }
}
</script>

<style>
  .hello{
    font-weight: normal;
    text-align: center;
    font-size: 8px;
  }
  h3 {
    text-align: center;
    font-size: 20px;
    color: red;

  }
</style>
