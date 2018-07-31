<template>
  <div>
    <banner></banner>
    <h1 style="font-size:45px;left:50px;position:absolute;color: #484C58;font-family:impact;top:150px">Parking Management</h1>
    <h4 style="left:650px;position:absolute;color: #484C58;font-family:impact;top:170px">Floor 1</h4>
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
    <a href="./parking-second-floor" class="button" style="left:975px;top:125px;background-color:#484C58;">Floor 2</a>
    <a href="./resident-table" class="button" style="left:1100px;top:125px;background-color: #484C58;">Residents</a>


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
            <form autocomplete="off" action="/action_page.php">
                <div class="autocomplete" style="width:300px;">
                     <input id="myInput" type="text" name="name" placeholder="Name" @click="nameComplete">
                     <br>
                     <br>
                     <input id="myPhoneInput" type="text" name="phone" placeholder="Phone #">
                     <br>
                     <br>
                     <input id="duration" type="text" name="duration" placeholder="Duration in minutes">
                     <br>
                     <br>
                     <input id="myRoomInput" type="text" name="room" placeholder="Room #">
                     <br>
                     <br>
                    <input id="id" type="text" name="room" style="background-color:#fff;color:#fff;">
                </div>
            </form>
          </div>
          <div class="footer">
            <button type="button" class="book-button" @click="bookSpot(currentSpot)">Book Spot</button>
            <br>
            <br>
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
      currentSpot: {},
      residents: [],
      names: [],
      phoneNumbers: [],
      roomNumbers: []
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
        var x = document.getElementById("duration").value
        x = parseInt(x)
        var occupiedSpot = {Space:spot.Space,Status:"occupied",Floor:"1",OccupantId:document.getElementById("id").value,duration:x}
        console.log(occupiedSpot)
        axios.post("/book-spot", occupiedSpot)
            .then(response => {
            console.log(response.data);
        })
        .catch(e => {
            console.error(e);
        })
        $('#myModal').modal('hide')
    },
    nameComplete(){
        autocomplete(document.getElementById("myInput"), this.names, this.residents);
    }

  },
  mounted() {
    axios.get("/spotstatus/1")
      .then(response => {
        console.log(response.data);
        this.spots = response.data;
      })
      .catch(e => {
        console.error(e);
      })
      axios.get("/residents")
      .then(response => {
        console.log(response.data);
        var i = 0;
        for(i = 0; i < response.data.length; i++){
            this.names[i] = response.data[i].firstName + " " + response.data[i].lastName;
        }
        this.residents = response.data;
        console.log(this.names);
      })
      .catch(e => {
        console.error(e);
      })
  }
}

