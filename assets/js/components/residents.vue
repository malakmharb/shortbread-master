<template>
<div>
<banner></banner>
<h1 style="font-size:45px;left:50px;position:absolute;color: #484C58;font-family:impact;top:150px">Residents</h1>
<a href="./parking-second-floor" class="button" style="left:1100px;top:125px;background-color:#484C58;">Floor 2</a>
    <a href="./parking-first-floor" class="button" style="left:975px;top:125px;background-color:#484C58;">Floor 1</a>
<table style="position:absolute;top: 225px;left:150px;">
  <tr>
    <th>Name</th>
    <th>Phone Number</th>
    <th>Room Number</th>
    <th>Spot Number</th>
  </tr>
  <tr v-for="resident in residents" :key="resident.Id">
    <td>{{resident.firstName}} {{resident.lastName}}</td>
    <td>{{resident.phone}}</td>
    <td>{{resident.room}}</td>
    <td>{{resident.spot}}</td>
  </tr>
</table>

</div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      residents: []
    }
  },
  methods: {
    
  },
  mounted() {
    axios.get("/residents")
      .then(response => {
        console.log(response.data);
        this.residents = response.data;
      })
      .catch(e => {
        console.error(e);
      })
  }
}
</script>

<style lang="scss">
table {
    font-family: arial, sans-serif;
    border-collapse: collapse;
    width: 80%;
}

td, th {
    border: 1px solid #dddddd;
    text-align: left;
    padding: 8px;
}

tr:nth-child(even) {
    background-color: #dddddd;
}
.button{
    color:#fff;
    position: absolute;
    text-align: center;
    font-family: impact;
    border-radius:5px;
    padding:5px;
    height:30px;
    width:100px;
}
a:hover{
    color:#fff;
}
</style>