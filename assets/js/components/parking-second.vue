<template>
  <div>
    <banner></banner>
    <h1 style="font-size:45px;left:50px;position:absolute;color: #484C58;font-family:impact;top:150px">Parking Management</h1>
    <h4 style="left:650px;position:absolute;color: #484C58;font-family:impact;top:170px">Floor 2</h4>
    <div class="box green" style="left:775px;top:175px;"></div>
    <div class="box yellow" style="left:925px;top:175px;"></div>
    <div class="box red" style="left:1075px;top:175px;"></div>
    <p class="key-items" style="left:800px;top:175px;">
    Available
    </p>
    <p class="key-items" style="left:950px;top:175px;">
    Occupied
    </p>
    <p class="key-items" style="left:1100px;top:175px;">
    Exceeded Time Limit
    </p>
    <a href="./parking-first-floor" class="button" style="left:975px;top:125px;background-color:#484C58;">Floor 1</a>
    <a href="./resident-table" class="button" style="left:1100px;top:125px;background-color:#484C58;">Residents</a>


    <div>
      <svg width="1201" height="500" style="position:absolute;left:100px;top:250px;" >
        <image xlink:href="/assets/images/shortbread-first-level-floorplan-edit.jpg" width="1200" height="500" style="position:absolute;left:0px;top:0px" alt="ShortBread logo"/>
        <rect v-for="spot in spots" :key="spot.Space" :width="spot.Width" :height="spot.Height" :class="spot.Status" :x="spot.X1" :y="spot.Y1" @click="seeSpot(spot)"/>
        <text v-for="spot in spots" :key="spot.Space + spot.X1" :x="spot.X1+8" :y="spot.Y1+spot.Height-5">{{spot.Space}}</text>
      </svg>
    </div>

    <div id="myModal" class="modal fade" role="dialog">
      <div class="modal-dialog">

        <!-- Modal content-->
        <div class="modal-content">
          <div class="header">
            <h3 class="modal-title" style="text-align:center">Spot {{currentSpot.Space}}</h3>
            <br>
          </div>
          <div style="text-align:center">
            Name <input type="text" id="nameText" name="name"><br>
            <br>
            Phone # <input type="text" id="phoneText" name="phone"><br>
            <br>
            Duration <input type="text" id="durationText" name="duration"><br>
            <br>
            Room # <input type="text" id="roomText" name="name"><br>
            <br>
          </div>
          <div class="footer">
            <button type="button" class="book-button" @click="bookSpot(currentSpot)">Book Spot</button>
            <br>
            <br>
            <p>Next Booking Begins: 10:10am</p>
          </div>
        </div>

      </div>
    </div>

  </div>
</template>


<script>
import axios from "axios";
export default {
  data() {
    return {
      spots: [],
      currentSpot: {}
    }
  },
  methods: {
    printSpots() {
      console.log(this.spots);
    },
    seeSpot(spot){
        this.currentSpot = spot;
        $('#myModal').modal()
    },
    bookSpot(spot){
        console.log("booked")
        $('#myModal').modal('hide')
    }
  },
  mounted() {
    axios.get("/second-floor-spots")
      .then(response => {
        console.log(response.data);
        this.spots = response.data;
      })
      .catch(e => {
        console.error(e);
      })
  }
}
$(document).ready(function() {
    $(window).on('hidden.bs.modal', function() {
        console.log("Closed")
         $("#nameText").val("");
        $("#phoneText").val("");
        $("#durationText").val("");
        $("#roomText").val("");
    });
});
</script>

<style lang="scss">
.key-items{
  position:absolute;
  top:230px;
  color:#484C58
}
#book-times{
  position:absolute;
  font-family:Montserrat;
  right: 50px;
  top: 150px;
  color: #484C58;
}
.box{
  position:absolute;
  top:230px;
  width:20px;
  height:20px;
}
.red{
  background-color: #DD9A9A
}
.yellow{
  background-color: #DCDD9B
}
.green{
  background-color: #B0DD9A
}
.header{
  text-align:center;
  color: #484C58;
  padding-top:20px;
  text-weight:bold;
  font-family: impact;
}
.footer{
  text-align:center;
}
.book-button{
  color: #fff;
  font-family: Montserrat;
  background-color: #91A6C6;
}
.Available{
    fill:#B0DD9A;
    stroke-width:3;
    stroke:rgb(140,140,140);
}
.Occupied{
    fill:#DCDD9B;
    stroke-width:3;
    stroke:rgb(140,140,140);
}
.Exceeded{
    fill:#DD9A9A;
    stroke-width:3;
    stroke:rgb(140,140,140);
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
