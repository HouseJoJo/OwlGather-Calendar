const btn = document.querySelector(".btn-select");
const items = document.querySelectorAll(".option");

const today = new Date();
const month = today.getMonth();

function getMonthName() {
  const monthNames = ["January", "February", "March", "April", "May", "June",
                      "July", "August", "September", "October", "November", "December"];
  return monthNames[month];
}

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