function autocomplete(inp, arr, res) {
  /*the autocomplete function takes two arguments,
  the text field element and an array of possible autocompleted values:*/
  var currentFocus;
  /*execute a function when someone writes in the text field:*/
  inp.addEventListener("input", function(e) {
      var a, b, i, val = this.value;
      /*close any already open lists of autocompleted values*/
      closeAllLists();
      if (!val) { return false;}
      currentFocus = -1;
      /*create a DIV element that will contain the items (values):*/
      a = document.createElement("DIV");
      a.setAttribute("id", this.id + "autocomplete-list");
      a.setAttribute("class", "autocomplete-items");
      /*append the DIV element as a child of the autocomplete container:*/
      this.parentNode.appendChild(a);
      /*for each item in the array...*/
      for (i = 0; i < arr.length; i++) {
        /*check if the item starts with the same letters as the text field value:*/
        if (arr[i].substr(0, val.length).toUpperCase() == val.toUpperCase()) {
          /*create a DIV element for each matching element:*/
          b = document.createElement("DIV");
          /*make the matching letters bold:*/
          b.innerHTML = "<strong>" + arr[i].substr(0, val.length) + "</strong>";
          b.innerHTML += arr[i].substr(val.length);
          /*insert a input field that will hold the current array item's value:*/
          b.innerHTML += "<input type='hidden' value='" + arr[i] + "'>";
          /*execute a function when someone clicks on the item value (DIV element):*/
          b.addEventListener("click", function(e) {
              /*insert the value for the autocomplete text field:*/
              inp.value = this.getElementsByTagName("input")[0].value;




            var i = 0;
            if(inp.value != ""){
                for(i=0; i < res.length; i++){
                    if(inp.value == res[i].firstName + " " + res[i].lastName){
                        document.getElementById("myPhoneInput").value = res[i].phone;
                        document.getElementById("myRoomInput").value = res[i].room;
                        document.getElementById("id").value = res[i].id;
                    }
                }
            }




              /*close the list of autocompleted values,
              (or any other open lists of autocompleted values:*/
              closeAllLists();
          });
          a.appendChild(b);
        }
      }
  });
  /*execute a function presses a key on the keyboard:*/
  inp.addEventListener("keydown", function(e) {
      var x = document.getElementById(this.id + "autocomplete-list");
      if (x) x = x.getElementsByTagName("div");
      if (e.keyCode == 40) {
        /*If the arrow DOWN key is pressed,
        increase the currentFocus variable:*/
        currentFocus++;
        /*and and make the current item more visible:*/
        addActive(x);
      } else if (e.keyCode == 38) { //up
        /*If the arrow UP key is pressed,
        decrease the currentFocus variable:*/
        currentFocus--;
        /*and and make the current item more visible:*/
        addActive(x);
      } else if (e.keyCode == 13) {
        /*If the ENTER key is pressed, prevent the form from being submitted,*/
        e.preventDefault();
        if (currentFocus > -1) {
          /*and simulate a click on the "active" item:*/
          if (x) x[currentFocus].click();
        }
      }
  });
  function addActive(x) {
    /*a function to classify an item as "active":*/
    if (!x) return false;
    /*start by removing the "active" class on all items:*/
    removeActive(x);
    if (currentFocus >= x.length) currentFocus = 0;
    if (currentFocus < 0) currentFocus = (x.length - 1);
    /*add class "autocomplete-active":*/
    x[currentFocus].classList.add("autocomplete-active");
  }
  function removeActive(x) {
    /*a function to remove the "active" class from all autocomplete items:*/
    for (var i = 0; i < x.length; i++) {
      x[i].classList.remove("autocomplete-active");
    }
  }
  function closeAllLists(elmnt) {
    /*close all autocomplete lists in the document,
    except the one passed as an argument:*/
    var x = document.getElementsByClassName("autocomplete-items");
    for (var i = 0; i < x.length; i++) {
      if (elmnt != x[i] && elmnt != inp) {
        x[i].parentNode.removeChild(x[i]);
      }
    }
  }
  /*execute a function when someone clicks in the document:*/
  document.addEventListener("click", function (e) {
      closeAllLists(e.target);
      });
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
  color: #ffffff;
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
.available{
    fill:#B0DD9A;
    stroke-width:3;
    stroke:rgb(140,140,140);
}
.occupied{
    fill:#DCDD9B;
    stroke-width:3;
    stroke:rgb(140,140,140);
}
.exceeded{
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


* { box-sizing: border-box; }
body {
  font: 16px Arial;
}
.autocomplete {
  /*the container must be positioned relative:*/
  position: relative;
  display: inline-block;
}
input {
  border: 1px solid transparent;
  background-color: #f1f1f1;
  padding: 10px;
  font-size: 16px;
}
input[type=text] {
  background-color: #f1f1f1;
  width: 100%;
}
input[type=submit] {
  background-color: DodgerBlue;
  color: #fff;
}
.autocomplete-items {
  position: absolute;
  border: 1px solid #d4d4d4;
  border-bottom: none;
  border-top: none;
  z-index: 99;
  /*position the autocomplete items to be the same width as the container:*/
  top: 100%;
  left: 0;
  right: 0;
}
.autocomplete-items div {
  padding: 10px;
  cursor: pointer;
  background-color: #fff;
  border-bottom: 1px solid #d4d4d4;
}
.autocomplete-items div:hover {
  /*when hovering an item:*/
  background-color: #e9e9e9;
}
.autocomplete-active {
  /*when navigating through the items using the arrow keys:*/
  background-color: DodgerBlue !important;
  color: #ffffff;
}

/* mouse over link */
a:hover {
    color: #fff;
}
</style>
