const btn = document.querySelector(".btn-select");
const items = document.querySelectorAll(".option");

const date = new Date();
const month = date.getMonth();

function getMonthName() {
  const monthNames = ["January", "February", "March", "April", "May", "June",
                      "July", "August", "September", "October", "November", "December"];
  return monthNames[month];
}

function populateMonth(){ //function to fill day numbers on calendar of current month.
  var date2 = new Date();
  var firstDay = new Date(date2.getFullYear(), date2.getMonth(), 1); //returns Date type of first day of month
  
  var weekNum = 1;
  var weekElement = document.querySelector(`.week${weekNum}`);
  
  var numFirstDay = firstDay.getDay(); // Returns int 0-6 : Sunday-Sat. -> MAP to Weekday of first week
  var dayElement = weekElement.querySelector(`.weekDay${numFirstDay}`);

  dayElement.innerHTML = firstDay.getDate();

  while(firstDay.getMonth() == month){
    firstDay.setDate(firstDay.getDate() + 1); //increment day
    if(firstDay.getDate() == 1) break;
    if(firstDay.getDay() == 1){ 
      weekNum += 1; //check if new week
      weekElement = document.querySelector(`.week${weekNum}`);
    }
    dayElement = weekElement.querySelector(`.weekDay${firstDay.getDay()}`);

    dayElement.innerHTML = firstDay.getDate();
  }
  console.log("Executed");
}

populateMonth();

//Initialize current month
document.querySelector(`.currentMonth`).innerHTML = getMonthName()

btn.addEventListener("click", ()=>{
  btn.classList.toggle("open");
});

items.forEach(item => {
  item.addEventListener("click", () => {
    item.classList.toggle("checked");
  });
});

const branches = document.querySelector(".btn-select1");
const branchItems = document.querySelectorAll(".option");
branches.addEventListener("click", ()=>{
  branches.classList.toggle("open");
});

const perk = document.querySelector(".btn-select2");
const perkItems = document.querySelectorAll(".option");
perk.addEventListener("click", ()=>{
  perk.classList.toggle("open");
});



